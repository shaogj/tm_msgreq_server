package proto

import (
	"2021New_BFLProjTotal/tm_msgreq_server/utils/bytes"
	"encoding/json"
	"net/http"
)

const (
	StatusSuccess = 200 // 调用成功
)

type Response struct {
	HTTPCode     int           `json:"-"`
	Code         string        `json:"errno"`
	Msg          string        `json:"message"`
	Header       http.Header   `json:"-"`
	Data         interface{}   `json:"data,omitempty"`
	IsGZip       bool          `json:"-"`
	IsResetToken bool          `json:"-"`
	MsgData      []interface{} `json:"-"`
}

type ErrorInfo struct {
	Code int
	Desc string
}

var (
	//请求参数错误
	ErrorRequest = ErrorInfo{Code: 1001, Desc: " 请求参数无效"}

	ErrorRequestWDCNodeRPC = ErrorInfo{Code: 7000, Desc: "请求节点RPC参数无效(WDC)"}
	//1030add

	ErrorCoinType = ErrorInfo{Code: 1004, Desc: "数字货币类型错误"}

	ErrorNodeRPCSuccess = ErrorInfo{Code: 2000, Desc: "调用节点RPC成功"}
	ErrorRequestWDCNode = ErrorInfo{Code: 5000, Desc: "请求Node错误(WDC)"}

	ErrorRequestWDCNodeJust = ErrorInfo{Code: 7000, Desc: "请求Node校验错误(WDC)"}
	/*
			2000 正确
		    2100 已确认
		    2200 未确认
		    5000 错误
		    6000 格式错误
		    7000 校验错误
		    8000 异常
	*/

	ErrorAddress = ErrorInfo{Code: 701, Desc: "无效的地址码"}
	ErrorHttpost = ErrorInfo{Code: 801, Desc: "http请求必须为POST方式"}

	ErrorSuccess       = ErrorInfo{Code: 200, Desc: "调用成功"}
	ErrorRequestWDCSDK = ErrorInfo{Code: 500, Desc: "调用SDK参数无效(WDC)"}

	ErrorGetPrivateKey = ErrorInfo{Code: 604, Desc: "无法获取用户私钥"}

	ErrDecryptFail    = ErrorInfo{Code: 3333, Desc: "解密json串失败"}
	ErrVerifyListFail = ErrorInfo{Code: 9189, Desc: "人工审核获取订单列表失败"}

	//1129 add
	ErrorRequestInfuraETHNode = ErrorInfo{Code: 7004, Desc: "请求Node错误(ETH)"}
	ErrorRequestInfuraBSCNode = ErrorInfo{Code: 7004, Desc: "请求Balance(BSC)"}

	ErrorRequestInfuraETHSend = ErrorInfo{Code: 7005, Desc: "请求Node错误(ETH)"}
)

func Success(status ErrorInfo) bool {
	if ErrorSuccess.Code == status.Code {
		return true
	}
	return false

}

//向serverin 请求分组

type ReturnInfo struct {
	ResultCode int    `json:"code"`
	ResultMsg  string `json:"msg"`
}

//20220601add
type ResultBroadcastTxCommit struct {
	CheckTx   ResponseCheckTx   `json:"check_tx"`
	DeliverTx ResponseDeliverTx `json:"deliver_tx"`
	Hash      bytes.HexBytes    `json:"hash"`
	Height    int64             `json:"height"`
}
type ResponseDeliverTx struct {
	Code      uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted int64  `protobuf:"varint,5,opt,name=gas_wanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed   int64  `protobuf:"varint,6,opt,name=gas_used,proto3" json:"gas_used,omitempty"`
	//Events    []Event `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace string `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
}

type ResponseCheckTx struct {
	Code      uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log       string `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info      string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted int64  `protobuf:"varint,5,opt,name=gas_wanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed   int64  `protobuf:"varint,6,opt,name=gas_used,proto3" json:"gas_used,omitempty"`
	//Events    []Event `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty"`
	Codespace string `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	Sender    string `protobuf:"bytes,9,opt,name=sender,proto3" json:"sender,omitempty"`
	Priority  int64  `protobuf:"varint,10,opt,name=priority,proto3" json:"priority,omitempty"`
	// mempool_error is set by Tendermint.
	// ABCI applictions creating a ResponseCheckTX should not set mempool_error.
	MempoolError string `protobuf:"bytes,11,opt,name=mempool_error,json=mempoolError,proto3" json:"mempool_error,omitempty"`
}

type RPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	//ID      string `json:"id"`
	CODE int `json:"code"`
	//Result	 []byte `json:"result,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}
