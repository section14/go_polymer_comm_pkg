package model

import (
    "net/http"
    //"log"
    "appengine"
    "appengine/datastore"
)

type User struct {
    Email string
    Name string
    Phone string
    Password string
    Role int
}

type UserReturn struct {
    Key *datastore.Key
    Id int64
    Email string
    Name string
    Phone string
    Role int
}

type LoginReturn struct {
    Id int64
    Email string
    Password string
    Role int
}

func (u *User) CreateUser(r *http.Request) (error) {
    //get context
    c := appengine.NewContext(r)

    //create new user entry
    key := datastore.NewIncompleteKey(c, "User", nil)

    _, err := datastore.Put(c, key, u)

    if err != nil {
        //hanle input error
        return err
    }

    return nil
}

func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) ([]UserReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //start query
    q := datastore.NewQuery("User")

    //populate user slices and get keys
    var users []User
    keys, err := q.GetAll(c, &users)

    if err != nil {
        //handle error
        return nil, err
    }

    //return array of user data
    results := make([]UserReturn, 0, 10)

    for i, r := range users {
        k := keys[i]
        y := UserReturn {
            Key: k,
            Id: k.IntID(),
            Email: r.Email,
        }

        results = append(results,y)
    }

    return results, nil
}

func (u *User) GetUser(r *http.Request, uid int64) (UserReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //get key
    k := datastore.NewKey(c, "User", "", uid, nil)

    //start query
    q := datastore.NewQuery("User").Filter("__key__ =", k)

    //populate user slices and get keys
    var users []User
    keys, err := q.GetAll(c, &users)

    if err != nil {
        //handle error
        return UserReturn{}, err
    }

    //return user data
    var result UserReturn

    for i, r := range users {
        k := keys[i]
        result = UserReturn {
            Key: k,
            Id: k.IntID(),
            Email: r.Email,
            Role: r.Role,
        }
    }

    return result, nil
}

func (u *User) GetLoginData(w http.ResponseWriter, r *http.Request) (LoginReturn, error) {
    //get context
    c := appengine.NewContext(r)

    //create query
    q := datastore.NewQuery("User").Filter("Email=", u.Email).Limit(1)

    //populate email slices
    var users []User
    keys, err := q.GetAll(c, &users)

    if err != nil {
        return LoginReturn{}, err
    }

    var l LoginReturn

    for i, r := range users {
        k := keys[i]
        l = LoginReturn {
            Id: k.IntID(),
            Email: r.Email,
            Password: r.Password,
            Role: r.Role,
        }
    }

    return l, nil
}

func (u *User) CheckEmail(w http.ResponseWriter, r *http.Request, email string) (bool, error) {
    //get context
    c := appengine.NewContext(r)

    //create query
    q := datastore.NewQuery("User").Filter("Email=", email).Limit(3)

    //populate user slices
    var emails []User
    _, err := q.GetAll(c, &emails)

    if err != nil {
        //handle error
        return  false, err
    }

    if emails != nil {
        return true, nil
    }

    return false, nil
}

func (u *User) UpdateEmail(w http.ResponseWriter, r *http.Request, uid int64, newEmail string) (error) {
    //get context
    c := appengine.NewContext(r)

    //set key
    k := datastore.NewKey(c, "User", "", uid, nil)

    //start query
    q := datastore.NewQuery("User").Filter("__key__ =", k)

    //populate user slices
    var users []User
    key, err := q.GetAll(c, &users)

    if err != nil {
        //handle error
        return  err
    }

    for i, r := range users {
        r.Email = newEmail

        //write to db
        _, err := datastore.Put(c, key[i], &r)

        if err != nil {
            //handle error
            return err
        }
    }

    return nil
}

/*

Update password needs to be in here

*/

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request, uid int64) (error) {
    //get context
    c := appengine.NewContext(r)

    //set key
    k := datastore.NewKey(c, "User", "", uid, nil)

    //start query
    q := datastore.NewQuery("User").Filter("__key__ =", k)

    //populate user slices
    var users []User
    key, err := q.GetAll(c, &users)

    if err != nil {
        //handle error
        return err
    }

    for i, _ := range users {

        //delete from db
        err := datastore.Delete(c, key[i])

        if err != nil {
            //handle error
            return err
        }
    }

    return nil
}
