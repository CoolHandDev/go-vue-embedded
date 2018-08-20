all: get-esc create-build-workspace build-vue-app build-go

.PHONY: get-esc
get-esc: 
	go get -u github.com/mjibson/esc

.PHONY: create-build-workspace
create-build-workspace: 
	rm -rf build-workspace
	mkdir -p build-workspace

.PHONY: build-vue-app
build-vue-app: 
	npm run --prefix front-end build
	cp -rf ./front-end/dist/* ./build-workspace/

.PHONY: build-go
build-go:
	cp main.go ./build-workspace
	cd ./build-workspace
	go generate ./build-workspace/main.go
	go build -o go-vue-embedded ./build-workspace/frontend-assets.go ./build-workspace/main.go