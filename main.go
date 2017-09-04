package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	fb "github.com/huandu/facebook"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Share struct {
	Count json.Number
}

func main() {
	r := httprouter.New()

	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		params := r.URL.Query()
		url := params.Get("url")

		total_share := GetCount(url)
		t := Share{
			Count: total_share,
		}

		ts, _ := json.Marshal(t)

		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", ts)
	})
	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func GetCount(url string) json.Number {
	res, _ := fb.Get("/", fb.Params{
		"id":           url,
		"access_token": "1034894286618560|wzqIO5AFDSXaGfp5yXVBpf7o9_M",
	})

	share := res["share"].(map[string]interface{})
	total_share := share["share_count"].(json.Number)
	return total_share
}
