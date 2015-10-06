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

func (cat *Category) CreateCategory(r *http.Request, parentKey string) error {
    //get context
    c := appengine.NewContext(r)

    /*

    If parentId is 0, it's the top level witout any parents. In this case,
    create an incomplete key. Otherwise, create your own key with it's
    ancestor reference key.

    */

    var k *datastore.Key
    var err error

    if parentKey == "0" {
        //create an incomplete key
        k = datastore.NewIncompleteKey(c, "Category", k)
        log.Println("orphan child")
    } else {
        //create Key
        k, err = datastore.DecodeKey(parentKey)
        log.Println("you have a mommy: ", parentKey)
        if err != nil {
            return err
        }
    }

    _, err = datastore.Put(c, k, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (cat *Category) GetCategories(r *http.Request, pk string) ([]CategoryReturn, error) {
    //get context
    c := appengine.NewContext(r)

    var q *datastore.Query
    var err error

    //if this isn't the top level, get ancestors
    if pk != "0" {
        //get parent key
        k, err := datastore.DecodeKey(pk)

        if err != nil {
            //handle error
            return []CategoryReturn{}, err
        }
        q = datastore.NewQuery("Category").Ancestor(k)
    } else {
        //query without ancestor
        q = datastore.NewQuery("Category")
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
