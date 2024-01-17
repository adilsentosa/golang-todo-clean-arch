package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GreetingServiceTestSuite struct {
	suite.Suite
	service     GreetingService
	mockService *GreetingServiceMock
}

func (s *GreetingServiceTestSuite) SetupTest() {
	s.mockService = new(GreetingServiceMock)
}

func (s *GreetingServiceTestSuite) TestGreeting_Success() {
	person := Person{"Beni"}
	s.mockService.On("Greeting", person).Return(person, nil)
	actual, err := s.service.Greeting(person)
	s.NoError(err)
	s.Equal(person, actual)
}

func (s *GreetingServiceTestSuite) TestGreeting_Failure() {
	person := Person{}
	s.mockService.On("Greeting", person).Return(person, fmt.Errorf("error"))
	actual, err := s.service.Greeting(person)
	s.NotNil(err)
	s.Error(err)
	s.Equal(person, actual)
}

func TestGreetingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GreetingServiceTestSuite))
}
