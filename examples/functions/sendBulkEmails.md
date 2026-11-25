# Send Bulk Emails

This guide walks you through sending a bulk email with per-recipient merge fields using the Elastic Email Go library.

*Required Access Level: SendHttp*

## What's a bulk email?

A bulk send is one job that goes to many recipients. Each recipient can have its own merge fields, so `{name}` in the content is replaced per person.

## Preparation

Install Go 1.18 or later.

Create a project and add the Elastic Email Go library:

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Create a file named `sendBulkEmails.go` and open it in your editor.

Create an API key in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with at least the access level above, and export it:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

## Let's dig into the code

The first line, `//go:build ignore`, lets this file live next to the other examples in the repository. `go run sendBulkEmails.go` still runs it.

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

Describe each recipient: the email address and the merge fields used in the content.

```go
recipient := ElasticEmail.NewEmailRecipient("johnsmith@domain.com")
recipient.SetFields(map[string]string{"name": "John"})
```

Prepare the body in HTML and in plain text. `{name}` is replaced with each recipient's `name` field.

```go
html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
html.SetContent("<strong>Hi {name}!</strong>")
html.SetCharset("utf-8")
text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
text.SetContent("Hi {name}!")
text.SetCharset("utf-8")
```

Set the sender (an address on your verified domain), the reply-to address, the subject and the body.

```go
content := ElasticEmail.NewEmailContent("myemail@domain.com")
content.SetReplyTo("myemail@domain.com")
content.SetSubject("Example email")
content.SetBody([]ElasticEmail.BodyPart{*html, *text})
```

Put the recipients and the content together.

```go
message := ElasticEmail.NewEmailMessageData([]ElasticEmail.EmailRecipient{*recipient}, *content)
```

Create an idempotency key for this send. If the call times out, retry with the same key and the same message: Elastic Email returns the first result instead of sending twice.

```go
// One key per logical send. Reuse it only to retry this same message.
key := newIdempotencyKey()
```

Call `EmailsAPI.EmailsPost` and print the TransactionID, MessageID and the key.

```go
result, httpRes, err := client.EmailsAPI.EmailsPost(ctx).
	EmailMessageData(*message).
	IdempotencyKey(key).
	Execute()
if err != nil {
	fail(err, httpRes)
}
fmt.Printf("Sent. TransactionID: %s, MessageID: %s, Idempotency-Key: %s\n",
	result.GetTransactionID(), result.GetMessageID(), key)
```

`fail` prints the error message Elastic Email returned (or the network error) and exits. `newIdempotencyKey` builds a random UUID with `crypto/rand` from the standard library, so the example needs no extra packages.

## The whole code to copy and paste

```go
//go:build ignore

// Send Bulk Emails
// Sends one bulk email job. Each recipient has its own merge fields, so {name} is
// replaced per recipient. Passes an Idempotency-Key so the send is safe to retry.
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

	recipient := ElasticEmail.NewEmailRecipient("johnsmith@domain.com")
	recipient.SetFields(map[string]string{"name": "John"})

	html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
	html.SetContent("<strong>Hi {name}!</strong>")
	html.SetCharset("utf-8")
	text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
	text.SetContent("Hi {name}!")
	text.SetCharset("utf-8")

	content := ElasticEmail.NewEmailContent("myemail@domain.com")
	content.SetReplyTo("myemail@domain.com")
	content.SetSubject("Example email")
	content.SetBody([]ElasticEmail.BodyPart{*html, *text})

	message := ElasticEmail.NewEmailMessageData([]ElasticEmail.EmailRecipient{*recipient}, *content)

	// One key per logical send. Reuse it only to retry this same message.
	key := newIdempotencyKey()

	result, httpRes, err := client.EmailsAPI.EmailsPost(ctx).
		EmailMessageData(*message).
		IdempotencyKey(key).
		Execute()
	if err != nil {
		fail(err, httpRes)
	}
	fmt.Printf("Sent. TransactionID: %s, MessageID: %s, Idempotency-Key: %s\n",
		result.GetTransactionID(), result.GetMessageID(), key)
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
```

## Run the code

```bash
go run sendBulkEmails.go
```
