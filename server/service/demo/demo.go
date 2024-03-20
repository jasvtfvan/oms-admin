package demo

type DemoService interface {
	Hello() string
}

type DemoServiceImpl struct{}

func (*DemoServiceImpl) Hello() string {
	return "Hello"
}
