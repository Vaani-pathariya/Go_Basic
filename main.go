package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	yes := "yes"
	//this command can only be used for var keywords
	var remaining = 50
	fmt.Printf("the types are %T", conferenceName)
	fmt.Println("welcome to our", conferenceName, " booking application")
	fmt.Println("please feel free to make a booking")
	fmt.Println(conferenceName)
	fmt.Printf("The number of remaining tickets is %v", remaining)
	var user string
	fmt.Scan(&user) //taking input
	//this is a comment
	var booking = [50]string{}
	booking[0] = "hello"
	fmt.Print(len(booking)) //prints the length of the array
	//the above thing is called an empty array
	var booking2 = [50]string{"hello"}
	//to create an array that has dynamic size ,we use slices
	var booking3 []string //defined a slice
	//how to append to a slice :
	booking3 = append(booking3, "my name is")
	// the above thing is called an array with one element
	user = "tom"
	fmt.Printf("user name : %v and the conferernce tickets are : %v", user, conferenceTickets)
}
