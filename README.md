## 简单学习记录

### 安装包
go get <package_path_or_url>

### 启动
- 进入 /src/main/
- go run main.go

### 打包
- 进入 /src/main/
- GOOS=linux GOARCH=amd64 go build -o hero-ranking-server main.go
