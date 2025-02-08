package rest

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"github.com/yaomianfa/go-okx/examples"
	rc "github.com/yaomianfa/go-okx/rest"
)

// 敏感信息申请的模拟盘KEY，不确定何时会删除
var TestClient = rc.New("", examples.Auth, &fasthttp.Client{
	Dial: fasthttpproxy.FasthttpSocksDialer("socks5://localhost:7890"),
})
