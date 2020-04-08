.PHONY: vet fmt lint build docker

vet:
	# go vet 静态错误检查
	go vet ./...

fmt:
	go fmt .

lint:
	# 检查代码规范
	@# 目测前期不规范处太多, 看看就好
	golint $$(go list ./...)

build:
    # 编译成可执行文件
	go build -mod vendor -o ./app/main .

docker: Dockerfile
    # 编译成镜像
	docker build -f ./Dockerfile .
