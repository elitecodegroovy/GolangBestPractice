package basic

import "fmt"

//TODO .... refactor the content.
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

//There are two critical points to make about subtyping in Go:
//
//1 We can use Anonymous fields to adhere to an interface. We can also adhere to many interfaces.
// By using Anonymous fields along with interfaces we are very close to true subtyping.
//
//2 Go does provide proper subtyping capabilities, but only in the using of a type. Interfaces can
// be used to ensure that a variety of different types can all be accepted as input into a function,
// or even as a return value from a function, but in reality they retain their distinct types. This is
// clearly displayed in the main function where we cannot set Name on Citizen directly because Name
// isn’t actually a property of Citizen, it’s a property of Person and consequently not yet present
// during the initialization of a Citizen.
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
