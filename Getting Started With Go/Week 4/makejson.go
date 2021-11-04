package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, addr string
	fmt.Print("Enter a name")
	fmt.Scan(&name)
	fmt.Print("Enter an address")
	fmt.Scan(&addr)

	var mp map[string]string
	mp = make(map[string]string)
	mp["name"] = name
	mp["addr"] = addr

	jsonobj, err := json.Marshal(mp)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(jsonobj)
}
