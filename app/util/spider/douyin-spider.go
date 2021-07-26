package spider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cilidm/dy-spider/app/global"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/cilidm/dy-spider/app/model/dao"
	"github.com/cilidm/dy-spider/app/util/emoji"
	"github.com/cilidm/dy-spider/app/util/zapLog"
	"github.com/cilidm/toolbox/file"
	"github.com/kirinlabs/HttpRequest"
	"github.com/tidwall/gjson"
)

var req *HttpRequest.Request
var downloadDir = "runtime/file/spider"

func init() {
	req = HttpRequest.NewRequest()
	req.SetHeaders(map[string]string{
		"User-Agent": "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36",
	})
	req.CheckRedirect(func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse /* 不进入重定向 */
	})
}

func SpiderDY(year bool, lines []string, down int) {
	for _, line := range lines {
		if line == "" {
			continue
		}
		zapLog.NewLog().Info("SpiderDY", line, "开始执行")
		reg := regexp.MustCompile(`[a-z]+://[\S]+`)
		url := reg.FindAllString(line, -1)[0]
		resp, err := req.Get(url)
		defer resp.Close()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if resp.StatusCode() != 302 {
			continue
		}
		location := resp.Headers().Values("location")[0]
		regNew := regexp.MustCompile(`(?:sec_uid=)[a-z,A-Z，0-9, _, -]+`)
		sec_uid := strings.Replace(regNew.FindAllString(location, -1)[0], "sec_uid=", "", 1)

		respIes, err := req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/user/info/?sec_uid=%s", sec_uid))
		defer respIes.Close()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		body, err := respIes.Body()
		var spUser model.SpiderUserForm
		err = json.Unmarshal(body, &spUser)
		if err != nil {
			//zapLog.NewLog().Error("spider", "json.Unmarshal", string(body))
			zapLog.NewLog().Error("spider", "json.Unmarshal", err.Error())
			continue
		}
		//spUid, err := SaveSpiderUser(spUser)
		//if err != nil {
		//	zapLog.NewLog().Error("spider", "SaveSpiderUser", err.Error())
		//	continue
		//}

		result := gjson.Get(string(body), "user_info.nickname").String()
		dirPath := fmt.Sprintf("%s/%s/", downloadDir, result)
		err = file.IsNotExistMkDir(dirPath)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		GetByMonthV2(line, spUser.UserInfo.Nickname, sec_uid, dirPath, year, down)
	}
}

