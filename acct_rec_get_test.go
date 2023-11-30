package accountviewnet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAcctRecGet(t *testing.T) {
	req := client.NewAcctRecGetRequest()
	req.QueryParams().PageSize = 1000
	// req.QueryParams().Fields = "CONTACT.SUB_NR,CONTACT.DEB_NRV,CONTACT.DEB_NR,CONTACT.ADDRESS1,CONTACT.ADDR_LINE2,CONTACT.BOX_CITY,CONTACT.ADDR_LINE1,CONTACT.MC_MAILB,CONTACT.MAIL_BUS,CONTACT.CNT_CODE,CONTACT.COUNTRY,CONTACT.POST_CODE,CONTACT.ACCT_NR,CONTACT.SRC_CODE"
	// req.QueryParams().FilterControlSource1 = "ACCT_NAME"
	// req.QueryParams().FilterOperator1 = "Equal"
	// req.QueryParams().FilterValueType1 = "C"
	// req.QueryParams().FilterValue1 = "Omniboost"
	// req.QueryParams().SortFields = "ACCT_NR"
	// req.QueryParams().SortOrder = "descending"
	// req.QueryParams().FilterIsListOfValues1 = false
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
