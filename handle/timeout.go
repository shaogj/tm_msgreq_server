package handle

import (
	"2021New_BFLProjTotal/tm_msgreq_server/proto"
	"context"
	"errors"
	"github.com/mkideal/log"
	//"runtime/debug"
	//"strings"
	"time"
)

var (
	// ErrCanceled is the error returned when the context is canceled.
	ErrCanceled = context.Canceled
	// ErrTimeout is the error returned when the context's deadline passes.
	ErrTimeout = context.DeadlineExceeded
)

// DoOption defines the method to customize a DoWithTimeout call.
type DoOption func() context.Context

// DoWithTimeout runs fn with timeout control.
func SendMsgWithTimeout(timeout time.Duration, nodeUrl, sendreqmsg string, opts ...DoOption)(height int64,geterr error) {
	parentCtx := context.Background()
	for _, opt := range opts {
		parentCtx = opt()
	}
	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	// create channel with buffer size 1 to avoid goroutine leak
	done := make(chan error,1)
	respdataChan := make(chan *proto.ResultBroadcastTxCommit, 1)
	var err error
	go func() {
		err = SendAsyncMsgToTMNode(nodeUrl, sendreqmsg, respdataChan)
		if err != nil {
			log.Error("cur exec SendAsyncMsgToTMNode() error! ,get err is:%v",  err)
			done <- err
		}
	}()

	select {
	case respdata := <-respdataChan:
		log.Info("cur check reqmsg broadcast_tx_sync_commit() good,get getresq hash is:%s,height is:%d", respdata.Hash, respdata.Height)
		return respdata.Height,nil
	case err := <-done:
		return 0,err
	case <-ctx.Done():
		log.Error("cur check reqmsg is timeout! in time:%d no resp to broadcast_tx_sync_commit(),req url is:%s", timeout,nodeUrl)
	}
	return 0,errors.New("cur reqmsg checked timeout!")
}

// WithContext customizes a DoWithTimeout call with given ctx.
func WithContext(ctx context.Context) DoOption {
	return func() context.Context {
		return ctx
	}
}
