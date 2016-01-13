package model

import (
    "net/http"
    "log"
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
    //return the key up here^^^^^, not the int64. Then call the GetId() method separately from the controller

    //get context
    c := appengine.NewContext(r)

    k := datastore.NewIncompleteKey(c, "Product", nil)

    //enter record
    newKey, err := datastore.Put(c, k, p)

    if err != nil {
        //hanle input error
        return 0, err
    }

    //get id of newly created record
    productId, err := GetId(r,newKey)

    log.Println("stupid id: ", productId)

    if err != nil {
        return 0, err
    }

    return productId, nil
}

func GetId(r *http.Request, k *datastore.Key) (int64, error) {
    p := ProductReturn{}

    //get context
    c := appengine.NewContext(r)

    err := datastore.Get(c, k, p)

    if err != nil {
        log.Println(err)
        return 0, err
    }

    return p.Id, nil
}

/*
func (p *Product) GetProducts(r *http.Request) ([]ProductReturn, error) {
    //get context
    c := appengine.NewContext(r)


}
*/
