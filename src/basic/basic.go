package basic

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"
	"syscall"
	"text/tabwriter"
	"time"
	"unicode"
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

	//error inteface
	DoError()

	//type assertion
	DoTypeAssertion()
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

//error interface
func DoError() {
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"

}

//A type assertion is an operation applied to an interface value
//A type assertion checks that the dynamic type of its operand matches the asserted type.
func DoTypeAssertion() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	fmt.Println("type assertion:", f)
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Println("panic type assertion:", c)

	//error and type assertion
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)

	DoChannel()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//channel
/**
conn, err := net.Dial("tcp", "127.0.0.1:8000")
if err != nil {
	log.Fatal(err)
}
done := make(chan struct{})
go func() {
	io.Copy(os.Stdout, conn) // NOTE: ignoring errors
	log.Println("done")
	done <- struct{}{}
	// signal the main goroutine
}()
mustCopy(conn, os.Stdin)
conn.Close()
<- done  // wait for background goroutine to finish
*/

func count(out chan<- int) {
	for i := 0; i < 101; i++ {
		out <- i
	}
	close(out)
}

func square(input <-chan int, output chan<- int) {
	for n := range input {
		output <- n * n
	}
	close(output)
}

func print(input <-chan int) {
	for p := range input {
		fmt.Print("\t", p)
	}
}

/**
// Counter
go func() {
	for x := 0; x < 101 ; x++ {
		naturals <- x
	}
	close(naturals)
}()
// Squarer
go func() {
	for {
		x, ok := <-naturals
		if !ok {
			break //channel was closed and drained
		}
		squares <- x * x
	}
	close(squares)
}()
// Printer (in main goroutine)
for x := range squares {
	fmt.Print("\t", x)
}
fmt.Println()
*/
func DoChannel() {
	naturals := make(chan int)
	squares := make(chan int)
	go count(naturals)
	go square(naturals, squares)
	print(squares)
	fmt.Println()
	DoConcurrent()
}

//Buffered Channels
var verbose = flag.Bool("v", true, "show verbose progress messages")

func DoConcurrent() {
	// ...start background goroutine...

	//!-
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	//!+
	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(1 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

// Package word provides utilities for word games.
// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudorandom  number generator rng.
func RandomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

var out io.Writer = os.Stdout // modified during testing

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func Display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			Display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			Display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			Display(fmt.Sprintf("%s[%s]", path,
				FormatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			Display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			Display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, FormatAtom(v))
	}
}
