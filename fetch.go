//Goes and gets a page source from a URL and prints it in the terminal window
//To run, compile then ./fetch <url>
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

func main (){
  for _, url := range os.Args[1:]{
     // make a http request and if no error return the response in a struct called resp
    resp, err := http.Get(url)
    if err != nil { // if error is not nil, then there is an error
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // so print the error to Stderr
      os.Exit(1) //and exit with status code 1
    }
    //next read the entire response and store it in a variable called b
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close() // close the body stream after reading it
    if err != nil { // if error is not nil, then there is an error
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err) // so print the error
      os.Exit(1) //and exit with status code 1
    }
    fmt.Printf("%s", b) //spam the whole page source to terminal window
  }
}
