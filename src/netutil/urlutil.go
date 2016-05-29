package netutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//Test the web handling time
func FetchUrl(url string, ch chan<- string) {
	startTime := time.Now()
	//Do the normal logic
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	takenTime := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%f s %7d %s", takenTime, nBytes, url)
}
