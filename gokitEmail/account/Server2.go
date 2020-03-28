package account

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex = &sync.Mutex{}

const Server1 string = "@here.com"

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, welcome to %s", Server1)
	fmt.Fprintf(w, "\nHenry@here.com")
}

func main() {
	http.HandleFunc("/", echoString)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8888", nil))

}
