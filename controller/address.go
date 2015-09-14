package controller

import (
    "net/http"
    "encoding/json"
    //"log"

    "github.com/section14/go_polymer_comm_pkg/model"
)

type Address struct {
    Id int64 `json:"id"`
    UserId int64 `json:"userid"`
    Street1 string `json:"street1"`
    Street2 string `json:"street2"`
    City string `json:"city"`
    State string `json:"state"`
    PostCode string `json:"postcode"`
    Country string `json:"country"`
}

type AddressReturn struct {
    Id int64 `json:"id"`
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

    //populate model
    addressModel.UserId = a.UserId
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

func (a *Address) UpdateAddress(r *http.Request) (bool, error) {
    addressModel := model.Address{}

    //grab json data from request
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&a)

    //populate model
    addressModel.UserId = a.UserId
    addressModel.Street1 = a.Street1
    addressModel.Street2 = a.Street2
    addressModel.City = a.City
    addressModel.State = a.State
    addressModel.PostCode = a.PostCode
    addressModel.Country = a.Country

    err = addressModel.UpdateAddress(r, a.Id)

    if err != nil {
        return false, err
    }

    return true, nil
}

func (a *Address) GetAllAddress(r *http.Request) ([]AddressReturn, error) {
    addressModel := model.Address{}
    addressModel.UserId = a.UserId
    addressData, err := addressModel.GetAllAddress(r)

    if err != nil {
        return []AddressReturn{}, err
    }

    //return address data slices
    results := make([]AddressReturn, 0, 10)

    for _, r := range addressData {
        y := AddressReturn {
            Id: r.Id,
            Street1: r.Street1,
            Street2: r.Street2,
            City: r.City,
            State: r.State,
            PostCode: r.PostCode,
            Country: r.Country,
        }

        results = append(results, y)
    }

    return results, nil
}

func (a *Address) GetAddress(r *http.Request) (AddressReturn, error) {
    addressModel := model.Address{}

    addressData,err := addressModel.GetAddress(r)

    if err != nil {
        return AddressReturn{}, err
    }

    address := AddressReturn {
        Id: addressData.Id,
        Street1: addressData.Street1,
        Street2: addressData.Street2,
        City: addressData.City,
        State: addressData.State,
        PostCode: addressData.PostCode,
        Country: addressData.Country,
    }

    return address, nil
}
