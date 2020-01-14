package objects

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodPost {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil || len(buf) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sqlModel := post(string(buf))
		if len(sqlModel) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(sqlModel)
	} else {
		fmt.Println("bad")
	}
}
