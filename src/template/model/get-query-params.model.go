package template_model

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	common_service "github.com/Rfluid/whatsapp-cloud-api/src/common/service"
)

type GetQueryParams struct {
	Fields *[]TemplateFields // Fields to be returned.
	After  *string           // Used to paginate.
	Before *string           // Used to paginate.
}

func (qp *GetQueryParams) BuildQuery(
	req *http.Request,
) {
	var queryStringWg sync.WaitGroup
	queryThreadsCh := make(chan int)
	fieldsCh := make(chan bool)
	beforeCh := make(chan bool)
	afterCh := make(chan bool)

	queryCh := make(chan url.Values)

	fieldsStr := ""

	go func() {
		queryThreadsCh <- qp.calculateQueryThreads(fieldsCh, afterCh, beforeCh)
	}()

	go func() {
		common_service.PropagateQueryParamsMultithread(
			queryCh,
			<-queryThreadsCh+1,
			req,
		)
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()

		if <-fieldsCh {
			fieldsStr += string((*qp.Fields)[0])
			for _, field := range (*qp.Fields)[1:] {
				fieldsStr += fmt.Sprintf(",%s", field)
				(<-queryCh).Add("fields", fieldsStr)
			}
		}
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()
		if <-afterCh {
			(<-queryCh).Add("after", *qp.After)
		}
	}()

	queryStringWg.Add(1)
	go func() {
		defer queryStringWg.Done()
		if <-beforeCh {
			(<-queryCh).Add("before", *qp.Before)
		}
	}()

	queryStringWg.Wait()

	req.URL.RawQuery = (<-queryCh).Encode()
}

// Calculates the quantity of threads that changes query params.
func (qp *GetQueryParams) calculateQueryThreads(
	fieldsCh chan<- bool,
	afterCh chan<- bool,
	beforeCh chan<- bool,
) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	quantity := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		var verifyWg sync.WaitGroup

		value := qp.Fields != nil

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()

			if value {
				mu.Lock()
				quantity++
				mu.Unlock()
			}
		}()

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()
			fieldsCh <- value
		}()

		verifyWg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var verifyWg sync.WaitGroup

		value := qp.Before != nil

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()

			if value {
				mu.Lock()
				quantity++
				mu.Unlock()
			}
		}()

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()
			beforeCh <- value
		}()

		verifyWg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var verifyWg sync.WaitGroup

		value := qp.After != nil

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()

			if value {
				mu.Lock()
				quantity++
				mu.Unlock()
			}
		}()

		verifyWg.Add(1)
		go func() {
			defer verifyWg.Done()
			afterCh <- value
		}()

		verifyWg.Wait()
	}()

	wg.Wait()

	return quantity
}
