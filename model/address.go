package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type Address struct {
    UserId int64
    Street1 string
    Street2 string
    City string
    State string
    PostCode string
    Country string
}

func (a *Address) CreateAddress(r *http.Request) (error) {
    //get context
    c := appengine.NewContext(r)

    //create new address
    key := datastore.NewIncompleteKey(c, "Address", nil)

    _, err := datastore.Put(c, key, a)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}
