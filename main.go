package main

import (
	"2021New_BFLProjTotal/tm_msgreq_server/config"
	"2021New_BFLProjTotal/tm_msgreq_server/curl_req"
	"fmt"
	"github.com/mkideal/log"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
)

func getIPV6Lan() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	//fmt.Println("InterfaceAddrs() ,---checking==addrs is:%s", addrs)
	for _, addr := range addrs {
		ipv6 := regexp.MustCompile(`(\w+:){7}\w+`).FindString(addr.String())
		if strings.Count(ipv6, ":") == 7 {
			return ipv6
		}
	}
	return ""
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return ip, err

}

func GetExternIP2() {
	responseClient, errClient := http.Get("http://ip.dhcp.cn/?ip") // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		panic(errClient)
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	clientIP := fmt.Sprintf("%s", string(body))
	print(clientIP)
}
func main() {
	getlocalIp := getIPV6Lan()
	fmt.Println("getlocalIp is:%s", getlocalIp)
	getlocalIp2, err := GetOutBoundIP()
	if err != nil {
		fmt.Println("getlocalIp2 error!,err is:%v", err)
	}
	fmt.Println("getlocalIp2 is:%s", getlocalIp2)
	//
	//curl_req.ReqTMCommitTx()
	if err := config.InitWithProviders("multifile/console", "./logs"); err != nil {
		panic("init log error: " + err.Error())
	}
	log.Info("log level: %v", log.SetLevelFromString("trace"))

	ch := make(chan int)
	reqProc := curl_req.VoteServer{RequestInterval: 10}
	go reqProc.ResetGroupVotesMap(reqProc.RequestInterval)
	getReq := <-ch
	fmt.Println(getReq)
	//curl_req.ReqRMcommitTxInterval()

}
