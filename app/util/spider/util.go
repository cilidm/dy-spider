package spider

import (
	"bufio"
	"fmt"
	"github.com/cilidm/toolbox/file"
	"github.com/dustin/go-humanize"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 下载文件显示进度条
type WriteCounter struct {
	Name  string
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\r 【%s】 Downloading... %s complete", wc.Name, humanize.Bytes(wc.Total))
}

func DownloadFile(fileName string, url string) error {
	err := BeginDownload(fileName, url)
	if err != nil {
		return err
	}
	if err = os.Rename(fileName+".tmp", fileName); err != nil {
		return err
	}
	return nil
}

func BeginDownload(fileName string, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36")
	resp, err := (&http.Client{Timeout: time.Second * 60}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fileNameSp := strings.Split(fileName, "/")

	if !file.CheckNotExist(fileName) {
		stat, err := os.Stat(fileName)
		if err != nil {
			return err
		}
		newSize := resp.Header.Get("Content-Length")
		if newSize == strconv.FormatInt(stat.Size(), 10) {
			log.Println(fileNameSp[len(fileNameSp)-1], "已存在，跳过")
			return nil
		}
	}

	out, err := os.Create(fileName + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	counter := &WriteCounter{}
	counter.Name = fileNameSp[len(fileNameSp)-1]
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		return err
	}
	fmt.Print("\n")
	return nil
}

//获取月份的第一天和最后一天
func GetMonthStart(myYear string, myMonth string) int64 {
	// 数字月份必须前置补零
	if len(myMonth) == 1 {
		myMonth = "0" + myMonth
	}
	yInt, _ := strconv.Atoi(myYear)

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, myYear+"-"+myMonth+"-01 00:00:00", loc)
	newMonth := theTime.Month()

	t1 := time.Date(yInt, newMonth, 1, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
	return t1
}

// 按行读取配置
func ReadLine(fileName string) (lines []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lines = append(lines, line)
		if err != nil {
			if err == io.EOF {
				return lines, nil
			}
			return nil, err
		}
	}
}