func GetByMonthV2(line, nickname, sec_uid, dirPath string, year bool, down int) {
	y := 2018
	nowY, _ := strconv.Atoi(time.Now().Format("2006"))
	nowM, _ := strconv.Atoi(time.Now().Format("01"))
	for i := y; i <= nowY; i++ {
		for m := 1; m <= 12; m++ {
			var (
				begin int64
				end   int64
			)
			if i == nowY && m > nowM {
				break
			}
			begin = GetMonthStart(strconv.Itoa(i), strconv.Itoa(m))
			if m == 12 {
				end = GetMonthStart(strconv.Itoa(i+1), "1") // 12月时年份+1 月份变为1月
			} else {
				end = GetMonthStart(strconv.Itoa(i), strconv.Itoa(m+1))
			}
			resp, err := req.Get(fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/post/?sec_uid=%s&count=200&min_cursor=%d&max_cursor=%d&aid=1128&_signature=PtCNCgAAXljWCq93QOKsFT7QjR",
				sec_uid, begin, end))
			defer resp.Close()
			if err != nil {
				log.Println(err.Error())
				continue
			}
			body, err := resp.Body()

			// 废弃gjson 使用结构体解析
			var form model.AwemeListForm
			err = json.Unmarshal(body, &form)
			if err != nil {
				//zapLog.NewLog().Error("GetByMonth", "json.Unmarshal", string(body))
				zapLog.NewLog().Error("GetByMonth", "json.Unmarshal", err.Error())
				continue
			}
			zapLog.NewLog().Info(line, nickname, fmt.Sprintf("%d年%d月，视频数量%d", i, m, len(form.AwemeList)))
			if len(form.AwemeList) > 0 {
				//CreateAwemeList(form)
				var newDirPath string
				if year {
					newDirPath = fmt.Sprintf("%s%d/", dirPath, i)
					err = file.IsNotExistMkDir(newDirPath)
					if err != nil {
						zapLog.NewLog().Error("GetByMonth", "IsNotExistMkDir", err.Error())
						continue
					}
				} else {
					newDirPath = dirPath
				}
				for n := 0; n < len(form.AwemeList); n++ {
					hasSpider, err := dao.NewSpiderDaoImpl().FindSpiderByAwemeId(form.AwemeList[n].AwemeID)
					if err != nil {
						global.ZapLog.Error(err.Error())
						continue
					}
					if hasSpider != nil && hasSpider.HasDown == 1 { // 数据库已存在此条记录并且已经下载
						if !file.CheckNotExist(hasSpider.SavePath) {
							global.ZapLog.Info("【" + hasSpider.Info + "】视频已存在")
							continue
						}
					}
					videotitle := form.AwemeList[n].Desc
					if videotitle == "" {
						videotitle = form.AwemeList[n].AwemeID
					}
					videotitle = ReplaceBeforeSave(videotitle)
					videourl := form.AwemeList[n].Video.PlayAddr.URLList[0]
					if down == 1 {
						err = DownloadFile(newDirPath+videotitle+".mp4", videourl)
						if err != nil {
							global.ZapLog.Error(err.Error())
							continue
						}
					}
					cover := form.AwemeList[n].Video.Cover.URLList[0]
					if hasSpider != nil {
						err = dao.NewSpiderDaoImpl().Update(hasSpider.ID, map[string]interface{}{
							"has_down":  down,
							"save_path": newDirPath + videotitle + ".mp4",
						})
						if err != nil {
							global.ZapLog.Error(err.Error())
							continue
						}
					} else {
						online := ""
						if len(form.AwemeList[n].Video.PlayAddrLowbr.URLList) > 3 {
							online = form.AwemeList[n].Video.PlayAddrLowbr.URLList[3]
						}
						_, err = dao.NewSpiderDaoImpl().InsertSpider(model.Spider{
							Stype:     model.DySpider,
							UserName:  emoji.Encode(nickname),
							Url:       line,
							SavePath:  newDirPath + videotitle + ".mp4",
							Info:      videotitle,
							Cover:     cover,
							OnlineUrl: online,
							AwemeID:   form.AwemeList[n].AwemeID,
							HasDown:   down,
						})
						if err != nil {
							zapLog.NewLog().Error("GetByMonth", "InsertSpider", err.Error())
							continue
						}
					}
				}
			}
		}
	}
}

