// MIT License
//
// Copyright (c) 2023 Krishna Miriyala
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var category string
var content string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Tips",
	Run: func(cmd *cobra.Command, args []string) {
		db := connectToDB()
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("Error getting database handle:", err)
			return
		}
		defer sqlDB.Close()
		if err = addTipToDB(db, content, category); err != nil {
			fmt.Println("Error dumping tips:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&category, "category", "vim", "Tip's category")
	addCmd.Flags().StringVar(&content, "content", "", "Tip's content")
}
