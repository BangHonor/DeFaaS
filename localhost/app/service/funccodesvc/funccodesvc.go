package funccodesvc

import (
	"defaas/localhost/app/model"
	"sync"

	"github.com/gogf/gf/frame/g"
)

var (
	service     *FunccodeSvc
	serviceOnce sync.Once
)

func Service() *FunccodeSvc {

	serviceOnce.Do(func() {

		service = NewFunccodeSvc()

	})

	return service
}

type FunccodeSvc struct {
	levels []model.FunccodeItem
}

func NewFunccodeSvc() *FunccodeSvc {

	svc := &FunccodeSvc{}
	svc.levels = []model.FunccodeItem{
		{
			Name :"first",
			Tag :"none",
			Files: struct{Filename string "json:\"filename\""; Language string "json:\"language\""; Code string "json:\"code\""}{
				Filename :"handler.py",
				Language :"python",
				Code :"123",
			},
		},
	}
	return svc
}

func (svc *FunccodeSvc) List() ([]model.FunccodeItem, error) {

	return svc.levels, nil
}

func (svc *FunccodeSvc) Add(item model.FunccodeItem) (model.FunccodeItem, error) {

	// add name of funccode

	svc.levels = append(svc.levels, item)

	g.Log().Printf("[funccodesvc] add funccode %v\n", item)

	return item, nil
}
