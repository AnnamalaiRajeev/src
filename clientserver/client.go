package main

import (
	"fmt"
	"io"
	"net/http"
)

// returns the function that can be called to get the dogname
func Getdogname() func() string {

	anony := func() string {
		resp, err := http.Get("http://127.0.0.1:8080/get/")
		if err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			} else {
				y := string(body)
				fmt.Printf("%T\n", body)
				return y
			}

		}
	}
	return anony

}
func main() {
	x := Getdogname()
	fmt.Println(x())
}
