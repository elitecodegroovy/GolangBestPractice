package main

import (
	"algo"
	"channel"
	"conc"
	"fmt"
	"io"
	"net/http"
	"os"
	"strcon"
	"time"
	"util"
	//	"web"
)

//There are several ways to create and initialize maps
func DefineMap() {
	refl := make(map[int]int)
	refl[1] = 10
	refl[2] = 20
	for k, r := range refl {
		fmt.Println("reflection k:", k, ", v:", r)
		if k == 1 {
			fmt.Println("get key :", k)
		}
	}

	if v, ok := refl[2]; ok {
		fmt.Println("get key 2 :", ok, ", value:", v)
	}
}

//Funfamental Slice struct operations
func DefineSlice() {
	//Specifying Length and Capacity in a Slice with the make Function
	//If the slice capacity is not specified, the capacity is the same as the length.
	p := make([]int, 3, 5)
	p[1] = 10
	p[2] = 20
	fmt.Println(p)

	p1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(p1)

	p2 := []string{"a", "b", 4: "eeee"}
	fmt.Println(p2)

	//append and copy
	p3 := append(p2, "f", "g")
	fmt.Println(p2, p3)

	p4 := []int{1, 2, 30, 4, 5}
	p5 := make([]int, 6)
	copy(p5, p4)
	fmt.Println(p4, p5, len(p4), cap(p5))

	//iterator over Slice type
	for k, v := range p5 {
		fmt.Println("index ", k, " value:", v)
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

//func StartSimpleWeb() {
//	http.ListenAndServe(":8080", util.SetUserRoutes())
//}

func DoCollection() {
	input, expected := "Hello, World", "dlroW ,olleH"
	fmt.Printf("%s was reversed %s\n", input, expected)

	//array demo
	var x1 [5]int
	x1[0] = 3
	x1[2] = 1
	fmt.Println(x1)
	//init array
	x2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x2)
	//specified element
	x3 := [3]int{9, 2: 4}
	fmt.Println(x3)

	x4 := [2]string{"Go", "Golang"}
	fmt.Println(x4)
	x5 := [3]string{"Beego", 1: ",", 2: "Gin"}
	fmt.Println(x5)

	//Slice
	DefineSlice()
	DefineMap()
}

func StartConcurrecy() {
	//	conc.StartSimple()
	//	conc.StartSyncSimple()
	conc.StartReadText()
	conc.StartChan()

	//start channel
	channel.StartChannel()
}

func main() {
	start := time.Now() // get current time

	//read file
	//	ReadFile()
	//	web.StartTLSHttp()
	//	DoCollection()
	StartConcurrecy()
	//	StartSimpleWeb()
	//	web.GetEtcdInfo()
	//web.SendShortEmail("liujigang@mama100.com", "Connect to the server, authenticate, set the sender and recipient.", "testGolangEmail")
	//	Init()
	//	StartupServer()
	//	strcon.StartSimpleServer()
	//	strcon.StartMuxServer()
	//	StartHtmlTemplateServer()
	//	StartLogMiddlewareServer()
	//	StartGorillaHandlersServer()
	//	web.StartNegroniServer()
	//TODO ...fixed the invalid arguments "invalid argument"
	//	StartGothServer()

	util.StartDemo()
	//	util.StartCac()
	elapsed := time.Since(start)
	fmt.Printf("time elapsed: %s\n", elapsed)
}
