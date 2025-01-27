# 版本号
version = ''

# 编译到 Linux
.PHONY: build-linux
build-linux:
	mkdir -p ./build/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/linux/go-pass ./src/main.go 

# 编译到 macOS
# make build-darwin
.PHONY: build-darwin
build-darwin:
	mkdir -p ./build/darwin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/darwin/go-pass ./src/main.go

# 编译到 windows
.PHONY: build-windows
build-windows:
	mkdir -p ./build/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/windows/go-pass.exe ./src/main.go 

# 编译到 全部平台
# make build
.PHONY: build
build:
	make clean
	make build-linux
	make build-darwin
	make build-windows


# 发布linux
.PHONY: release-linux
release-linux:
	mkdir -p ./release/linux
	cp ./env.example ./release/linux/
	cp ./build/linux/go-pass ./release/linux/
	tar -zcvf release/go-pass-$(version)-linux-amd64.tar.gz -C ./release/linux/ .

# 发布darwin
.PHONY: release-darwin
release-darwin:
	mkdir -p ./release/darwin
	cp ./env.example ./release/darwin/
	cp ./build/darwin/go-pass ./release/darwin/
	tar -zcvf release/go-pass-$(version)-darwin-amd64.tar.gz -C ./release/darwin/ .

# 发布windows
.PHONY: release-windows
release-windows:
	mkdir -p ./release/windows
	cp ./env.example ./release/windows/
	cp ./build/windows/go-pass.exe ./release/windows/
	zip -j release/go-pass-$(version)-windows-amd64.zip ./release/windows/*

# 发布全平台
# make release
.PHONY: release
release:
	make release-linux
	make release-darwin
	make release-windows


.PHONY: clean
clean:
	rm -rf ./build ./release

.PHONY: dev
dev:
	go run ./src/main.go
	# air