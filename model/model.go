package model

type RequestFormHeader struct {
	FormHeaderID string `json:"form_header_id"`
	FormOwner    string `json:"form_owner"`
	FormType     string `json:"form_type"`
	FormSheetOf  string `json:"form_sheet_of"`
	FormHeader   string `json:"form_header"`
}

type RequestFormDetail struct {
	FormDetailID             string `json:"form_detail_id"`
	FormHeaderID             string `json:"form_header_id"`
	FormDetailInspectionWhat string `json:"form_detail_inspection_what"`
	FormDetailInspectionHow  string `json:"form_detail_inspection_how"`
	FormDetailSTD            string `json:"form_detail_std"`
	FormDetailResult         bool   `json:"form_detail_result"`
	FormDetailComment        string `json:"form_detail_comment"`
	FormDetailEvidence       string `json:"form_detail_evidence"`
}

type LeaderCreateForm struct {
	Leader_id   string            `json:"leader_id"`
	Sheet_of    string            `json:"sheet_of"`
	Header      string            `json:"form_header"`
	Inspections []InspectionPoint `json:"inspection_points"`
}

type InspectionPoint struct {
	What string      `json:"what"`
	Hows []HowDetail `json:"hows"`
}

type HowDetail struct {
	How string `json:"how"`
	Std string `json:"std"`
}

type SubmitForm struct {
	RefID            string             `json:"refID"`
	Creator          string             `json:"creator"`
	Timestamp        string             `json:"timestamp"`
	Line             string             `json:"line"`
	InspectionPoints []InspectionPoint2 `json:"inspection_points"`
}

type InspectionPoint2 struct {
	What string       `json:"what"`
	Hows []HowDetail2 `json:"hows"`
}

type HowDetail2 struct {
	How      string `json:"how"`
	Std      string `json:"std"`
	Result   bool   `json:"result"`
	Comment  string `json:"comment"`
	Evidence string `json:"evidence"`
}
