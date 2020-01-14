package main

import (
	"flag"
	"fmt"
	"net/http"
	"slowlog_elk_http/objects"
)

var (
	host, scritps string
	port int

)

func init()  {
	flag.StringVar(&host, "h", "0.0.0.0", "host")
	flag.StringVar(&scritps, "s", "", "php脚本位置")
	flag.IntVar(&port, "P", 6000, "port")
}

func main()  {
	flag.Parse()
	if scritps == "" {
		panic("scripts不能为空")
	}
	objects.Scripts = scritps
	// 启动一个http作为接入层，对外提供服务
	httpEndPoint := fmt.Sprintf("%s:%d", host, port)
	http.HandleFunc("/sql/", objects.Handler)
	http.ListenAndServe(httpEndPoint, nil)
}
