package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestContactGet(t *testing.T) {
	req := client.NewContactGetRequest()
	req.QueryParams().PageSize = 1000
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
