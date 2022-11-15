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
	//0919add,for tmurl
	NodeUrlList []string
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

// 清理历史数据
func (this *VoteServer) ResetGroupVotesMap(interval int) {
	log.Info("TM reqmsg server start!interval is:%d", interval)
	ticker := time.NewTicker(time.Second * time.Duration(this.RequestInterval))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:

			curNodeIp := this.NodeUrlList[this.CurSerionNum%3]
			//ipsevre :
			log.Info("run ticker cur auto UrlSwitch===  is:%s，", curNodeIp)
			curheight, err := ReqRMcommitTxInterval(curNodeIp)
			if err != nil {
				log.Error("run ticker task ReqRMcommitTxInterval()，curheight is:%d，this.RequestInterval is:%d", curheight, this.RequestInterval)
			}
			log.Info("run ticker task ReqRMcommitTxInterval()，curheight is:%d，err is:%v", curheight, err)
			this.CurSerionNum++

		}
	}
}

func ReqRMcommitTxInterval(curip string) (curheight int64, err error) {
	//tri_broadcast_tx_commit
	//nodeUrl := fmt.Sprintf("http://%s:46657/tri_bc_tx_async?", curip)
	nodeUrl := fmt.Sprintf("http://%s:46657/tri_broadcast_tx_commit?", curip)
	//end update
	getrandnumstr := handle.RandStr(5)
	fmt.Println("get randnumstr7777 is:%s,nodeUrl is:%s", getrandnumstr, nodeUrl)

	//sendmsgNew := fmt.Sprintf("%s:%s%22","ssA1710BBsssCCaaCC11",getrandnumstr)
	sendmsgNew := "tx=%22" + "ssA1710BBsssCCaadCCs11" + getrandnumstr + "%22"
	//getblockInfo, getresq, err := handle.SendRMSVoteMsgToNode(nodeUrl, sendmsgNew)
	//11.10 set req to Timeout
	timeoout := time.Second * 30
	getblockheight, err := handle.SendMsgWithTimeout(timeoout,nodeUrl, sendmsgNew)

	if err != nil {
		log.Error("cur after broadcast_tx_commit(),get error! ,getRespInfo is :%v,err is:%v", "getblockInfo", err)
		return 0, err
	}
	fmt.Println("cur broadcast_tx_commit(),get getresq hash is:%s,height is:%d", "getresq.Hash", getblockheight)
	return getblockheight, nil
}
func ReqSTDTMcommitValidatorTx(curip string, newnodeip, pubkey string) (curheight int64, err error) {
	if pubkey == "" {
		newnodeip = "101.251.223.189"
		pubkey = "EA52F8DA1710C4ABF51B725C79F69D90E2F3CE38AA4194B0979AF762DA6D4259"
	}
	nodeUrl := fmt.Sprintf("http://%s:46657/tri_broadcast_tx_commit?", curip)
	commitTx := fmt.Sprint("[addvalidator]val:%s/%s/%d", newnodeip, pubkey, 123400)
	sendmsgNew := "tx=%22" + commitTx + "%22"
	getblockInfo, getresq, err := handle.SendRMSVoteMsgToNode(nodeUrl, sendmsgNew)
	if err != nil {
		log.Error("cur after broadcast_tx_commit(),get error! ,getRespInfo is :%v,err is:%v", getblockInfo, err)
		return 0, err
	}
	fmt.Println("cur broadcast to validator_tx_commit(),get getresq hash is:%s,height is:%d", getresq.Hash, getresq.Height)
	return getresq.Height, nil

	//http://210.73.218.172:46657/tri_broadcast_tx_commit?tx=%22[addvalidator]val:101.251.223.189/EA52F8DA1710C4ABF51B725C79F69D90E2F3CE38AA4194B0979AF762DA6D4259/304321%22
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
