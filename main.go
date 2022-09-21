package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	var port string
	var ok bool

	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "8000"
	}

	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("New request:", r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
	})

	//Define routes
	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := getParamsAandB(w, r)
		if err != nil {
			return
		}
		result := add(a, b)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprint(result))
	})

	r.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := getParamsAandB(w, r)
		if err != nil {
			return
		}
		result := subtract(a, b)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprint(result))
	})

	r.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := getParamsAandB(w, r)
		if err != nil {
			return
		}
		result := multiply(a, b)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprint(result))
	})

	r.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := getParamsAandB(w, r)
		if err != nil {
			return
		}
		result := divide(a, b)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprint(result))
	})

	r.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		a, b, err := getParamsAandB(w, r)
		if err != nil {
			return
		}
		result := divide(a, b)
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, fmt.Sprint(result))
	})

	//UNCOMMENT FOR V2

	/*
		r.HandleFunc("/modulo", func(w http.ResponseWriter, r *http.Request) {
			a, b, err := getParamsAandB(w, r)
			if err != nil {
				return
			}
			result := modulo(a, b)
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, fmt.Sprint(result))
		})

		r.HandleFunc("/pow", func(w http.ResponseWriter, r *http.Request) {
			a, b, err := getParamsAandB(w, r)
			if err != nil {
				return
			}
			result := pow(a, b)
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, fmt.Sprint(result))
		})
	*/

	log.Println("Staring server in port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getParamsAandB(w http.ResponseWriter, r *http.Request) (int, int, error) {
	error_msg := "Parameters are not correct"
	a, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, error_msg)
		return -1, -1, err
	}

	b, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, error_msg)
		return -1, -1, err
	}

	return a, b, nil

}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) float32 {
	return float32(a) / float32(b)
}

func modulo(a, b int) int {
	return a % b
}

func pow(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}
