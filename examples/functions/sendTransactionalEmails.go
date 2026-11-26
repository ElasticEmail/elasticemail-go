//go:build ignore

// Send Transactional Emails
// Sends a transactional email with an HTML and a plain-text body (up to 50 recipients).
// Passes an Idempotency-Key so the send is safe to retry.
//
// Required access level: SendHttp
package main

import (
	"context"
	"crypto/rand"
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

	// README:BEGIN quickstart
	html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
	html.SetContent("<strong>Mail content.</strong>")
	html.SetCharset("utf-8")
	text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
	text.SetContent("Mail content.")
	text.SetCharset("utf-8")

	content := ElasticEmail.NewEmailContent("myemail@domain.com")
	content.SetReplyTo("myemail@domain.com")
	content.SetSubject("Example transactional email")
	content.SetBody([]ElasticEmail.BodyPart{*html, *text})

	recipients := ElasticEmail.NewTransactionalRecipient([]string{"johnsmith@domain.com"})
	message := ElasticEmail.NewEmailTransactionalMessageData(*recipients, *content)

	// One key per logical send. Reuse it only to retry this same message.
	key := newIdempotencyKey()

	result, httpRes, err := client.EmailsAPI.EmailsTransactionalPost(ctx).
		EmailTransactionalMessageData(*message).
		IdempotencyKey(key).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	fmt.Printf("Sent. TransactionID: %s, MessageID: %s, Idempotency-Key: %s\n",
		result.GetTransactionID(), result.GetMessageID(), key)
	// README:END quickstart
}

// fail prints the error returned by Elastic Email (or the network error) and exits.
func fail(err error, httpRes *http.Response) {
	var apiErr *ElasticEmail.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpRes != nil {
		log.Fatalf("Elastic Email API error %d: %s", httpRes.StatusCode, apiErr.Body())
	}
	log.Fatal(err)
}

// newIdempotencyKey returns a random UUID v4. Create one per logical send and
// reuse it only to retry that same send.
func newIdempotencyKey() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Fatal(err)
	}
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
