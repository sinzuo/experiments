package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//var sfd = make([]byte, 0, 10*1024)
	rr, _ := ioutil.ReadAll(r.Body)
	m1, err := r.Cookie("wori")
	if err == nil {
		fmt.Println(m1.Value)
	}

	//	c1 := &http.Cookie{Name: "wori", Value: "ok", MaxAge: 100000}
	//	http.SetCookie(w, c1)

	fmt.Println(string(rr))
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
