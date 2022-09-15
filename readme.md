# 中小学卷子生成程序

## 编译指南

个人项目（命令行）

1. 下载 Go SDK 1.19 及之后版本
2. 进入 cmd/cli 目录
3. 执行 `go build` 命令
4. 运行编译出的 `cli` 可执行文件

结对项目（Web，未完成）

1. 下载 Go SDK 1.19 及之后版本、Node.js 最新 LTS、yarn
2. 进入 frontend 目录
3. 备份 frontend/dist/frontend.go 文件
4. 执行 `yarn build` 命令
5. 恢复 frontend/dist/frontend.go 文件
6. 进入 cmd/server 目录
7. 执行 `go build` 命令
8. 创建 conf.ini 文件
9. 运行编译出的 `server` 可执行文件
10. 打开 `http://localhost:8000`

## 配置文件格式

```ini
[mail]
identity = "默认为空即可"
username = "邮箱地址"
password = "邮箱密码"
host = "邮件服务器地址"

[sms]
access_key_id = "阿里云短信服务 AccessKeyId"
access_key_secret = "阿里云短信服务 AccessKeySecret"

[server]
port = 8000 # 服务器端口
```
