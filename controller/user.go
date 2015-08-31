package controller

import (
    "net/http"
    "encoding/json"
    "log"

    "golang.org/x/crypto/bcrypt"
    "github.com/section14/go_polymer_comm_pkg/model"
)

type User struct {
    Email string
    Password string
    Role int
}

type Return struct {
    Email string
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) (bool, error) {
    type Message struct {
    Email string `json:"email"`
    Password string `json:"password"`
    }

    var m Message
    decoder := json.NewDecoder(r.Body)
    err:= decoder.Decode(&m)

    if err != nil {
        //handle err
    }

    //populate user data
    userModel := model.User{}
    userModel.Email = m.Email

    //encrypt password
    password, err := bcrypt.GenerateFromPassword([]byte(m.Password), 10)

    if err != nil {
        panic(err)
    }

    userModel.Password = string(password)
    userModel.Role = 1

    //make sure email doesn't exist
    emailStatus, err := u.CheckEmail(w,r)

    if emailStatus == true || err != nil {
        return false, nil
    }

    //create new user
    err = userModel.CreateUser(w,r)

    if err != nil {
        return false, err
    }

    return true, nil
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) (bool, error) {
    /*
    *
    You need to return the id and role here, so it can go into the jwt
    *
    */

    type Message struct {
    Email string `json:"email"`
    Password string `json:"password"`
    }

    var m Message
    decoder := json.NewDecoder(r.Body)
    err:= decoder.Decode(&m)

    if err != nil {
        //handle err
        log.Println(err)
    }

    userModel := model.User{}
    userModel.Email = m.Email

    //find user
    user, err := userModel.GetLoginData(w,r)

    if err != nil {
        //handle err
        log.Println("no user")
        return false, err
    }

    //match passwords
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))

    if err != nil {
        //passwords don't match
        log.Println("no passwords")
        return false, err
    }

    //everything checks out!
    log.Println("everything matches")
    return true, nil
}

func (u *User) CheckEmail(w http.ResponseWriter, r *http.Request) (bool, error) {
    type Message struct {
    Email string `json:"email"`
    }

    var m Message
    decoder := json.NewDecoder(r.Body)
    err:= decoder.Decode(&m)

    if err != nil {
        //handle err
    }

    userModel := model.User{}

    emailStatus, err:= userModel.CheckEmail(w,r,m.Email)

    if err != nil {
        return false, err
    }

    return emailStatus, nil
}

func (u *User) TestHit() {

}
