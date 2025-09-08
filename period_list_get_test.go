package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPeriodListGet(t *testing.T) {
	req := client.NewPeriodListGetRequest()
	req.QueryParams().PageSize = 10
	req.QueryParams().Fields = "ADM_PER.START_DATE,ADM_PER.END_DATE,ADM_PER.INP_DATE,ADM_PER.INP_USR,ADM_PER.CNG_USR,ADM_PER.PER_NR,ADM_PER.REC_ID,ADM_PER.CNG_DATE,ADM_PER.CNG_NR"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

