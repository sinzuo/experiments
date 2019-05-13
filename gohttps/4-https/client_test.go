package main

//go test -v -bench . client_test.go
//go test -v -cpu 4 -count 1 -bench  . client_test.go
import (
	"bytes"
	"crypto/tls"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

type A struct {
	uu   int
	Name string
}

func BenchmarkA1(t1 *testing.B) {
	for i := 1; i < t1.N; i++ {
		fmt.Sprintln("ok")
		time.Sleep(time.Second * 1)

	}

}

func TestA1(t1 *testing.T) {
	fmt.Println("ok")

}

func main() {

	a1 := &A{uu: 2, Name: "jiang"}
	a2 := &bytes.Buffer{}
	a3 := gob.NewEncoder(a2)
	a3.Encode(a1)
	a4, e := os.OpenFile("jiang.txt", os.O_RDWR, os.ModePerm)
	if e != nil {
		fmt.Println("kkkk")
	}
	io.Copy(a4, a2)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8081")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
