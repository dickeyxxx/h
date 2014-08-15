package status

import "fmt"

type Status struct{}

func (s *Status) Name() string {
	return "status"
}

func (s *Status) Run(...string) int {
	fmt.Println("FOOB")
	return 0
}

func New() *Status {
	return &Status{}
}
