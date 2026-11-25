# Export Contacts

This guide walks you through exporting your contacts to a downloadable file using the Elastic Email Go library.

*Required Access Level: Export*

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

Create a file named `exportContacts.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run exportContacts.go` still runs it.

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

Call `ContactsAPI.ContactsExportPost` with the file format (`Csv`, `Xml` or `Json`), the compression (`None` or `Zip`) and a file name. To export only some contacts, add `Emails([]string{...})` or `Rule("Status = Engaged")`. The response holds the download link and the export ID.

```go
result, httpRes, err := client.ContactsAPI.ContactsExportPost(ctx).
	FileFormat(ElasticEmail.EXPORTFILEFORMATS_CSV).
	CompressionFormat(ElasticEmail.COMPRESSIONFORMAT_NONE).
	FileName("exported.csv").
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

// Export Contacts
// Exports your contacts to a CSV file and prints the download link.
// Optional filters: Emails (a list of addresses) or Rule (a query such as "Status = Engaged").
//
// Required access level: Export
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

	result, httpRes, err := client.ContactsAPI.ContactsExportPost(ctx).
		FileFormat(ElasticEmail.EXPORTFILEFORMATS_CSV).
		CompressionFormat(ElasticEmail.COMPRESSIONFORMAT_NONE).
		FileName("exported.csv").
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
go run exportContacts.go
```
