package utils

import (
	"encoding/json"
	"io"
	"net/http"
)


func ParseBody(req *http.Request, data interface {}){
	body, err := io.ReadAll(req.Body)

	if(err == nil){
		err := json.Unmarshal([]byte(body), data)

		if(err != nil){
			return
		}
	}
}