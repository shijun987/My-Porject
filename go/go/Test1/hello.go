package main

import (
	"encoding/json"
	"fmt"
    "net/http"
    "github.com/tarm/goserial"
	"strings"
)

func main() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", loginTask)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("192.168.1.192:8083", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func loginTask(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    // 用于存放参数key=value数据
    var params map[string]string
    // 解析参数 存入map
    decoder.Decode(&params)
    fmt.Fprintf(w, `{"code":0 ,`)

    bytes, _ := json.Marshal(params)
	fmt.Fprint(w, string(bytes))
    fmt.Fprintf( w,`}`)
}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}