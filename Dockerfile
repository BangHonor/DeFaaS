# Dockerfile 实验中...

# 构建 cli
FROM tetafro/golang-gcc:1.13-alpine AS build
WORKDIR /defaas-cli-build
COPY . .
# Go 交叉编译
# - 目标操作系统 GOOS=linux
# - 目标架构     GOARCH=amd64
RUN GO111MODULE=on GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o defaas cli/app/main.go

# 在运行时挂载一个主机目录，该主机目录是使用者想保存 defaas 内容的目录
# docker run --rm --mount type=bind,source=/host/path,target=/defaas kitchen233/defaas:v0.0.1
FROM alpine AS final
# /defaas-pre 目录保存预准备的多个文件，相当于文件夹打包
WORKDIR /defaas-pre
# 工具
COPY ./tools/get_account.sh tools/get_account.sh
# cli 二进制文件
COPY --from=build /defaas-cli-build/defaas .
# 还需要 copy 配置和证书
# COPY TODO
# 把文件复制到挂载的主机目录中
CMD ["cp", "-r", ".", "/defaas"]