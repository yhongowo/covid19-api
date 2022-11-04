#交叉编译，目标平台linux
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o covid19_api main.go