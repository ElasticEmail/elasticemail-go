//go:build ignore

// Upload Contacts
// Uploads contacts from files/contacts.csv and adds them to the list "example list".
// Required CSV column: Email. Suggested: AllowUnsubscribe, Status, ConsentDate, ConsentIP, ConsentTracking.
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

	file, err := os.Open("files/contacts.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	httpRes, err := client.ContactsAPI.ContactsImportPost(ctx).
		ListName("example list").
		EncodingName("utf-8").
		File(file).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	fmt.Println("Contacts uploaded.")
}

// fail prints the error returned by Elastic Email (or the network error) and exits.
func fail(err error, httpRes *http.Response) {
	var apiErr *ElasticEmail.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpRes != nil {
		log.Fatalf("Elastic Email API error %d: %s", httpRes.StatusCode, apiErr.Body())
	}
	log.Fatal(err)
}
