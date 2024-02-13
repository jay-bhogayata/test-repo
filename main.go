package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {

		type response struct {
			Message string `json:"message"`
		}

		res, _ := json.Marshal(&response{Message: "OK"})

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	r.HandleFunc("POST /echo/{name}", func(w http.ResponseWriter, r *http.Request) {

		name := r.PathValue("name")

		type Response struct {
			Message string `json:"message"`
		}

		rr, err := json.Marshal(&Response{Message: name})
		if err != nil {
			log.Println(err)
			http.Error(w, "error in marshaling json", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(rr)

	})

	h := slog.NewJSONHandler(os.Stdin, nil)
	l := slog.New(h)
	log.Println("server is starting on port 8080")
	l.Info("Server is starting on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("error in starting the server %v", err.Error())
	}

}
