package ws

type Handler func(message []byte)
type HandlerError func(err error)

type Operate struct {
	Request      *Request
	Response     *Response
	Handler      func(message []byte)
	HandlerError HandlerError
}

// new subscribe
func NewOperateSubscribe(args interface{}, handler Handler, handlerError HandlerError) *Operate {
	return &Operate{
		Request:      NewRequestSubscribe(args),
		Handler:      handler,
		HandlerError: handlerError,
	}
}

// unsubscribe
func NewOperateUnsubscribe(args interface{}) *Operate {
	return &Operate{
		Request: NewRequestUnSubscribe(args),
	}
}

// new login
func NewOperateLogin(args *ArgsLogin) *Operate {
	return &Operate{
		Request: NewRequestLogin(args),
	}
}
