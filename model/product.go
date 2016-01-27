package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type Product struct {
    Title string
    Sku string
    Desc string
    Image string
}

type ProductReturn struct {
    Id int64
    Key string
    Title string
    Sku string
    Desc string
    Image string
}

func (p *Product) CreateProduct(r *http.Request) (int64, error) {
    //get context
    c := appengine.NewContext(r)

    k := datastore.NewIncompleteKey(c, "Product", nil)

    //enter record
    newKey, err := datastore.Put(c, k, p)

    if err != nil {
        //hanle input error
        return 0, err
    }

    //return new id
    return newKey.IntID(), nil
}

func (p *Product) GetProduct(r *http.Request, id int64) (ProductReturn, error) {
    //return struct
    var pr ProductReturn

    //get context
    c := appengine.NewContext(r)

    //new query
    k := datastore.NewKey(c, "Product", "", id, nil)

    //get single product
    err := datastore.Get(c, k, pr)

    if err != nil {
        return ProductReturn{}, err
    }

    return pr, nil
}

/*
func (p *Product) GetProducts(r *http.Request, id int64) ([]ProductReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //id is a category id which holds an array of product id's


}
*/
