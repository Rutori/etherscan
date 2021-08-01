package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"etherscan_parse/etherscanio/entities"
	"github.com/pkg/errors"
)

const queryTimeout = 20 * time.Second

// API queries the etherscan API
type API struct {
	apiKey            string
	httpClient        *http.Client
	connectionControl *throttler
}

// NewAPI creates new API instance
func NewAPI(key string, rps int) *API {
	return &API{
		apiKey:            key,
		httpClient:        getHTTPClient(),
		connectionControl: newThrottler(rps, time.Second),
	}
}

// Query makes a request to api
func (c *API) Query(ctx context.Context, endpoint string) ([]byte, error) {
	c.connectionControl.allow()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	default:
	}

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

	// We can try and override rate limit error (that happens to unknown reasons)
	err = checkError(body)
	switch {
	case err == nil:

	case err.Error() == "Max rate limit reached":
		waitRandomTimeout(c.connectionControl.timeout)
		return c.Query(ctx, endpoint)

	default:
		return nil, err
	}

	return body, nil
}

// getHTTPClient creates a new http.Client without keepalives and reasonable timeout (unlike the default http client)
func getHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: queryTimeout,
	}
}

func checkError(body []byte) error {
	err := new(entities.Error)
	jsonErr := json.Unmarshal(body, &err)
	if jsonErr != nil {
		return nil
	}

	if err.Status == "0" {
		return err
	}

	return nil
}

// waitRandomTimeout takes max timeout t and sleeps a random amount of time between 0 and t
func waitRandomTimeout(t time.Duration) {
	time.Sleep(time.Duration(float32(t) * rand.Float32()))
}
