package main

import (
	"log"

	"github.com/yaomianfa/go-okx/examples/rest"
	"github.com/yaomianfa/go-okx/rest/api/asset"
)

func main() {
	param := &asset.GetTransferStateParam{
		TransId: "123456",
	}
	req, resp := asset.NewGetTransferState(param)
	if err := rest.TestClient.Do(req, resp); err != nil {
		panic(err)
	}
	log.Println(req, resp.(*asset.GetTransferStateResponse))
}
