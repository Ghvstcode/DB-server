package controllers

import (
	"fmt"
	"net/http"

	"github.com/Ghvstcode/RC/utils"
)

type mt map[string]string

var m mt

var kvs []mt

func Set(w http.ResponseWriter, r *http.Request) {
	beforeSet := len(kvs)
	m = make(mt)
	//Extract Out the Query
	rawQuery := r.URL.Query()
	//fmt.Println("RawQueryTest", rawQuery["nothing"])
	if len(rawQuery) < 1 {
		utils.Response(false, "Unable To Parse Request Query", http.StatusBadRequest).Send(w)
	}

	//fmt.Println("rq", rawQuery)
	//fmt.Println("Leng", len(rawQuery))

	for rq, v := range rawQuery {
		for _, ks := range kvs {
			if ks[rq] != "" {
				utils.Response(false, "A Value with this Key has already been set", http.StatusBadRequest).Send(w)
				return
			}

		}
		// Get the value
		val := v[0]
		//Check for duplicates in the existing KVS store.
		m = map[string]string{
			rq: val,
		}

		kvs = append(kvs, m)

		fmt.Println("KVS", kvs)
	}

	afterSet := len(kvs)

	if beforeSet >= afterSet {
		utils.Response(false, "An error occurred, unable to set Value", http.StatusBadRequest).Send(w)
		return
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	var retrievedVal string

	keys, ok := r.URL.Query()["key"]
	key := keys[0]
	if !ok {
		utils.Response(false, "Unable To Parse Request Query", http.StatusBadRequest).Send(w)
		return
	}

	for _, ks := range kvs {
		if ks[key] != "" {
			retrievedVal = ks[key]
			//utils.Response(true, ks[key], http.StatusOK).Send(w)
			//return
		}
	}

	res := utils.Response(true, "Unable to Retrieve Value", http.StatusBadRequest)
	fmt.Println(map[string]interface{}{key: retrievedVal})
	res.Data = map[string]interface{}{key: retrievedVal}
	res.Send(w)
	return //key := mux.Vars(r)["key"]
	//Pkeys := r.URL.Query()
	//Rkeys := r.URL.RawQuery
	//key := mux.Vars(r)["key"]
	//fmt.Println("KEY", value)
	//user := &models.User{}
	//err := json.NewDecoder(r.Body).Decode(user)
	//if err != nil {
	//	l.ErrorLogger.Println(err)
	//	utils.Response(false, "Invalid request", http.StatusBadRequest).Send(w)
	//	return
	//}
	//res := user.Create()
	//res.Send(w)
}
