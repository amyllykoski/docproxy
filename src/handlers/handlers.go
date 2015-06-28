// handlers.go
package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request...")
	io.WriteString(w, "Hello schmooworld!")
}

func HandleRequest1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request1...")
	io.WriteString(w, "Hello world1!")
}

func HandleRequest2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request2...")
	io.WriteString(w, "Hello world2!")
}