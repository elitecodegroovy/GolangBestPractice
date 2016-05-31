package basic

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

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

func ThinkingInInterface() {
	var w io.Writer
	w = os.Stdout
	n, err := w.Write([]byte("hello"))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(",len: ", n) // OK: io.Writer has Write method
	defineInte()

	//Parsing Flags with flag.Value
	doParseFlag()

	//sort.interface
	DoSort()
}

//interface satisfaction
//This may seem useless, but in fac t the typ e interface{}, which is
//called the empty interface type, is indispensable.
func defineInte() {
	var any interface{}
	any = true
	fmt.Print(any, "\t")
	fmt.Printf("%T\n", any) // "<nil>"
	any = 12.34
	fmt.Print(any, "\t")
	any = "hello"
	fmt.Print(any, "\t")
	any = map[string]int{"one": 1}
	fmt.Print(any, "\t")
	any = new(bytes.Buffer)
	fmt.Print(any, "\t")

	// *bytes.Buffer must satisfy io.Writer
	var w io.Writer = new(bytes.Buffer)
	fmt.Print(w, "\t")
}

//Parsing Flags with flag.Value
var period = flag.Duration("period", 1*time.Second, "sleep period")

func doParseFlag() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

const debug = true

func doInterfaceValue() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect if define the variable 'buf *bytes.Buffer'
	if debug {
		// ...use buf...
	}
}

// If out is nonnil,
//output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

//sort.interface
func DoSort() {
	DoSort4Int()
	DoObjSort()
}
func DoObjSort() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	//!-customcall
	printTracks(tracks)
}
func DoSort4Int() {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!+customcode
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
