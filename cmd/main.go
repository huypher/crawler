package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://voz.vn"
	fmt.Println("Downloading ", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0))

	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(buf.String())

	return
}
