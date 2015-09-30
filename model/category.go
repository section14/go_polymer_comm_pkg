package model

import (
    "net/http"
    "log"
    "appengine"
    "appengine/datastore"
)

type Category struct {
    Name string
    ParentId int64
}

/*

Figure out what you want to do with parent Id's. Use datastore's "Ancestor" method, or manage it yourself.
Managing it yourself seems easier, but the ancestor method will allow for easier querying of nested blocks.
For example, Fender Guitars would pull all the associated down the line. Just go with fucking datastore's method. sheesh.

*/

func (cat *Category) CreateCategory(r *http.Request) error {
    //get context
    c := appengine.NewContext(r)

    log.Println("fuck name: ", cat.Name)

    //create new category entry
    key := datastore.NewIncompleteKey(c, "Category", nil)

    _, err := datastore.Put(c, key, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}
