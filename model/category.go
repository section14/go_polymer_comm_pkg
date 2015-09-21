package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type Category struct {
    Name string
    ParentId int64
}

func (cat *Category) CreateCategory(w http.ResponseWriter, r *http.Request) error {
    //get context
    c := appengine.NewContext(r)

    //create new category entry
    key := datastore.NewIncompleteKey(c, "Category", nil)

    _, err := datastore.Put(c, key, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}
