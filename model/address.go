package model

import (
    "net/http"
    "log"
    "appengine"
    "appengine/datastore"
)

type Address struct {
    UserId int64
    Street1 string
    Street2 string
    City string
    State string
    PostCode string
    Country string
}

type AddressReturn struct {
    Key *datastore.Key
    Id int64
    Street1 string
    Street2 string
    City string
    State string
    PostCode string
    Country string
}

func (a *Address) CreateAddress(r *http.Request) (error) {
    //get context
    c := appengine.NewContext(r)

    //create new address
    key := datastore.NewIncompleteKey(c, "Address", nil)

    _, err := datastore.Put(c, key, a)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (a *Address) GetAllAddress(r *http.Request) ([]AddressReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //get key
    //k := datastore.NewKey(c, "Address", "", a.UserId, nil)

    //start query
    q := datastore.NewQuery("Address").Filter("UserId=", a.UserId)

    log.Println("model user id: ", a.UserId)

    //populate address slices and get keys
    var addresses []AddressReturn
    keys, err := q.GetAll(c, &addresses)

    if err != nil {
        //handle error
        return []AddressReturn{}, err
    }

    //return address data slices
    results := make([]AddressReturn, 0, 10)

    for i, r := range addresses {
        k := keys[i]
        y := AddressReturn {
            Key: k,
            Id: k.IntID(),
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
    //get context
    c := appengine.NewContext(r)

    //get key
    k := datastore.NewKey(c, "Address", "", a.UserId, nil)

    //start query
    q := datastore.NewQuery("Address").Filter("__key__ =", k)

    //populate address slices and get keys
    var addresses []AddressReturn
    keys, err := q.GetAll(c, &addresses)

    if err != nil {
        //handle error
        return AddressReturn{}, err
    }

    //return address data
    var result AddressReturn

    for i, r := range addresses {
        k := keys[i]
        result = AddressReturn {
            Key: k,
            Id: k.IntID(),
            Street1: r.Street1,
            Street2: r.Street2,
            City: r.City,
            State: r.State,
            PostCode: r.PostCode,
            Country: r.Country,
        }
    }

    return result, nil
}
