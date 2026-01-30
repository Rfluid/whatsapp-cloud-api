package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Rfluid/whatsapp-cloud-api/src/bootstrap"
	"github.com/Rfluid/whatsapp-cloud-api/src/common"
)

// SafeSpams 
//
// SafeSpam service retries to send message at most maxTries times if the request gets ignored byu cloud API. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is safer than Spam service.
func SafeSpam(
	api bootstrap.WhatsAppAPI,
	data Message,
	maxTries int,
) (Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return Response{}, err
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// SafeSpams message with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// SafeSpam service retries to send message at most maxTries times if the request gets ignored byu cloud API. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is safer than Spam service.
func SafeSpamWithCacheControll(
	api bootstrap.WhatsAppAPI,
	data Message,
	cacheControl MediaCacheControl,
	maxTries int,
) (Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/messages", api.WABAIDURL),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respErr common.ErrorResponse
		if decodeErr := json.NewDecoder(resp.Body).Decode(&respErr); decodeErr != nil {
			err = decodeErr
		} else {
			err = &respErr
		}
		return Response{}, err
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// SafeSpams many messages.
//
// All messages are sent using parallelism.
//
// SafeSpam service retries to send message at most maxTries times if the request gets ignored byu cloud API. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is safer than Spam service.
func SafeSpamMany(
	api bootstrap.WhatsAppAPI,
	data []Message,
	maxTries int,
) ([](Response), []error) {
	respCh := make(chan Response)
	errCh := make(chan error)
	responses := []Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SafeSpam(api, msg, maxTries)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}

	go func() {
		wg.Wait()
		close(respCh)
		close(errCh)
	}()

	for response := range respCh {
		responses = append(responses, response)
	}

	for err := range errCh {
		errs = append(errs, err)
	}

	return responses, errs
}

// SafeSpams many messages with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// All messages are sent using parallelism.
//
// SafeSpam service retries to send message at most maxTries times if the request gets ignored byu cloud API. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is safer than Spam service.
func SafeSpamManyWithCacheControll(
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
	maxTries int,
) ([](Response), []error) {
	respCh := make(chan Response)
	errCh := make(chan error)
	responses := []Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SafeSpamWithCacheControll(api, msg, cacheControl, maxTries)

			if err == nil {
				respCh <- response
			} else {
				errCh <- err
			}
		}(msg)
	}

	go func() {
		wg.Wait()
		close(respCh)
		close(errCh)
	}()

	for response := range respCh {
		responses = append(responses, response)
	}
	for err := range errCh {
		errs = append(errs, err)
	}

	return responses, errs
}

// Same as SafeSpamMany but applies a callback for each result.
func SafeSpamManyWithCallback(
	api bootstrap.WhatsAppAPI,
	data []Message,
	maxTries int,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SafeSpam(api, msg, maxTries)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SafeSpamWithCacheControll but applies a callback for each result.
func SafeSpamManyWithCacheControllAndCallback(
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
	maxTries int,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SafeSpamWithCacheControll(api, msg, cacheControl, maxTries)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
