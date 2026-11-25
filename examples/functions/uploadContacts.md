# Upload Contacts

This guide walks you through uploading contacts from a CSV file using the Elastic Email Go library.

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

Create a file named `uploadContacts.go` and open it in your editor.

Create `files/contacts.csv` next to the example, for example:

```csv
Email
john@domain.com
```

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run uploadContacts.go` still runs it.

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

Open the CSV file. It needs an `Email` column; `AllowUnsubscribe`, `Status`, `ConsentDate`, `ConsentIP` and `ConsentTracking` are optional.

```go
file, err := os.Open("files/contacts.csv")
if err != nil {
	log.Fatal(err)
}
defer file.Close()
```

Call `ContactsAPI.ContactsImportPost` with the list to add the contacts to, the file encoding and the file.

```go
httpRes, err := client.ContactsAPI.ContactsImportPost(ctx).
	ListName("example list").
	EncodingName("utf-8").
	File(file).
	Execute()
if err != nil {
	fail(err, httpRes)
}
fmt.Println("Contacts uploaded.")
```

`fail` prints the error message Elastic Email returned (or the network error) and exits.

## The whole code to copy and paste

```go
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
```

## Run the code

```bash
go run uploadContacts.go
```
