package funccodesvc

import (
	"defaas/localhost/app/model"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"io/ioutil"
	"log"
	"os"
	"path"
	//"sync"
)

var (
	service *FunccodeSvc
	//serviceOnce sync.Once
)

func Service() *FunccodeSvc {
	// serviceOnce.Do(func() { //只执行一次
	// })
	funccodesDirPath := "/home/dds/kitchen/defaas/tmp/funccodes"
	service, _ = NewFunccodeSvc(funccodesDirPath)
	//开发的时候放在外面，方便调试，不然只响应一次，改变了无法察觉

	return service
}

type FunccodeSvc struct {
	dirPath   string
	funccodes []model.FunccodeItem
}

func NewFunccodeSvc(funccodesDirPath string) (*FunccodeSvc, error) {
	svc := &FunccodeSvc{}
	svc.dirPath = funccodesDirPath
	svc.loadJSONAll()
	return svc, nil
}

func (svc *FunccodeSvc) loadJSONAll() {
	//遍历所有文件
	dir, err := ioutil.ReadDir("/home/dds/kitchen/defaas/tmp/funccodes/json")
	if err != nil {
		fmt.Println(err)
	}

	for _, fi := range dir {

		if fi.IsDir() { // 忽略目录
		} else {
			fmt.Println(fi.Name() + " ")
			var item model.FunccodeItem

			jsonFilePath := path.Join(svc.dirPath, "json", fi.Name())

			f, err := os.OpenFile(jsonFilePath, os.O_RDONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}

			dec := json.NewDecoder(f)
			if err := dec.Decode(&item); err != nil {
				log.Fatal(err)
			}
			svc.funccodes = append(svc.funccodes, item)

		}
	}
	g.Log().Println("[funccodesvc] load loadJSONAll done")
}

func storeFunccodeInJSON(jsonFilePath string, item *model.FunccodeItem) error {

	f, err := os.OpenFile(jsonFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")

	if err := enc.Encode(item); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (svc *FunccodeSvc) storeJSON(item *model.FunccodeItem) error {

	jsonFilePath := path.Join(svc.dirPath, "json", item.Name+".json")

	fmt.Println("FuncName", item.Name)
	if err := storeFunccodeInJSON(jsonFilePath, item); err != nil {
		return err
	}

	return nil
}

func (service *FunccodeSvc) List() ([]model.FunccodeItem, error) {
	//连接数据库
	return service.funccodes, nil
}

func (service *FunccodeSvc) Add(item model.FunccodeItem) (model.FunccodeItem, error) {

	// add name of funccode
	service.storeJSON(&item)

	service.funccodes = append(service.funccodes, item)

	g.Log().Printf("[funccodesvc] add funccode %v\n", item)

	return item, nil
}

func (service *FunccodeSvc) Delete(item model.FunccodeItem) (model.FunccodeItem, error) {
	logFile := "/home/dds/kitchen/defaas/tmp/funccodes/json/" + item.Name + ".json"
	err := os.Remove(logFile)
	if err != nil {
		fmt.Println("delete file failed!", err)
	}
	fmt.Println("delete file succeeded!")

	for index, value := range service.funccodes {
		if value.Name == item.Name {
			service.funccodes = append(service.funccodes[:index], service.funccodes[index+1:]...)
			break
			//只需要找到一次然后合并就行了
		}
	}

	return item, nil
}
