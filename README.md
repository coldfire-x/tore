tore
====

extract information for web page

    

How to use
====

    package main
    
    import (
        "log"
        "net/http"
        "time"
    
        "github.com/pengfei-xue/tore"
        "github.com/pengfei-xue/tore/libs"
    )
    
    func main() {
        log.Println("Runing on 0.0.0.0:8080")
        s := &http.Server{
            Addr:         "0.0.0.0:8080",
            Handler:      tore.NewHttpHandler(),
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 10 * time.Second,
        }
        log.Fatal(s.ListenAndServe())
    }
    
    func init() {
        // tore.SetAlg("simple")
        tore.SetAlg("ttr")
    }
