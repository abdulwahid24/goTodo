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

func getDBCollection() (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial("mongodb://abdul:abdul123@ds049641.mongolab.com:49641/gotodo")
        if err != nil {
                panic(err)
        }

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("gotodo").C("todos")
	return c, session
}

func (todo *Todo) Create(){
        c, session := getDBCollection()
	defer session.Close()
	err := c.Insert(todo)
        if err != nil {
                log.Fatal(err)
        }


	err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
        if err != nil {
                log.Fatal(err)
        }	
}

func (todos Todos) List() Todos{
        c, session := getDBCollection()
	defer session.Close()
	err := c.Find(nil).All(&todos)
        if err != nil {
                log.Fatal(err)
        }
	return todos	
}


func (todo *Todo) Get(){
       	c, session := getDBCollection()
	defer session.Close()
        err := c.Find(bson.M{"_id": todo.Id}).One(&todo)
        if err != nil {
                log.Fatal(err)
        }
}

func (todo *Todo) Update() {
        c, session := getDBCollection()
	defer session.Close()
	err := c.Update(bson.M{"_id":todo.Id}, bson.M{"$set": todo})
	//err = c.Find(bson.M{"_id": todo.Id}).One(&todo)
	if err != nil {
		panic(err)
	}
}



func (todo *Todo) Delete() {	
        c, session := getDBCollection()
	defer session.Close()
	err := c.RemoveId(todo.Id)
	if err != nil {
		panic(err)
	}
		
}
