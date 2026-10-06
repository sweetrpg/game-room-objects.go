package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestLoanRoundTripLinkedBorrower(t *testing.T) {
	lentAt := time.Date(2026, 9, 21, 12, 0, 0, 0, time.UTC)
	borrowerUserID := "user-2"
	loan := Loan{
		ID:             "loan-1",
		LenderUserID:   "user-1",
		VolumeID:       "vol-1",
		BorrowerUserID: &borrowerUserID,
		BorrowerName:   "Alex",
		Status:         LoanStatusLent,
		LentAt:         lentAt,
	}

	data, err := json.Marshal(loan)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Loan
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.ID != loan.ID || got.LenderUserID != loan.LenderUserID || got.VolumeID != loan.VolumeID {
		t.Fatalf("round-trip mismatch: got %+v, want %+v", got, loan)
	}
	if got.BorrowerUserID == nil || *got.BorrowerUserID != borrowerUserID {
		t.Fatalf("BorrowerUserID did not round-trip: got %v, want %s", got.BorrowerUserID, borrowerUserID)
	}
	if got.BorrowerName != "Alex" {
		t.Fatalf("BorrowerName = %s, want Alex", got.BorrowerName)
	}
	if got.Status != LoanStatusLent {
		t.Fatalf("Status = %s, want %s", got.Status, LoanStatusLent)
	}
	if !got.LentAt.Equal(lentAt) {
		t.Fatalf("LentAt did not round-trip: got %v, want %v", got.LentAt, lentAt)
	}
	if got.ReturnedAt != nil {
		t.Fatalf("ReturnedAt = %v, want nil", got.ReturnedAt)
	}
}

func TestLoanRoundTripFreeFormBorrower(t *testing.T) {
	lentAt := time.Date(2026, 9, 20, 9, 30, 0, 0, time.UTC)
	returnedAt := time.Date(2026, 9, 21, 18, 0, 0, 0, time.UTC)
	loan := Loan{
		ID:           "loan-2",
		LenderUserID: "user-1",
		VolumeID:     "vol-2",
		BorrowerName: "Sam, from game night",
		Status:       LoanStatusReturned,
		LentAt:       lentAt,
		ReturnedAt:   &returnedAt,
	}

	data, err := json.Marshal(loan)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	var got Loan
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	if got.BorrowerUserID != nil {
		t.Fatalf("BorrowerUserID = %v, want nil", got.BorrowerUserID)
	}
	if got.BorrowerName != "Sam, from game night" {
		t.Fatalf("BorrowerName = %s, want %q", got.BorrowerName, "Sam, from game night")
	}
	if got.Status != LoanStatusReturned {
		t.Fatalf("Status = %s, want %s", got.Status, LoanStatusReturned)
	}
	if got.ReturnedAt == nil || !got.ReturnedAt.Equal(returnedAt) {
		t.Fatalf("ReturnedAt did not round-trip: got %v, want %v", got.ReturnedAt, returnedAt)
	}
}

func TestNewLoanDefaultsLent(t *testing.T) {
	loan := NewLoan("loan-3", "user-1", "vol-3", nil, "Jordan")
	if loan.Status != LoanStatusLent {
		t.Errorf("Status = %s, want %s", loan.Status, LoanStatusLent)
	}
	if loan.BorrowerUserID != nil {
		t.Errorf("BorrowerUserID = %v, want nil", loan.BorrowerUserID)
	}
	if loan.ReturnedAt != nil {
		t.Errorf("ReturnedAt = %v, want nil", loan.ReturnedAt)
	}
}
