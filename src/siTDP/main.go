package main

import (
	"github.com/fiorix/wsdl2go/soap"
)

func main() {
	cli := soap.Client{
		URL: "http://sipo.kemendag.go.id/ws/server.php?wsdl",
	}
	soapService := example.NewEchoService(&cli)
	echoReply, err := soapService.Echo(&example.EchoRequest{Data: "hello world"})
	panic(echoReply)
}
