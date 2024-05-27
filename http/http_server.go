package main

import (
	"fmt"
	"net/http"
	"os"
)

const STATIC_PRE = "/static/"

const PUB_DIR = "./static"

func main() {
	//新建一个请求多路复用器
	serverMux := http.NewServeMux()
	// 注册静态地址的多路复用器,用来访问静态文件
	serverMux.HandleFunc(STATIC_PRE, staticHandler)
	// 注册动态接口部分的路由
	serverMux.HandleFunc("/hello", helloHandle)

	err := http.ListenAndServe(":8080", serverMux)
	if err != nil {
		panic(err)
	}

}

// 用来处理静态的文件部分
func staticHandler(w http.ResponseWriter, r *http.Request) {
	file := PUB_DIR + r.URL.Path[len(STATIC_PRE)-1:]
	if ok := isFileExist(file); !ok {
		w.Write([]byte("下载的文件不存在!"))
		return
	}
	http.ServeFile(w, r, file)
}

// 处理 /hello接口
func helloHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func processFormData(w http.ResponseWriter, r *http.Request) {

}

func isFileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("要下载的文件不存在%s", filename)
		return false
	}
	return true
}
