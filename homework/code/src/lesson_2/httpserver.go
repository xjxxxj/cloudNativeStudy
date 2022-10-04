package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
当访问 localhost/healthz 时，应返回 200
*/

// HealthService 处理healthz请求,直接返回200
func HealthService(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	io.WriteString(response, "200")
}

/**
处理其他请求：
1. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
2. 将 request 中带的 header 写入 response header
3. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
*/

func CommonService(response http.ResponseWriter, request *http.Request) {

	requestHeader := request.Header

	for key, value := range requestHeader {
		response.Header().Set(key, value[0])
	}

	response.Header().Set("VERSION", os.Getenv("VERSION"))

	fmt.Printf("client host:%s, request url:%s, status :%s", request.Host, request.RequestURI, "200")

	io.WriteString(response, "you request url is not found!")

}

func main() {
	http.HandleFunc("/", CommonService)
	http.HandleFunc("/healthz", HealthService)

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
