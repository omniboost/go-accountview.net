package accountviewnet

import "github.com/omniboost/go-accountview.net/utils"

func (c *Client) NewPeriodListGetRequest() PeriodListGetRequest {
	r := PeriodListGetRequest{
		AccountviewDataGetRequest: c.NewAccountviewDataGetRequest(),
	}
	r.AccountviewDataGetRequest.QueryParams().BusinessObject = "PER"
	return r
}

type PeriodListGetRequest struct {
	AccountviewDataGetRequest
}

func (r *PeriodListGetRequest) NewResponseBody() *PeriodListGetResponseBody {
	return &PeriodListGetResponseBody{}
}

type PeriodListGetResponseBody struct {
	PERIOD []PERIOD `json:"ADM_PER"`
}

func (r *PeriodListGetRequest) Do() (PeriodListGetResponseBody, error) {
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

