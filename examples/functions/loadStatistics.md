# Load Statistics

This guide walks you through loading your account statistics using the Elastic Email Go library.

*Required Access Level: ViewReports*

## What are statistics?

Elastic Email keeps delivery statistics for everything you send: emails sent, delivered, opened, clicked, bounced, unsubscribed and more.

## Preparation

Install Go 1.18 or later.

Create a project and add the Elastic Email Go library:

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Create a file named `loadStatistics.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run loadStatistics.go` still runs it.

Read the API key from the environment and create the client. The key travels in the request context, so every call takes `ctx`:

```go
apiKey := os.Getenv("ELASTICEMAIL_API_KEY")
if apiKey == "" {
	log.Fatal("ELASTICEMAIL_API_KEY is not set. Export your Elastic Email API key first.")
}

client := ElasticEmail.NewAPIClient(ElasticEmail.NewConfiguration())
ctx := context.WithValue(context.Background(), ElasticEmail.ContextAPIKeys,
	map[string]ElasticEmail.APIKey{
		"apikey": {Key: apiKey},
	})
```

Pick the dates. `From` is required, `To` is optional.

```go
from := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
to := time.Date(2022, 1, 30, 0, 0, 0, 0, time.UTC)
```

Call `StatisticsAPI.StatisticsGet` with only a start date, then again with a start and an end date, and print both results.

```go
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
```

`fail` prints the error message Elastic Email returned (or the network error) and exits.

## The whole code to copy and paste

```go
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
```

## Run the code

```bash
go run loadStatistics.go
```
