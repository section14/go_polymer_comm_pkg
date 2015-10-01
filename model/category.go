package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type Category struct {
    Name string
}

type CategoryReturn struct {
    Name string
    ParentId *datastore.Key
}

func (cat *Category) CreateCategory(r *http.Request, p *datastore.Key) error {
    //get context
    c := appengine.NewContext(r)

    //create new category entry
    key := datastore.NewIncompleteKey(c, "Category", p)

    _, err := datastore.Put(c, key, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (cat *Category) GetCategories(r *http.Request, p *datastore.Key) (CategoryReturn, error) {
    
}
