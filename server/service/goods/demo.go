package goods

type DemoService interface {
	Hello()
}

type DemoServiceImpl struct{}

func (*DemoServiceImpl) Hello() {}
