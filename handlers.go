package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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
	todo.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
        todo.Get()
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
	todo.Id = bson.NewObjectId()
        todo.Uri = "http://"+r.Host+"/todos/"+string(todo.Id.Hex())+"/"
	todo.Create()	
	if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
        }
}


func TodoUpdate(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var todo Todo
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	
	todo.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
        todo.Uri = "http://"+r.Host+r.URL.Path
	todo.Update()
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}	

}

func TodoDelete(w http.ResponseWriter, r *http.Request){
	var todo Todo
	todo.Id = bson.ObjectIdHex(mux.Vars(r)["id"])
        todo.Delete()
        if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
        }
}
