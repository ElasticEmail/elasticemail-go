<div align="center">

<img src=".github/ee-logo.png" alt="Elastic Email" width="96" />

# Elastic Email Go SDK

The official Go client library for the [Elastic Email](https://elasticemail.com) REST API v4.

[![Go module](https://img.shields.io/github/v/tag/ElasticEmail/elasticemail-go?logo=go&logoColor=white&label=Go&color=00ADD8)](https://pkg.go.dev/github.com/elasticemail/elasticemail-go/v4)
[![Go Reference](https://pkg.go.dev/badge/github.com/elasticemail/elasticemail-go/v4.svg)](https://pkg.go.dev/github.com/elasticemail/elasticemail-go/v4)
[![Go version](https://img.shields.io/badge/Go-1.18%2B-00ADD8?logo=go&logoColor=white)](https://go.dev/doc/devel/release)
[![API](https://img.shields.io/badge/API-v4-0A7BBB)](https://elasticemail.com/developers/api-documentation/rest-api)
[![OpenAPI Generator](https://img.shields.io/badge/generated%20by-OpenAPI%20Generator-6BA539?logo=openapiinitiative&logoColor=white)](https://openapi-generator.tech)
[![License: MIT](https://img.shields.io/github/license/ElasticEmail/elasticemail-go?color=yellow)](LICENSE)

[![Latest release](https://img.shields.io/github/v/release/ElasticEmail/elasticemail-go?logo=github&label=release)](https://github.com/ElasticEmail/elasticemail-go/releases)
[![Last commit](https://img.shields.io/github/last-commit/ElasticEmail/elasticemail-go?logo=github)](https://github.com/ElasticEmail/elasticemail-go/commits/master)
[![Open issues](https://img.shields.io/github/issues/ElasticEmail/elasticemail-go?logo=github)](https://github.com/ElasticEmail/elasticemail-go/issues)
[![GitHub stars](https://img.shields.io/github/stars/ElasticEmail/elasticemail-go?style=flat&logo=github)](https://github.com/ElasticEmail/elasticemail-go/stargazers)

[Installation](#installation) •
[Quick start](#quick-start) •
[Examples](#more-examples) •
[API reference](#api-reference) •
[Models](#models) •
[Contributing](#contributing)

</div>

---

## Features

- **Transactional and bulk email.** Send single messages, bulk campaigns or CSV merge-file sends.
- **Contacts, lists and segments.** Add, update, import, export and bulk-delete contacts.
- **Campaigns and automations.** Create, update, pause and trigger automations for a contact.
- **Templates, files and attachments.** Manage templates and uploaded files.
- **Domains.** Verify sending domains and check SPF, DKIM, tracking and certificate status.
- **Webhooks and inbound routes.** Receive delivery events and route incoming mail.
- **Statistics, events and suppressions.** Track delivery, bounces, complaints and unsubscribes.
- **Subaccounts and security.** Manage subaccounts and API keys.
- **Idiomatic Go.** Every call takes a `context.Context`, uses a fluent request builder ending in `Execute()`, and has no third-party runtime dependencies.

## Requirements

| Platform | Version |
| --- | --- |
| Go | 1.18 or later, with Go modules |
| Dependencies | None beyond the Go standard library |

You'll also need an Elastic Email **API key**. You can create one in your [API settings](https://app.elasticemail.com/marketing/settings/new/manage-api). Each endpoint's documentation lists the access level it needs.

## Installation

Add the [`github.com/elasticemail/elasticemail-go/v4`](https://pkg.go.dev/github.com/elasticemail/elasticemail-go/v4) module to your project:

```bash
go get github.com/elasticemail/elasticemail-go/v4@latest
```

Then import it. The package name is `ElasticEmail`:

```go
import ElasticEmail "github.com/elasticemail/elasticemail-go/v4"
```

> [!NOTE]
> Since 4.2.0 the module path ends in `/v4`. If you imported `github.com/elasticemail/elasticemail-go` before, update your imports and run `go mod tidy`.

## Quick start

> [!IMPORTANT]
> Elastic Email only sends from verified domains. Before your first send, [verify your sending domain](https://help.elasticemail.com/en/articles/4934400-how-to-verify-your-domain) and use an address on that domain as the sender.

### Configure the client

```go
package main

import (
	"context"
	"os"

	ElasticEmail "github.com/elasticemail/elasticemail-go/v4"
)

func main() {
	client := ElasticEmail.NewAPIClient(ElasticEmail.NewConfiguration())

	ctx := context.WithValue(context.Background(), ElasticEmail.ContextAPIKeys,
		map[string]ElasticEmail.APIKey{
			"apikey": {Key: os.Getenv("ELASTICEMAIL_API_KEY")},
		})

	// use client and ctx, see below
}
```

The API key travels in the request context, so pass that `ctx` to every call.

> [!TIP]
> Keep your API key out of source code. Load it from an environment variable or a secrets manager.

### Send a transactional email

```go
html := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_HTML)
html.SetContent("<h1>Hello!</h1><p>Thanks for signing up.</p>")
text := ElasticEmail.NewBodyPart(ElasticEmail.BODYCONTENTTYPE_PLAIN_TEXT)
text.SetContent("Hello! Thanks for signing up.")

content := ElasticEmail.NewEmailContent("My App <no-reply@yourdomain.com>")
content.SetSubject("Welcome aboard!")
content.SetBody([]ElasticEmail.BodyPart{*html, *text})

recipients := ElasticEmail.NewTransactionalRecipient([]string{"john.doe@example.com"})
message := ElasticEmail.NewEmailTransactionalMessageData(*recipients, *content)

result, httpRes, err := client.EmailsAPI.EmailsTransactionalPost(ctx).
	EmailTransactionalMessageData(*message).
	Execute()
if err != nil {
	var apiErr *ElasticEmail.GenericOpenAPIError
	if errors.As(err, &apiErr) && httpRes != nil {
		log.Fatalf("Elastic Email API error %d: %s", httpRes.StatusCode, apiErr.Body())
	}
	log.Fatal(err)
}
fmt.Printf("Sent. TransactionID: %s, MessageID: %s\n", result.GetTransactionID(), result.GetMessageID())
```

The `from` address must use a domain you've [verified in your Elastic Email account](https://help.elasticemail.com/en/articles/4934400-how-to-verify-your-domain).

### Send from a template with merge fields

```go
content := ElasticEmail.NewEmailContent("My App <no-reply@yourdomain.com>")
content.SetTemplateName("welcome-template")
content.SetMerge(map[string]string{"firstname": "John"})

recipients := ElasticEmail.NewTransactionalRecipient([]string{"john.doe@example.com"})
message := ElasticEmail.NewEmailTransactionalMessageData(*recipients, *content)

_, _, err := client.EmailsAPI.EmailsTransactionalPost(ctx).
	EmailTransactionalMessageData(*message).
	Execute()
```

### Using a proxy or a custom HTTP client

`Configuration.HTTPClient` accepts any `*http.Client`, so you can set timeouts, a proxy or your own transport:

```go
proxyURL, _ := url.Parse("http://myProxyUrl:80/")

cfg := ElasticEmail.NewConfiguration()
cfg.HTTPClient = &http.Client{
	Timeout:   30 * time.Second,
	Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
}
client := ElasticEmail.NewAPIClient(cfg)
```

The default client (`http.DefaultClient`) also honours the `HTTP_PROXY` and `HTTPS_PROXY` environment variables.

### Pointer helpers

Optional model fields are pointers. Use the setters (`SetSubject`, `SetMerge`, …) or the `Ptr*` helpers (`PtrString`, `PtrInt32`, `PtrBool`, `PtrTime`, …) to build values inline, and the `Get*` accessors to read them without nil checks.

## More examples

More complete, runnable samples are in the **[Elastic Email examples repository](https://github.com/ElasticEmail/elasticemail-examples)**. It covers transactional email, SMTP, webhooks, inbound email, contacts and serverless platforms across 20+ languages and frameworks.

- 🐹 [Go examples](https://github.com/ElasticEmail/elasticemail-examples/tree/main/go-elasticemail-examples) (plain `net/http`, Chi and Gin)
- 📂 [All examples](https://github.com/ElasticEmail/elasticemail-examples)

## Authentication

| Scheme | Header | Used for |
| --- | --- | --- |
| `apikey` | `X-ElasticEmail-ApiKey` | All standard API calls |
| `ApiKeyAuthCustomBranding` | `X-Auth-Token` | Custom-branding (white-label) accounts |

Both are passed as entries in the `map[string]ElasticEmail.APIKey` stored under `ElasticEmail.ContextAPIKeys`, keyed by the scheme name.

## API limits

- Up to **20 concurrent connections** per account
- A hard timeout of **600 seconds** per request

## API reference

All URIs are relative to `https://api.elasticemail.com/v4`. The SDK covers **114 endpoints** across 16 API services, exposed as fields on `APIClient`: `CampaignsAPI`, `ContactsAPI`, `DomainsAPI`, `EmailsAPI`, `EventsAPI`, `FilesAPI`, `InboundRouteAPI`, `ListsAPI`, `SecurityAPI`, `SegmentsAPI`, `StatisticsAPI`, `SubAccountsAPI`, `SuppressionsAPI`, `TemplatesAPI`, `VerificationsAPI` and `WebhookAPI`.

<details>
<summary><strong>Show all endpoints</strong></summary>

All URIs are relative to *https://api.elasticemail.com/v4*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CampaignsAPI* | [**CampaignsAutomationByNameTriggerPost**](docs/CampaignsAPI.md#campaignsautomationbynametriggerpost) | **Post** /campaigns/automation/{name}/trigger | Trigger Automation for Contact
*CampaignsAPI* | [**CampaignsByNameDelete**](docs/CampaignsAPI.md#campaignsbynamedelete) | **Delete** /campaigns/{name} | Delete Campaign
*CampaignsAPI* | [**CampaignsByNameGet**](docs/CampaignsAPI.md#campaignsbynameget) | **Get** /campaigns/{name} | Load Campaign
*CampaignsAPI* | [**CampaignsByNamePausePut**](docs/CampaignsAPI.md#campaignsbynamepauseput) | **Put** /campaigns/{name}/pause | Pause Campaign
*CampaignsAPI* | [**CampaignsByNamePut**](docs/CampaignsAPI.md#campaignsbynameput) | **Put** /campaigns/{name} | Update Campaign
*CampaignsAPI* | [**CampaignsGet**](docs/CampaignsAPI.md#campaignsget) | **Get** /campaigns | Load Campaigns
*CampaignsAPI* | [**CampaignsPost**](docs/CampaignsAPI.md#campaignspost) | **Post** /campaigns | Add Campaign
*ContactsAPI* | [**ContactsByEmailDelete**](docs/ContactsAPI.md#contactsbyemaildelete) | **Delete** /contacts/{email} | Delete Contact
*ContactsAPI* | [**ContactsByEmailGet**](docs/ContactsAPI.md#contactsbyemailget) | **Get** /contacts/{email} | Load Contact
*ContactsAPI* | [**ContactsByEmailPut**](docs/ContactsAPI.md#contactsbyemailput) | **Put** /contacts/{email} | Update Contact
*ContactsAPI* | [**ContactsDeletePost**](docs/ContactsAPI.md#contactsdeletepost) | **Post** /contacts/delete | Delete Contacts Bulk
*ContactsAPI* | [**ContactsExportByIdStatusGet**](docs/ContactsAPI.md#contactsexportbyidstatusget) | **Get** /contacts/export/{id}/status | Check Export Status
*ContactsAPI* | [**ContactsExportPost**](docs/ContactsAPI.md#contactsexportpost) | **Post** /contacts/export | Export Contacts
*ContactsAPI* | [**ContactsGet**](docs/ContactsAPI.md#contactsget) | **Get** /contacts | Load Contacts
*ContactsAPI* | [**ContactsImportPost**](docs/ContactsAPI.md#contactsimportpost) | **Post** /contacts/import | Upload Contacts
*ContactsAPI* | [**ContactsPost**](docs/ContactsAPI.md#contactspost) | **Post** /contacts | Add Contact
*DomainsAPI* | [**DomainsByDomainDelete**](docs/DomainsAPI.md#domainsbydomaindelete) | **Delete** /domains/{domain} | Delete Domain
*DomainsAPI* | [**DomainsByDomainGet**](docs/DomainsAPI.md#domainsbydomainget) | **Get** /domains/{domain} | Load Domain
*DomainsAPI* | [**DomainsByDomainPut**](docs/DomainsAPI.md#domainsbydomainput) | **Put** /domains/{domain} | Update Domain
*DomainsAPI* | [**DomainsByDomainRestrictedGet**](docs/DomainsAPI.md#domainsbydomainrestrictedget) | **Get** /domains/{domain}/restricted | Check for domain restriction
*DomainsAPI* | [**DomainsByDomainVerificationPut**](docs/DomainsAPI.md#domainsbydomainverificationput) | **Put** /domains/{domain}/verification | Verify Domain
*DomainsAPI* | [**DomainsByEmailDefaultPatch**](docs/DomainsAPI.md#domainsbyemaildefaultpatch) | **Patch** /domains/{email}/default | Set Default
*DomainsAPI* | [**DomainsGet**](docs/DomainsAPI.md#domainsget) | **Get** /domains | Load Domains
*DomainsAPI* | [**DomainsPost**](docs/DomainsAPI.md#domainspost) | **Post** /domains | Add Domain
*EmailsAPI* | [**EmailsByMsgidViewGet**](docs/EmailsAPI.md#emailsbymsgidviewget) | **Get** /emails/{msgid}/view | View Email
*EmailsAPI* | [**EmailsByTransactionidStatusGet**](docs/EmailsAPI.md#emailsbytransactionidstatusget) | **Get** /emails/{transactionid}/status | Get Status
*EmailsAPI* | [**EmailsMergefilePost**](docs/EmailsAPI.md#emailsmergefilepost) | **Post** /emails/mergefile | Send Bulk Emails CSV
*EmailsAPI* | [**EmailsPost**](docs/EmailsAPI.md#emailspost) | **Post** /emails | Send Bulk Emails
*EmailsAPI* | [**EmailsTransactionalPost**](docs/EmailsAPI.md#emailstransactionalpost) | **Post** /emails/transactional | Send Transactional Email
*EventsAPI* | [**EventsByTransactionidGet**](docs/EventsAPI.md#eventsbytransactionidget) | **Get** /events/{transactionid} | Load Email Events
*EventsAPI* | [**EventsChannelsByNameExportPost**](docs/EventsAPI.md#eventschannelsbynameexportpost) | **Post** /events/channels/{name}/export | Export Channel Events
*EventsAPI* | [**EventsChannelsByNameGet**](docs/EventsAPI.md#eventschannelsbynameget) | **Get** /events/channels/{name} | Load Channel Events
*EventsAPI* | [**EventsChannelsExportByIdStatusGet**](docs/EventsAPI.md#eventschannelsexportbyidstatusget) | **Get** /events/channels/export/{id}/status | Check Channel Export Status
*EventsAPI* | [**EventsExportByIdStatusGet**](docs/EventsAPI.md#eventsexportbyidstatusget) | **Get** /events/export/{id}/status | Check Export Status
*EventsAPI* | [**EventsExportPost**](docs/EventsAPI.md#eventsexportpost) | **Post** /events/export | Export Events
*EventsAPI* | [**EventsGet**](docs/EventsAPI.md#eventsget) | **Get** /events | Load Events
*FilesAPI* | [**FilesByNameDelete**](docs/FilesAPI.md#filesbynamedelete) | **Delete** /files/{name} | Delete File
*FilesAPI* | [**FilesByNameGet**](docs/FilesAPI.md#filesbynameget) | **Get** /files/{name} | Download File
*FilesAPI* | [**FilesByNameInfoGet**](docs/FilesAPI.md#filesbynameinfoget) | **Get** /files/{name}/info | Load File Details
*FilesAPI* | [**FilesGet**](docs/FilesAPI.md#filesget) | **Get** /files | List Files
*FilesAPI* | [**FilesPost**](docs/FilesAPI.md#filespost) | **Post** /files | Upload File
*InboundRouteAPI* | [**InboundrouteByIdDelete**](docs/InboundRouteAPI.md#inboundroutebyiddelete) | **Delete** /inboundroute/{id} | Delete Route
*InboundRouteAPI* | [**InboundrouteByIdGet**](docs/InboundRouteAPI.md#inboundroutebyidget) | **Get** /inboundroute/{id} | Get Route
*InboundRouteAPI* | [**InboundrouteByIdPut**](docs/InboundRouteAPI.md#inboundroutebyidput) | **Put** /inboundroute/{id} | Update Route
*InboundRouteAPI* | [**InboundrouteGet**](docs/InboundRouteAPI.md#inboundrouteget) | **Get** /inboundroute | Get Routes
*InboundRouteAPI* | [**InboundrouteOrderPut**](docs/InboundRouteAPI.md#inboundrouteorderput) | **Put** /inboundroute/order | Update Sorting
*InboundRouteAPI* | [**InboundroutePost**](docs/InboundRouteAPI.md#inboundroutepost) | **Post** /inboundroute | Create Route
*ListsAPI* | [**ListsByListnameContactsGet**](docs/ListsAPI.md#listsbylistnamecontactsget) | **Get** /lists/{listname}/contacts | Load Contacts in List
*ListsAPI* | [**ListsByNameContactsPost**](docs/ListsAPI.md#listsbynamecontactspost) | **Post** /lists/{name}/contacts | Add Contacts to List
*ListsAPI* | [**ListsByNameContactsRemovePost**](docs/ListsAPI.md#listsbynamecontactsremovepost) | **Post** /lists/{name}/contacts/remove | Remove Contacts from List
*ListsAPI* | [**ListsByNameDelete**](docs/ListsAPI.md#listsbynamedelete) | **Delete** /lists/{name} | Delete List
*ListsAPI* | [**ListsByNameGet**](docs/ListsAPI.md#listsbynameget) | **Get** /lists/{name} | Load List
*ListsAPI* | [**ListsByNamePut**](docs/ListsAPI.md#listsbynameput) | **Put** /lists/{name} | Update List
*ListsAPI* | [**ListsGet**](docs/ListsAPI.md#listsget) | **Get** /lists | Load Lists
*ListsAPI* | [**ListsPost**](docs/ListsAPI.md#listspost) | **Post** /lists | Add List
*SecurityAPI* | [**SecurityApikeysByNameDelete**](docs/SecurityAPI.md#securityapikeysbynamedelete) | **Delete** /security/apikeys/{name} | Delete ApiKey
*SecurityAPI* | [**SecurityApikeysByNameGet**](docs/SecurityAPI.md#securityapikeysbynameget) | **Get** /security/apikeys/{name} | Load ApiKey
*SecurityAPI* | [**SecurityApikeysByNamePut**](docs/SecurityAPI.md#securityapikeysbynameput) | **Put** /security/apikeys/{name} | Update ApiKey
*SecurityAPI* | [**SecurityApikeysGet**](docs/SecurityAPI.md#securityapikeysget) | **Get** /security/apikeys | List ApiKeys
*SecurityAPI* | [**SecurityApikeysPost**](docs/SecurityAPI.md#securityapikeyspost) | **Post** /security/apikeys | Add ApiKey
*SecurityAPI* | [**SecuritySmtpByNameDelete**](docs/SecurityAPI.md#securitysmtpbynamedelete) | **Delete** /security/smtp/{name} | Delete SMTP Credential
*SecurityAPI* | [**SecuritySmtpByNameGet**](docs/SecurityAPI.md#securitysmtpbynameget) | **Get** /security/smtp/{name} | Load SMTP Credential
*SecurityAPI* | [**SecuritySmtpByNamePut**](docs/SecurityAPI.md#securitysmtpbynameput) | **Put** /security/smtp/{name} | Update SMTP Credential
*SecurityAPI* | [**SecuritySmtpGet**](docs/SecurityAPI.md#securitysmtpget) | **Get** /security/smtp | List SMTP Credentials
*SecurityAPI* | [**SecuritySmtpPost**](docs/SecurityAPI.md#securitysmtppost) | **Post** /security/smtp | Add SMTP Credential
*SegmentsAPI* | [**SegmentsByNameDelete**](docs/SegmentsAPI.md#segmentsbynamedelete) | **Delete** /segments/{name} | Delete Segment
*SegmentsAPI* | [**SegmentsByNameGet**](docs/SegmentsAPI.md#segmentsbynameget) | **Get** /segments/{name} | Load Segment
*SegmentsAPI* | [**SegmentsByNamePut**](docs/SegmentsAPI.md#segmentsbynameput) | **Put** /segments/{name} | Update Segment
*SegmentsAPI* | [**SegmentsGet**](docs/SegmentsAPI.md#segmentsget) | **Get** /segments | Load Segments
*SegmentsAPI* | [**SegmentsPost**](docs/SegmentsAPI.md#segmentspost) | **Post** /segments | Add Segment
*StatisticsAPI* | [**StatisticsCampaignsByNameGet**](docs/StatisticsAPI.md#statisticscampaignsbynameget) | **Get** /statistics/campaigns/{name} | Load Campaign Stats
*StatisticsAPI* | [**StatisticsCampaignsGet**](docs/StatisticsAPI.md#statisticscampaignsget) | **Get** /statistics/campaigns | Load Campaigns Stats
*StatisticsAPI* | [**StatisticsChannelsByNameGet**](docs/StatisticsAPI.md#statisticschannelsbynameget) | **Get** /statistics/channels/{name} | Load Channel Stats
*StatisticsAPI* | [**StatisticsChannelsGet**](docs/StatisticsAPI.md#statisticschannelsget) | **Get** /statistics/channels | Load Channels Stats
*StatisticsAPI* | [**StatisticsGet**](docs/StatisticsAPI.md#statisticsget) | **Get** /statistics | Load Statistics
*SubAccountsAPI* | [**SubaccountsByEmailApikeyGet**](docs/SubAccountsAPI.md#subaccountsbyemailapikeyget) | **Get** /subaccounts/{email}/apikey | Get SubAccount ApiKey
*SubAccountsAPI* | [**SubaccountsByEmailCreditsPatch**](docs/SubAccountsAPI.md#subaccountsbyemailcreditspatch) | **Patch** /subaccounts/{email}/credits | Add, Subtract Email Credits
*SubAccountsAPI* | [**SubaccountsByEmailDelete**](docs/SubAccountsAPI.md#subaccountsbyemaildelete) | **Delete** /subaccounts/{email} | Delete SubAccount
*SubAccountsAPI* | [**SubaccountsByEmailGet**](docs/SubAccountsAPI.md#subaccountsbyemailget) | **Get** /subaccounts/{email} | Load SubAccount
*SubAccountsAPI* | [**SubaccountsByEmailSettingsEmailPut**](docs/SubAccountsAPI.md#subaccountsbyemailsettingsemailput) | **Put** /subaccounts/{email}/settings/email | Update SubAccount Email Settings
*SubAccountsAPI* | [**SubaccountsGet**](docs/SubAccountsAPI.md#subaccountsget) | **Get** /subaccounts | Load SubAccounts
*SubAccountsAPI* | [**SubaccountsPost**](docs/SubAccountsAPI.md#subaccountspost) | **Post** /subaccounts | Add SubAccount
*SuppressionsAPI* | [**SuppressionsBouncesGet**](docs/SuppressionsAPI.md#suppressionsbouncesget) | **Get** /suppressions/bounces | Get Bounce List
*SuppressionsAPI* | [**SuppressionsBouncesImportPost**](docs/SuppressionsAPI.md#suppressionsbouncesimportpost) | **Post** /suppressions/bounces/import | Add Bounces Async
*SuppressionsAPI* | [**SuppressionsBouncesPost**](docs/SuppressionsAPI.md#suppressionsbouncespost) | **Post** /suppressions/bounces | Add Bounces
*SuppressionsAPI* | [**SuppressionsByEmailDelete**](docs/SuppressionsAPI.md#suppressionsbyemaildelete) | **Delete** /suppressions/{email} | Delete Suppression
*SuppressionsAPI* | [**SuppressionsByEmailGet**](docs/SuppressionsAPI.md#suppressionsbyemailget) | **Get** /suppressions/{email} | Get Suppression
*SuppressionsAPI* | [**SuppressionsComplaintsGet**](docs/SuppressionsAPI.md#suppressionscomplaintsget) | **Get** /suppressions/complaints | Get Complaints List
*SuppressionsAPI* | [**SuppressionsComplaintsImportPost**](docs/SuppressionsAPI.md#suppressionscomplaintsimportpost) | **Post** /suppressions/complaints/import | Add Complaints Async
*SuppressionsAPI* | [**SuppressionsComplaintsPost**](docs/SuppressionsAPI.md#suppressionscomplaintspost) | **Post** /suppressions/complaints | Add Complaints
*SuppressionsAPI* | [**SuppressionsGet**](docs/SuppressionsAPI.md#suppressionsget) | **Get** /suppressions | Get Suppressions
*SuppressionsAPI* | [**SuppressionsUnsubscribesGet**](docs/SuppressionsAPI.md#suppressionsunsubscribesget) | **Get** /suppressions/unsubscribes | Get Unsubscribes List
*SuppressionsAPI* | [**SuppressionsUnsubscribesImportPost**](docs/SuppressionsAPI.md#suppressionsunsubscribesimportpost) | **Post** /suppressions/unsubscribes/import | Add Unsubscribes Async
*SuppressionsAPI* | [**SuppressionsUnsubscribesPost**](docs/SuppressionsAPI.md#suppressionsunsubscribespost) | **Post** /suppressions/unsubscribes | Add Unsubscribes
*TemplatesAPI* | [**TemplatesByNameDelete**](docs/TemplatesAPI.md#templatesbynamedelete) | **Delete** /templates/{name} | Delete Template
*TemplatesAPI* | [**TemplatesByNameGet**](docs/TemplatesAPI.md#templatesbynameget) | **Get** /templates/{name} | Load Template
*TemplatesAPI* | [**TemplatesByNamePut**](docs/TemplatesAPI.md#templatesbynameput) | **Put** /templates/{name} | Update Template
*TemplatesAPI* | [**TemplatesGet**](docs/TemplatesAPI.md#templatesget) | **Get** /templates | Load Templates
*TemplatesAPI* | [**TemplatesPost**](docs/TemplatesAPI.md#templatespost) | **Post** /templates | Add Template
*VerificationsAPI* | [**VerificationsByEmailDelete**](docs/VerificationsAPI.md#verificationsbyemaildelete) | **Delete** /verifications/{email} | Delete Email Verification Result
*VerificationsAPI* | [**VerificationsByEmailGet**](docs/VerificationsAPI.md#verificationsbyemailget) | **Get** /verifications/{email} | Get Email Verification Result
*VerificationsAPI* | [**VerificationsByEmailPost**](docs/VerificationsAPI.md#verificationsbyemailpost) | **Post** /verifications/{email} | Verify Email
*VerificationsAPI* | [**VerificationsFilesByIdDelete**](docs/VerificationsAPI.md#verificationsfilesbyiddelete) | **Delete** /verifications/files/{id} | Delete File Verification Result
*VerificationsAPI* | [**VerificationsFilesByIdResultDownloadGet**](docs/VerificationsAPI.md#verificationsfilesbyidresultdownloadget) | **Get** /verifications/files/{id}/result/download | Download File Verification Result
*VerificationsAPI* | [**VerificationsFilesByIdResultGet**](docs/VerificationsAPI.md#verificationsfilesbyidresultget) | **Get** /verifications/files/{id}/result | Get Detailed File Verification Result
*VerificationsAPI* | [**VerificationsFilesByIdVerificationPost**](docs/VerificationsAPI.md#verificationsfilesbyidverificationpost) | **Post** /verifications/files/{id}/verification | Start verification
*VerificationsAPI* | [**VerificationsFilesPost**](docs/VerificationsAPI.md#verificationsfilespost) | **Post** /verifications/files | Upload File with Emails
*VerificationsAPI* | [**VerificationsFilesResultGet**](docs/VerificationsAPI.md#verificationsfilesresultget) | **Get** /verifications/files/result | Get Files Verification Results
*VerificationsAPI* | [**VerificationsGet**](docs/VerificationsAPI.md#verificationsget) | **Get** /verifications | Get Emails Verification Results
*WebhookAPI* | [**WebhookByPublicidDelete**](docs/WebhookAPI.md#webhookbypubliciddelete) | **Delete** /webhook/{publicid} | Delete Webhook
*WebhookAPI* | [**WebhookByPublicidGet**](docs/WebhookAPI.md#webhookbypublicidget) | **Get** /webhook/{publicid} | Load Webhook
*WebhookAPI* | [**WebhookByPublicidPut**](docs/WebhookAPI.md#webhookbypublicidput) | **Put** /webhook/{publicid} | Update Webhook
*WebhookAPI* | [**WebhookGet**](docs/WebhookAPI.md#webhookget) | **Get** /webhook | Load Webhooks
*WebhookAPI* | [**WebhookPost**](docs/WebhookAPI.md#webhookpost) | **Post** /webhook | Add Webhook


</details>

## Models

<details>
<summary><strong>Show all 98 models</strong></summary>


 - [AccessLevel](docs/AccessLevel.md)
 - [AccountStatusEnum](docs/AccountStatusEnum.md)
 - [ApiKey](docs/ApiKey.md)
 - [ApiKeyPayload](docs/ApiKeyPayload.md)
 - [BodyContentType](docs/BodyContentType.md)
 - [BodyPart](docs/BodyPart.md)
 - [Campaign](docs/Campaign.md)
 - [CampaignOptions](docs/CampaignOptions.md)
 - [CampaignRecipient](docs/CampaignRecipient.md)
 - [CampaignStatus](docs/CampaignStatus.md)
 - [CampaignTemplate](docs/CampaignTemplate.md)
 - [CertificateValidationStatus](docs/CertificateValidationStatus.md)
 - [ChannelLogStatusSummary](docs/ChannelLogStatusSummary.md)
 - [CompressionFormat](docs/CompressionFormat.md)
 - [ConsentData](docs/ConsentData.md)
 - [ConsentTracking](docs/ConsentTracking.md)
 - [Contact](docs/Contact.md)
 - [ContactActivity](docs/ContactActivity.md)
 - [ContactPayload](docs/ContactPayload.md)
 - [ContactSource](docs/ContactSource.md)
 - [ContactStatus](docs/ContactStatus.md)
 - [ContactUpdatePayload](docs/ContactUpdatePayload.md)
 - [ContactsList](docs/ContactsList.md)
 - [DKIMRecord](docs/DKIMRecord.md)
 - [DeliveryOptimizationType](docs/DeliveryOptimizationType.md)
 - [DomainData](docs/DomainData.md)
 - [DomainDetail](docs/DomainDetail.md)
 - [DomainOwner](docs/DomainOwner.md)
 - [DomainPayload](docs/DomainPayload.md)
 - [DomainUpdatePayload](docs/DomainUpdatePayload.md)
 - [EmailContent](docs/EmailContent.md)
 - [EmailData](docs/EmailData.md)
 - [EmailJobFailedStatus](docs/EmailJobFailedStatus.md)
 - [EmailJobStatus](docs/EmailJobStatus.md)
 - [EmailMessageData](docs/EmailMessageData.md)
 - [EmailPredictedValidationStatus](docs/EmailPredictedValidationStatus.md)
 - [EmailRecipient](docs/EmailRecipient.md)
 - [EmailSend](docs/EmailSend.md)
 - [EmailStatus](docs/EmailStatus.md)
 - [EmailTransactionalMessageData](docs/EmailTransactionalMessageData.md)
 - [EmailValidationResult](docs/EmailValidationResult.md)
 - [EmailValidationStatus](docs/EmailValidationStatus.md)
 - [EmailView](docs/EmailView.md)
 - [EmailsPayload](docs/EmailsPayload.md)
 - [EncodingType](docs/EncodingType.md)
 - [EventType](docs/EventType.md)
 - [EventsOrderBy](docs/EventsOrderBy.md)
 - [ExportFileFormats](docs/ExportFileFormats.md)
 - [ExportLink](docs/ExportLink.md)
 - [ExportStatus](docs/ExportStatus.md)
 - [FileInfo](docs/FileInfo.md)
 - [FilePayload](docs/FilePayload.md)
 - [FileUploadResult](docs/FileUploadResult.md)
 - [InboundPayload](docs/InboundPayload.md)
 - [InboundRoute](docs/InboundRoute.md)
 - [InboundRouteActionType](docs/InboundRouteActionType.md)
 - [InboundRouteFilterType](docs/InboundRouteFilterType.md)
 - [ListPayload](docs/ListPayload.md)
 - [ListUpdatePayload](docs/ListUpdatePayload.md)
 - [LogJobStatus](docs/LogJobStatus.md)
 - [LogStatusSummary](docs/LogStatusSummary.md)
 - [MergeEmailPayload](docs/MergeEmailPayload.md)
 - [MessageAttachment](docs/MessageAttachment.md)
 - [MessageCategory](docs/MessageCategory.md)
 - [MessageCategoryEnum](docs/MessageCategoryEnum.md)
 - [NewApiKey](docs/NewApiKey.md)
 - [NewSmtpCredentials](docs/NewSmtpCredentials.md)
 - [Options](docs/Options.md)
 - [RecipientEvent](docs/RecipientEvent.md)
 - [Segment](docs/Segment.md)
 - [SegmentPayload](docs/SegmentPayload.md)
 - [SmtpCredentials](docs/SmtpCredentials.md)
 - [SmtpCredentialsPayload](docs/SmtpCredentialsPayload.md)
 - [SortOrderItem](docs/SortOrderItem.md)
 - [SplitOptimizationType](docs/SplitOptimizationType.md)
 - [SplitOptions](docs/SplitOptions.md)
 - [SubAccountInfo](docs/SubAccountInfo.md)
 - [SubaccountEmailCreditsPayload](docs/SubaccountEmailCreditsPayload.md)
 - [SubaccountEmailSettings](docs/SubaccountEmailSettings.md)
 - [SubaccountEmailSettingsPayload](docs/SubaccountEmailSettingsPayload.md)
 - [SubaccountPayload](docs/SubaccountPayload.md)
 - [SubaccountSettingsInfo](docs/SubaccountSettingsInfo.md)
 - [SubaccountSettingsInfoPayload](docs/SubaccountSettingsInfoPayload.md)
 - [Suppression](docs/Suppression.md)
 - [Template](docs/Template.md)
 - [TemplatePayload](docs/TemplatePayload.md)
 - [TemplateScope](docs/TemplateScope.md)
 - [TemplateType](docs/TemplateType.md)
 - [TrackingType](docs/TrackingType.md)
 - [TrackingValidationStatus](docs/TrackingValidationStatus.md)
 - [TransactionalRecipient](docs/TransactionalRecipient.md)
 - [Utm](docs/Utm.md)
 - [VerificationFileResult](docs/VerificationFileResult.md)
 - [VerificationFileResultDetails](docs/VerificationFileResultDetails.md)
 - [VerificationStatus](docs/VerificationStatus.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookCreatePayload](docs/WebhookCreatePayload.md)
 - [WebhookUpdatePayload](docs/WebhookUpdatePayload.md)


</details>

## Versioning

The SDK follows the Elastic Email API v4 and uses [semantic import versioning](https://go.dev/blog/v2-go-modules), so the major version is part of the module path (`/v4`). Available versions are listed on [pkg.go.dev](https://pkg.go.dev/github.com/elasticemail/elasticemail-go/v4?tab=versions) and in [GitHub Releases](https://github.com/ElasticEmail/elasticemail-go/releases).

<details>
<summary>Build details</summary>

- API version: 4.0.0
- SDK version: 4.2.0
- Generator version: 7.11.0
- Build package: `org.openapitools.codegen.languages.GoClientCodegen`

</details>

## Contributing

Contributions are welcome! Most of this SDK is generated from the [OpenAPI specification](api/openapi.yaml), so please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

- 🐛 [Report a bug](https://github.com/ElasticEmail/elasticemail-go/issues/new?template=bug_report.md)
- 💡 [Request a feature](https://github.com/ElasticEmail/elasticemail-go/issues/new?template=feature_request.md)
- 🔒 [Report a security issue](SECURITY.md)

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).

## Support

> [!IMPORTANT]
> The fastest way to get help is the **chat widget on [elasticemail.com](https://elasticemail.com)**. Our support team can help with your account, sending, deliverability and API questions.

- 💬 [Chat with support on elasticemail.com](https://elasticemail.com) (preferred)
- 📚 [API documentation](https://elasticemail.com/developers/api-documentation/rest-api)
- 🧪 [Examples repository](https://github.com/ElasticEmail/elasticemail-examples)
- 🐛 [GitHub issues](https://github.com/ElasticEmail/elasticemail-go/issues), for bugs in this SDK only

## License

Released under the [MIT License](LICENSE). Copyright © 2021–2026 Elastic Email.
