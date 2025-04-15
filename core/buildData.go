package core

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"
)

// 构建阿里云盘挂载数据
func (a *AlistBatch) BuildAliPushData(mountPath string, aliUrl string) []byte {
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
	addition.RefreshToken = a.config.Aliyun.RefreshToken
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

// 构建 PikPak 挂载数据
func (a *AlistBatch) BuildPikPakData(mountPath string, pikPakUrl string) []byte {
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
		RootFolderId:          shareFolder,
		ShareId:               shareId,
		SharePwd:              sharePwd,
		OrderBy:               "",
		OrderDirection:        "",
		Platform:              "android",
		UseTranscodingAddress: true,
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

// 构建阿里云盘批量更新 RefreshToken 数据
func (a *AlistBatch) BuildUpdateAliRefreshToken(aliShareData models.StorageListContent, refreshToken string) []byte {
	var oldAddition models.AliAddition
	json.Unmarshal([]byte(aliShareData.Addition), &oldAddition)
	oldAddition.RefreshToken = refreshToken
	aliShareData.Status = "work"
	newAddition, _ := json.Marshal(oldAddition)
	aliShareData.Addition = string(newAddition)
	pushJson, _ := json.Marshal(aliShareData)
	return pushJson
}

// 构建 Onedrive APP 挂载数据
func (a *AlistBatch) BuildOnedriverApp(mountPath string, emailInfo string) []byte {
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
		Region:         a.config.OneDriveApp.Region,
		ClientId:       a.config.OneDriveApp.Tenants[tid-1].ClientId,
		ClientSecret:   a.config.OneDriveApp.Tenants[tid-1].ClientSecret,
		TenantId:       a.config.OneDriveApp.Tenants[tid-1].TenantId,
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
