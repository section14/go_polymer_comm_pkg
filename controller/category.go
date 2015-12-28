package controller

import (
    "net/http"
    "encoding/json"
    "strconv"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Category struct {
    Name string `json:"name"`
    ParentId int64 `json:"parentid"`
}

type CategoryReturn struct {
    Name string
    Id int64
    ParentId int64
    Products []int64
    Key string
}

func (cat *Category) CreateCategory(r *http.Request) (bool, error) {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&cat)

    if err != nil {
        //handle err
        return false, err
    }

    //populate category data
    categoryModel := model.Category{}
    categoryModel.Name = cat.Name
    err = categoryModel.CreateCategory(r, cat.ParentId)

    if err != nil {
        log.Println(err)
    }

    return true, nil
}

func (cat *Category) GetCategories(r *http.Request) ([]CategoryReturn, error) {
    parentCat := r.Header.Get("Parent-Category")

    //convert header string to int64
    parentId, err := strconv.ParseInt(parentCat, 10, 64)

    if err != nil {
        return []CategoryReturn{}, err
    }

    categoryModel := model.Category{}
    categories, err := categoryModel.GetCategories(r, parentId)

    if err != nil {
        log.Println(err)
    }

    results := make([]CategoryReturn, 0, 20)

    for _, r := range categories {
        y := CategoryReturn {
            Name: r.Name,
            Id: r.Id,
            ParentId: r.ParentId,
            Products: r.Products,
            Key: r.Key,
        }

        results = append(results, y)
    }

    return results, nil
}
