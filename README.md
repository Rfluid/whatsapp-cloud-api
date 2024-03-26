# About

This package provides a sdk for WhatsApp Cloud API using golang. To install, hit
```
go get github.com/Rfluid/whatsapp-cloud-api
```

# Bootstrap

To use an API you first need to bootstrap your API. You will need the bootstrap module to do so. The task is as simple as creating a `WhatsAppAPI` instance and setting your variables.
```
// Define your package
package packageName

// Import what you need

var Api bootstrap_model.WhatsAppAPI

func main() {
    Api = bootstrap_service.GenerateWhatsAppAPI(
        accessToken, // Your meta API access token.
        version, // Optional pointer to a version string. Default is v19.0.
        customMainURL, // Optional pointer to string that if provided will replace https://graph.facebook.com/{version} in the requests with the pointed string.
        customWABAIdURL, // Optional pointer to string that if provided will replace https://graph.facebook.com/{version}/{waba_id} in the requests with the pointed string.
    )

    // To send json through your API you might use
    Api.SetJSONHeaders()

    // And to send forms
    Api.SetFormHeaders()

    // Now you are ready to send messages as
    resp, err := message_service.Send(api, msg) // Where your msg is an Object like message_model.Message.
    // Sending messages only require the json headers.
}
```

# Modules

All the logic is condensed in the service level. Structs are defined to validate the object one may send to WhatsApp API via service args.

You can check the functionality of each module by comments but I believe the module names are very straight forward.
