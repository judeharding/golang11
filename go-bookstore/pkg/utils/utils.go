package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Unmarshal decodes JSON array elements into corresponding Go array elements
func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
} 
