package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Organization struct {
	name    string
	country string
}

type SecretAgent struct {
	profile       Person
	licenseToKill bool
	organization  Organization
}

func main() {
	mossad := Organization{
		name:    "Mossad",
		country: "Israel",
	}

	person_1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       100,
	}

	secret_agent_1 := SecretAgent{
		profile: Person{
			firstName: "Tamar",
			lastName:  "Rabinya",
			age:       27,
		},
		licenseToKill: true,
		organization:  mossad,
	}

	fmt.Printf("%T, %#v, %v\n", person_1, person_1, person_1)
	fmt.Printf("%T, %#v, %v\n", secret_agent_1, secret_agent_1, secret_agent_1)

	fmt.Println(person_1.firstName, person_1.age)
	fmt.Println(secret_agent_1.profile.firstName, secret_agent_1.organization.name)
}
