package main

import (
	"net/http"

	"github.com/YashashwiniDixit/appointy-task.git/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	pc := controllers.NewPostsController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.GET("/posts/users/:uid", pc.GetUserPosts)
	r.POST("/posts", pc.CreatePosts)
	r.GET("/post/:id", pc.GetPosts)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
