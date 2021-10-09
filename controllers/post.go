package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YashashwiniDixit/appointy-task.git/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostsController struct {
	session *mgo.Session
}

func NewPostsController(s *mgo.Session) *PostsController {
	return &PostsController{s}
}

func (uc PostsController) GetPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Posts{}

	if err := uc.session.DB("appointy-task").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}

func (uc PostsController) CreatePosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.Posts{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()
	u.PostedTimeStamp = time.Now()
	uc.session.DB("appointy-task").C("posts").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc PostsController) GetUserPosts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("uid")
	u := make([]models.Posts, 0, 10)

	if err := uc.session.DB("appointy-task").C("posts").Find(bson.M{"userid": id}).All(&u); err != nil {
		w.WriteHeader(500)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)

}
