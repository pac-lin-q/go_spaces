package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	//c, python, java := true, false, "no!"
	//fmt.Println(c, python, java)
	//res := http.ListenAndServe(":80", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	res := http.ListenAndServe(":80", mux)
	if res != nil {
		log.Fatal(res)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	fmt.Println("user:" + user)
	if user != "" {
		io.WriteString(w, fmt.Sprintf("url params is : [%s]\n", user))
	} else {
		io.WriteString(w, "this test is no params\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))

	}
	fmt.Println("exit root handler")
}
