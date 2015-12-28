package controller

import (
    "net/http"
    "encoding/json"
    //"strconv"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Product struct {
    Title string
    Sku string
    Desc string
    Image string
    Category int64
}

type ProductReturn struct {
    Id int64
    Key string
    Title string
    Sku string
    Desc string
    Image string
    Category int64
}

func (p *Product) CreateProduct(r *http.Request) error {
    //get json request body
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&p)

    if err != nil {
        //handle err
        return err
    }

    //populate product data
    productModel := model.Product{}
    productModel.Title = p.Title
    productModel.Sku = p.Sku
    productModel.Desc = p.Desc
    productModel.Image = p.Image
    productModel.Category = p.Category //this can probably go

    err = productModel.CreateProduct(r)

    if err != nil {
        log.Println(err)
    }

    //add product to selected category .... needs updated to handle multiple categories
    categoryModel := model.Category{}
    

    return nil
}
