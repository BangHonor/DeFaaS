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

	devutils "defaas/dev-cmd/utils"

	devutilscmd "github.com/go-cmd/cmd"

	"github.com/spf13/cobra"
)

const (
	constAccountsDirPath = "./accounts"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account",
	Long:  `account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("account called")
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate [no options!]",
	Short: "generate an address",
	Long:  `generate an address`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		// 使用 FISCO BCOS 官方提供的 shell 脚本创建
		// https://fisco-bcos-documentation.readthedocs.io/zh_CN/latest/docs/manual/account.html
		devutils.RunCmd(devutilscmd.NewCmd(
			"./tools/get_account.sh"))

		fmt.Println("new address generated")
	},
}

var switchCmd = &cobra.Command{
	Use:   "switch [account]",
	Short: "switch current account",
	Long:  `switch current account to the specified account`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		var addressHex string = args[0]
		if len(addressHex) != 42 || addressHex[0:2] != "0x" {
			return fmt.Errorf("invalid format, a 42-length string like '0x00000123456789ABCDEF00000123456789ABCDEF' is expected")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := checkAccountsDir(); err != nil {
			log.Fatal(err)
		}
		pemFilePath := getPemFilePath(args[0])
		if err := checkPemFile(pemFilePath); err != nil {
			log.Fatal(err)
		}

		// TODO switch

		fmt.Printf("load pem file [%s]\n", pemFilePath)
		fmt.Printf("switch current account to [%s]\n", args[0])
	},
}

func init() {

	// 添加到根命令
	rootCmd.AddCommand(accountCmd)

	// 添加子命令
	accountCmd.AddCommand(generateCmd)
	accountCmd.AddCommand(switchCmd)
}

// ------------------------------------------------------------------------------------------------

func getPemFilePath(addressHex string) string {
	return path.Join(constAccountsDirPath, addressHex+".pem")
}

func checkAccountsDir() error {
	stat, err := os.Stat(constAccountsDirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", constAccountsDirPath)
	}
	if !stat.IsDir() {
		return fmt.Errorf("'%s' should be a directory", constAccountsDirPath)
	}
	return nil
}

func checkPemFile(pemFilePath string) error {
	stat, err := os.Stat(pemFilePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", pemFilePath)
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' should be a file", pemFilePath)
	}
	return nil
}
