package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDjPageGet(t *testing.T) {
	req := client.NewDjPageGetRequest()
	req.QueryParams().PageSize = 20
	req.QueryParams().FilterControlSource1 = "DJ_CODE"
	req.QueryParams().FilterOperator1 = "Equal"
	req.QueryParams().FilterValueType1 = "C"
	req.QueryParams().FilterValue1 = "MEM"
	// req.QueryParams().FilterControlSource1 = "TRN_DATE"
	// req.QueryParams().FilterOperator1 = "Contains"
	// req.QueryParams().FilterValueType1 = "D"
	// req.QueryParams().FilterValue1 = "2022-01-05"
	// req.QueryParams().FilterIsListOfValues1 = false
	req.QueryParams().FilterControlSource2 = "TRN_DATE"
	req.QueryParams().FilterOperator2 = "Equal"
	req.QueryParams().FilterValueType2 = "C"
	req.QueryParams().FilterValue2 = "8000"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
