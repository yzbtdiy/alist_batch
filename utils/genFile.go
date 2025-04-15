package utils

import (
	_ "embed" // 导入 embed 包
	"os"
)

// 使用 embed 嵌入模板文件
//
//go:embed template/config.yaml
var configYamlTemplate []byte

//go:embed template/ali_share.yaml
var aliShareYamlTemplate []byte

//go:embed template/pik_share.yaml
var pikShareYamlTemplate []byte

//go:embed template/onedrive_app.yaml
var onedriveAppYamlTemplate []byte

// 生成配置模板文件
// 通过 embed 读取模板内容并生成 config.yaml 文件
func GenConfFile(fileName string) {
	err := os.WriteFile("./"+fileName, configYamlTemplate, 0777)
	if err != nil {
		panic(err)
	}
}

// 生成阿里云盘分享链接模板文件
// 通过 embed 读取模板内容并生成 ali_share.yaml 文件
func GenAliShareFile(fileName string) {
	err := os.WriteFile("./"+fileName, aliShareYamlTemplate, 0777)
	if err != nil {
		panic(err)
	}
}

// 生成 PikPak 分享链接模板文件
// 通过 embed 读取模板内容并生成 pik_share.yaml 文件
func GenPikShareFile(fileName string) {
	err := os.WriteFile("./"+fileName, pikShareYamlTemplate, 0777)
	if err != nil {
		panic(err)
	}
}

// 生成 Onedrive APP 分享链接模板文件
// 通过 embed 读取模板内容并生成 onedrive_app.yaml 文件
func GenOnedriveAppFile(fileName string) {
	err := os.WriteFile("./"+fileName, onedriveAppYamlTemplate, 0777)
	if err != nil {
		panic(err)
	}
}
