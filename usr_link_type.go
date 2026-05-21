package accountviewnet

type UsrLink struct {
	LinkFile  string `json:"LINK_FILE" field_type:"C"`
	FileData  string `json:"FILE_DATA" field_type:"C"`
	DocDesc   string `json:"DOC_DESC" field_type:"C"`
	RowAction int    `json:"RowAction,omitempty" field_type:"N"`
	RowID     int    `json:"RowId,omitempty" field_type:"C"`

	fields *fields
}

func (usrLink *UsrLink) BusinessObject() string {
	return ""
}

func (usrLink *UsrLink) Table() string {
	return "USR_LINK"
}

func (usrLink *UsrLink) Fields() *fields {
	if usrLink.fields == nil {
		usrLink.fields = &fields{}
		usrLink.fields.Set(
			"RowAction",
			"LinkFile",
			"FileData",
			"DocDesc",
		)
	}

	return usrLink.fields
}

func (usrLink *UsrLink) Values() ([]interface{}, error) {
	return FieldsToValues(usrLink, *usrLink.Fields())
}
