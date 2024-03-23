package bootstrap_service

import (
	"os"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
)

func synchSetApiValue(
	wg *sync.WaitGroup,
	cbk func(string) *bootstrap_model.ApiBootstrap,
	dotEnvGetter string,
) *bootstrap_model.ApiBootstrap {
	defer wg.Done()
	return cbk(os.Getenv(dotEnvGetter))
}

func BootstrapApiValues(
	api *bootstrap_model.ApiBootstrap,
) {
	ch := make(chan *bootstrap_model.ApiBootstrap, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go synchSetApiValue(
		&wg,
		api.SetAccessToken,
		"META_WHATSAPP_CLOUD_API_KEY",
	)

	wg.Add(1)
	go func() {
		ch <- synchSetApiValue(
			&wg,
			api.SetPrefix,
			"META_WHATSAPP_CLOUD_API_PREFIX",
		)
	}()

	wg.Add(1)
	go func() {
		ch <- synchSetApiValue(
			&wg,
			api.SetVersion,
			"META_WHATSAPP_CLOUD_API_VERSION",
		)
	}()

	<-ch
	<-ch
	api.SetMainURL()

	wg.Wait()
}
