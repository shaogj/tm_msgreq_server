package curl_req

import (
	"2021New_BFLProjTotal/tm_msgreq_server/handle"
	"fmt"
	"github.com/mkideal/log"
	"io/ioutil"
	"net/http"
	"time"
)

type VoteServer struct {
	RequestInterval int
	CurSerionNum    int
}

func UrlSwitchLocal(n int) string { //-----
	var url string
	switch n {
	case 0:
		url = "192.168.1.221"
	case 1:
		url = "192.168.1.222"
	case 2:
		url = "192.168.1.223"
	}
	return url
}

func UrlSwitchBJ(n int) string { //-----
	var url string
	switch n {
	case 0:
		url = "106.3.133.179"
	case 1:
		url = "106.3.133.180"
	case 2:
		url = "210.73.218.172"
	}
	return url
}

//清理历史数据
func (this *VoteServer) ResetGroupVotesMap(interval int) {
	log.Info("启动投票服务，interval is:%d", interval)
	//20 *RequestInterval = 20个分组周期清理历史数据
	ticker := time.NewTicker(time.Second * time.Duration(this.RequestInterval))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			//curmin := time.Now().Minute()
			curNodeIp := UrlSwitchLocal(this.CurSerionNum % 3)
			//ipsevre :
			log.Info("run ticker cur UrlSwitch IP  is:%s，", curNodeIp)
			curheight, err := ReqRMcommitTxInterval(curNodeIp)
			if err != nil {
				log.Error("run ticker task ReqRMcommitTxInterval()，curheight is:%d，this.RequestInterval is:%d", curheight, this.RequestInterval)
			}
			log.Info("run ticker task ReqRMcommitTxInterval()，curheight is:%d，err is:%v", curheight, err)
			this.CurSerionNum++
			//to do,,获取上一个周期时间的投票分组信息,to single队列任务
			//case <-this.qtChan:
			//	ticker.Stop()
			//return
		}
	}
}

func ReqRMcommitTxInterval(curip string) (curheight int64, err error) {
	//nodeUrl := "http://101.251.211.201:21630/tri_broadcast_tx_commit?"
	//nodeUrl := "http://192.168.1.222:46657/tri_broadcast_tx_commit?"
	nodeUrl := fmt.Sprintf("http://%s:46657/tri_broadcast_tx_commit?", curip)
	getrandnumstr := handle.RandStr(5)
	fmt.Println("get randnumstr7777 is:%s", getrandnumstr)

	//sendmsgNew := fmt.Sprintf("%s:%s%22","ssA1710BBsssCCaaCC11",getrandnumstr)
	sendmsgNew := "tx=%22" + "ssA1710BBsssCCaadCCs11" + getrandnumstr + "%22"
	getblockInfo, getresq, err := handle.SendRMSVoteMsgToNode(nodeUrl, sendmsgNew)
	if err != nil {
		log.Error("cur after broadcast_tx_commit(),get error! ,getRespInfo is :%v,err is:%v", getblockInfo, err)
		return 0, err
	}
	fmt.Println("cur broadcast_tx_commit(),get getresq hash is:%s,height is:%d", getresq.Hash, getresq.Height)
	return getresq.Height, nil
}
func ReqTMCommitTx() (succ int, err error) {

	//var jsonRet jsonResult
	//stdstr := "http://192.168.1.221:46657/tri_broadcast_tx_commit?tx=%22aapppppppppss888889333888%22"
	//stdstr := "http://192.168.1.223:46657/tri_broadcast_tx_commit?tx='%s'"
	txinfo := "curtxinfo +randnum"
	url := fmt.Sprintf("http://192.168.1.223:46657/tri_broadcast_tx_commit?tx='%s'", txinfo)

	ret, err := get(url)
	if err != nil {
		return 0, err
	}
	fmt.Printf("return is :%v", ret)
	/*
		err = json.Unmarshal(ret, &jsonRet)
		if err != nil {
			return 0, err
		}
	*/
	return 1, nil
}

func get(url string) ([]byte, error) {
	// url := "http://106.3.133.179:46657/tri_block_info?height=104360"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
