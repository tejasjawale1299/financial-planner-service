package FinancialPlanner

import (
	dao "FinancialPlannerGo/app/domain/dao/financialPlanner"
	repo "FinancialPlannerGo/app/repositories/FinancialPlanner"
	"FinancialPlannerGo/app/utils"
	"FinancialPlannerGo/config"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FinancialPlannerService interface {
	GenerateFinancialReport(ctx context.Context, userID uint, req utils.FinancialPlannerRequest) (*dao.FinancialPlannerReport, *utils.PolicyPlannerAPIResponse, error)
	GetReportByID(id uint, userID uint) (*dao.FinancialPlannerReport, error)
}

type financialPlannerService struct {
	repo       repo.FinancialPlannerRepository
	httpClient *http.Client
	reportURL  string
}

func NewFinancialPlannerService(financialPlannerRepo repo.FinancialPlannerRepository) FinancialPlannerService {
	timeoutSeconds := config.GetEnvAsInt("POLICY_PLANNER_TIMEOUT_SECONDS", 30)

	return &financialPlannerService{
		repo: financialPlannerRepo,
		httpClient: &http.Client{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
		reportURL: config.GetEnv(
			"POLICY_PLANNER_REPORT_URL",
			"https://api.policyplanner.com/individual/api/report",
		),
	}
}

func (s *financialPlannerService) GenerateFinancialReport(
	ctx context.Context,
	userID uint,
	req utils.FinancialPlannerRequest,
) (*dao.FinancialPlannerReport, *utils.PolicyPlannerAPIResponse, error) {

	requestBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	report := &dao.FinancialPlannerReport{
		UserID:         userID,
		FullName:       req.FullName,
		Contact:        req.Contact,
		Email:          req.Email,
		DOB:            req.DOB,
		City:           req.City,
		Status:         "report_requested",
		RequestPayload: requestBytes,
	}

	if err := s.repo.Save(report); err != nil {
		return report, nil, err
	}

	apiResp, responseBytes, err := s.callPolicyPlannerReportAPI(ctx, requestBytes)
	if err != nil {
		report.Status = "report_failed"
		report.ErrorMessage = err.Error()
		report.ResponsePayload = responseBytes

		_ = s.repo.Update(report)
		return report, nil, err
	}

	report.ResponsePayload = responseBytes
	report.ReportPDFURL = apiResp.Data.ReportPDFURL
	report.ReportPDFPath = apiResp.Data.ReportPDFPath
	report.ScoreValue = apiResp.Data.Score.Value
	report.ScoreLabel = apiResp.Data.Score.Label

	if !apiResp.Success {
		report.Status = "report_failed"

		if apiResp.Message != "" {
			report.ErrorMessage = apiResp.Message
		} else if apiResp.Error != "" {
			report.ErrorMessage = apiResp.Error
		} else {
			report.ErrorMessage = "policy planner API returned unsuccessful response"
		}

		_ = s.repo.Update(report)
		return report, apiResp, errors.New(report.ErrorMessage)
	}

	report.Status = "report_generated"
	report.ErrorMessage = ""

	if err := s.repo.Update(report); err != nil {
		return report, apiResp, err
	}

	return report, apiResp, nil
}

func (s *financialPlannerService) callPolicyPlannerReportAPI(
	ctx context.Context,
	requestBytes []byte,
) (*utils.PolicyPlannerAPIResponse, []byte, error) {

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		s.reportURL,
		bytes.NewReader(requestBytes),
	)
	if err != nil {
		return nil, nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("User-Agent", "FinancialPlannerGo/1.0")

	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, body, fmt.Errorf(
			"policy planner API failed with status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	apiResp, err := utils.ParsePolicyPlannerAPIResponse(body)
	if err != nil {
		return nil, body, err
	}

	return apiResp, body, nil
}

func (s *financialPlannerService) GetReportByID(
	id uint,
	userID uint,
) (*dao.FinancialPlannerReport, error) {

	if id == 0 {
		return nil, errors.New("invalid report id")
	}

	return s.repo.GetByIDForUser(id, userID)
}
