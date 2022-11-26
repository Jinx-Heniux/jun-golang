package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// 抽取单个server对象
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	Name    xml.Name `xml:"servers"`
	Version int      `xml:"version"`
	Servers []Server `xml:"server"`
}

func Xml1Example1() {

	data, err := ioutil.ReadFile("./libs/encoding/xml/xml1-my.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var servers Servers
	err = xml.Unmarshal(data, &servers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("xml: %#v\n", servers)
}

/*
xml: xml.Servers{Name:xml.Name{Space:"", Local:""}, Version:0, Servers:[]xml.Server{xml.Server{ServerName:"Shanghai_VPN", ServerIP:"127.0.0.1"}, xml.Server{ServerName:"Beijing_VPN", ServerIP:"127.0.0.2"}}}
*/
