package controller

import (
    "net/http"
    //"log"
    "github.com/section14/go_polymer_comm_pkg/model"
)

type User struct {
    Email string
    Role int
}

func (u *User) CheckEmail(w http.ResponseWriter, r *http.Request) (bool, error) {
    userModel := model.User{}

    emailStatus, err:= userModel.CheckEmail(w,r,u.Email)

    if err != nil {
        return false, err
    }

    return emailStatus, nil
}

func (u *User) TestHit() {

}
