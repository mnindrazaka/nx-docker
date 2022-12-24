package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func main() {
	dsn := "root:((root))@tcp(127.0.0.1:3306)/biodata?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var students []Student
		result := db.Find(&students)
		if result.Error != nil {
			w.Write([]byte("Failed to fetch data"))
			return
		}

		jsonData, err := json.Marshal(students)
		if err != nil {
			w.Write([]byte("Failed to fetch data"))
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		w.Write(jsonData)
	})
	http.ListenAndServe(":3000", router)
}
