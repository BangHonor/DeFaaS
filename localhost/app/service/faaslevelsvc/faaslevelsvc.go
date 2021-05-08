package faaslevelsvc

import (
	"defaas/localhost/app/model"
	"strconv"
	"sync"
)

var (
	service     *FaaslevelSvc
	serviceOnce sync.Once
)

func Service() *FaaslevelSvc {

	serviceOnce.Do(func() {

		service = NewFaaslevelSvc()

	})

	return service
}

type FaaslevelSvc struct {
	levels []model.FaaslevelItem
}

func NewFaaslevelSvc() *FaaslevelSvc {

	svc := &FaaslevelSvc{}

	svc.levels = []model.FaaslevelItem{

		{
			ID:  "0",
			CPU: "1",
			Mem: "512",
		},

		{
			ID:  "1",
			CPU: "2",
			Mem: "1024",
		},

		{
			ID:  "3",
			CPU: "4",
			Mem: "2048",
		},
	}

	return svc
}

func (svc *FaaslevelSvc) List() ([]model.FaaslevelItem, error) {

	return svc.levels, nil

}

func (svc *FaaslevelSvc) Add(item model.FaaslevelItem) (model.FaaslevelItem, error) {

	// add ID of faas level
	item.ID = strconv.FormatInt(int64(len(svc.levels)), 10)

	svc.levels = append(svc.levels, item)

	return item, nil
}
