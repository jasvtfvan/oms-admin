package system

type DemoService interface {
	Hello()
}

type DemoServiceImpl struct{}

func (*DemoServiceImpl) Hello() {}
