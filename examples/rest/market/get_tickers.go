package main

import (
	"log"

	"github.com/yaomianfa/go-okx/examples/rest"
	"github.com/yaomianfa/go-okx/rest/api"
	"github.com/yaomianfa/go-okx/rest/api/market"
)

func main() {
	param := &market.GetTickersParam{
		InstType: api.InstTypeSPOT,
	}
	req, resp := market.NewGetTickers(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*market.GetTickersResponse))
}
