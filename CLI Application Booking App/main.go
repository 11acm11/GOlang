package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

//Initializing global variables
const conferencetickets int = 50

var conferencename = "Go Conference"
var remainingtickets uint = 50
var bookings = make([]userdata, 0)

//Initializing Stucture to store variables for userdata
type userdata struct {
	firstname       string
	lastname        string
	email           string
	numberoftickets uint
}

//Initializing function for delay during program runtime
var wg = sync.WaitGroup{}

func main() {

	greetusers()//Function for "Welcome" message
	for {
		firstname, lastname, email, usertickets := getuserinput()//Function to get user input
		isvalidemail, isvalidname, isvalidticnum := Validateuserinput(firstname, lastname, email, usertickets, remainingtickets)//Function to validate user input
		if isvalidname && isvalidemail && isvalidticnum {
			booktickets(usertickets, firstname, lastname, email)//Funtion to store user data and update remaining tickets
			wg.Add(1)
			go sendticket(usertickets, firstname, lastname, email)//Running concurrency
			firstnames := getfirstname()//Storing first names of users
			fmt.Printf("The firstnames of the bookings are:%v\n", firstnames)
			if remainingtickets == 0 {
				fmt.Println("No tickets left!!")
				break
			}
		} else {//Error messages if user input invalid
			if !isvalidname {
				fmt.Println("First or Last name is too short")
			}
			if !isvalidemail {
				fmt.Println("Email is invalid")
			}
			if !isvalidticnum {
				fmt.Println("Entered ticket number is invalid")
			}
		}
	}
	wg.Wait()
}

//Welcome message function
func greetusers() {
	fmt.Printf("Welcome to %v booking app\n", conferencename)
	fmt.Printf("We have %v tickets for %v with remaining %v left\n", conferencetickets, conferencename, remainingtickets)
	fmt.Println("Get your tickets")
}

//Funtion to store firstnames of users
func getfirstname() []string {	
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
	}
	return firstnames
}

//Function to get user input
func getuserinput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var usertickets uint
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstname)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastname)
	fmt.Println("Enter email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&usertickets)
	return firstname, lastname, email, usertickets
}

//Function ot book tickets and storing user data in structure with a conformation message
func booktickets(usertickets uint, firstname string, lastname string, email string) {
	remainingtickets = remainingtickets - usertickets
	var userdata = userdata{
		firstname:       firstname,
		lastname:        lastname,
		email:           email,
		numberoftickets: usertickets,
	}
	bookings = append(bookings, userdata)
	fmt.Printf("BOOKINGS:%v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a conformation email at %v.\n", firstname, lastname, usertickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingtickets, conferencename)
}

//Funtion to check user input
func Validateuserinput(firstname string, lastname string, email string, usertickets uint, remainingtickets uint) (bool, bool, bool) {
	isvalidname := len(firstname) >= 2 && len(lastname) >= 2
	isvalidemail := strings.Contains(email, "@")
	isvalidticnum := usertickets > 0 && usertickets <= remainingtickets
	return isvalidemail, isvalidname, isvalidticnum
}

//Conformational/Ticket sending message
func sendticket(usertickets uint, firstname string, lastname string, email string) {
	time.Sleep(10 * time.Second)//10 sec delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", usertickets, firstname, lastname)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
