package main

import "fmt"

//Insert variables declarations here based on activity

var (
	text             string  = "The following is the account information."
	firstname        string  = "Luke"
	lastname         string  = "Skywalker"
	age              int     = 20
	weight           float64 = 73.0
	height           float64 = 1.72
	remainingcredits float64 = 123.55
	accountname      string  = "admin"
	accountpassword  string  = "password"
	subscribeduser   string  = "true"
)

func tellMeTypes() {
	//insert code here to print out types of variables
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T", text, firstname, lastname, age, weight, height, remainingcredits, accountname, accountpassword, subscribeduser)
}

func main() {
	tellMeTypes()

}