func SaveSpiderUser(spUser model.SpiderUserForm) (uint, error) {
	u, err := dao.NewSpiderUserDaoImpl().FindUserByUid(spUser.UserInfo.UID)
	if u != nil || err != nil {
		return u.ID, nil
	}
	var sp model.SpiderUserInfo
	sp.UID = spUser.UserInfo.UID
	sp.ShortID = spUser.UserInfo.ShortID
	sp.UniqueID = spUser.UserInfo.UniqueID
	sp.Nickname = spUser.UserInfo.Nickname
	sp.AwemeCount = spUser.UserInfo.AwemeCount
	sp.TotalFavorited = spUser.UserInfo.TotalFavorited
	sp.FollowersDetail = spUser.UserInfo.FollowersDetail
	sp.Region = spUser.UserInfo.Region
	sp.FollowerCount = spUser.UserInfo.FollowerCount
	sp.CustomVerify = spUser.UserInfo.CustomVerify
	sp.FollowingCount = spUser.UserInfo.FollowingCount
	sp.FavoritingCount = spUser.UserInfo.FavoritingCount
	sp.VerificationType = spUser.UserInfo.VerificationType
	sp.PlatformSyncInfo = spUser.UserInfo.PlatformSyncInfo
	sp.Geofencing = spUser.UserInfo.Geofencing
	if spUser.UserInfo.IsGovMediaVip == false {
		sp.IsGovMediaVip = 0
	} else {
		sp.IsGovMediaVip = 1
	}
	avatarLarger, _ := json.Marshal(spUser.UserInfo.AvatarLarger)
	sp.AvatarLarger = string(avatarLarger)
	avatarThumb, _ := json.Marshal(spUser.UserInfo.AvatarThumb)
	sp.AvatarThumb = string(avatarThumb)
	policyVersion, _ := json.Marshal(spUser.UserInfo.PolicyVersion)
	sp.PolicyVersion = string(policyVersion)
	typeLabel, _ := json.Marshal(spUser.UserInfo.TypeLabel)
	sp.TypeLabel = string(typeLabel)
	avatarMedium, _ := json.Marshal(spUser.UserInfo.AvatarMedium)
	sp.AvatarMedium = string(avatarMedium)
	originalMusician, _ := json.Marshal(spUser.UserInfo.OriginalMusician)
	sp.OriginalMusician = string(originalMusician)
	sp.Secret = spUser.UserInfo.Secret
	sp.Signature = spUser.UserInfo.Signature
	id, err := dao.NewSpiderUserDaoImpl().InsertUser(sp)
	return id, err
}

func CreateAwemeList(form model.AwemeListForm) {
	for _, aw := range form.AwemeList {
		u, err := dao.NewSpiderAwemeListDaoImpl().FindByAwemeId(aw.AwemeID)
		if err != nil {
			zapLog.NewLog().Error("CreateAwemeList", "NewSpiderAwemeListDaoImpl", err.Error())
			continue
		}
		if u != nil {
			zapLog.NewLog().Error("视频信息已存在", "FindByAwemeId:"+aw.AwemeID, aw)
			continue
		}
		var al model.AwemeList
		al.AwemeID = aw.AwemeID
		al.AwemeType = aw.AwemeType
		al.Desc = aw.Desc
		al.ChaList, _ = ChangeToJson(aw.ChaList)
		al.CommentList, _ = ChangeToJson(aw.CommentList)
		al.VideoLabels, _ = ChangeToJson(aw.VideoLabels)
		al.ImageInfos, _ = ChangeToJson(aw.ImageInfos)
		al.Geofencing, _ = ChangeToJson(aw.Geofencing)
		al.Video, _ = ChangeToJson(aw.Video)
		al.Statistics, _ = ChangeToJson(aw.Statistics)
		al.TextExtra, _ = ChangeToJson(aw.TextExtra)
		al.Promotions, _ = ChangeToJson(aw.Promotions)
		al.LongVideo, _ = ChangeToJson(aw.LongVideo)
		al.LabelTopText, _ = ChangeToJson(aw.LabelTopText)
		al.Images, _ = ChangeToJson(aw.Images)
		al.LabelTopText, _ = ChangeToJson(aw.LabelTopText)
		al.Author, _ = ChangeToJson(aw.Author)
		al.VideoText, _ = ChangeToJson(aw.VideoText)
		_, err = dao.NewSpiderAwemeListDaoImpl().InsertAwemeList(al)
		if err != nil {
			zapLog.NewLog().Error("CreateAwemeList", "InsertAwemeList", err.Error())
			continue
		}
	}
}

func ChangeToJson(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	return string(b), err
}

func ReplaceBeforeSave(desc string) string {
	temp := strings.TrimSpace(desc)
	strList := []string{"~", "～", "#", "!", "$", "@", "%", "^", "&", "*", "(", ")", " ", "+", "！", ":", "/", "\\", "?", ">", "<", "|", "'", "\""}
	for _, v := range strList {
		temp = strings.ReplaceAll(temp, v, "")
	}
	return temp
}
