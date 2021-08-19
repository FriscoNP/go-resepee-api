package repository

import (
	"encoding/json"
	"errors"
	"go-resepee-api/entity"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	emailValidationURL = "https://emailvalidation.abstractapi.com/v1/"
)

type EmailValidationResponse struct {
	Email          string `json:"email"`
	Autocorrect    string `json:"autocorrect"`
	Deliverability string `json:"deliverability"`
	QualityScore   string `json:"quality_score"`
	IsValidFormat  struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_valid_format"`
	IsFreeEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_free_email"`
	IsDisposableEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_disposable_email"`
	IsRoleEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_role_email"`
	IsCatchallEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_catchall_email"`
	IsMxFound struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_mx_found"`
	IsSMTPValid struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_smtp_valid"`
}

type AbstractApiRepository struct {
	httpClient http.Client
	apiKey     string
}

type AbstractApiRepositoryInterface interface {
	ValidateEmail(email string) (res entity.AbstractEmailValidation, err error)
}

func NewAbstractApiRepository(httpClient http.Client, apiKey string) AbstractApiRepositoryInterface {
	return &AbstractApiRepository{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

func (repo *AbstractApiRepository) ToEmailValidationEntity(rec *EmailValidationResponse) entity.AbstractEmailValidation {
	return entity.AbstractEmailValidation{
		Email:             rec.Email,
		QualityScore:      rec.QualityScore,
		IsValidFormat:     rec.IsValidFormat.Value,
		IsFreeEmail:       rec.IsFreeEmail.Value,
		IsDisposableEmail: rec.IsDisposableEmail.Value,
		IsRoleEmail:       rec.IsRoleEmail.Value,
		IsCatchallEmail:   rec.IsCatchallEmail.Value,
		IsMxFound:         rec.IsMxFound.Value,
		IsSMTPValid:       rec.IsSMTPValid.Value,
	}
}

func (repo *AbstractApiRepository) ValidateEmail(email string) (res entity.AbstractEmailValidation, err error) {
	if repo.apiKey == "" {
		return res, errors.New("invalid api key")
	}

	fullUrl := emailValidationURL + "?api_key=" + repo.apiKey + "&email=" + email
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}
	defer resp.Body.Close()

	data := EmailValidationResponse{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Warn(err.Error())
		return res, err
	}

	return repo.ToEmailValidationEntity(&data), err
}
