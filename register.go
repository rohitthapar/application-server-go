package main

import (
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"regexp"
	"io/ioutil"
    "github.com/gin-gonic/gin"
	"crypto/rand"
    "io"
)

type userData struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
type userOTP struct {
	OTP int`json:"otp"`
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
    b := make([]byte, max)
    n, err := io.ReadAtLeast(rand.Reader, b, max)
    if n != max {
        panic(err)
    }
    for i := 0; i < len(b); i++ {
        b[i] = table[int(b[i])%len(table)]
    }
    return string(b)
}


func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}

func existDatabase () {


}

func passwordValid(e string) bool {
	passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return passwordRegex.MatchString(e)
} 

func pushToSQL () {
	
}

func getEmail(c *gin.Context) {
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
	if isEmailValid(newUserData.Email) == false {
		emailError := errors.New("Enter Valid Email")
		fmt.Println(emailError)
		return	
	}
	if existDatabase(newUserData.Email) == true {
		dataExistError := errors.New("Email Already Exists")
		fmt.Println(dataExistError)	
		return 
	}
	if passwordValid(newUserData.Password) == false {
		passwordError := errors.New("Enter Correct Password")
		fmt.Println(passwordError)	
		return
	} else {
		otp := EncodeToString(6)
		pushToSQL()
	}
}

func verifyEmail (c *gin.Context) {
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

	reqBody2, err2 := ioutil.ReadAll(c.Request.Body)
	if err2 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "error"})
		return
	}
	var toCheckOTP userOTP 
	err3 := json.Unmarshal(reqBody2, &toCheckOTP)
	if err3 != nil {
		fmt.Println(err3)
	}
	uOTP := getOTP(newUserData.Email)
	if uOTP != toCheckOTP.OTP {
		wrongOTP := errors.New("Enter the correct OTP")
		fmt.Println(wrongOTP)
		return
	} else {
		// set flag of sql table to 1 
	}
}


func main() {        
	router := gin.Default()
	router.POST("/user/signup",getEmail)
	router.POST("/user/login",login)
	router.POST("user/signup/verifyEmail",verifyEmail)
	router.Run("localhost:8080")
}



// if isEmailValid(newUserData.Email) == true {
	// 	if existDatabase(newUserData.Email) == false {
	// 		if passwordValid(newUserData.Password) == true {
	// 			pushToSQL(newUserData)
	// 		} else {
	// 			passwordError := errors.New("Enter Correct Password")
	// 			fmt.Println(passwordError)	
	// 		}
	// 	} else {
	// 		dataExistError := errors.New("Email Already Exists")
	// 		fmt.Println(dataExistError)
	// 	} 
	// } else {
	// 	emailError := errors.New("Enter Valid Email")
	// 	fmt.Println(emailError)
	// }