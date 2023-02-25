## Alist批量添加阿里云资源(Golang)

**此项目为 [https://github.com/yzbtdiy/alist_batch_add](https://github.com/yzbtdiy/alist_batch_add) 的 Golang实现**

* 使用 Golang 实现了 Alist 批量添加阿里云链接
* 自动获取并保存 token
* 操作前验证 cookie 有效性, cookie 无效自动更新
* 配置文件和阿里云资源文件使用 yaml 文件保存

网友**DayoWong0**提供了一个油猴脚本可以从浏览器打开的阿里云盘链接抓取资源名和链接, 大家可以试试 [脚本地址](https://greasyfork.org/zh-CN/scripts/457223-%E5%A4%8D%E5%88%B6%E4%B8%BA%E6%B7%BB%E5%8A%A0%E5%88%B0alist%E9%98%BF%E9%87%8C%E4%BA%91%E7%9B%98%E5%88%86%E4%BA%AB%E9%93%BE%E6%8E%A5%E7%9A%84%E6%A0%BC%E5%BC%8F)

#### 如果您不了解Alist, 请查看官网 [https://alist.nn.ci/zh/](https://alist.nn.ci/zh/)

### 用法说明

[Bilibili视频介绍](https://www.bilibili.com/video/BV1uP411K747)

* Golang 编译的二进制文件可直接运行(alist_batch.exe)
  * 初次运行会自动生成配置模板 config.yaml
  * 在 config.yaml 文件中添加 alist 地址, url 结尾不需要 /
  * 在 config.yaml 文件中 username 和 password 字段后添加 alist 登录账号和密码
  * 添加阿里云盘分享链接请设置aliyun下的enable为true, 添加pikpak分享链接请设置pikpak下的enable为true
    * 添加阿里云盘分享在配置文件 aliyun 下 refresh_token 字段后添加阿里云盘的 refresh_token
    * 添加pikpak分享在配置文件 pikpak 下的 username 和 password 字段后添加pikpak用户和密码
    * 支持有提取码的分享链接, 在链接结尾添加 `?pwd=提取码` 即可(阿里云盘和pikpak分享链接都支持)
  * 在 ali_share.yaml 或者 pik_share.yaml 文件中添加资源的分类
  * 在 ali_share.yaml 或 pik_share.yaml 文件分类下级添加 `资源名: 分享资源链接`
    * 阿里云盘分享链接需要包含 folder 后的 id, pikpak 分享链接可以不添加 root_folder_id
  * 修改后运行 alist_batch.exe 即可, 推荐命令行执行, 双击运行不会输出信息
  * `alist_batch.exe -delete dis` 删除已禁用存储
  * `alist_batch.exe -delete all` 删除所有添加的存储(慎用)

* 下载源码编译
  * `git clone https://github.com/yzbtdiy/alist_batch.git`
  * `cd alist_batch`
  * `go mod tidy`
  * `go build .`

* 使用 `go install` 安装
  * `go install github.com/yzbtdiy/alist_batch@latest`

### other

* alist 的登录用户和密码仅用于自动获取 cookie, 手动获取有效cookie填入config.yaml可以不用添加用户和密码
* 目前实现了阿里云盘分享链接和 PikPak 分享链接的批量添加
* 此工具只是批量挂载工具, 挂载后视频不能播放请关注alist的github issues
