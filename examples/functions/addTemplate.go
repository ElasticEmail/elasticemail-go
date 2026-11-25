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
