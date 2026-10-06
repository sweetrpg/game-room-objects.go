package vo

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/vo"
)

// Loan value object.
// This value object is a serializable representation of the Loan model.
type LoanVO struct {
	ID             string     `json:"id" jsonapi:"primary,loan"`
	LenderUserID   string     `json:"lender_user_id" jsonapi:"relation,user"`
	VolumeID       string     `json:"volume_id" jsonapi:"attr,volume_id"`
	BorrowerUserID *string    `json:"borrower_user_id,omitempty" jsonapi:"attr,borrower_user_id"`
	BorrowerName   string     `json:"borrower_name" jsonapi:"attr,borrower_name"`
	Status         string     `json:"status" jsonapi:"attr,status"`
	LentAt         time.Time  `json:"lent_at" jsonapi:"attr,lent_at"`
	ReturnedAt     *time.Time `json:"returned_at,omitempty" jsonapi:"attr,returned_at"`
	modelcore.AuditableVO
}
