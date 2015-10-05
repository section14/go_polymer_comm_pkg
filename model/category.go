package model

import (
    "net/http"
    "log"
    "appengine"
    "appengine/datastore"
)

type Category struct {
    Name string
}

type CategoryReturn struct {
    Name string
    Id int64
    Key string
}

func (cat *Category) CreateCategory(r *http.Request, parentId int64) error {
    //get context
    c := appengine.NewContext(r)

    /*

    If parentId is 0, it's the top level witout any parents. In this case,
    create an incomplete key. Otherwise, create your own key with it's
    ancestor reference key.

    */

    var k *datastore.Key

    if parentId == 0 {
        //create an incomplete key
        k = datastore.NewIncompleteKey(c, "Category", k)
    } else {
        //create Key
        k = datastore.NewKey(c, "Category", "", parentId, nil)
    }

    _, err := datastore.Put(c, k, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

/*
func getKey(r *http.Request, id int64) datastore.Key {
    //get record key for ancestor based queries

    //get context
    c := appengine.NewContext(r)

    //get key
    k := datastore.NewKey(c, "Category", "", a.UserId, nil)

    //start query
    q := datastore.NewQuery("Address").Filter("__key__ =", k)
}
*/

func (cat *Category) GetCategories(r *http.Request, pid int64) ([]CategoryReturn, error) {
    //get context
    c := appengine.NewContext(r)

    var q *datastore.Query

    //if this isn't the top level, get ancestors
    if pid != 0 {
        //make ancestor key
        k := datastore.NewKey(c, "Category", "", pid, nil)

        //start query
        q = datastore.NewQuery("Category").Ancestor(k)
        log.Println("------------------dat id", pid)
    } else {
        q = datastore.NewQuery("Category")
        log.Println("------------------------------------------------------------------------gamed last yall bitch made fools")
    }

    //populate category slices
    var categories []CategoryReturn
    keys, err := q.GetAll(c, &categories)

    if err != nil {
        //handle error
        return []CategoryReturn{}, err
    }

    //create return object
    results := make([]CategoryReturn, 0, 20)

    for i, r := range categories {
        k := keys[i]
        y := CategoryReturn {
            Name: r.Name,
            Id: k.IntID(),
            Key: k.Encode(),
        }

        results = append(results, y)
    }

    return results, nil

}
