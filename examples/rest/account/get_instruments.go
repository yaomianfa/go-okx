package main

import (
	"log"

	"github.com/yaomianfa/go-okx/examples/rest"
	"github.com/yaomianfa/go-okx/rest/api/account"
)

func main() {
	param := &account.GetInstrumentsParam{InstType: "SPOT"}
	req, resp := account.NewGetInstruments(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*account.GetInstrumentsResponse))
}
