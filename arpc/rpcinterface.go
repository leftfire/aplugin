package arpc

type Arpc interface {
	UseFunc(Req, Resp) error
}

type Req struct {
	FuncName string
	Args     map[string]string
}

type Resp struct {
	Code string
	Data string
}
