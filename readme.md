# 中小学卷子生成程序

## 个人项目（命令行）

### 编译指南

1. 下载 Go SDK 1.19 及之后版本
2. 进入 cmd/cli 目录
3. 执行 `go build` 命令
4. 运行编译出的 `cli` 可执行文件

## 结对项目（Web）

可用实例请见 [Online Demo](https://math.cyp0633.icu/)。能发送邮件，但不能正常发短信（原因见下）。

### 编译指南

1. 下载 Go SDK 1.19 及之后版本、Node.js 最新 LTS、yarn
2. 进入 frontend 目录
3. 执行 `yarn install` 和 `yarn build` 命令
4. 进入 cmd/server 目录
5. 执行 `go build` 命令
6. 创建 conf.ini 文件，填入对应的数据
7. 运行编译出的 `server` 可执行文件
8. 打开 `http://localhost:port`，其中 port 为 conf.ini 中的端口号

### 配置文件格式

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
cookie_auth = "cookie 加密密钥"
```

### 注意事项

0. 如不编写配置文件，程序将在启动时直接 panic，这是正常现象。如不需要发送邮件或短信，也可以不配置对应内容。
1. 根据厂商限制，短信发送模板为阿里云测试模板，且只有绑定过的测试手机号才可接收短信。可以在 `internal/models/sms.go` 中修改。
2. 邮件服务器使用 SMTP TLS 465 端口。
