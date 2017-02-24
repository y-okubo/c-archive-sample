package main

import (
	"C"
	"fmt"

	"gopkg.in/resty.v0"
)

//export GetUsers
func GetUsers(p *C.char) *C.char {
	fmt.Println(C.GoString(p))
	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return C.CString("")
	}

	return C.CString(resp.String())
}

func init() {}
func main() {}
