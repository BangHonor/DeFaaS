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

	"defaas/faas/market"

	"github.com/spf13/cobra"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// levelCmd represents the level command
var levelCmd = &cobra.Command{
	Use:   "level",
	Short: "show FaaS level information",
	Long:  `show FaaS level information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("level called")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list FaaS levels",
	Long:  `list FaaS levels`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		levels, _ := market.GetFaaSLevels()
		tablePrintFaaSLevels(levels)
	},
}

func init() {
	rootCmd.AddCommand(levelCmd)

	levelCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// levelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// levelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func tablePrintFaaSLevels(levels []market.FaaSLevel) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Core", "Mem")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, level := range levels {
		tbl.AddRow(level.ID, level.Core, level.Mem)
	}

	tbl.Print()
}
