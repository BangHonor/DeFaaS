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
	funccodes []model.FunccodeItem
}

func NewFunccodeSvc() *FunccodeSvc {

	svc := &FunccodeSvc{}
	svc.funccodes = []model.FunccodeItem{
		{
			Name: "first",
			Tag:  "666",
			Files:[]struct{
				Filename string `json:"filename"`
				Language string `json:"language"`
				Code string `json:"code"`
			  }{
				  {
					Filename: "handler.py",
					Language: "python",
					Code:     "import { NzCodeEditorModule } from 'ng-zorro-antd/code-editor",
					},//数组
					{
						Filename: "handler.ts",
						Language: "typescript",
						Code:     "import { NzCodeEditorModule } from 'ng-zorro-antd/code-editor",
					},
			},
		},

	}
	return svc
}

func (svc *FunccodeSvc) List() ([]model.FunccodeItem, error) {

	return svc.funccodes, nil
}

func (svc *FunccodeSvc) Add(item model.FunccodeItem) (model.FunccodeItem, error) {

	// add name of funccode

	svc.funccodes = append(svc.funccodes, item)

	g.Log().Printf("[funccodesvc] add funccode %v\n", item)

	return item, nil
}

func (svc *FunccodeSvc) Delete(item model.FunccodeItem)  (model.FunccodeItem, error) {

	// add name of funccode
	for index, value := range svc.funccodes {
		if(value.Name==item.Name){
			svc.funccodes = append(svc.funccodes[:index], svc.funccodes[index+1:]...)
			break;
		}
	}
	return item, nil
}