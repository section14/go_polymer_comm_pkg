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
    Products []int64
}

type CategoryReturn struct {
    Name string
    Id int64
    ParentId int64
    Products []int64
    Key string
}

func (cat *Category) CreateCategory(r *http.Request, pid int64) error {
    //get context
    c := appengine.NewContext(r)

    //set parent id
    cat.ParentId = pid

    k := datastore.NewIncompleteKey(c, "Category", nil)

    //enter record
    _, err := datastore.Put(c, k, cat)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (cat *Category) GetCategories(r *http.Request, pid int64) ([]CategoryReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //new query
    q := datastore.NewQuery("Category").Filter("ParentId=", pid)

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
            ParentId: r.ParentId,
            Products: r.Products
            Key: k.Encode(),
        }

        results = append(results, y)
    }

    return results, nil

}

func (cat *Category) UpdateProductList(r *http.Request, prodId int64, catId int64) error {
    //method to update product list associated to a category

    //get context
    c := appengine.NewContext(r)

    //new query
    q := datastore.NewQuery("Category").Filter("Id=", id)

    var categories []CategoryReturn
    keys, err := q.GetAll(c, &categories)

    if err != nil {
        return err
    }
}
