package testing

import "fmt"

type Person struct {
	Name string
}

type GreetingService interface {
	Greeting(person Person) (Person, error)
}

type greetingService struct{}

func (g *greetingService) Greeting(person Person) (Person, error) {
	if person.Name == "" {
		return Person{}, fmt.Errorf("name cannot be empty")
	}

	return Person{Name: person.Name}, nil
}

func NewGreetingService() GreetingService {
	return &greetingService{}
}
