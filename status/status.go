package status

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type statusResponse struct {
	Status struct {
		Production  string
		Development string
	} `json:"status"`
	Issues []statusIssue `json:"issues"`
}

type statusIssue struct {
	Resolved   bool   `json:"resolved"`
	StatusDev  string `json:"status_dev"`
	StatusProd string `json:"status_prod"`
	Title      string `json:"title"`
	Upcoming   bool   `json:"upcoming"`
	Href       string `json:"href"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Status struct{}

func (s *Status) Name() string {
	return "status"
}

func (s *Status) Run(args ...string) int {
	if len(args) != 0 {
		log.Fatalln("USAGE")
	}
	resp, err := http.Get("https://status.heroku.com/api/v3/current-status.json")
	if err != nil {
		log.Fatalf("Error connecting to status server. HTTP Code %d\n", resp.StatusCode)
	}
	var sr statusResponse
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		log.Fatalln("Error parsing json", err)
	}

	fmt.Println("=== Heroku Status")
	fmt.Println("Development: ", statusText(sr.Status.Development))
	fmt.Println("Production:  ", statusText(sr.Status.Production))
	return 0
}

func statusText(status string) string {
	if status == "green" {
		return "No known issues at this time."
	}
	return status
}

func New() *Status {
	return &Status{}
}
