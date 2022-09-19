package main

import (
	"2021New_BFLProjTotal/tm_msgreq_server/config"
	"2021New_BFLProjTotal/tm_msgreq_server/curl_req"
	"fmt"
	"github.com/mkideal/log"
	"io/ioutil"
	"net"
	"net/http"
	"os"
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

func InitNodeList(curcfg *config.ConfigInfomation) []string {
	curNodeUrlList := make([]string, 0)
	if curcfg.TMNodeUrl1 == "" {
		curcfg.TMNodeUrl1 = "106.3.133.179"
	}
	if curcfg.TMNodeUrl1 == "" {
		curcfg.TMNodeUrl1 = "106.3.133.179"
	}
	if curcfg.TMNodeUrl1 == "" {
		curcfg.TMNodeUrl1 = "106.3.133.179"
	}
	curNodeUrlList = append(curNodeUrlList, curcfg.TMNodeUrl1)
	curNodeUrlList = append(curNodeUrlList, curcfg.TMNodeUrl2)
	curNodeUrlList = append(curNodeUrlList, curcfg.TMNodeUrl3)
	return curNodeUrlList

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
	log.Info("log level: %v\r\n", log.SetLevelFromString("trace"))
	//0917add
	err = config.InitConfigInfo()
	if nil != err {
		log.Error("from config.json,get json conf err!")
		os.Exit(0)
	}
	gbConf := &config.GbConf

	strHost := fmt.Sprintf(":%d", gbConf.WebPort)
	fmt.Printf("strHost is :%s,loglevel is:%s", strHost, gbConf.LogLevel)
	curNodeUrlList := InitNodeList(gbConf)
	fmt.Printf("checking==strHost is :%v\n", curNodeUrlList)

	//0917
	ch := make(chan int)
	curRequestInterval := gbConf.SendTxInterval
	if curRequestInterval == 0 {
		curRequestInterval = 10
	}
	reqProc := curl_req.VoteServer{RequestInterval: curRequestInterval, NodeUrlList: curNodeUrlList} //10
	go reqProc.ResetGroupVotesMap(reqProc.RequestInterval)
	//curl_req.ReqRMcommitTxInterval()
	//curl_req.ReqSTDTMcommitValidatorTx("106.3.133.179", "", "")
	getReq := <-ch
	fmt.Println(getReq)

}
