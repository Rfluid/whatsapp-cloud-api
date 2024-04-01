package message_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	bootstrap_model "github.com/Rfluid/whatsapp-cloud-api/src/bootstrap/model"
	common_enum "github.com/Rfluid/whatsapp-cloud-api/src/common/enum"
	message_model "github.com/Rfluid/whatsapp-cloud-api/src/message/model"
)

// SafeSpams message.
//
// SafeSpam service retries to send message at most maxTries times if the request gets ignored byu cloud API. It can happen if you are trying to send 100 messages
// to the same user. If you don't want to retry on fail, use send service.
//
// This is safer than Spam service.
func SafeSpam(
	api bootstrap_model.WhatsAppAPI,
	data message_model.Message,
	maxTries int,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", api.WABAIdURL, common_enum.Messages),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	resp, err := api.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return message_model.Response{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return message_model.Response{}, err
		}
		return message_model.Response{}, errors.New(string(errMsgBytes))
	}

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	newMaxTries := maxTries - 1
	if body.MessagingProduct.MessagingProduct == "" && newMaxTries > 0 {
		return SafeSpam(api, data, newMaxTries)
	}

	return body, nil
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
	api bootstrap_model.WhatsAppAPI,
	data message_model.Message,
	cacheControl message_model.MediaCacheControl,
	maxTries int,
) (message_model.Response, error) {
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/messages", api.WABAIdURL),
		bytes.NewBuffer(jsonData),
	)
	req.Header = api.JSONHeaders

	for key, value := range cacheControl.ToMap() {
		req.Header.Add(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return message_model.Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errInt map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errInt); err != nil {
			return message_model.Response{}, err
		}
		errMsgBytes, err := json.MarshalIndent(errInt, "", "    ")
		if err != nil {
			return message_model.Response{}, err
		}
		return message_model.Response{}, errors.New(string(errMsgBytes))
	}

	var body message_model.Response

	json.NewDecoder(resp.Body).Decode(&body)

	newMaxTries := maxTries - 1
	if body.MessagingProduct.MessagingProduct == "" && newMaxTries > 0 {
		return SafeSpamWithCacheControll(api, data, cacheControl, newMaxTries)
	}

	return body, nil
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
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	maxTries int,
) ([](message_model.Response), []error) {
	respCh := make(chan message_model.Response)
	errCh := make(chan error)
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
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
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
	maxTries int,
) ([](message_model.Response), []error) {
	respCh := make(chan message_model.Response)
	errCh := make(chan error)
	responses := []message_model.Response{}
	errs := []error{}
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
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
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	maxTries int,
	callback func(message_model.Message, message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SafeSpam(api, msg, maxTries)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}

// Same as SafeSpamWithCacheControll but applies a callback for each result.
func SafeSpamManyWithCacheControllAndCallback(
	api bootstrap_model.WhatsAppAPI,
	data []message_model.Message,
	cacheControl message_model.MediaCacheControl,
	maxTries int,
	callback func(message_model.Message, message_model.Response, error),
) {
	var wg sync.WaitGroup

	for _, msg := range data {
		wg.Add(1)

		go func(msg message_model.Message) {
			defer wg.Done()

			response, err := SafeSpamWithCacheControll(api, msg, cacheControl, maxTries)

			callback(msg, response, err)
		}(msg)
	}
	wg.Wait()
}
