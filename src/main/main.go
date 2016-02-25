package main

import (
	"algo"
	"fmt"
	"io"
	"net/http"
	"os"
	"strcon"
	"time"
	//	"web"
)

func DefineMap() {
	dict := make(map[string]string)
	dict["go"] = "Golang"
	dict["cs"] = "CSharp"
	dict["rb"] = "Ruby"
	dict["py"] = "Python"
	dict["js"] = "JavaScript"
	fmt.Print("\n")
	for k, v := range dict {
		fmt.Printf("Key: %s Value: %s\t", k, v)
	}
	fmt.Print("\n")
	if lan, ok := dict["go"]; ok {
		fmt.Println(lan, ok)
	}
}
func DefineSlice() {
	//x := make([]int, 5,10)
	//A Slice Initializes for a Specific Length Without Providing Elements
	//x2 := []int{4: 0}
	//append and copy
	x := []int{10, 20, 30}
	y := append(x, 40, 50)
	fmt.Println(x, y)

	x1 := []int{1, 2, 3, 4, 5}
	y1 := make([]int, 6)
	copy(y1, x1)
	fmt.Println(x1, y1)

	x2 := []int{10, 20, 30}
	for k, v := range x2 {
		fmt.Printf("Index: %d Value: %d\t ", k, v)
	}
}

func DefineArray() {
	//define array, default is zero.
	x1 := [5]int{0: 101, 2: 12, 4: 22}
	fmt.Println("define array:", x1)
}

func ShowPersonInfo() {
	p := strcon.Person{
		"Shiju",
		"Varghese",
		time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		"shiju@email.com",
		"Kochi",
	}
	p.PrintName()
	p.PrintDetails()
}

func ShowAdminInfo() {
	alex := strcon.Admin{
		strcon.Person{
			"Alex",
			"John",
			time.Date(1970, time.January, 10,
				0, 0, 0, 0, time.UTC),
			"alex@email.com",
			"New York"},
		[]string{"Manage Team", "Manage Tasks"},
	}
	shiju := strcon.Member{
		strcon.Person{
			"Shiju",
			"Varghese",
			time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
			"shiju@email.com",
			"Kochi"},
		[]string{"Go", "Docker", "Kubernetes"},
	}
	//call methods for alex
	alex.PrintName()
	alex.PrintDetails()
	//call methods for shiju
	shiju.PrintName()
	shiju.PrintDetails()
}
func Init() {
	if algo.Fibonacci(6) == 8 {
		fmt.Println("fibonacci(6)== 8 ")
	}
	var _s = "Golang world, I are coming!"
	s := strcon.SwapCase(_s)
	fmt.Println("Converted string is :", s)
	DefineArray()
	DefineSlice()
	DefineMap()
	ShowPersonInfo()
	ShowAdminInfo()
}

type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func ChartHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func ReadFile() {
	file, err := os.Open("h5/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// create a buffer to keep chunks that are read

	buffer := make([]byte, 1024)
	for {
		// read a chunk
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// out the file content
		fmt.Println(string(buffer[:n]))

	}
}

func StartupServer() {
	mux := http.NewServeMux()
	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)

	mh2 := &messageHandler{"net/http web service component"}
	mux.Handle("/message", mh2)

	tag := "North Korea hurls insults at South Korea"
	mux.Handle("/chart-1", ChartHandler(tag))

	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)

	http.ListenAndServe(":8080", mux)
}

func main() {
	//read file
	ReadFile()
	//	Init()
	//	StartupServer()
	//	strcon.StartSimpleServer()
	//	strcon.StartMuxServer()
	//	StartHtmlTemplateServer()
	//	StartLogMiddlewareServer()
	//	StartGorillaHandlersServer()
	//	web.StartNegroniServer()
	StartGothServer()
}
