package main

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	//b := httplib.Post("http://oss.cartenz.co.id/api/third-party/complaint/track")
	//b := httplib.Get("http://localhost:8080/oss/ppm")
	b := httplib.Get("http://oss.ekon.go.id:8080/oss/ppm")
	b.Header("Accept-Encoding", "gzip,deflate,sdch")
	b.Header("Host", "beego.me")
	b.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	str, err := b.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
