package utils

import "encoding/json"

type PolicyPlannerAPIResponse struct {
	Success bool                    `json:"success"`
	Data    PolicyPlannerReportData `json:"data"`
	Message string                  `json:"message,omitempty"`
	Error   string                  `json:"error,omitempty"`
	Raw     map[string]interface{}  `json:"-"`
}

type PolicyPlannerReportData struct {
	ReportPDFURL   string                 `json:"report_pdf_url"`
	ReportPDFPath  string                 `json:"report_pdf_path"`
	Score          PolicyPlannerScore     `json:"score"`
	DatabaseStatus PolicyPlannerDBStatus  `json:"database_status"`
	Customer       map[string]interface{} `json:"customer,omitempty"`
	Summary        map[string]interface{} `json:"summary,omitempty"`
}

type PolicyPlannerScore struct {
	Value int    `json:"value"`
	Label string `json:"label"`
}

type PolicyPlannerDBStatus struct {
	ID            int    `json:"id"`
	Saved         bool   `json:"saved"`
	ReportPDFURL  string `json:"report_pdf_url"`
	ReportPDFPath string `json:"report_pdf_path"`
}

func ParsePolicyPlannerAPIResponse(body []byte) (*PolicyPlannerAPIResponse, error) {
	var apiResp PolicyPlannerAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	_ = json.Unmarshal(body, &raw)
	apiResp.Raw = raw

	if apiResp.Data.ReportPDFURL == "" {
		apiResp.Data.ReportPDFURL = apiResp.Data.DatabaseStatus.ReportPDFURL
	}
	if apiResp.Data.ReportPDFPath == "" {
		apiResp.Data.ReportPDFPath = apiResp.Data.DatabaseStatus.ReportPDFPath
	}

	return &apiResp, nil
}
