# Description

Example of embedding a Vue application inside a Go binary. Here we explore how a project might be structured and built. The Vue application is bootstrapped using Vue CLI (v3).

## Vue

Development is done in the /front-end directory

## Go

Development is done in the root directory

# Build Process

Run `make` to start the build process

## Vue

Vue was bootstrapped using the the Vue cli. The makefile target, build-vue-app, runs `npm run build` and the resulting output is placed in the `front-end/dist` directory

## Go

[Packr](https://github.com/gobuffalo/packr) is used to embedd the front-end assets to the Go binary. In `main.go` the line `box := packr.NewBox("./front-end/dist")` refers to the distribution directory of the Vue front-end. The build-go makefile target will generate the Go source code containing the front-end assets and then run the Go build to generate the static binary executable.

# How to run the application

Run the `go-vue-embedded` binary. Open up the browser to http://127.0.0.1:4022/#/. The `#` is due to Vue router's default [mode](https://router.vuejs.org/guide/essentials/history-mode.html).

# Conclusion

Embedding a Vue application in Go is possible and straightforward thanks to [Pkger](https://github.com/markbates/pkger).
