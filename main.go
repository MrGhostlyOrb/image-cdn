package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	date := fmt.Sprint("[CDN startup] The date is ", day, " ", month, " ", year)
	log.Output(1, date)
	http.HandleFunc("/coke_promotion_banner.webp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/webp")
		file := prepareFile(month, year)
		w.Write(file)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func prepareFile(month time.Month, year int) []byte {
	buf, err := os.ReadFile(fmt.Sprint(year, "/", strings.ToLower(month.String()), ".webp"))

	if err != nil {
		buf, err = os.ReadFile("default.webp")
		if err != nil {
			log.Output(1, "[CDN] no default file found")
		} else {
			log.Output(1, "[CDN] serving default.webp")
		}
	} else {
		log.Output(1, fmt.Sprint("[CDN] serving ", year, "/", strings.ToLower(month.String()), ".webp"))
	}

	return buf
}
