package funcs

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/yzbtdiy/alist_batch/flags"
	"github.com/yzbtdiy/alist_batch/models"
)

// 根据传入参数判断是否执行删除操作
func DeleteStorageIfHaveFlag(storageListApi, delStorageApi, token string) {
	storageData := getStorgeData(storageListApi, token)
	needDeleteIds := GetDeleteStroageIds(storageData)
	DeleteStorageById(delStorageApi, needDeleteIds, token)
	os.Exit(0)
}

// 获取需要删除的存储 id, 生成列表
func GetDeleteStroageIds(storageListData *models.StorageListData) []string {
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
func getStorgeData(storageListApi string, token string) *models.StorageListData {
	storageListRes := HttpGet(storageListApi, token)
	data, _ := json.Marshal(storageListRes.Data)
	var storageData *models.StorageListData
	json.Unmarshal(data, &storageData)
	return storageData
}

// 根据 id 列表发送请求删除存储
func DeleteStorageById(delStorageApi string, deleteIds []string, token string) {
	for _, id := range deleteIds {
		resData := HttpPost(delStorageApi+"?id="+id, token, []byte(""))
		if resData.Code == 200 {
			log.Println("Id为 " + id + " 的存储已删除")
		}
	}
}
