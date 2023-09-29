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

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tip struct {
	ID        string `gorm:"autoincrement"`
	Category  string `gorm:"primaryKey"`
	Content   string `gorm:"primaryKey"`
	UpdatedAt time.Time
}

//go:embed tips.db
var dbContent string

//go:embed LICENSE
var license string

func main() {
	// Command-line flags
	addTip := flag.String("add", "", "Add a new tip to the specified category")
	category := flag.String("category", "vim", "Category for reading tips")
	dump := flag.Bool("dump", false, "Dump tips to files under 'tips' directory")
	printLicense := flag.Bool("license", false, "Print license")
	flag.Parse()
	if *printLicense {
		fmt.Println(license)
		return
	}

	// Load the embedded SQLite database
	dbPath := "/tmp/_totd_tips.db"
	if err := os.WriteFile(dbPath, []byte(dbContent), 0644); err != nil {
		fmt.Println("Error extracting database:", err)
		return
	}

	// Connect to the SQLite database using GORM
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Error getting database handle:", err)
		return
	}
	defer sqlDB.Close()

	// Auto Migrate the schema
	err = db.AutoMigrate(&Tip{})
	if err != nil {
		fmt.Println("Error migrating database:", err)
		return
	}

	// Check if the user wants to add a tip
	if *addTip != "" {
		err := addTipToDB(db, *addTip, *category)
		if err != nil {
			fmt.Println("Error adding tip:", err)
			return
		}
		fmt.Println("Tip added successfully!")
		return
	}

	// Check if the user wants to dump tips
	if *dump {
		err := dumpTips(db)
		if err != nil {
			fmt.Println("Error dumping tips:", err)
		}
		return
	}

	tips, err := readTipsFromDB(db)
	if err != nil {
		fmt.Println("Error reading tips:", err)
		return
	}

	if len(tips) == 0 {
		fmt.Println("No tips found.")
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(tips))
	tip := tips[randomIndex]

	fmt.Println(strings.Repeat("*", len(tip.Category)+len(tip.Content)+3))
	fmt.Println(tip.Category, tip.Content)
	fmt.Println(strings.Repeat("*", len(tip.Category)+len(tip.Content)+3))
}

func readTipsFromDB(db *gorm.DB) ([]Tip, error) {
	var tips []Tip

	result := db.Find(&tips)
	if result.Error != nil {
		return nil, result.Error
	}
	return tips, nil
}

func addTipToDB(db *gorm.DB, tipContent, category string) error {
	newTip := Tip{
		Category: category,
		Content:  tipContent,
	}
	result := db.Create(&newTip)
	return result.Error
}

func dumpTips(db *gorm.DB) error {
	var tips []Tip

	result := db.Find(&tips)
	if result.Error != nil {
		return result.Error
	}

	files := make(map[string]*os.File)
	for _, tip := range tips {
		if _, ok := files[tip.Category]; ok {
			continue
		}
		// Create the file for writing
		file, err := os.OpenFile(fmt.Sprintf("tips/%s.md", tip.Category), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		files[tip.Category] = file
	}

	for _, tip := range tips {
		file := files[tip.Category]

		// Write the tip to the file
		if _, err := file.WriteString("- " + tip.Content + "\n"); err != nil {
			return err
		}
	}

	fmt.Println("Tips dumped successfully.")
	return nil
}
