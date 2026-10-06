package models

import (
	"time"

	modelcore "github.com/sweetrpg/model-core.go/models"
)

// LoanStatus is the lifecycle state of a Loan.
type LoanStatus string

const (
	LoanStatusLent     LoanStatus = "lent"
	LoanStatusReturned LoanStatus = "returned"
)

// Loan records a catalog volume lent by one player to another. The borrower
// is either a platform-linked User (BorrowerUserID set) or a free-form name
// for a player not on the platform; BorrowerName is always populated, either
// the linked user's display name at lend time or the typed free-form text.
type Loan struct {
	ID             string     `bson:"_id" json:"id" jsonapi:"primary,loan"`
	LenderUserID   string     `bson:"lender_user_id" json:"lender_user_id" jsonapi:"relation,user"`
	VolumeID       string     `bson:"volume_id" json:"volume_id" jsonapi:"attr,volume_id"`
	BorrowerUserID *string    `bson:"borrower_user_id,omitempty" json:"borrower_user_id,omitempty" jsonapi:"attr,borrower_user_id"`
	BorrowerName   string     `bson:"borrower_name" json:"borrower_name" jsonapi:"attr,borrower_name"`
	Status         LoanStatus `bson:"status" json:"status" jsonapi:"attr,status"`
	LentAt         time.Time  `bson:"lent_at" json:"lent_at" jsonapi:"attr,lent_at"`
	ReturnedAt     *time.Time `bson:"returned_at,omitempty" json:"returned_at,omitempty" jsonapi:"attr,returned_at"`
	modelcore.Auditable
}

// NewLoan creates a loan in the lent state, lent now.
func NewLoan(id, lenderUserID, volumeID string, borrowerUserID *string, borrowerName string) Loan {
	return Loan{
		ID:             id,
		LenderUserID:   lenderUserID,
		VolumeID:       volumeID,
		BorrowerUserID: borrowerUserID,
		BorrowerName:   borrowerName,
		Status:         LoanStatusLent,
		LentAt:         time.Now().UTC(),
	}
}
