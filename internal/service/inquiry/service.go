package inquiry

import (
	"context"

	mic "tapera.integrasi/grpc/client/mitraintegrasi/v1"
)

type (
	Service interface {
		InquirySubs(ctx context.Context, parm *InqSubsParam) (*InqSubsResponse, error)
		// ConvertStringtoJson(parm *InquiryConfirmSubResponse) *InquiryConfirmSubResponseSuccessJson
	}

	TokenResult struct {
		OrganizationName string `json:"organization_name"`
		IssuedAt         string `json:"issued_at"`
		ExpiresIn        string `json:"expires_in"`
		ClientID         string `json:"client_id"`
		AccessToken      string `json:"access_token"`
		ApplicationName  string `json:"application_name"`
	}

	service struct {
		miClientMgr mic.GrpcClientManager
	}

	InqSubsParam struct {
		RefNo string `json:"reference_no"`
	}

	//InqSubsResponse struct
	InqSubsResponse struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		ResponseStatus      string `json:"response_status"`
		Data                Data
	}

	//Data struct
	Data struct {
		RefNo                string `json:"reference_no"`
		PaymentPoolID        string `json:"payment_pool_id"`
		BankBicCode          string `json:"payment_bank_bic_code"`
		NavDate              string `json:"nav_date"`
		InvstrFundUnitAcNo   string `json:"investor_fund_unit_ac_no"`
		InvstrFundUnitAcName string `json:"investor_fund_unit_ac_name"`
		Sid                  string `json:"sid"`
		Status               string `json:"status"`
		FundCode             string `json:"fund_code"`
		AmmountNominal       string `json:"amount_nominal"`
		Unit                 string `json:"unit"`
		Price                string `json:"price"`
	}

	//InqSubsResponseError struct
	InqSubsResponseError struct {
		ResponseCode        string `json:"response_code"`
		ResponseDescription string `json:"response_description"`
		Status              Status
	}

	//Status struct
	Status struct {
		Code string `json:"code"`
		Desc string `json:"desc"`
	}
)

func NewService(miClientMgr mic.GrpcClientManager) Service {
	return &service{miClientMgr: miClientMgr}
}
