package funcs

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"
)

// 生成添加阿里云盘挂载json字符串
func BuildAliPushData(mountPath string, aliUrl string, config *models.Config) []byte {
	// reId, _ := regexp.Compile("https://www.aliyundrive.com/s/(.*)/folder")
	// reFolder, _ := regexp.Compile("/folder/(.*)$")
	// reShareId := reId.FindStringSubmatch(aliUrl)
	// reShareFolder := reFolder.FindStringSubmatch(aliUrl)
	// shareId := reShareId[len(reShareId)-1]
	// shareFolder := reShareFolder[len(reShareFolder)-1]
	params, err := url.Parse(aliUrl)
	if err != nil {
		log.Fatal(err)
	}

	var shareId, shareFolder, sharePwd string
	if params.Query().Has("pwd") {
		sharePwd = params.Query()["pwd"][0]
	} else {
		sharePwd = ""
	}
	pathArray := strings.Split(params.Path, "/")
	shareId = pathArray[2]
	shareFolder = pathArray[4]

	addition := new(models.AliAddition)
	addition.RefreshToken = config.Aliyun.RefreshToken
	addition.ShareId = shareId
	addition.SharePwd = sharePwd
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
		OrderBy:         "",
		OrderDirection:  "",
		ExtractFolder:   "",
		EnableSign:      false,
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
		OrderBy:         "",
		OrderDirection:  "",
		ExtractFolder:   "",
		EnableSign:      false,
		Driver:          "PikPakShare",
		Addition:        additionData,
	}
	pushJson, _ := json.Marshal(data)
	return pushJson
}

// 生成添加阿里云盘分享批量更新 RefreshToken json字符串
func BuildUpdateAliRefreshToken(aliShareData models.StorageListContent, refreshToken string) []byte {
	var oldAddition models.AliAddition
	json.Unmarshal([]byte(aliShareData.Addition), &oldAddition)
	oldAddition.RefreshToken = refreshToken
	aliShareData.Status = "work"
	newAddition, _ := json.Marshal(oldAddition)
	aliShareData.Addition = string(newAddition)
	pushJson, _ := json.Marshal(aliShareData)
	return pushJson
}

func BuildOnedriverApp(mountPath string, emailInfo string, config *models.Config) []byte {
	params := strings.Split(emailInfo, ":")
	var tid int
	var email, folderPath string
	if len(params) == 3 {
		tid, _ = strconv.Atoi(params[0])
		email = strings.Split(emailInfo, ":")[1]
		folderPath = strings.Split(emailInfo, ":")[2]
	} else if len(params) == 2 {
		tid, _ = strconv.Atoi(params[0])
		email = strings.Split(emailInfo, ":")[1]
		folderPath = "/"
	}

	addition := models.OnedriveAppAddition{
		RootFolderPath: folderPath,
		Region:         config.OneDriveApp.Region,
		ClientId:       config.OneDriveApp.Tenant[tid-1].ClientId,
		ClientSecret:   config.OneDriveApp.Tenant[tid-1].ClientSecret,
		TenantId:       config.OneDriveApp.Tenant[tid-1].TenantId,
		Email:          email,
		ChunkSize:      5,
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
		OrderBy:         "",
		OrderDirection:  "",
		ExtractFolder:   "",
		EnableSign:      false,
		Driver:          "OnedriveAPP",
		Addition:        additionData,
	}
	pushJson, _ := json.Marshal(data)
	return pushJson
}
