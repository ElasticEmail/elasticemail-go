# Add Campaign

This guide walks you through creating a campaign using the Elastic Email Go library.

*Required Access Level: ModifyCampaigns*

## What's a campaign?

When you send an email to a group of contacts, Elastic Email calls that a campaign. A campaign needs a template (the email body) and recipients (lists or segments).

## Preparation

Install Go 1.18 or later.

Create a project and add the Elastic Email Go library:

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Create a file named `addCampaign.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run addCampaign.go` still runs it.

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

Describe the campaign content: the sender (an address on your verified domain), reply-to address, subject and the name of an existing template.

```go
template := ElasticEmail.NewCampaignTemplate("myemail@domain.com")
template.SetReplyTo("myemail@domain.com")
template.SetSubject("Hello")
template.SetTemplateName("hello_template")
```

Choose the recipients: one or more list names (or segment names).

```go
recipients := ElasticEmail.NewCampaignRecipient()
recipients.SetListNames([]string{"my list name"})
```

Create the campaign with a name and the recipients, then add the content and a status. `Draft` saves it without sending.

```go
campaign := ElasticEmail.NewCampaign("hello campaign", *recipients)
campaign.SetContent([]ElasticEmail.CampaignTemplate{*template})
campaign.SetStatus(ElasticEmail.CAMPAIGNSTATUS_DRAFT)
```

Call `CampaignsAPI.CampaignsPost` and print the created campaign.

```go
result, httpRes, err := client.CampaignsAPI.CampaignsPost(ctx).
	Campaign(*campaign).
	Execute()
if err != nil {
	fail(err, httpRes)
}
out, _ := json.MarshalIndent(result, "", "  ")
fmt.Println(string(out))
```

`fail` prints the error message Elastic Email returned (or the network error) and exits.

## The whole code to copy and paste

```go
//go:build ignore

// Add Campaign
// Creates a campaign that sends the template "hello_template" to the list "my list name".
// It is saved as a Draft. Set another status to send it right away or on schedule.
//
// Required access level: ModifyCampaigns
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

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

	template := ElasticEmail.NewCampaignTemplate("myemail@domain.com")
	template.SetReplyTo("myemail@domain.com")
	template.SetSubject("Hello")
	template.SetTemplateName("hello_template")

	recipients := ElasticEmail.NewCampaignRecipient()
	recipients.SetListNames([]string{"my list name"})

	campaign := ElasticEmail.NewCampaign("hello campaign", *recipients)
	campaign.SetContent([]ElasticEmail.CampaignTemplate{*template})
	campaign.SetStatus(ElasticEmail.CAMPAIGNSTATUS_DRAFT)

	result, httpRes, err := client.CampaignsAPI.CampaignsPost(ctx).
		Campaign(*campaign).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
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
go run addCampaign.go
```
