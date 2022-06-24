// "{'name':'jude','address':'456 main'}"

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func enableCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "*")
	(*res).Header().Set("Access-Control-Allow-Headers", "*")
}


func formHandler(res http.ResponseWriter, req *http.Request){
	//request is the user sending and response is coming FROM the server.
	// star is a pointer. 
	// var err = r.ParseForm()
	// if err != nil {
	enableCors(&res)


	// if err := req.ParseForm(); err != nil {
	// 	fmt. Fprintf(res, "ParseForm() err: %v ", err)
	// 	return
	// }

	fmt.Fprintf(res, "POST request successful...\n")

  type Member struct {
		name string `json:"name"`
		address string `json:"address"`
	}
	
	jsonBody, _ := ioutil.ReadAll(req.Body)
	// maybe, _ := strconv.Unquote(fmt.Sprintf(`"%s"`, jsonBody))
	// if err != nil{
	// 	log.Fatal("FATAL ERROR IN formhandler-JSONBODY", err)
	// }


	fmt.Printf("%s", jsonBody)
	// fmt.Sprintf("%s", maybe)



	// var response Member 
	// json.Unmarshal(jsonBody, &response)

	// fmt.Println("PRINTING.....x ", jsonBody)
	// fmt.Println("RESPONSE???", response.name)




	// var name = req.FormValue("name")
	// var address = req.FormValue("address")

	// prints to the screen 
	// fmt.Fprintf(res, "Name = %s\n", name)
	// fmt.Fprintf(res, "Address = %s\n", address)

	// prints to the terminal
	// fmt.Printf("NAME is %v ... \n", name)
	// fmt.Printf("ADDRESS is %v ... \n", address)
}

func helloHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path !="/hello"{ // checking the path
		http.Error(res, "404 Not Found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" { // if hacking to another method, we will catch them
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "hello.....") // print to the screen 


}

func main(){

	
// starting webserver
	// var fileServer = http.FileServer(http.Dir("./static"))  // go look in the STATIC folder for the index.html
	// http.Handle("/", fileServer) //handles the initializer ROUTE 
	http.HandleFunc("/form", formHandler)  // handles the form
	// http.HandleFunc("/hello", helloHandler)  // print hello

	fmt.Printf("Starting Server at port 8080 ... \n")

	var err = http.ListenAndServe(":8080", nil) // listen and serve creates the server

	if err !=nil{
		log.Fatal(err)
	}
}