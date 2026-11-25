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
