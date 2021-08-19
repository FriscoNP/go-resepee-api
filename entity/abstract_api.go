package entity

type AbstractEmailValidation struct {
	Email             string `json:"email"`
	QualityScore      string `json:"quality_score"`
	IsValidFormat     bool   `json:"is_valid_format"`
	IsFreeEmail       bool   `json:"is_free_email"`
	IsDisposableEmail bool   `json:"is_disposable_email"`
	IsRoleEmail       bool   `json:"is_role_email"`
	IsCatchallEmail   bool   `json:"is_catchall_email"`
	IsMxFound         bool   `json:"is_mx_found"`
	IsSMTPValid       bool   `json:"is_smtp_valid"`
}
