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


func (todo *Todo) Create(){
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
	
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


func (todo *Todo) Get(){
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()
        
        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

       	c := session.DB("TodoDB").C("todos")
 
        err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
        if err != nil {
                log.Fatal(err)
        }
}

func (todo *Todo) Update() {
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("TodoDB").C("todos")
	err = c.Update(bson.M{"_id":todo.Id}, bson.M{"$set": todo})
	//err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
	if err != nil {
		panic(err)
	}
}



func (todo *Todo) Delete() {	
	session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("TodoDB").C("todos")
	err = c.RemoveId(todo.Id)
	if err != nil {
		panic(err)
	}
		
}
