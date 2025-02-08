package main

import (
	"log"

	"github.com/yaomianfa/go-okx/examples/rest"
	"github.com/yaomianfa/go-okx/rest/api/account"
	"github.com/yaomianfa/go-okx/rest/api/subaccount"
)

func main() {
	param := &subaccount.GetBalancesParam{
		SubAcct: "test-01",
	}
	req, resp := subaccount.NewGetBalances(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetBalanceResponse))
}
