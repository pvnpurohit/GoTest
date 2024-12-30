package main

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println("INIT function")
}

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	s := `[{"First":"Praveen","Last":"Purohit","Age":42},{"First":"Sangu","Last":"Balakundi","Age":41}]`
	bs1 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs1)

	p1 := person{
		First: "Praveen",
		Last:  "Purohit",
		Age:   42,
	}
	p2 := person{
		First: "Sangu",
		Last:  "Balakundi",
		Age:   41,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	var people2 []person

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))

	err = json.Unmarshal(bs1, &people2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(people)
	for i, v := range people2 {
		fmt.Println("Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
