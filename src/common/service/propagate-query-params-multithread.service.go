package common_service

import (
	"net/http"
	"net/url"
)

// Propagates the query params through a channel.
func PropagateQueryParamsMultithread(
	ch chan<- url.Values, // Channel that will be user to propagate query params.
	quantity int, // Amount of times to propagate if.
	req *http.Request, // Request that query params will be mounted on top off.
) {
	query := req.URL.Query()

	for i := 0; i < quantity; i++ {
		go func() {
			ch <- query
		}()
	}
}
