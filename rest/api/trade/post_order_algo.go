package trade

import "github.com/yaomianfa/go-okx/rest/api"

func NewPostOrderAlgo(param *PostOrderAlgoParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/order-algo",
		Method: api.MethodPost,
		Param:  param,
	}, &PostOrderAlgoResponse{}
}

type PostOrderAlgoParam struct {
	InstId          string `json:"instId"`
	TdMode          string `json:"tdMode"`
	Side            string `json:"side"`
	OrdType         string `json:"ordType"`
	Sz              string `json:"sz"`
	TgtCcy          string `json:"tgtCcy"`
	AlgoClOrdId     string `json:"algoClOrdId"`
	Tag             string `json:"tag"`
	TpTriggerPx     string `json:"tpTriggerPx"`
	TpTriggerPxType string `json:"tpTriggerPxType"`
	TpOrdPx         string `json:"tpOrdPx"`
	SlTriggerPx     string `json:"slTriggerPx"`
	SlTriggerPxType string `json:"slTriggerPxType"`
	SlOrdPx         string `json:"slOrdPx"`
	CallbackRatio   string `json:"callbackRatio"`
	CallbackSpread  string `json:"callbackSpread"`
	ActivePx        string `json:"activePx"`
}

type PostOrderAlgoResponse struct {
	api.Response
	Data []PostOrderAlgo `json:"data"`
}

type PostOrderAlgo struct {
	AlgoId      string `json:"algoId"`
	ClOrdId     string `json:"clOrdId"`
	AlgoClOrdId string `json:"algoClOrdId"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
	Tag         string `json:"tag"`
}
