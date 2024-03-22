package initialize

import "errors"

type UpdateDBService interface {
	UpdateDB() error
	ClearInitializer()
}

type UpdateDBServiceImpl struct{}

// 已经升级，重启服务后，清除 initializers
func (s *UpdateDBServiceImpl) ClearInitializer() {
	// initializers = initSlice{}
	// cache = map[string]*orderedInitializer{}
}

// 升级
// 升级的前提是，部署了代码，部署代码一定会重新启动server，需要
func (s *UpdateDBServiceImpl) UpdateDB() (err error) {
	// ctx := context.Background()
	if len(initializers) == 0 {
		return errors.New("升级任务列表为空，请检查是否已执行完成")
	}

	return err
}
