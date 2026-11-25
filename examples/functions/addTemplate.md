# Add Template

This guide walks you through creating an email template using the Elastic Email Go library.

*Required Access Level: ModifyTemplates*

## What's a template?

A template is an email body saved under a name. Until you delete it, you can reuse it to send any number of emails and campaigns.

## Preparation

Install Go 1.18 or later.

Create a project and add the Elastic Email Go library:

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Create a file named `addTemplate.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run addTemplate.go` still runs it.

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

Prepare the body parts: one HTML and one plain-text version.

```go
html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
html.SetContent("<strong>My template</strong>")
html.SetCharset("utf-8")
text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
text.SetContent("My template")
text.SetCharset("utf-8")
```

Describe the template: name, default subject, body and scope (`TEMPLATESCOPE_PERSONAL` or `TEMPLATESCOPE_GLOBAL`).

```go
template := ElasticEmail.NewTemplatePayload("My new template")
template.SetSubject("Default subject")
template.SetBody([]ElasticEmail.BodyPart{*html, *text})
template.SetTemplateScope(ElasticEmail.TEMPLATESCOPE_PERSONAL)
```

Call `TemplatesAPI.TemplatesPost` and print the new template.

```go
result, httpRes, err := client.TemplatesAPI.TemplatesPost(ctx).
	TemplatePayload(*template).
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

// Add Template
// Creates the template "My new template" with an HTML and a plain-text body.
//
// Required access level: ModifyTemplates
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

	html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
	html.SetContent("<strong>My template</strong>")
	html.SetCharset("utf-8")
	text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
	text.SetContent("My template")
	text.SetCharset("utf-8")

	template := ElasticEmail.NewTemplatePayload("My new template")
	template.SetSubject("Default subject")
	template.SetBody([]ElasticEmail.BodyPart{*html, *text})
	template.SetTemplateScope(ElasticEmail.TEMPLATESCOPE_PERSONAL)

	result, httpRes, err := client.TemplatesAPI.TemplatesPost(ctx).
		TemplatePayload(*template).
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
go run addTemplate.go
```
