package track

import (
	"encoding/json"
	"path"
	"github.com/aabizri/laposte"
	"net/http"
	"net/url"
	"strings"
)

type ResourceResponse struct{
	Data *Response `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type intermediateResponseRepresentation []*ResourceResponse

type BatchResponse map[string]*ResourceResponse

func (cl *Client) TrackBatch(codes []string) (BatchResponse, error) {
	// Create the query values
	values := url.Values{
		"codes": []string{strings.Join(codes,",")},
	}

	// Build the url
	dest := url.URL{
		Scheme: "https",
		Host: laposte.APIHost,
		Path: path.Join(root, version, "list"),
		RawQuery: values.Encode(),
	}

	// Create the request
	req, err := http.NewRequest("GET", dest.String(), nil)
	if err != nil {
		return nil, err
	}

	// Authenticate
	cl.auth(req)

	// Execute
	httpResp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return an ErrorResponse in case of error
	if httpResp.StatusCode != 200 {
		dec := json.NewDecoder(httpResp.Body)
		errResp := &ErrorResponse{}
		err = dec.Decode(errResp)
		if err != nil{
			return nil, err
		}
		return nil, errResp
	}

	// Parse intermediate representation
	dec := json.NewDecoder(httpResp.Body)
	intermediate := make(intermediateResponseRepresentation,0)
	err = dec.Decode(&intermediate)
	if err != nil {
		return nil, err
	}

	// Beautify
	br := make(BatchResponse, len(intermediate))
	for index, resource := range intermediate {
		br[codes[index]] = resource
	}
	intermediate = nil

	// We're done and return
	return br, nil
}