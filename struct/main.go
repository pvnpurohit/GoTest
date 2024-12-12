package main

import "fmt"

func init() {
	fmt.Println("INIT function")
}

type person struct {
	firstname string
	lastname  string
	age       int
}

// embedded struct
type secretPerson struct {
	person
	ltk       bool
	firstname string //this field is for outer struct
}

type person2 struct {
	firstname string
	lastname  string
	favIC     []string
}

func main() {
	p1 := person{
		firstname: "Praveen",
		lastname:  "Purohit",
		age:       42,
	}
	p2 := person{
		firstname: "Sangeetha",
		lastname:  "Balakundi",
		age:       41,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)

	sp1 := secretPerson{
		person: person{
			firstname: "Praveen",
			lastname:  "Purohit",
			age:       42,
		},
		ltk:       true,
		firstname: "Richie",
	}
	fmt.Println(sp1)
	fmt.Println(sp1.person)
	fmt.Printf("%#v\n", sp1)
	fmt.Println(sp1.firstname, sp1.person.firstname) // see the difference in output

	//anonymous struct; you define the struct followed by values of the struct
	p3 := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "Advika",
		lastname:  "Purohit",
		age:       11,
	}

	fmt.Printf("%T \n %#v\n", p3, p3)

	p4 := person2{
		firstname: "Tom",
		lastname:  "Hanks",
		favIC:     []string{"choco", "vanilla", "strawberry"},
	}

	p5 := person2{
		firstname: "Brad",
		lastname:  "Pitt",
		favIC:     []string{"pista", "gauva", "custordapple"},
	}

	fmt.Println(p4)
	fmt.Println(p5)

	for _, v := range p4.favIC {
		fmt.Println(p4.firstname, "fav IC is", v)
	}

	for _, v := range p5.favIC {
		fmt.Println(p5.firstname, "fav IC is", v)
	}

	m := map[string]person2{
		p4.lastname: p4,
		p5.lastname: p5,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.firstname, v.lastname, v2)
		}
	}

	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Dwayne",
		friends: map[string]int{
			"Men":   5,
			"Women": 5,
		},
		favDrinks: []string{"cola", "juice", "water"},
	}
	fmt.Printf("%#v\n", as)

	for i, v := range as.favDrinks {
		fmt.Println("Fav drink", i, "of", as.first, "is :", v)
	}

}
