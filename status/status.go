package status

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dickeyxxx/h/cli"
)

var url = "https://status.heroku.com/api/v3/current-status.json"

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

var Topic = &cli.Topic{
	Name: "status",
	Run:  Run,
}

func Run(ctx *cli.Context) int {
	if len(ctx.Args) != 0 {
		ctx.ErrPrintln("USAGE")
		return 1
	}
	var response statusResponse
	getStatus(&response)
	ctx.Println("=== Heroku Status")
	ctx.Println("Development: ", statusText(response.Status.Development))
	ctx.Println("Production:  ", statusText(response.Status.Production))
	for _, issue := range response.Issues {
		ctx.Println("=== ", issue.Title, issue.CreatedAt)
		ctx.Println(issue)
	}
	return 0
}

var getStatus = func(response *statusResponse) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error connecting to status server. HTTP Code %d\n", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalln("Error parsing json", err)
	}
}

func statusText(status string) string {
	if status == "green" {
		return "No known issues at this time."
	}
	return status
}
