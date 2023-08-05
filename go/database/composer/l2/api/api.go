package api

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	fmt.Println("You win bro!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
	})
	http.ListenAndServe(":8000", nil)
}
