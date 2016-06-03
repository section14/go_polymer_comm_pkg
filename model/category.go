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

    //exmpty product array
    var products []int64 = []int64{0}

    //set parent id
    cat.ParentId = pid
    cat.Products = products

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

    //log.Println("parent id: ", pid)

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
            Products: r.Products,
            Key: k.Encode(),
        }

        results = append(results, y)
    }

    return results, nil
}

func (cat *Category) GetAllCategories(r *http.Request) ([]CategoryReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //new query
    q := datastore.NewQuery("Category")

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
            Products: r.Products,
            Key: k.Encode(),
        }

        results = append(results, y)
    }

    return results, nil
}

func (cat *Category) UpdateProductList(r *http.Request, catIds []int64, prodId int64, add bool) error {
    //method to update product list associated to a category

    /*

    Make this a projection query for the id

    */

    //get context
    c := appengine.NewContext(r)

    /*

    Something is not right, it's writing too much data to each category

    */

    //loop over each category that needs this product associated with it
    for _,id := range catIds {
        //new query
        k := datastore.NewKey(c, "Category", "", id, nil)

        //get category
        err := datastore.Get(c, k, cat)

        if err != nil {
            return err
        }

        //get product list for this category
        prodList := cat.Products
        var newProdList []int64

        //update list of products
        if add == true {
            newProdList, err = AddProduct(prodList, prodId)
        } else {
            newProdList, err = RemoveProduct(prodList, prodId)
        }

        if err != nil {
            return err
        }

        //update struct
        cat.Products = newProdList

        //insert into database
        _, err = datastore.Put(c, k, cat)

        if err != nil {
            return err
        }
    }

    return nil
}

//add a product to the list for this category
func AddProduct(products []int64, id int64) ([]int64, error) {
    for _, r := range products {
        if r == id {
            return products, nil
        }
    }
    products = append(products, id)

    return products, nil
}

//remove a product from the list for this category
//needs completed
func RemoveProduct(products []int64, id int64) ([]int64, error){
    return products, nil
}
