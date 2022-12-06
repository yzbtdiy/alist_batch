## Alist批量添加阿里云资源(Golang)

**此项目为 [https://github.com/yzbtdiy/alist_batch_add](https://github.com/yzbtdiy/alist_batch_add) 的 Golang实现**

* 使用 Golang 实现了 Alist 批量添加阿里云链接
* 自动获取并保存 token
* 操作前验证 cookie 有效性, cookie 无效自动更新
* 配置文件和阿里云资源文件使用 yaml 文件保存

#### 如果您不了解Alist, 请查看官网 [https://alist.nn.ci/zh/](https://alist.nn.ci/zh/)

### 用法说明



* Golang 编译的二进制文件可直接运行(alist_batch.exe)
  * 初次运行会自动生成配置模板和阿里云资源模板
  * 在 config.yaml 文件中添加 alist 地址, url 结尾不需要 /
  * 在 config.yaml 文件中 username 和 password 字段后添加 alist 登录账号和密码
  * 在 config.yaml 文件中 refresh_token 字段后添加阿里云盘的 refresh_token
  * 在 ali_share.yaml 文件中添加资源的分类
  * 在 ali_share.yaml 文件分类下级添加 `资源名: 阿里云资源链接` , 链接需要需要包含 folder
  * 修改后运行 alist_batch.exe 即可, 推荐命令行执行, 双击运行不会输出信息

* 下载源码编译
  * 代码目录 `go build -o alist_batch.exe main.go`

### other

* alist 的登录用户和密码仅用于自动获取 cookie, 手动获取有效cookie填入config.yaml可以不用添加用户和密码
* 目前只实现了不带提取码的阿里云盘链接批量添加
