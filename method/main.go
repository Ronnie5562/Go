package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// Disclaimer: Golang does not have inheritance; No super or parent.
	// It uses a concept called composition instead

	Ronald := User{"Ronald", "abimbola@google.dev", true, 17}
	fmt.Println(Ronald)
	fmt.Printf("%+v", Ronald)
	fmt.Printf("My name is %v, I am %v years old. You can reach out to me on %v\n", Ronald.Name, Ronald.Age, Ronald.Email)

	Ronald.GetStatus()
	Ronald.SetEmail("r.abimbola@neuralink.eng")

	// When you print this, you'll realise that the email did not actually change and this is why you need a pointer.
	fmt.Println(Ronald)
}

type User struct {
	Name     string
	Email    string
	verified bool
	Age      int
}

// Methods for the struct User

func (u User) GetStatus() {
	fmt.Println("The user verificatation status is:", u.verified)
}

func (u User) SetEmail(email string) {
	u.Email = email
	fmt.Println("The user's new email is:", u.Email)
}
