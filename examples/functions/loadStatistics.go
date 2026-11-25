//go:build ignore

// Load Statistics
// Loads account statistics, first from a start date only, then for a date range.
//
// Required access level: ViewReports
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	ElasticEmail "github.com/elasticemail/elasticemail-go/v4"
)

func main() {
	apiKey := os.Getenv("ELASTICEMAIL_API_KEY")
	if apiKey == "" {
		log.Fatal("ELASTICEMAIL_API_KEY is not set. Export your Elastic Email API key first.")
	}

	client := ElasticEmail.NewAPIClient(ElasticEmail.NewConfiguration())
	ctx := context.WithValue(context.Background(), ElasticEmail.ContextAPIKeys,
		map[string]ElasticEmail.APIKey{
			"apikey": {Key: apiKey},
		})

	from := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2022, 1, 30, 0, 0, 0, 0, time.UTC)

	// From a start date until now.
	result, httpRes, err := client.StatisticsAPI.StatisticsGet(ctx).
		From(from).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("From %s\n%s\n", from.Format("2006-01-02"), out)

	// Between two dates.
	result, httpRes, err = client.StatisticsAPI.StatisticsGet(ctx).
		From(from).
		To(to).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	out, _ = json.MarshalIndent(result, "", "  ")
	fmt.Printf("From %s to %s\n%s\n", from.Format("2006-01-02"), to.Format("2006-01-02"), out)
}

// fail prints the error returned by Elastic Email (or the network error) and exits.
func fail(err error, httpRes *http.Response) {
	var apiErr *ElasticEmail.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpRes != nil {
		log.Fatalf("Elastic Email API error %d: %s", httpRes.StatusCode, apiErr.Body())
	}
	log.Fatal(err)
}
