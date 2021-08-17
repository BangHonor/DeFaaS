package cmd

/*
Copyright © 2021 kitchen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize the template of defaas",
	Long:  `Initialize the template of defaas`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initFaaSTemplate(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initFaaSTemplate() (err error) {

	var workingDir string = "."

	// defer func() {
	// 	// 如果出错，执行回滚，删除创建的文件
	// 	if err != nil {
	// 		devutils.ClearDir(workingDir)
	// 	}
	// }()

	// 要求工作目录是一个空目录
	// var isEmpty bool
	// isEmpty, err = devutils.IsDirEmpty(workingDir)
	// if err != nil {
	// 	return err
	// }
	// if !isEmpty {
	// 	err = fmt.Errorf("working directory is not empty, please use 'defaas init' in an empty directory")
	// 	return err
	// }

	// 创建文件夹
	// if err = os.Mkdir(path.Join(workingDir, "tools"), os.ModePerm); err != nil {
	// 	return err
	// }
	// // TODO
	// fmt.Println("generate tools/")

	if err := os.Mkdir(path.Join(workingDir, "accounts"), os.ModePerm); err != nil {
		return err
	}
	// TODO
	fmt.Println("generate directory accounts/")

	if err := os.Mkdir(path.Join(workingDir, "functions"), os.ModePerm); err != nil {
		return err
	}
	// TODO
	fmt.Println("generate directory functions/")

	return nil
}
