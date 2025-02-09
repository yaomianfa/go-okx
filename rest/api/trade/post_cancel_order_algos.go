package trade

import "github.com/yaomianfa/go-okx/rest/api"

func NewPostCancelOrderAlgos(param *PostCancelOrderAlgosParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/cancel-algos",
		Method: api.MethodPost,
		Param:  param,
	}, &PostCancelOrderAlgosResponse{}
}

type PostCancelOrderAlgosParam []CancelOrderAlgosParam

type CancelOrderAlgosParam struct {
	AlgoId string `json:"algoId"`
	InstId string `json:"instId"`
}

type PostCancelOrderAlgosResponse struct {
	api.Response
	Data []CancelOrderAlgo `json:"data"`
}

type CancelOrderAlgo struct {
	AlgoClOrdId string `json:"algoClOrdId"`
	AlgoId      string `json:"algoId"`
	ClOrdId     string `json:"clOrdId"`
	SCode       string `json:"sCode"`
	SMsg        string `json:"sMsg"`
	Tag         string `json:"tag"`
}
