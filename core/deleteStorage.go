package core

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/yzbtdiy/alist_batch/flags"
	"github.com/yzbtdiy/alist_batch/models"
)

// 根据参数判断是否执行删除操作
func (a *AlistBatch) DeleteStorageIfHaveFlag() {
	storageData := a.GetStorgeData()
	needDeleteIds := a.GetDeleteStroageIds(storageData)
	a.DeleteStorageById(needDeleteIds)
	os.Exit(0)
}

// 获取需要删除的存储 id 列表
func (a *AlistBatch) GetDeleteStroageIds(storageListData *models.StorageListData) []string {
	var deleteIds []string
	if *flags.DeleteFlag == "dis" {
		println("尝试删除禁用存储")
		for _, mountInfo := range storageListData.Content {
			if mountInfo.Disabled {
				deleteIds = append(deleteIds, strconv.Itoa(mountInfo.Id))
			}
		}
	} else if *flags.DeleteFlag == "all" {
		println("尝试删除所有存储")
		for _, mountInfo := range storageListData.Content {
			deleteIds = append(deleteIds, strconv.Itoa(mountInfo.Id))
		}
	}
	return deleteIds
}

// 获取存储列表信息
func (a *AlistBatch) GetStorgeData() *models.StorageListData {
	storageListRes := a.client.Get(a.storageListApi)
	data, _ := json.Marshal(storageListRes.Data)
	var storageData *models.StorageListData
	json.Unmarshal(data, &storageData)
	return storageData
}

// 根据 id 列表删除存储
func (a *AlistBatch) DeleteStorageById(deleteIds []string) {
	for _, id := range deleteIds {
		resData := a.client.Post(a.delStorageApi+"?id="+id, []byte(""))
		if resData.Code == 200 {
			log.Println("Id为 " + id + " 的存储已删除")
		}
	}
}
