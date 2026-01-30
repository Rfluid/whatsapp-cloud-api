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

// Spams 
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func Spam(
	api bootstrap.WhatsAppAPI,
	data Message,
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
		return Spam(api, data)
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Spams message with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamWithCacheControll(
	api bootstrap.WhatsAppAPI,
	data Message,
	cacheControl MediaCacheControl,
) (Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIDURL, common.Messages),
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
		return SpamWithCacheControll(api, data, cacheControl)
	}

	var body Response

	err = json.NewDecoder(resp.Body).Decode(&body)

	return body, err
}

// Spams many messages.
//
// All messages are sent using parallelism.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamMany(
	api bootstrap.WhatsAppAPI,
	data []Message,
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

			response, err := Spam(api, msg)

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

// Spams many messages with media cache control.
//
// You may use this to control WhatsApp caching sending media via url.
//
// All messages are sent using parallelism.
//
// Spam service retries to send message until WhatsApp accepts if the request gets ignored
// by WhatsApp. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is not safe at all because of the lack of time limit and try limit.
func SpamManyWithCacheControll(
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
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

			response, err := SpamWithCacheControll(api, msg, cacheControl)

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

// Same as SpamMany but applies a callback for each result.
func SpamManyWithCallback(
	api bootstrap.WhatsAppAPI,
	data []Message,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := Spam(api, msg)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SpamWithCacheControll but applies a callback for each result.
func SpamManyWithCacheControllAndCallback(
	api bootstrap.WhatsAppAPI,
	data []Message,
	cacheControl MediaCacheControl,
	callback func(Message, Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg Message) {
			defer wg.Done()

			response, err := SpamWithCacheControll(api, msg, cacheControl)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
