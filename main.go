package main

import (
	"fmt"
	"net/http"

	"github.com/DasoTD/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	r.PUT("/user:id", uc.Updateuser)
	http.ListenAndServe("localhost:8080", r)
}
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost/mongo-golang")
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to DB")
	return s
}
