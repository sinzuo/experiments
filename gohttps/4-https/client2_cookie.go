package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	bb := bytes.NewBuffer([]byte("jiangyibo"))
	mm := &http.Cookie{Name: "jiang", Value: "sfdfsafds", MaxAge: 33333}
	req, _ := http.NewRequest("GET", "https://localhost:8081", bb)
	req.AddCookie(mm)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
