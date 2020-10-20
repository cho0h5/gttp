package main

import (
        "log"
        "flag"
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
    // port
    port := flag.String("p", "8080", "port")
    flag.Parse()

    // server
    http.Handle("/", logging(http.FileServer(http.Dir("."))))

    log.Println("Start: http://localhost:" + *port)
    log.Fatal(http.ListenAndServe(":" + *port, nil))
}
