package main

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	type test struct {
		Msg string `json:"msg"`
	}
	var a = test{
		Msg: "test",
	}
	fmt.Println(a)
	fmt.Println(c)
	c.JSON(http.StatusOK, a)
}

func password(c *gin.Context) {
	type requestStruct struct { //เก็บค่า Post
		Init_password string `json:"init_password"`
	}
	type responseStruct struct { //เก็บจำนวน steps
		Num_of_steps int `json:"num_of_steps"`
	}
	var request requestStruct   //  requestStruct -> request
	var response responseStruct // responseStruct -> response
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// response.Num_of_steps = count
	arString := []rune(request.Init_password)
	countPassword := len(arString) // นับจำนวน password ที่รับเข้ามา
	stringContainUppercase := false
	stringContainLowercase := false
	stringContainIsnumber := false
	stringContainRepeating := false
	response.Num_of_steps = 3

	if countPassword >= 6 && countPassword < 20 { //เช็ค password ที่รับเข้ามา ผ่านเงื่อนไขที่ 1 (-1)
		response.Num_of_steps -= 1
		fmt.Println("pass correct")
	} else {
		c.JSON(http.StatusOK, response)
		return
	}
	for index, char := range arString { //เช็ค character ทีละตัว
		// if char.IsUpper(())
		// fmt.Printf("%c ", char)
		isUpperCase := unicode.IsUpper(char)
		isLowerCase := unicode.IsLower(char)
		isDigit := unicode.IsDigit(char)
		if stringContainUppercase == false && isUpperCase == true { // เช็คตัวพิมพ์ใหญ่
			stringContainUppercase = true
		}
		if stringContainLowercase == false && isLowerCase == true { // เช็คตัวพิมพ์เล็ก
			stringContainLowercase = true
		}
		if stringContainIsnumber == false && isDigit == true { // เช็คตัวเลข
			stringContainIsnumber = true
		}
		if index+2 <= len(arString)-1 && char == arString[index+1] && char == arString[index+2] { // เช็คตัวอักษรซ้ำติดกัน 3 ตัว
			stringContainRepeating = true
		}

	}
	if stringContainUppercase == true && stringContainLowercase == true && stringContainIsnumber == true { // ผ่านเงื่อนไขที่ 2 (-1)
		response.Num_of_steps -= 1
		fmt.Println("upper and lower and number correct")
	} else {
		c.JSON(http.StatusOK, response)
		return
	}
	if stringContainRepeating == false { // ผ่านเงื่อนไขที่ 3 (-1)
		response.Num_of_steps -= 1
		fmt.Println("repeating correct")
	} else {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/strong_password_steps", password)
	router.Run("localhost:8080")

}
