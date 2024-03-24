package message_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp/src/bootstrap/model"
	message_model "github.com/Rfluid/whatsapp/src/message/model"
)

func Send(
	api bootstrap_model.CloudApi,
	data message_model.Message,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/messages", api.MainURL),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

func SendWithCacheControll(
	api bootstrap_model.CloudApi,
	data message_model.Message,
	cacheControl message_model.MediaCacheControl,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/messages", api.MainURL),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.Headers

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	return body, nil
}

func SendMany(
	api bootstrap_model.CloudApi,
	data []message_model.Message,
) ([](message_model.Response), []error) {
	var mu sync.Mutex
	var errMu sync.Mutex
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			if err == nil {
				mu.Lock()
				responses = append(responses, response)
				mu.Unlock()
			} else {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
			}
		}(msg)
	}
	wg.Wait()

	return responses, errs
}

func SendManyWithCacheControll(
	api bootstrap_model.CloudApi,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
) ([](message_model.Response), []error) {
	var mu sync.Mutex
	var errMu sync.Mutex
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			if err == nil {
				mu.Lock()
				responses = append(responses, response)
				mu.Unlock()
			} else {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
			}
		}(msg)
	}
	wg.Wait()

	return responses, errs
}

func SendManyWithCallback(
	api bootstrap_model.CloudApi,
	data []message_model.Message,
	cbk func(message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := Send(api, msg)

			cbk(response, err)
		}(msg)
	}
	wg.Wait()
}

func SendManyWithCacheControllAndCallback(
	api bootstrap_model.CloudApi,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
	cbk func(message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SendWithCacheControll(api, msg, cacheControl)

			cbk(response, err)
		}(msg)
	}
	wg.Wait()
}

func SendInitialTemplate(api bootstrap_model.CloudApi) {
	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/messages", api.MainURL), // 5513991050558
		bytes.NewBufferString("{ \"messaging_product\": \"whatsapp\", \"to\": \"5591984288778\", \"type\": \"template\", \"template\": { \"name\": \"hello_world\", \"language\": { \"code\": \"en_US\" } } }"),
	)
	req.Header = api.Headers

	_, err := api.Client.Do(req)
	if err != nil {
		fmt.Println("an error occurred")
	}
}
