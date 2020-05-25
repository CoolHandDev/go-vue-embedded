all: build-vue-app build-go

.PHONY: build-vue-app
build-vue-app: 
	npm run --prefix front-end build	

.PHONY: build-go
build-go:
	pkger
	go build -o go-vue-embedded .