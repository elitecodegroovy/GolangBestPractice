package basic

import "fmt"

type rectangle struct {
	width  int
	height int
}
type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

type Citizen struct {
	Country string
	Person
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Human interface {
	Talk()
}

func SpeakTo(h Human) {
	h.Talk()
}

func ShowMeObject() {
	r := rectangle{width: 8, height: 5}
	fmt.Printf("rectangle %v， area：%d", r, r.area())
	fmt.Println()
}

func StartComposeGo() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gotham",
			State:  "NY",
			Zip:    "01313",
		},
	}

	p.Talk()
	p.Location()

	c := Citizen{}
	c.Name = "Steve"
	c.Country = "America"
	c.Talk()
	c.Nationality()

	//Subtype with interface method.
	SpeakTo(&p)
	SpeakTo(&c)
}
