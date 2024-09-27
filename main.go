package main

import (
	"fmt"
	"log"
	"math/rand"
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
		file := prepareFile(month)
		w.Write(file)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func prepareFile(month time.Month) []byte {
	var buf []byte
	entries, err := os.ReadDir(fmt.Sprint(strings.ToLower(month.String())))
	if err != nil || len(entries) < 1 {
		buf, err = os.ReadFile("default.webp")
		if err != nil {
			log.Output(1, "[CDN] no default file found")
		} else {
			log.Output(1, "[CDN] serving default.webp")
		}
	} else {
		buf, err = os.ReadFile(fmt.Sprint(strings.ToLower(month.String()), "/", entries[rand.Intn(len(entries))].Name()))
		if err != nil {
			log.Output(1, fmt.Sprint("[CDN] serving ", strings.ToLower(month.String()), "/", entries[rand.Intn(len(entries))].Name()))
		}
	}

	return buf
}
