package business

import (
	"github.com/yaomianfa/go-okx/ws"
)

type Business struct {
	C *ws.Client
}

func NewBusiness(simulated bool) *Business {
	business := &Business{
		C: ws.DefaultClientBusiness,
	}
	if simulated {
		business.C = ws.DefaultClientBusinessSimulated
	}
	return business
}

// subscribe
func (p *Business) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}

// unsubscribe
func (p *Business) UnSubscribe(args interface{}) error {
	UnSubscribe := ws.NewOperateUnsubscribe(args)
	return p.C.Operate(UnSubscribe, nil)
}
