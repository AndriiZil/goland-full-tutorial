package main

import "fmt"

func main() {
	var conferenceName = "Go Conference" // Changable
	const conferenceTickets = 50         // Constant
	var remainingTickets = 50

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
