package controller

import (
    "net/http"
    "encoding/json"
    "log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Address struct {
    UserId int64
    Street1 string `json:"street1"`
    Street2 string `json:"street2"`
    City string `json:"city"`
    State string `json:"state"`
    PostCode string `json:"postcode"`
    Country string `json:"country"`
}

func (a *Address) CreateAddress(r *http.Request) (bool, error) {
    addressModel := model.Address{}

    //grab json data from request
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&a)

    log.Println(a)

    //populate model
    addressModel.Street1 = a.Street1
    addressModel.Street2 = a.Street2
    addressModel.City = a.City
    addressModel.State = a.State
    addressModel.PostCode = a.PostCode
    addressModel.Country = a.Country

    err = addressModel.CreateAddress(r)

    if err != nil {
        return false, err
    }

    return true, nil
}
