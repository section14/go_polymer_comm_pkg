package controller

import (
    "net/http"
    "encoding/json"
    "strconv"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Product struct {
    Title string
    Sku string
    Desc string
    Image string
    Category []string
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

    pid, err := productModel.CreateProduct(r) //this needs to return an id to use below

    if err != nil {
        log.Println(err)
    }

    log.Println("categories: ", p.Category)

    //add product to selected category
    categoryModel := model.Category{}
    categoryList, err := StringsToInts(p.Category)

    if err != nil {
        log.Println(err)
    }
    
    categoryModel.UpdateProductList(r,categoryList,pid,true)

    return nil
}

func StringsToInts(s []string) ([]int64, error) {
    var ints []int64

    for _,r := range s {
        convInt, err := strconv.ParseInt(r,0,64)

        if err != nil {
            return ints, err
        }

        ints = append(ints, convInt)
    }

    return ints, nil
}

func (p *Product) GetProduct(r *http.Request, id int64) (ProductReturn, error) {

    //populate return struct
    productModel := model.Product{}
    product, err := productModel.GetProduct(r, id)

    if err != nil {
        return ProductReturn{}, err
    }

    result := ProductReturn {
        Id: product.Id,
        Key: product.Key,
        Title: product.Title,
        Sku: product.Sku,
        Desc: product.Desc,
        Image: product.Image,
    }

    return result, nil

    /*
    productModel := model.Product{}
    products, err := productModel.GetProduct(r,pid.Id)

    if err != nil {
        return []ProductReturn{}, err
    }

    results := make([]ProductReturn, 0, 20)

    for _, pr := range products {
        y := ProductReturn {
            Id: pr.Id,
            Key: pr.Key,
            Title: pr.Title,
            Sku: pr.Sku,
            Desc: pr.Desc,
            Image: pr.Image,
        }

        results = append(results, y)
    }

    return results, nil
    */
}
