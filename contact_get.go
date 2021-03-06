package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewContactGetRequest() ContactGetRequest {
	r := ContactGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "CON"
	return r
}

type ContactGetRequest struct {
	AccountviewDataGetRequest
}

func (r *ContactGetRequest) NewResponseBody() *ContactGetResponseBody {
	return &ContactGetResponseBody{}
}

type ContactGetResponseBody struct {
	Contact []Contact `json:"CONTACT"`
}

func (r *ContactGetRequest) Do() (ContactGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
