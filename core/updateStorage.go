package core

import (
	"log"
	"os"
	"strconv"

	"github.com/yzbtdiy/alist_batch/flags"
	"github.com/yzbtdiy/alist_batch/models"
)

// 更新阿里云盘 refresh token
func (a *AlistBatch) UpdateAliStorage() {
	if *flags.UpdateFlag == "ali" {
		storageListData := a.GetStorgeData()
		a.updateAliRefreshToken(a.updateStorageApi, storageListData)
	}
	os.Exit(0)
}

// 批量更新阿里云盘存储 refresh token
func (a *AlistBatch) updateAliRefreshToken(updateStorageApi string, storageListData *models.StorageListData) {
	for _, mountInfo := range storageListData.Content {
		if mountInfo.Driver == "AliyundriveShare" && mountInfo.Status != "work" {
			pushData := a.BuildUpdateAliRefreshToken(mountInfo, a.config.Aliyun.RefreshToken)
			resData := a.client.Post(updateStorageApi, pushData)
			if resData.Code == 200 {
				log.Println("Id 为 " + strconv.Itoa(mountInfo.Id) + " 存储的 refresh_token 已更新")
			} else {
				log.Println("Id 为 " + strconv.Itoa(mountInfo.Id) + " 存储的 refresh_token 更新失败, 请检查是否有效")
			}
		}
	}
}
