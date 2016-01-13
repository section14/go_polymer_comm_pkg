package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type DataId struct {
    ID int64
    Image string
}

func (d *DataId) GetId(r *http.Request, k *datastore.Key) (int64, error) {
    //get context
    c := appengine.NewContext(r)

    err := datastore.Get(c, k, d)

    if err != nil {
        return 0, err
    }

    return d.ID, nil
}
