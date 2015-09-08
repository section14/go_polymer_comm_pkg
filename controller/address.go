package controller

import (
    "net/http"
    //"encoding/json"
    //"log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Address struct {
    Line1 string
    Line2 string
    City string
    State string
    PostCode string
    Country string
}

func (a *Address) CreateAddress(r *http.Request) (bool, error) {
    addressModel := model.Address{}

    /*
    addressModel.Line1 = a.Line1
    addressModel.Line2 = a.Line2
    addressModel.City = a.City
    addressModel.State = a.State
    addressModel.PostCode = a.PostCode
    addressModel.Country = a.Country
    */

    err := addressModel.CreateAddress(r)

    if err != nil {
        return false, err
    }

    return true, nil
}
