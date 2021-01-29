package main

import (
	"encoding/json"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		//w.Write([]byte("Hello " + name))
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("about"))
	})
	
	err := http.ListenAndServe(":3000", nil)
	
	if err != nil {
		log.Fatal(err)
	}
}