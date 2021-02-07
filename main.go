package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
)

func main() {
    s := &http.Server{
        Addr: ":8081",
        Handler: serverHandler(),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
    }
    fmt.Printf("server start listening on %s\n", s.Addr)
    err := s.ListenAndServe()
    if err != nil {
        fmt.Println(err)
    }
}

func serverHandler() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        i := 1
        for {
            w.Write([]byte(strconv.Itoa(i) + "\n"))
            if f, ok := w.(http.Flusher); ok {
                f.Flush()
                fmt.Println(i)
            }
            <-time.After(1 * time.Second)
            i += 1
        }
    })
}

