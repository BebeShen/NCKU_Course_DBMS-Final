package utils

import "errors"

var ErrorUserNotExist = errors.New("Not Exist")
func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}