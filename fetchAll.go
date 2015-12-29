package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //create a channel of strings using make

	for _, url := range os.Args[1:] { //for each url passed to commandline
		go fetch(url, ch) //start a goroutine that calls fetch() asynchronously
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive summary lines from channel ch and print them
	}
	// %.2f = default width, precision 2 (formats % value to two decimal places)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()        //keep track of the time
	resp, err := http.Get(url) //fetch the url using http.get
	if err != nil {            // Do the following if there is an error...
		ch <- fmt.Sprint(err) // send the error to channel ch
		// Note, Sprint formats using default formats and returns the resulting string
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	//the io.Copy function reads the response and discards it by writing it to ioutil.Discard output stream
	// io.Copy returns the byte count and any error that happens
	resp.Body.Close() // close the Body stream after reading it so don't leak resources
	if err != nil {   //if there is an error...
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send the error to channel ch
		return
	}
	secs := time.Since(start).Seconds()
	// as each result arrives, send a summary line on channel ch
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
