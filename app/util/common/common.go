package common

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// base64转img
func Base2Img(baseCode, fileName string) error {
	newBase := strings.Replace(baseCode, "data:image/png;base64,", "", 1)
	dist, err := base64.StdEncoding.DecodeString(newBase)
	if err != nil {
		return err
	}
	f, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	return nil
}

// 根据链接地址获取网址
func GetDomain(url string) string {
	head := strings.Split(url, "//")[0]
	link := strings.Split(url, "//")[1]
	return head + "//" + strings.Split(link, "/")[0]
}

// 判断网址是否是ip地址 不是ip返回nil
func ParseIp(str string) net.IP {
	shortIp := strings.Split(str, "//")[1]   // 去掉http://
	noPort := strings.Split(shortIp, ":")[0] // 去掉端口
	return net.ParseIP(noPort)
}

//	读取文件中的数据
func ReadJson(filePath string) (result string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	buf := bufio.NewReader(file)
	for {
		s, err := buf.ReadString('\n')
		result += s
		if err != nil {
			if err == io.EOF {
				fmt.Println("Read is ok")
				break
			} else {
				fmt.Println("ERROR:", err)
				return
			}
		}
	}
	return result
}

func ArrayToStr(arrayOrSlice interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(arrayOrSlice), "[]"), " ", ",", -1)
}
