/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"english-in-use/models"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Get a text and save useful words",
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")

		buf, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		r, _ := regexp.Compile(`\w{4,}`)
		words := r.FindAll(buf, -1)

		maxLen := 0
		for i := 0; i < len(words); i++ {
			if maxLen < len(words[i]) {
				maxLen = len(words[i])
			}
		}

		fmt.Println("Press <- for go back to previous word")
		fmt.Println("Press -> for skip current word")
		fmt.Println("Press Enter for save word")

		savedWords := make(map[string]bool)

		for i := 0; i < len(words); i++ {
			word := strings.ToLower(string(words[i]))

			fmt.Print("\r")
			for j := 0; j < maxLen; j++ {
				fmt.Print(" ")
			}
			fmt.Printf("\r%s", word)

			err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
				switch key.Code {
				case keys.Right:
				case keys.Enter:
					if !savedWords[word] {
						savedWords[word] = true
					}
				case keys.Left:
					if i == 0 {
						return false, nil
					}
					i -= 2
				case keys.CtrlC:
					return false, errors.New("cancel")
				default:
					return false, nil
				}
				return true, nil
			})
			if err != nil {
				break
			}
		}
		fmt.Println()

		var mWords []models.Word
		for word := range savedWords {
			mWords = append(mWords, models.Word{
				Value:  word,
				Source: source,
			})
		}

		db, err := gorm.Open(sqlite.Open(models.GetDbPath()))
		if err != nil {
			log.Fatal("Failed to open database:", err)
		}

		db.AutoMigrate(&models.Word{})

		err = db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(mWords, len(mWords)).Error
		if err != nil {
			log.Fatal("Failed to create in batches:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.PersistentFlags().StringP("source", "s", "general", "Source of this word. e.g a music name, web, book name, movie and etc.")
}
