package main

import (
        //"fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Todo struct{
	Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Uri 		string     `json:"uri"`
	Name		string	   `json:"name"`
	Completed	bool	   `json:"completed"`
}

type Todos []Todo


func (todo *Todo) Create(domain string){
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
	
	todo.Id = bson.NewObjectId()
	todo.Uri = domain+"/todos/"+string(todo.Id.Hex())
        c := session.DB("TodoDB").C("todos")
	err = c.Insert(todo)
        if err != nil {
                log.Fatal(err)
        }


	err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
        if err != nil {
                log.Fatal(err)
        }	
}

func (todos Todos) List() Todos{
	session, err := mgo.Dial("localhost:27017")
        if err != nil { 
                panic(err)
        }       
        defer session.Close()
        
        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true) 
        
        c := session.DB("TodoDB").C("todos")
	err = c.Find(nil).All(&todos)
        if err != nil {
                log.Fatal(err)
        }
	return todos	
}


func (todo *Todo) Get(todo_id string){
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        
        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

       	c := session.DB("TodoDB").C("todos")
 
        todo.Id = bson.ObjectIdHex(todo_id)
        
        err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
        if err != nil {
                log.Fatal(err)
        }
}
