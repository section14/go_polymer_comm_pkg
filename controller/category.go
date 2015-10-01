package controller

import (
    "net/http"
    "encoding/json"
    "appengine/datastore"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Category struct {
    Name string `json:"name"`
    ParentId *datastore.Key `json:"parentid"`
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
