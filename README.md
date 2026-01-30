# WhatsApp Cloud API SDK (Go)

Lightweight Go SDK for the WhatsApp Cloud API with helpers to bootstrap multiple senders and keep the code organized by domain.

## Install

```bash
go get github.com/Rfluid/whatsapp-cloud-api
```

## Quick start (single sender)

```go
package main

import (
    "log"

    "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
    "github.com/Rfluid/whatsapp-cloud-api/src/message"
    "github.com/Rfluid/whatsapp-cloud-api/src/message/content"
)

func main() {
    cfg := bootstrap.SenderConfig{
        AccessToken:   "YOUR_ACCESS_TOKEN",
        WABAID:        "YOUR_PHONE_NUMBER_ID",
        WABAAccountID: "YOUR_WABA_ACCOUNT_ID",
    }

    api, err := bootstrap.FromConfig(cfg)
    if err != nil {
        log.Fatal(err)
    }

    msg := message.Message{
        Direction: message.Direction{
            To:   "RECIPIENT_PHONE",
            Type: content.Text,
        },
        Content: message.Content{
            Text: &content.TextData{
                Body:       "Hello from Go",
                PreviewURL: false,
            },
        },
    }
    msg.Default.SetDefault()

    _, err = message.Send(*api, msg)
    if err != nil {
        log.Fatal(err)
    }
}
```

## Multi-sender / multi-API setup

`FromConfigWithClient` lets you share a single `http.Client` across multiple APIs for better connection pooling. This is useful when your application sends on behalf of many WABA numbers.

```go
package main

import (
    "log"
    "net/http"
    "time"

    "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
)

type Sender struct {
    Name string
    Cfg  bootstrap.SenderConfig
}

func main() {
    sharedClient := &http.Client{
        Timeout: 15 * time.Second,
    }

    senders := []Sender{
        {
            Name: "support",
            Cfg: bootstrap.SenderConfig{
                AccessToken:   "SUPPORT_TOKEN",
                WABAID:        "SUPPORT_PHONE_ID",
                WABAAccountID: "SUPPORT_ACCOUNT_ID",
            },
        },
        {
            Name: "sales",
            Cfg: bootstrap.SenderConfig{
                AccessToken:   "SALES_TOKEN",
                WABAID:        "SALES_PHONE_ID",
                WABAAccountID: "SALES_ACCOUNT_ID",
            },
        },
    }

    apis := map[string]*bootstrap.WhatsAppAPI{}
    for _, sender := range senders {
        api, err := bootstrap.FromConfigWithClient(sender.Cfg, sharedClient)
        if err != nil {
            log.Fatal(err)
        }
        apis[sender.Name] = api
    }

    _ = apis
}
```

## Configuration

`bootstrap.SenderConfig` fields:

- `AccessToken`: Meta access token for the sender.
- `WABAID`: Phone number ID.
- `WABAAccountID`: WhatsApp Business Account ID.
- `Version`: Graph API version (defaults to `v24.0` when `nil`).
- `CustomMainURL`: Optional override for the Graph API base URL.

## Packages overview

- `src/bootstrap`: Create configured API clients (`GenerateWhatsAppAPI`, `FromConfig`, `FromConfigWithClient`).
- `src/message`: Send messages, mark as read, check status, identity checks, and helpers for message payloads.
- `src/media`: Upload media, retrieve media info, use media in messages.
- `src/template`: Create and fetch templates, build template components.
- `src/business`: Fetch business account data.
- `src/profile`: Manage business profile.
- `src/phone`: Request/verify registration codes and PIN flows.
- `src/typing`: Send typing indicators.
- `src/webhook`: Webhook payload models.
- `src/compliance`: Compliance and customer care details.
- `src/validators`: Request payload validators.
- `src/common`: Shared models, errors, and endpoints.

## Notes

- `FromConfig` creates a new `http.Client` for every API instance.
- `FromConfigWithClient` reuses a shared client for multiple senders.
