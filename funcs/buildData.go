package funcs

import (
	"github.com/yzbtdiy/alist_batch/models"

	"encoding/json"
	"regexp"
)

// 生成添加阿里云盘挂载json字符串
func BuildPushData(mountPath string, aliUrl string, config *models.Config) []byte {
	reId, _ := regexp.Compile("https://www.aliyundrive.com/s/(.*)/folder")
	reFolder, _ := regexp.Compile("/folder/(.*)$")
	reShareId := reId.FindStringSubmatch(aliUrl)
	reShareFolder := reFolder.FindStringSubmatch(aliUrl)
	shareId := reShareId[len(reShareId)-1]
	shareFolder := reShareFolder[len(reShareFolder)-1]
	addition := new(models.Addition)
	addition.RefreshToken = config.RefreshToken
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
