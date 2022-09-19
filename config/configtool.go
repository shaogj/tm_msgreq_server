package config

import (
	"encoding/json"
	"flag"
	"github.com/mkideal/log"
	"io/ioutil"
)

var GbConf ConfigInfomation = ConfigInfomation{}

// sgj 1019 add
type ConfigInfomation struct {
	//TM大账户地址.默认有钱
	TMFromAddress string `json:"TMFromAddress"`
	//TransAmount    int64  `json:"TransAmount"`

	//0907add
	LogLevel       string `json:"LogLevel"`
	SendTxInterval int    `json:"SendTxInterval"`
	TMNodeUrl1     string `json:"TMNodeUrl1"`
	TMNodeUrl2     string `json:"TMNodeUrl2"`
	TMNodeUrl3     string `json:"TMNodeUrl3"`
	WebPort        int    `json:"WebPort"`
}

func InitWithProviders(providers, dir string) error {
	return log.Init(providers, log.M{
		"rootdir":     dir,
		"suffix":      ".txt",
		"date_format": "%04d-%02d-%02d",
	})
}

func InitConfigInfo() error {
	//*good conf:
	//log.SetFlags(log.Lshortfile | log.Ltime)
	var strConf string
	flag.StringVar(&strConf, "conf", "config.json", "config <file>")
	flag.Parse()
	byData, err := ioutil.ReadFile(strConf)
	if nil != err {
		log.Error("Read config file :::%v", err)
		return err
	}
	err = json.Unmarshal(byData, &GbConf)
	if nil != err {
		log.Error("Unmarshal config file :::%v", err)
		return err
	}
	log.Info("ConfigInfo:::%+v", GbConf)
	return nil
}
