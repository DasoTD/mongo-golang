package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DasoTD/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *userController {
	return &userController{s}
}

func (uc userController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
func (uc userController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("mongo-golang").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}
func (uc userController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}

func (uc userController) Updateuser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
}

// id := p.ByName("id")
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// 	u := models.User{}
// 	json.NewDecoder(r.Body).Decode(&u)
// 	oid := bson.ObjectIdHex(id)

// 	c := uc.session.DB("mongo-golang").C("users")
// 	colQuerier := bson.M{"_id": oid}
// 	update := bson.M{"title": u.Name, "body": u.Gender, "author": u.Age}
// 	if err := c.Update(colQuerier, update); err != nil {
// 		w.WriteHeader(404)
// 		return
// 	}

// if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
// 	w.WriteHeader(404)
// 	return
// }
// update := bson.M{"title": u.Name, "body": u.Gender, "author": u.Age}
// c := uc.session.DB("mongo-golang").C("users")
// err := c.UpdateId(id, update)
// if err != nil {
// 	panic(err)
// }
//}

// func(m * MoviesDAO) Update(movie Movie) error {
//     err: = db.C(COLLECTION).UpdateId(movie.ID, & movie)
//     return err
// }

// func UpdateMovieEndPoint(w http.ResponseWriter, r * http.Request) {
//     defer r.Body.Close()
//     var movie Movie
//     if err: = json.NewDecoder(r.Body).Decode( & movie);
//     err != nil {
//         respondWithError(w, http.StatusBadRequest, "Invalid request payload")
//         return
//     }
//     if err: = dao.Update(movie);
//     err != nil {
//         respondWithError(w, http.StatusInternalServerError, err.Error())
//         return
//     }
//     respondWithJson(w, http.StatusOK, map[string] string {
//         "result": "success"
//     })
// }

// c := session.DB("test").C("people")
// colQuerier := bson.M{"name": "Ale"}
// 	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
// 	err = c.Update(colQuerier, change)
// 	if err != nil {
// 		panic(err)
// 	}
// update := bson.M{"title": post.Title, "body": post.Body, "author": post.Author}

// c := uc.session.DB("mongo-golang").C("users")
// 	update := bson.M{"name": u.Name, "gender": u.Gender, "age": u.Age}
// 	err := c.UpdateId(bson.M{"id": oid}, bson.M{"$set": update})
// 	if err != nil {
// 		w.WriteHeader(404)
// 		return
// 	}
