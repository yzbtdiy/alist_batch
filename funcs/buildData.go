package funcs

import (
	"log"
	"net/url"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"

	"encoding/json"
	"regexp"
)

// 生成添加阿里云盘挂载json字符串
func BuildAliPushData(mountPath string, aliUrl string, config *models.Config) []byte {
	reId, _ := regexp.Compile("https://www.aliyundrive.com/s/(.*)/folder")
	reFolder, _ := regexp.Compile("/folder/(.*)$")
	reShareId := reId.FindStringSubmatch(aliUrl)
	reShareFolder := reFolder.FindStringSubmatch(aliUrl)
	shareId := reShareId[len(reShareId)-1]
	shareFolder := reShareFolder[len(reShareFolder)-1]

	addition := new(models.AliAddition)
	addition.RefreshToken = config.Aliyun.RefreshToken
	addition.ShareId = shareId
	addition.SharePwd = ""
	addition.RootFolderId = shareFolder
	addition.OrderBy = ""
	addition.OrderDirection = ""

	additionJson, _ := json.Marshal(addition)
	additionData := string(additionJson)

	data := models.PushData{
		MountPath:       mountPath,
		Order:           0,
		Remark:          "",
		CacheExpiration: 30,
		WebProxy:        false,
		WebdavPolicy:    "302_redirect",
		DownProxyUrl:    "",
		ExtractFolder:   "",
		Driver:          "AliyundriveShare",
		Addition:        additionData,
	}
	pushJson, _ := json.Marshal(data)
	return pushJson
}

// 生成添加pikpak分享链接挂载json字符串
func BuildPikPakData(mountPath string, pikPakUrl string, config *models.Config) []byte {
	params, err := url.Parse(pikPakUrl)
	if err != nil {
		log.Fatal(err)
	}

	var sharePwd, shareFolder string
	if params.Query().Has("pwd") {
		sharePwd = params.Query()["pwd"][0]
	} else {
		sharePwd = ""
	}
	pathArray := strings.Split(params.Path, "/")
	shareId := pathArray[2]
	if len(pathArray) < 4 {
		shareFolder = ""
	} else {
		shareFolder = pathArray[3]
	}

	addition := models.PikPakAddition{
		RootFolderId:   shareFolder,
		Username:       config.PikPak.Username,
		Password:       config.PikPak.Password,
		ShareId:        shareId,
		SharePwd:       sharePwd,
		OrderBy:        "",
		OrderDirection: "",
	}
	additionJson, _ := json.Marshal(addition)
	additionData := string(additionJson)

	data := models.PushData{
		MountPath:       mountPath,
		Order:           0,
		Remark:          "",
		CacheExpiration: 30,
		WebProxy:        false,
		WebdavPolicy:    "302_redirect",
		DownProxyUrl:    "",
		ExtractFolder:   "",
		Driver:          "PikPakShare",
		Addition:        additionData,
	}
	pushJson, _ := json.Marshal(data)
	return pushJson
}
