package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	templateFile = template.Must(template.ParseFiles("Templates/index.html"))
)

type Promotion struct {
	ID             string `json:"id"`
	Price          string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/promotions", getPromotionID)

	log.Println("Starting server on port 1321")
	if err := http.ListenAndServe(":1321", nil); err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleUpload(w, r)
		return
	}
	err := templateFile.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("uploadedCSV")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fileName := path.Base("PreProcessed_" + fileHeader.Filename)
	dest, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Println("reading file " + fileHeader.Filename + " Start : " + time.Now().String())
	defer writeCSVtoDB(fileName)
}

func writeCSVtoDB(filePath string) {
	db := getDbConnection()

	mysql.RegisterLocalFile(filePath)

	query := "LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE promotions " +
		"FIELDS TERMINATED BY ',' " +
		"ENCLOSED BY '\"' " +
		"LINES TERMINATED BY '\\n' " +
		"IGNORE 0 LINES " +
		"(id, price, expiration_date);"

	_, err := db.Exec(query)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("All Records Inserted" + " End : " + time.Now().String())
}

func getDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	return db
}

func getPromotionID(w http.ResponseWriter, req *http.Request) {
	searchedId := req.FormValue("")

	if searchedId != "" {
		db := getDbConnection()

		result := db.QueryRow("SELECT * FROM promotions WHERE id=?", searchedId)

		var p Promotion
		err := result.Scan(&p.ID, &p.Price, &p.ExpirationDate)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(&p)
		if err != nil {
			return
		}
	}
}
