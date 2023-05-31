package main

import (
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"io/ioutil"
    "github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	reqBody, err1 := ioutil.ReadAll(c.Request.Body)
	if err1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "error"})
		return
	}
	var newUserData userData 
	err := json.Unmarshal(reqBody, &newUserData)
	if err != nil {
		fmt.Println(err)
	}
	if !existDatabase(newUserData.Email)  {
		doNotExist := errors.New("Email does not exist. Register")
		fmt.Println(doNotExist)
		return
	}
	if getPassword(newUserData.Email) != newUserData.Password {
		passwordMatch := errors.New("Password do not match")
		fmt.Println(passwordMatch)
		return 
	} else {
		otpLogin := EncodeToString(6)
		// Update the OTP in Database
		// check OTP 
		// if OTP is correct  -- 
		// Login Successful
	}
	
}

func existDatabase () {

}
func getPassword () {

}

// func main () {
// 	router := gin.Default()
// 	router.POST("/user/login",login)
// 	router.Run("localhost:8080")
// }