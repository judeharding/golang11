





// "{'name':'jude','address':'456 main'}"

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )


// func enableCors(res *http.ResponseWriter) {
// 	(*res).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*res).Header().Set("Access-Control-Allow-Methods", "*")
// 	(*res).Header().Set("Access-Control-Allow-Headers", "*")
// }


// func formHandler(res http.ResponseWriter, req *http.Request){

// 	enableCors(&res)
// 	// fmt.Fprintf(res, "POST request successful...xx\n")

//   type Member struct {
// 		Name string `json:"name"`
// 		Address string `json:"address"`
// 	}
// 	// instead of a hard coded string, we need to run off to a mysql and turns into a result insteadof a hard coded string
// 	jsonBody := []byte(`{"name":"jude","address":"456 main"}`)

// 	member := &Member{}
//     err := json.Unmarshal(jsonBody, member)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

//     fmt.Println(member.Name)
//     fmt.Println(member.Address)

// 		type ResponseObj struct {
// 			Sex string 
// 			Skill string 
// 		}
// 		// memberJSON := Member{member.Name, member.Address}
// 		responsObjJSON := ResponseObj{member.Name, member.Address}

// 		// jsonRes, err := json.Marshal(memberJSON)
// 		jsonRes, err := json.Marshal(responsObjJSON)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		// res.Header().Set("Content-Type", "application/json")
// 		res.Write(jsonRes)
		

// }

// // func helloHandler(res http.ResponseWriter, req *http.Request){
// // 	if req.URL.Path !="/hello"{ // checking the path
// // 		http.Error(res, "404 Not Found", http.StatusNotFound)
// // 		return
// // 	}
// // 	if req.Method != "GET" { // if hacking to another method, we will catch them
// // 		http.Error(res, "Method is not supported", http.StatusNotFound)
// // 		return
// // 	}
// // 	fmt.Fprintf(res, "hello.....") // print to the screen 


// // }

// func main(){

	
// // starting webserver
// 	// var fileServer = http.FileServer(http.Dir("./static"))  // go look in the STATIC folder for the index.html
// 	// http.Handle("/", fileServer) //handles the initializer ROUTE 
// 	http.HandleFunc("/form", formHandler)  // handles the form
// 	// http.HandleFunc("/hello", helloHandler)  // print hello

// 	fmt.Printf("Starting Server at port 8080 ... \n")

// 	var err = http.ListenAndServe("0.0.0.0:8080", nil) // listen and serve creates the server

// 	if err !=nil{
// 		log.Fatal(err)
// 	}
// }