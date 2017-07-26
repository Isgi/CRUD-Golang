package main

import (
	"net/http"
	"log"
)
//func handleIndex(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "Hello, this is simple restful in golang")
//}
func main()  {
	r := Routes()
	log.Fatal(http.ListenAndServe(":8000", r))
}
