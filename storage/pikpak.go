package storage

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/yzbtdiy/alist_batch/models"
)

// BuildPikPakData 生成添加pikpak分享链接挂载的JSON数据
func BuildPikPakData(mountPath string, pikPakUrl string, config *models.Config) []byte {
	// 解析分享链接
	params, err := url.Parse(pikPakUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 提取分享密码
	sharePwd := ""
	if params.Query().Has("pwd") {
		sharePwd = params.Query().Get("pwd")
	}

	// 提取分享ID和文件夹路径
	pathArray := strings.Split(params.Path, "/")
	if len(pathArray) < 3 {
		log.Fatal("Invalid PikPak share URL format")
	}

	shareId := pathArray[2]
	shareFolder := ""
	if len(pathArray) >= 4 {
		shareFolder = pathArray[3]
	}

	// 构建附加信息
	addition := models.PikPakAddition{
		RootFolderId:          shareFolder,
		ShareId:               shareId,
		SharePwd:              sharePwd,
		OrderBy:               "",
		OrderDirection:        "",
		Platform:              "android",
		UseTranscodingAddress: true,
	}

	additionJson, err := json.Marshal(addition)
	if err != nil {
		log.Fatal("Failed to marshal PikPak addition:", err)
	}
	additionData := string(additionJson)

	// 构建完整的挂载数据
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

	pushJson, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to marshal push data:", err)
	}

	return pushJson
}
