package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	year, month, day := time.Now().Date()
	date := fmt.Sprintf("[CDN startup] The date is %d %s %d", day, month, year)
	log.Output(1, date)
	http.HandleFunc("/coke_promotion_banner.webp", func(w http.ResponseWriter, r *http.Request) {
		_, month, _ := time.Now().Date()
		w.Header().Set("Content-Type", "image/webp")
		file, err := prepareFileBanner(month)
		if err != nil {
			log.Output(1, fmt.Sprint("[CDN] ", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(file)
	})

	http.HandleFunc("/product-images", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		product := queryParams.Get("product")
		w.Header().Set("Content-Type", "image/webp")
		file, err := prepareFileProduct(product)
		if err != nil {
			log.Output(1, fmt.Sprint("[CDN] ", err))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		w.Write(file)
	})

	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")
	log.Printf("Starting server on http://%s:%s\n", address, port)
	err := http.ListenAndServe(address+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func prepareFileBanner(month time.Month) ([]byte, error) {
	var buf []byte
	monthLower := strings.ToLower(month.String())
	entries, err := os.ReadDir(fmt.Sprint(monthLower))
	if err != nil || len(entries) < 1 {
		buf, err = os.ReadFile("default.webp")
		if err != nil {
			return nil, err
		} else {
			log.Output(1, "[CDN] serving default.webp")
		}
	} else {
		randomImageIndex := rand.Intn(len(entries))
		buf, err = os.ReadFile(fmt.Sprint(monthLower, "/", entries[randomImageIndex].Name()))
		if err != nil {
			return nil, err
		} else {
			log.Output(1, fmt.Sprint("[CDN] serving ", monthLower, "/", entries[randomImageIndex].Name()))
		}
	}

	return buf, nil
}

func prepareFileProduct(product string) ([]byte, error) {
	var buf []byte
	entries, err := os.ReadDir("products")
	if err != nil || len(entries) < 1 {
		return nil, err
	} else {
		buf, err = os.ReadFile(fmt.Sprint("products/", product, ".webp"))
		if err != nil {
			return nil, err
		} else {
			log.Output(1, fmt.Sprint("[CDN] serving ", "products/", product, ".webp"))
		}
	}

	return buf, nil
}
