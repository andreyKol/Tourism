package users

import "Tourism/internal/domain"

type Patient struct {
	domain.User  `json:"user"`
	*Policy      `json:"policy"`
	*MedicalCard `json:"medical_card"`
}
