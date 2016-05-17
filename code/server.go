package main
import (
    "fmt"
    "golang.org/x/net/http2"
    "html"
    "log"
    "net/http"
    "encoding/json"
)

type jsonData struct {
    User string
    Sql  string
}


func ShowRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Method: %s\n", r.Method)
    fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
    fmt.Fprintf(w, "Host: %s\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
    fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
    fmt.Fprintf(w, "URL: %#v\n", r.URL)
    fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
    fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
    fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
    fmt.Fprintf(w, "\nHeaders:\n")
    r.Header.Write(w)
}

func main() {
    var srv http.Server
    http2.VerboseLogs = true
    srv.Addr = ":8443"
    // This enables http2 support
    http2.ConfigureServer(&srv, nil)
 
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "你在使用HTTP2协议访问 %q\n", html.EscapeString(r.URL.Path))
        ShowRequestInfoHandler(w, r)
    })
    
    http.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {  
        //get "POST" request
        if r.Method == "POST" {  
             decoder := json.NewDecoder(r.Body)
             var t jsonData   
             err := decoder.Decode(&t)
             if err != nil {
                 panic(err)
             }
             log.Println(t.User, t.Sql)
        } 
    })
    
    // Listen as https ssl server
    // To self generate a test ssl cert/key you could go to
    // http://www.selfsignedcertificate.com/
    // or use the openssl manual
    // or create with openssl or let's encypt
    log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}
