package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YadavPrasad/fileship/go-backend/models"
	"github.com/YadavPrasad/fileship/go-backend/utils"
)

func main() {
	utils.InitDB()

	// Run auto migration
	err := utils.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Failed to migrate DB:", err)
	}

	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong from Go + NeonDB + GORM 🚀")
	})

	fmt.Println("🚀 Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
