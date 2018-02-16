package track

import (
	"path"
	"github.com/aabizri/laposte"
	"net/http"
	"encoding/json"
)

func (cl *Client) Track(id string) (*Response, error) {
	endpoint := "https://" + path.Join(laposte.APIHost, root, version, id)

	// Create the request
	req, err := http.NewRequest("GET", endpoint, nil)
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

	// Parse
	dec := json.NewDecoder(httpResp.Body)
	resp := &Response{}
	err = dec.Decode(resp)
	if err != nil {
		return nil, err
	}

	// Return an ErrorResponse in case of error
	if httpResp.StatusCode != 200 {
		return resp, ErrorResponse{
			Code: ErrorCode(resp.Code),
			Message: resp.Message,
		}
	}

	// Else, we're done and we return the response directly
	return resp, nil
}
