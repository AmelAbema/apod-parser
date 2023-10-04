package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	db *sql.DB
)

type APODResponse struct {
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"explanation"`
	URL         string `json:"url"`
	MediaType   string `json:"media_type"`
}

func initDB() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	db = conn
	return nil
}

func fetchAPODData() error {

	return nil
}

func getAPODData(w http.ResponseWriter, r *http.Request) {

}

func getAlbumEntries(w http.ResponseWriter, r *http.Request) {

}

func main() {
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/apod", getAPODData).Methods("GET")
	r.HandleFunc("/album", getAlbumEntries).Methods("GET")

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
