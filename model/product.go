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

func (p *Product) CreateProduct(r *http.Request) error {
    //get context
    c := appengine.NewContext(r)

    k := datastore.NewIncompleteKey(c, "Product", nil)

    //enter record
    _, err := datastore.Put(c, k, p)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

/*
func (p *Product) GetProducts(r *http.Request) ([]ProductReturn, error) {
    //get context
    c := appengine.NewContext(r)


}
*/
