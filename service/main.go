package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

const (
	DEFAULT_DISTANCE = "200km"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Product struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string `json:"user"`
	Description  string  `json:"description"`
	Category string `json:"category"`
	Location Location `json:"location"`
}

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/addproduct", handlerAddProduct)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerAddProduct(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body)
	var p Product
	if err := decoder.Decode(&p); err != nil {
		panic(err)
		return
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Description)
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")

	lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	// range is optional
	ran := DEFAULT_DISTANCE
	if val := r.URL.Query().Get("range"); val != "" {
		ran = val + "km"
	}

	fmt.Println("range is ", ran)

	// Return a fake post
	p := &Product{
		User:"1111",
		Description:"it's a smart phone",
		Category:"phone",
		Location: Location{
			Lat:lat,
			Lon:lon,
		},
	}

	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	fmt.Fprintf(w, "Search received: %s %s", lat, lon)
}


