package bootstrap_service

import (
	"os"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
)

// As a lib, this service is not well suited, but we need it for testing.

func synchSetApiValue(
	wg *sync.WaitGroup,
	callback func(string) *bootstrap_model.CloudApi,
	dotEnvGetter string,
) *bootstrap_model.CloudApi {
	defer wg.Done()
	return callback(os.Getenv(dotEnvGetter))
}

func BootstrapApiValues(
	api *bootstrap_model.CloudApi,
) {
	var wg sync.WaitGroup

	wg.Add(1)
	go bootstrapHeaders(api, &wg)

	wg.Add(1)
	go bootstrapUrl(api, &wg)

	wg.Wait()

	// fmt.Println(api.PhoneIdURL)
	// fmt.Println(api.Headers)
}

func bootstrapUrl(
	api *bootstrap_model.CloudApi,
	globalBtpWg *sync.WaitGroup,
) {
	defer globalBtpWg.Done()

	var localBtp sync.WaitGroup

	localBtp.Add(1)
	go synchSetApiValue(
		&localBtp,
		api.SetMainURL,
		"META_WHATSAPP_CLOUD_API_PREFIX",
	)

	localBtp.Add(1)
	go synchSetApiValue(
		&localBtp,
		api.SetVersion,
		"META_WHATSAPP_CLOUD_API_VERSION",
	)

	localBtp.Add(1)
	go synchSetApiValue(
		&localBtp,
		api.SetWABAId,
		"META_WHATSAPP_CLOUD_API_WABA_ID",
	)

	localBtp.Wait()

	api.SetPhoneIdURL()
}

func bootstrapHeaders(
	api *bootstrap_model.CloudApi,
	globalBtpWg *sync.WaitGroup,
) {
	defer globalBtpWg.Done()
	ch := make(chan *bootstrap_model.CloudApi)

	go func() {
		ch <- api.SetAccessToken(os.Getenv("META_WHATSAPP_CLOUD_API_KEY"))
	}()

	<-ch
	api.SetHeaders()
}
