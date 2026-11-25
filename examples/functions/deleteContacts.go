//go:build ignore

// Delete Contacts
// Deletes contacts in bulk, by a list of email addresses.
// To delete a single contact, see deleteContact.go.
//
// Required access level: ModifyContacts
package main

import (
	"context"
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

	payload := ElasticEmail.NewEmailsPayload()
	payload.SetEmails([]string{"johnsmith@domain.com"})

	httpRes, err := client.ContactsAPI.ContactsDeletePost(ctx).
		EmailsPayload(*payload).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	fmt.Println("Contacts deleted.")
}

// fail prints the error returned by Elastic Email (or the network error) and exits.
func fail(err error, httpRes *http.Response) {
	var apiErr *ElasticEmail.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpRes != nil {
		log.Fatalf("Elastic Email API error %d: %s", httpRes.StatusCode, apiErr.Body())
	}
	log.Fatal(err)
}
