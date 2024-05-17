package users

import "Tourism/internal/domain"

type Doctor struct {
	ID              int64 `json:"id"`
	domain.User     `json:"user"`
	*Specialization `json:"specialization"`
	*Portfolio      `json:"portfolio"`
}
