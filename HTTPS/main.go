package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	/* bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs)) */
	lw := logWritter{}
	io.Copy(lw, resp.Body)

}

//Custom writer func
func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
