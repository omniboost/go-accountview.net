package accountviewnet_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	accountviewnet "github.com/omniboost/go-accountview.net"
)

func TestDjPageTypeTest(t *testing.T) {
	subNr := "00005"
	invNr := "00005"
	o := accountviewnet.DjPage{
		DjCode:  "600",
		HdrDesc: "TEST",
		TrnDate: accountviewnet.Date{time.Date(1983, 4, 12, 0, 0, 0, 0, time.UTC)},
		Period:  4,
		SubNr:   &subNr,
		InvNr:   &invNr,
	}
	lines := []accountviewnet.DjLine{
		{
			AcctNr: "9999",
			Amount: 12.1,
			RecID:  "REC_ID",
		},
		{
			AcctNr: "9999",
			Amount: -12.1,
			RecID:  "REC_ID",
		},
	}
	o.Fields().Del("TrnDate")
	req, err := o.ToAccountviewDataPostRequest(client, lines)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
