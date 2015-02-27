package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome")
}


func TodoList(w http.ResponseWriter, r *http.Request){
	todos := Todos{}
	todos = todos.List()
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}


func TodoDetail(w http.ResponseWriter, r *http.Request){
	var todo Todo
        todo.Get(mux.Vars(r)["id"])
        if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
        }
}

func TodoCreate(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    	var todo Todo   
    	err := decoder.Decode(&todo)
    	if err != nil {
        	panic(err)
   	}
	todo.Create("asdas")	
	if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
        }
}
