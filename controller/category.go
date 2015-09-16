package controller

import (
    "net/http"
    "encoding/json"
    //"log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Category struct {
    Name string `json:"name"`
}

func (cat *Category) CreateCategory(w http.ResponseWriter, r *http.Request) (bool, error) {
    decoder := json.NewDecoder(r.Body)
    err:= decoder.Decode(&cat)

    if err != nil {
        //handle err
        return false, err
    }

    //populate category data
    categoryModel := model.Category{}
    categoryModel.Name = cat.Name

    return true, nil
}
