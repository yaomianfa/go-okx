package private

import (
	"encoding/json"

	"github.com/yaomianfa/go-okx/common"
	"github.com/yaomianfa/go-okx/ws"
)

type HandlerOrders func(EventOrders)

type EventOrders struct {
	Arg  ws.Args `json:"arg"`
	Data []Order `json:"data"`
}

type Order struct {
	AccFillSz         string        `json:"accFillSz"`
	AmendResult       string        `json:"amendResult"`
	AvgPx             string        `json:"avgPx"`
	CTime             string        `json:"cTime"`
	Category          string        `json:"category"`
	Ccy               string        `json:"ccy"`
	ClOrdId           string        `json:"clOrdId"`
	Code              string        `json:"code"`
	ExecType          string        `json:"execType"`
	Fee               string        `json:"fee"`
	FeeCcy            string        `json:"feeCcy"`
	FillFee           string        `json:"fillFee"`
	FillFeeCcy        string        `json:"fillFeeCcy"`
	FillNotionalUsd   string        `json:"fillNotionalUsd"`
	FillPx            string        `json:"fillPx"`
	FillSz            string        `json:"fillSz"`
	FillPnl           string        `json:"fillPnl"`
	FillTime          string        `json:"fillTime"`
	FillPxVol         string        `json:"fillPxVol"`
	FillPxUsd         string        `json:"fillPxUsd"`
	FillMarkVol       string        `json:"fillMarkVol"`
	FillFwdPx         string        `json:"fillFwdPx"`
	FillMarkPx        string        `json:"fillMarkPx"`
	InstId            string        `json:"instId"`
	InstType          string        `json:"instType"`
	Lever             string        `json:"lever"`
	Msg               string        `json:"msg"`
	NotionalUsd       string        `json:"notionalUsd"`
	OrdId             string        `json:"ordId"`
	OrdType           string        `json:"ordType"`
	Pnl               string        `json:"pnl"`
	PosSide           string        `json:"posSide"`
	Px                string        `json:"px"`
	PxUsd             string        `json:"pxUsd"`
	PxVol             string        `json:"pxVol"`
	PxType            string        `json:"pxType"`
	Rebate            string        `json:"rebate"`
	RebateCcy         string        `json:"rebateCcy"`
	ReduceOnly        string        `json:"reduceOnly"`
	ReqId             string        `json:"reqId"`
	Side              string        `json:"side"`
	AttachAlgoClOrdId string        `json:"attachAlgoClOrdId"`
	SlOrdPx           string        `json:"slOrdPx"`
	SlTriggerPx       string        `json:"slTriggerPx"`
	SlTriggerPxType   string        `json:"slTriggerPxType"`
	Source            string        `json:"source"`
	State             string        `json:"state"`
	StpId             string        `json:"stpId"`
	StpMode           string        `json:"stpMode"`
	Sz                string        `json:"sz"`
	Tag               string        `json:"tag"`
	TdMode            string        `json:"tdMode"`
	TgtCcy            string        `json:"tgtCcy"`
	TpOrdPx           string        `json:"tpOrdPx"`
	TpTriggerPx       string        `json:"tpTriggerPx"`
	TpTriggerPxType   string        `json:"tpTriggerPxType"`
	TradeId           string        `json:"tradeId"`
	LastPx            string        `json:"lastPx"`
	QuickMgnType      string        `json:"quickMgnType"`
	AlgoClOrdId       string        `json:"algoClOrdId"`
	AttachAlgoOrds    []interface{} `json:"attachAlgoOrds"`
	AlgoId            string        `json:"algoId"`
	AmendSource       string        `json:"amendSource"`
	CancelSource      string        `json:"cancelSource"`
	IsTpLimit         string        `json:"isTpLimit"`
	UTime             string        `json:"uTime"`
	LinkedAlgoOrd     struct {
		AlgoId string `json:"algoId"`
	} `json:"linkedAlgoOrd"`
}

// default subscribe
func SubscribeOrders(args *ws.Args, auth common.Auth, handler HandlerOrders, handlerError ws.HandlerError) error {
	args.Channel = "orders"

	h := func(message []byte) {
		var event EventOrders
		if err := json.Unmarshal(message, &event); err != nil {
			handlerError(err)
			return
		}
		handler(event)
	}

	return NewPrivate(auth).Subscribe(args, h, handlerError)
}
