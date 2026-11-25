# Delete Contacts

This guide walks you through deleting several contacts in one call using the Elastic Email Go library.

*Required Access Level: ModifyContacts*

## What's a contact?

When using Elastic Email, you send emails to contacts: the recipients who receive your emails. Contacts can be grouped into lists and segments.

## Preparation

Install Go 1.18 or later.

Create a project and add the Elastic Email Go library:

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Create a file named `deleteContacts.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run deleteContacts.go` still runs it.

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

List the emails of the contacts to delete. You can pass a rule (`SetRule`) instead of emails, but not both.

```go
payload := ElasticEmail.NewEmailsPayload()
payload.SetEmails([]string{"johnsmith@domain.com"})
```

Call `ContactsAPI.ContactsDeletePost`.

```go
httpRes, err := client.ContactsAPI.ContactsDeletePost(ctx).
	EmailsPayload(*payload).
	Execute()
if err != nil {
	fail(err, httpRes)
}
fmt.Println("Contacts deleted.")
```

`fail` prints the error message Elastic Email returned (or the network error) and exits.

## The whole code to copy and paste

```go
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
```

## Run the code

```bash
go run deleteContacts.go
```
