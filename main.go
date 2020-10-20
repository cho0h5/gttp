package main

import (
        "log"
        "net/http"
       )

func logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)

        // log
		method := r.Method
		uri := r.URL.String()
        proto := r.Proto
        log.Println(method, uri, proto)
	})
}

func main() {
    http.Handle("/", logging(http.FileServer(http.Dir("."))))

    log.Println("Start: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
