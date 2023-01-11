package funcs

import (
    "github.com/yzbtdiy/alist_batch/models"

	"os"
	"fmt"

	"gopkg.in/yaml.v3"
)

// 判断文件是否存在
func CheckFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 读取配置文件
func GetConfig(fileName string) *models.Config {
	// _conf := models.Config{}
	var _conf *models.Config
	content, err := os.ReadFile("./" + fileName)
	if err != nil {
		fmt.Println("读取配置文件出错")
		fmt.Println(err)
	}
	err = yaml.Unmarshal(content, &_conf)
	if err != nil {
		fmt.Println(err)
	}
	return _conf
}

// 修改配置文件, 添加token
func ModConfig(fileName string, oldConf *models.Config, token string) {
	oldConf.Token = token
	newConf, err := yaml.Marshal(oldConf)
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile("./"+fileName, newConf, 0o777)
}

// 读取 ali_share.yaml 文件
func GetShareList(fileName string) map[string]map[string]string {
	shareListContent := make(map[string]map[string]string)
	content, err := os.ReadFile("./" + fileName)
	if err != nil {
		fmt.Println("读取阿里分享列表文件出错")
		fmt.Println(err)
	}
	yaml.Unmarshal(content, &shareListContent)
	return shareListContent
}