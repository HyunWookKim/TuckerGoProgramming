package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var files embed.FS //static 폴더 내 모든 파일 추가

func main() {
	http.Handle("/", http.FileServer(http.FS(files)))
	http.ListenAndServe(":3000", nil)
}
