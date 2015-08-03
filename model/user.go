package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type User struct {
    Email string `json:"email"`
    Password string `json:"password"`
    Role int `json:"role"`
}

func (u *User) AddUser(w http.ResponseWriter, r *http.Request) (err error) {
    //get context
    c := appengine.NewContext(r)

    //create new user entry
    key := datastore.NewIncompleteKey(c, "User", nil)

    _, err = datastore.Put(c, key, u)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) (users []User, err error) {
    //get context
    c := appengine.NewContext(r)

    //start query
    q := datastore.NewQuery("User")

    //populate user slices
    _, err = q.GetAll(c, &users)

    if err != nil {
        //handle error
        return nil, err
    }

    return users, nil
}
