package main

import (
	"log"

	"github.com/yaomianfa/go-okx/examples/rest"
	"github.com/yaomianfa/go-okx/rest/api"
	"github.com/yaomianfa/go-okx/rest/api/account"
)

func main() {
	param := &account.GetLeverageInfoParam{
		InstId:  "BTC-USDT",
		MgnMode: api.MgnModeCross,
	}
	req, resp := account.NewGetLeverageInfo(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetLeverageInfoResponse))
}
