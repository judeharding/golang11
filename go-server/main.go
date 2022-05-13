package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	// var err = r.ParseForm()
	// if err != nil {
	if err := r.ParseForm(); err != nil {
		fmt. Fprintf(w, "ParseForm() err: %v ", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n...")

	var name = r.FormValue("name")
	var address = r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	fmt.Printf("NAME is %v ... \n", name)
	fmt.Printf("ADDRESS is %v ... \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello.....")
}

func main(){
// starting prog
	var fileServer = http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer) //handles the initializer
	http.HandleFunc("/form", formHandler)  // handles the form
	http.HandleFunc("/hello", helloHandler)  // print hello

	fmt.Printf("Starting Server at port 8080 ... \n")

	var err = http.ListenAndServe(":8080", nil)

	if err !=nil{
		log.Fatal(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"time"
// )
// func main(){

// var start  = time.Now()

// var sum = 0

// 	for i := 0; i <= 1000000; i++ {
// 		sum = i 
// 		fmt.Printf("%v\n", sum)
// 	}

// t := time.Now()
// elapsed := t.Sub(start) * 1000


// fmt.Printf("\nENDING NOW...  %v", elapsed)

// }