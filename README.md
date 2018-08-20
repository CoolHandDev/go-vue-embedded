# Description

Example of embedding a Vue application inside a Go binary. Here we explore how a project might be structured and built. The Vue application is bootstrapped using Vue CLI (v3).

## Vue

Development is done in the /front-end directory

## Go

Development is done in the root directory

# Build Process

**Update**: a Makefile has been added so just run `make` to build the project and produce a binary.

The following steps are necessary to make the Vue build work as an embedded resource inside the Go binary. We have to make it so that the paths that will be in the frontend-assets.go are located off the same directory as the Go source. For example

```
"/index.html": {
		local:   "index.html",
		size:    772,
		modtime: 1534702974,
		compressed: `
H4sIAAAAAAAC/4ySz24TMRDGX2XIpRd2nZQ2bSTbEio9wIVKgIDjxJ6N3Ti2sWc3ytuj7ob0jwD1YmnG
329mPJ/lmw+fb77+vLsFx7ug5cMJAeNGUdTSEVotd8QIxmGpxKrnrrk+5hxzbuhX7wf1o/n2vrlJu4zs
14HApMgUWc0+3iqyG5odkYg7UoOnfU6FH1V7b9kpS4M31IzBWx89ewxNNRhILWZaBh+3UCgob1IEV6hT
osPhIWq9SVqy50C6KylyQ9FKMSUmcNLfV4Hr1HN7sVqdz1dX8/a+jjVzoY7YuGdqU6vAnNvl4vxisbjC
1tSTOiS0gFVVPvylR84tXa4MXS/Nkw4nxhSf+SVkXB+3zUDRplJbnHfzy3dL+zr8n5OO41VHxFqKyc91
sgctY/pTpnJJcaO/01khqKmUA6x7htMewSaq8Yxhn8oWckmZSjjA3rNLPcMnHPDLWAoo4jqQbeEuEFY6
xuAZOI1m+9hTK8WxoxSPQ1g/gLcKc9ZSWD9oOd1ALeb/29FSnF7yHHnhwhOhmHYgxk//OwAA///aDgFu
BAMAAA==
`,
```

`index.html` is in the same directory as main.go. Assets are in `/css`, `/img`, and `/js` respectively.

## Build Workspace

Create a `build-workspace` directory.

## Vue

Vue is built using `npm run build`. The resulting `/dist` directory is copied to the `/build-workspace` directory.

```
/build-workspace
    /dist
```

## Go

[esc](https://github.com/mjibson/esc) is used to embed the resources into the Go binary. Go get it before starting.

```
go get github.com/mjibson/esc
```

-   copy the \*.go files and relevant go subdirectories to the `/build-workspace/dist` created from the Vue build process. The objective is to make sure the main entrypoint to the Go program is in the same directory as the entry point (index.html) of the Vue application.
-   run `go generate` to generate the `frontend-assets.go` source that will contain the Vue application assets that will be embedded into the binary. Specifics of the embedding and code generation are in the main.go file. For example, a `//go:generate esc -o frontend-assets.go -pkg main index.html favicon.ico js css img` appears on the top of main.go.
-   run `go build -o go-vue-embedded frontend-assets.go main.go` to create the binary

# How to run the application

Run the `go-vue-embedded` binary. Open up the browser to http://127.0.0.1:4022/#/. The `#` is due to Vue router's default [mode](https://router.vuejs.org/guide/essentials/history-mode.html).

# Conclusion

This experiment suggests that maybe interleaving Go source into the Vue project simplifies the build process. We will not have to create a build workspace and copy files to it to perform the build. An outsider will find the project hard to understand.
