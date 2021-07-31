package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const queryTimeout = 20 * time.Second

type API struct {
	apiKey            string
	httpClient        *http.Client
	connectionControl *throttler
}

func NewAPI(key string, rps int) *API {
	return &API{
		apiKey:            key,
		httpClient:        getHTTPClient(),
		connectionControl: newThrottler(rps, time.Second),
	}
}

func (c *API) Query(ctx context.Context, endpoint string) ([]byte, error) {
	c.connectionControl.allow()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s&apikey=%s", endpoint, c.apiKey), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

func getHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: queryTimeout,
	}
}
