package remotekeyvalue

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	// "golang.org/x/time/rate"
)

type apiClientOpt struct {
	uri                  string
	api_key_header_name  string
	api_key_header_value string
	method               string
	path                 string
	timeout              int
}

type api_client struct {
	http_client *http.Client
	uri         string
	method      string
	headers     map[string]string
	path        string
}

// Make a new api client for RESTful calls
func NewAPIClient(opt *apiClientOpt) (*api_client, error) {
	if opt.uri == "" {
		return nil, errors.New("uri must be set to construct an API client")
	}

	/* Remove any trailing slashes since we will append
	   to this URL with our own root-prefixed location */
	if strings.HasSuffix(opt.uri, "/") {
		opt.uri = opt.uri[:len(opt.uri)-1]
	}

	if opt.timeout <= 0 {
		opt.timeout = 100
	}

	headers := make(map[string]string)

	if opt.api_key_header_name != "" && opt.api_key_header_value != "" {
		headers[opt.api_key_header_name] = opt.api_key_header_value
	}

	tlsConfig := &tls.Config{
		/* Disable TLS verification if requested */
		InsecureSkipVerify: false,
		Renegotiation:      tls.RenegotiateOnceAsClient,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
		Proxy:           http.ProxyFromEnvironment,
	}

	client := api_client{
		http_client: &http.Client{
			Timeout:   time.Second * time.Duration(opt.timeout),
			Transport: tr,
		},
		uri:     opt.uri,
		method:  opt.method,
		headers: headers,
		path:    opt.path,
	}

	return &client, nil
}

/*
Helper function that handles sending/receiving and handling

	of HTTP data in and out.
*/
func (client *api_client) send_request() ([]byte, error) {
	full_uri := client.uri + client.path
	var req *http.Request
	var err error

	if client.path == "" {
		return nil, errors.New("path must be set to sent HTTP request.")
	}

	if client.method == "" {
		client.method = "GET"
	}

	req, err = http.NewRequest(client.method, full_uri, nil)

	if err != nil {
		log.Fatal(err)
		return make([]byte, 0), err
	}

	/* Allow for tokens or other pre-created secrets */
	if len(client.headers) > 0 {
		for n, v := range client.headers {
			req.Header.Set(n, v)
		}
	}

	resp, err := client.http_client.Do(req)

	if err != nil {
		return make([]byte, 0), err
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err2 != nil {
		return make([]byte, 0), err2
	}

	body := bodyBytes

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return body, errors.New(fmt.Sprintf("Unexpected response code '%d': %s", resp.StatusCode, body))
	}

	return body, nil

}
