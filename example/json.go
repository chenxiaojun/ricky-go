package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
}

type Response struct {
	Page   int `json:"page"`
	Fruits []string
}

func main() {
	booB, _ := json.Marshal(true)
	fmt.Println(string(booB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcB, _ := json.Marshal([]string{"A", "B"})
	fmt.Println(string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apple": 3})
	fmt.Println(string(mapB))

	structB, _ := json.Marshal(Person{"Ricky"})
	fmt.Println(string(structB))

	structC, _ := json.Marshal(&Person{"Ricky"})
	fmt.Println(string(structC))

	fmt.Println("")

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	fmt.Println(json.Marshal(dat))
	fmt.Println("")

	num := dat["num"]
	num1 := dat["num"].(float64)
	fmt.Printf("%T,%T\n", num, num1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	var res = &Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%+v\n", res)
	fmt.Println(res.Fruits[0])

	var s ServerSlice
	server := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(server), &s)
	fmt.Println(s)

}

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}
