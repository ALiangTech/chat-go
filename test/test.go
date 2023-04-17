package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apikeyResponse, err := http.Get("https://ny2qtwmutr.hk.aircode.run/getapikey")
	if err != nil {

	}
	apikeyBody, _ := io.ReadAll(apikeyResponse.Body)

	fmt.Print(string(apikeyBody))
}
