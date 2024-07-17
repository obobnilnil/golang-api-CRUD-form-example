package utility

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mars_git/model"
	"strings"
)

func CountTables(db *sql.DB) {
	var count int
	query := `SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public'`
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatalf("Failed to query table count: %s", err.Error())
	}
	fmt.Printf("There are %d tables in the database.\n", count)
}

func ValidateLeaderCreateForm(form model.LeaderCreateForm) error {
	// Validate LeaderCreateForm fields are not empty
	if strings.TrimSpace(form.Leader_id) == "" {
		return errors.New("leader_id must not be empty")
	}
	if strings.TrimSpace(form.Sheet_of) == "" {
		return errors.New("sheet_of must not be empty")
	}
	if strings.TrimSpace(form.Header) == "" {
		return errors.New("form_header must not be empty")
	}
	if len(form.Inspections) == 0 {
		return errors.New("at least one inspection must be provided")
	}

	// Validate each InspectionPoint
	for _, inspection := range form.Inspections {
		if strings.TrimSpace(inspection.What) == "" {
			return errors.New("what field in inspection points must not be empty")
		}
		if len(inspection.Hows) == 0 {
			return errors.New("at least one how must be provided for each inspection")
		}
		// Validate each HowDetail
		for _, how := range inspection.Hows {
			if strings.TrimSpace(how.How) == "" {
				return errors.New("how field in how details must not be empty")
			}
			if strings.TrimSpace(how.Std) == "" {
				return errors.New("std field in how details must not be empty")
			}
		}
	}
	return nil
}

func ValidateSubmitForm(form model.SubmitForm) error {
	// Validate SubmitForm fields are not empty
	if strings.TrimSpace(form.RefID) == "" {
		return errors.New("refID must not be empty")
	}
	if strings.TrimSpace(form.Creator) == "" {
		return errors.New("creator must not be empty")
	}
	if strings.TrimSpace(form.Timestamp) == "" {
		return errors.New("timestamp must not be empty")
	}
	if strings.TrimSpace(form.Line) == "" {
		return errors.New("line must not be empty")
	}
	if len(form.InspectionPoints) == 0 {
		return errors.New("at least one inspection point must be provided")
	}

	// Validate each InspectionPoint
	for _, inspection := range form.InspectionPoints {
		if strings.TrimSpace(inspection.What) == "" {
			return errors.New("what field in inspection points must not be empty")
		}
		if len(inspection.Hows) == 0 {
			return errors.New("at least one how must be provided for each inspection")
		}
		// Validate each HowDetail
		for _, how := range inspection.Hows {
			if strings.TrimSpace(how.How) == "" {
				return errors.New("how field in how details must not be empty")
			}
			if strings.TrimSpace(how.Std) == "" {
				return errors.New("std field in how details must not be empty")
			}
			// No need to check bool for emptiness but ensure strings are not empty
			if strings.TrimSpace(how.Comment) == "" {
				return errors.New("comment field must not be empty")
			}
			if strings.TrimSpace(how.Evidence) == "" {
				return errors.New("evidence field must not be empty")
			}
		}
	}
	return nil
}
