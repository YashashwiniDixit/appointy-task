package controllers

import(
"fmt"
"encoding/json"
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"github.com/YashashwiniDixit/appointy-task.git/models"
"net/http"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

  if err := uc.session.DB("appointy-task").C("users").FindId(oid).One(&u); err != nil{
	w.WriteHeader(404)
	return
   }

   uj, err :=json.Marshal(u)
   if err!= nil{
	   fmt.Println(err)
   }

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprintf(w, "%s\n", uj)

}


func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("appointy-task").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

