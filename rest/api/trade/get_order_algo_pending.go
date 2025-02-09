package trade

import (
	"github.com/yaomianfa/go-okx/rest/api"
)

func NewGetOrderAlgoPending(param *GetOrderAlgoPendingParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-algo-pending",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderAlgoPendingResponse{}
}

type GetOrderAlgoPendingParam struct {
	AlgoId      string `url:"algoId,omitempty"`
	InstType    string `url:"instType,omitempty"`
	InstId      string `url:"instId,omitempty"`
	OrdType     string `url:"ordType"`
	After       string `url:"after,omitempty"`
	Before      string `url:"before,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	AlgoClOrdId string `url:"algoClOrdId,omitempty"`
}

type GetOrderAlgoPendingResponse struct {
	api.Response
	Data []OrderAlgo `json:"data"`
}

type OrderAlgoPending struct {
	ActivePx             string        `json:"activePx"`
	ActualPx             string        `json:"actualPx"`
	ActualSide           string        `json:"actualSide"`
	ActualSz             string        `json:"actualSz"`
	AlgoClOrdId          string        `json:"algoClOrdId"`
	AlgoId               string        `json:"algoId"`
	AmendPxOnTriggerType string        `json:"amendPxOnTriggerType"`
	AttachAlgoOrds       []interface{} `json:"attachAlgoOrds"`
	CTime                string        `json:"cTime"`
	CallbackRatio        string        `json:"callbackRatio"`
	CallbackSpread       string        `json:"callbackSpread"`
	Ccy                  string        `json:"ccy"`
	ChaseType            string        `json:"chaseType"`
	ChaseVal             string        `json:"chaseVal"`
	ClOrdId              string        `json:"clOrdId"`
	CloseFraction        string        `json:"closeFraction"`
	FailCode             string        `json:"failCode"`
	InstId               string        `json:"instId"`
	InstType             string        `json:"instType"`
	IsTradeBorrowMode    string        `json:"isTradeBorrowMode"`
	Last                 string        `json:"last"`
	Lever                string        `json:"lever"`
	LinkedOrd            struct {
		OrdId string `json:"ordId"`
	} `json:"linkedOrd"`
	MaxChaseType    string        `json:"maxChaseType"`
	MaxChaseVal     string        `json:"maxChaseVal"`
	MoveTriggerPx   string        `json:"moveTriggerPx"`
	OrdId           string        `json:"ordId"`
	OrdIdList       []interface{} `json:"ordIdList"`
	OrdPx           string        `json:"ordPx"`
	OrdType         string        `json:"ordType"`
	PosSide         string        `json:"posSide"`
	PxLimit         string        `json:"pxLimit"`
	PxSpread        string        `json:"pxSpread"`
	PxVar           string        `json:"pxVar"`
	QuickMgnType    string        `json:"quickMgnType"`
	ReduceOnly      string        `json:"reduceOnly"`
	Side            string        `json:"side"`
	SlOrdPx         string        `json:"slOrdPx"`
	SlTriggerPx     string        `json:"slTriggerPx"`
	SlTriggerPxType string        `json:"slTriggerPxType"`
	State           string        `json:"state"`
	Sz              string        `json:"sz"`
	SzLimit         string        `json:"szLimit"`
	Tag             string        `json:"tag"`
	TdMode          string        `json:"tdMode"`
	TgtCcy          string        `json:"tgtCcy"`
	TimeInterval    string        `json:"timeInterval"`
	TpOrdPx         string        `json:"tpOrdPx"`
	TpTriggerPx     string        `json:"tpTriggerPx"`
	TpTriggerPxType string        `json:"tpTriggerPxType"`
	TriggerPx       string        `json:"triggerPx"`
	TriggerPxType   string        `json:"triggerPxType"`
	TriggerTime     string        `json:"triggerTime"`
	UTime           string        `json:"uTime"`
}
