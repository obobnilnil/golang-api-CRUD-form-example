package repository

import (
	"database/sql"
	"fmt"
	"log"
	"mars_git/model"
)

type RepositoryPort interface {
	LeaderCreateFormRepository(leaderCreateForm model.LeaderCreateForm, departmentID string) error
	RoleValidateRepository(id string) (bool, string, error)
	FormValidateRepository(refID string) (bool, string, error)
	SubmitFormRepository(submitDetail model.SubmitForm, formID string) error
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) LeaderCreateFormRepository(leaderCreateForm model.LeaderCreateForm, departmentID string) error {

	var existingHeaderId string
	checkQuery := `
    SELECT form_header_id FROM form_header
    WHERE form_sheet_of = $1 AND form_header = $2
    `
	err := r.db.QueryRow(checkQuery, leaderCreateForm.Sheet_of, leaderCreateForm.Header).Scan(&existingHeaderId)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error checking existing form header:", err)
		return err
	}
	if existingHeaderId != "" {
		log.Println("A form with the same sheet and header already exists.")
		return fmt.Errorf("a form with the same sheet '%s' and header '%s' already exists", leaderCreateForm.Sheet_of, leaderCreateForm.Header)
	}

	var formHeaderId string
	queryHeader := `
    INSERT INTO form_header (form_owner, form_type, form_sheet_of, form_header, createtime, updatetime)
    VALUES ($1, 'Your Form Type Here', $2, $3, NOW(), NOW())
    RETURNING form_header_id
    `
	if err := r.db.QueryRow(queryHeader, departmentID, leaderCreateForm.Sheet_of, leaderCreateForm.Header).Scan(&formHeaderId); err != nil {
		log.Println("Error inserting form header:", err)
		return err
	}

	for _, inspection := range leaderCreateForm.Inspections {
		for _, howDetail := range inspection.Hows {
			queryDetail := `
            INSERT INTO form_detail (form_header_id, form_detail_inspection_what, form_detail_inspection_how, form_detail_std)
            VALUES ($1, $2, $3, $4)
            `
			if _, err := r.db.Exec(queryDetail, formHeaderId, inspection.What, howDetail.How, howDetail.Std); err != nil {
				log.Println("Error inserting form detail:", err)
				return err
			}
		}
	}
	return nil
}

func (r *repositoryAdapter) RoleValidateRepository(id string) (bool, string, error) {
	var roleName string
	var departmentId string
	query := `
    SELECT r.orgrl_name_en, r.orgrl_orgdp_id
    FROM organize_member m
    JOIN organize_role r ON m.orgmb_role = r.orgrl_id
    WHERE m.orgmb_id = $1 AND r.orgrl_name_en = 'Leader'
    `
	err := r.db.QueryRow(query, id).Scan(&roleName, &departmentId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", err
	}
	return true, departmentId, nil
}

func (r *repositoryAdapter) FormValidateRepository(refID string) (bool, string, error) {
	var formDetailID string
	// query := `
	// SELECT fd.form_detail_id
	// FROM form_header fh
	// JOIN form_detail fd ON fh.form_header_id = fd.form_header_id
	// WHERE fh.form_header_id = $1
	// `
	// query := `
	// SELECT fd.form_header_id
	// FROM form_header fh
	// JOIN form_detail fd ON fh.form_header_id = fd.form_header_id
	// WHERE fh.form_header_id = $1
	// `
	query := `
    SELECT form_header_id
    FROM form_header 
    WHERE form_header_id = $1
    `
	err := r.db.QueryRow(query, refID).Scan(&formDetailID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, "", nil
		}
		return false, "", err
	}
	return true, formDetailID, nil
}

func (r *repositoryAdapter) SubmitFormRepository(submitDetail model.SubmitForm, formID string) error {
	var formOptHeaderId string
	queryHeader := `
    INSERT INTO form_operation_header (form_ref_id, form_opt_creator, form_opt_timestamp, form_opt_line)
    VALUES ($1, $2, $3, $4)
    RETURNING form_opt_detail_id
    `
	if err := r.db.QueryRow(queryHeader, formID, submitDetail.Creator, submitDetail.Timestamp, submitDetail.Line).Scan(&formOptHeaderId); err != nil {
		log.Println("Error inserting form operation header:", err)
		return err
	}

	for _, inspection := range submitDetail.InspectionPoints {
		for _, how := range inspection.Hows {
			// queryDetail := `
			// INSERT INTO form_operation_detail (form_operation_id, form_detail_id, form_inspection_how, form_inspection_what, form_std, form_result, form_comment, form_evidence)
			// VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			// `
			// if _, err := r.db.Exec(queryDetail, formOptHeaderId, formID, how.How, inspection.What, how.Std, how.Result, how.Comment, how.Evidence); err != nil {
			queryDetail := `
            INSERT INTO form_operation_detail (form_header_id, form_inspection_how, form_inspection_what, form_std, form_result, form_comment, form_evidence)
            VALUES ($1, $2, $3, $4, $5, $6, $7)
            `
			if _, err := r.db.Exec(queryDetail, formID, how.How, inspection.What, how.Std, how.Result, how.Comment, how.Evidence); err != nil {
				log.Println("Error inserting form operation detail:", err)
				return err
			}
		}
	}
	return nil
}
