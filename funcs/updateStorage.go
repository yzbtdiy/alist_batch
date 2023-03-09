package funcs

import (
	"log"
	"os"
	"strconv"

	"github.com/yzbtdiy/alist_batch/flags"
	"github.com/yzbtdiy/alist_batch/models"
)

// 更新阿里云盘 refresh token
func UpdateAliStorage(storageListApi, updateStorageApi string, conf *models.Config) {
	if *flags.UpdateFlag == "ali" {
		storageListData := getStorgeData(storageListApi, conf.Token)
		updateAliRefreshToken(updateStorageApi, storageListData, conf)
	}
	os.Exit(0)
}

// 发送请求更新 status 不是 work 的存储
func updateAliRefreshToken(updateStorageApi string, storageListData *models.StorageListData, conf *models.Config) {
	for _, mountInfo := range storageListData.Content {
		if mountInfo.Driver == "AliyundriveShare" && mountInfo.Status != "work" {
			pushData := BuildUpdateAliRefreshToken(mountInfo, conf.Aliyun.RefreshToken)
			resData := HttpPost(updateStorageApi, conf.Token, pushData)
			if resData.Code == 200 {
				log.Println("Id 为 " + strconv.Itoa(mountInfo.Id) + " 存储的 refresh_token 已更新")
			} else {
				log.Println("Id 为 " + strconv.Itoa(mountInfo.Id) + " 存储的 refresh_token 更新失败, 请检查是否有效")
			}
		}
	}
}
