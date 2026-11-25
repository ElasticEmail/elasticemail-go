# Go Examples

Small, runnable programs for the Elastic Email Go library. Each one calls a single API operation and has a walkthrough (`.md`) next to it.

## Requirements

- Go 1.18 or later
- An Elastic Email API key. Create one in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api) with the access level listed for the example.
- For the send examples, a sender address on a domain you've [verified in your account](https://help.elasticemail.com/en/articles/4934400-how-to-verify-your-domain).

## Set your API key

Every example reads the key from the `ELASTICEMAIL_API_KEY` environment variable and stops with a message if it's missing:

```bash
export ELASTICEMAIL_API_KEY="your-api-key"
```

Other values (sender, recipients, list, template and campaign names) are placeholders in the code, such as `myemail@domain.com` and `johnsmith@domain.com`. Replace them before you run an example.

## Run an example

### From a clone of this repository

The examples import the module this repository defines, so they build against the code in your clone with no extra steps:

```bash
git clone https://github.com/ElasticEmail/elasticemail-go.git
cd elasticemail-go/examples/functions
go run sendTransactionalEmails.go
```

Run them from `examples/functions`: `uploadContacts.go` reads `files/contacts.csv` and `sendTransactionalEmailsWithAttachment.go` reads `invoice.pdf` from the current directory.

### In your own project

```bash
mkdir ee-example && cd ee-example
go mod init example.com/ee-example
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Copy an example file into the folder and run it:

```bash
go run sendTransactionalEmails.go
```

### Why every file starts with `//go:build ignore`

Each example is its own `package main` with its own `main` function, so they can't be built together as one package. The `//go:build ignore` line keeps them out of `go build ./...`, `go vet ./...` and `go test ./...`, and the library builds as before. Naming a file explicitly still works: `go run file.go`, `go build file.go` and `go vet file.go` all build it. You can delete the line when you copy an example into your own project.

## Examples

| Example | What it does | Access level | Docs |
| --- | --- | --- | --- |
| [addCampaign](functions/addCampaign.go) | Create a campaign | ModifyCampaigns | [walkthrough](functions/addCampaign.md) |
| [addContacts](functions/addContacts.go) | Add contacts, optionally to a list | ModifyContacts | [walkthrough](functions/addContacts.md) |
| [addList](functions/addList.go) | Create a contact list | ModifyContacts | [walkthrough](functions/addList.md) |
| [addTemplate](functions/addTemplate.go) | Create a template | ModifyTemplates | [walkthrough](functions/addTemplate.md) |
| [deleteCampaign](functions/deleteCampaign.go) | Delete a campaign | ModifyCampaigns | [walkthrough](functions/deleteCampaign.md) |
| [deleteContact](functions/deleteContact.go) | Delete one contact by email | ModifyContacts | [walkthrough](functions/deleteContact.md) |
| [deleteContacts](functions/deleteContacts.go) | Delete contacts in bulk | ModifyContacts | [walkthrough](functions/deleteContacts.md) |
| [deleteList](functions/deleteList.go) | Delete a contact list | ModifyContacts | [walkthrough](functions/deleteList.md) |
| [deleteTemplate](functions/deleteTemplate.go) | Delete a template | ModifyTemplates | [walkthrough](functions/deleteTemplate.md) |
| [exportContacts](functions/exportContacts.go) | Export contacts to a file | Export | [walkthrough](functions/exportContacts.md) |
| [loadCampaign](functions/loadCampaign.go) | Load a campaign | ViewCampaigns | [walkthrough](functions/loadCampaign.md) |
| [loadCampaignsStats](functions/loadCampaignsStats.go) | Load campaign statistics | ViewChannels | [walkthrough](functions/loadCampaignsStats.md) |
| [loadChannelsStats](functions/loadChannelsStats.go) | Load channel statistics | ViewChannels | [walkthrough](functions/loadChannelsStats.md) |
| [loadList](functions/loadList.go) | Load a contact list | ViewContacts | [walkthrough](functions/loadList.md) |
| [loadStatistics](functions/loadStatistics.go) | Load account statistics | ViewReports | [walkthrough](functions/loadStatistics.md) |
| [loadTemplate](functions/loadTemplate.go) | Load a template | ViewTemplates | [walkthrough](functions/loadTemplate.md) |
| [sendBulkEmails](functions/sendBulkEmails.go) | Send a bulk email with merge fields | SendHttp | [walkthrough](functions/sendBulkEmails.md) |
| [sendTransactionalEmails](functions/sendTransactionalEmails.go) | Send a transactional email | SendHttp | [walkthrough](functions/sendTransactionalEmails.md) |
| [sendTransactionalEmailsWithAttachment](functions/sendTransactionalEmailsWithAttachment.go) | Send a transactional email with an attachment | SendHttp | [walkthrough](functions/sendTransactionalEmailsWithAttachment.md) |
| [updateCampaign](functions/updateCampaign.go) | Update a campaign | ModifyCampaigns | [walkthrough](functions/updateCampaign.md) |
| [uploadContacts](functions/uploadContacts.go) | Upload contacts from a CSV file | ModifyContacts | [walkthrough](functions/uploadContacts.md) |

The send examples pass an `Idempotency-Key`, created with the standard library, and print it next to the `TransactionID`. If a send times out, retrying with the same key and the same message returns the first result instead of sending twice. See [Retry a send safely](../README.md#retry-a-send-safely-idempotency-key) in the main README.
