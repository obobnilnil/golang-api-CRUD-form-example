package service

import (
	"errors"
	"log"
	"mars_git/model"
	"mars_git/repository"
	"mars_git/utility"
)

type ServicePort interface {
	LeaderCreateFormService(leaderCreateForm model.LeaderCreateForm) error
	SubmitFormService(submitDetail model.SubmitForm) error
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewServiceAdapter(r repository.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) LeaderCreateFormService(leaderCreateForm model.LeaderCreateForm) error {
	if err := utility.ValidateLeaderCreateForm(leaderCreateForm); err != nil {
		log.Println(err)
		return err
	}
	roleValid, departmentID, err := s.r.RoleValidateRepository(leaderCreateForm.Leader_id)
	if err != nil {
		log.Println("Error validating role:", err)
		return err
	}
	if !roleValid {
		return errors.New("user is not a leader")
	}
	err = s.r.LeaderCreateFormRepository(leaderCreateForm, departmentID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *serviceAdapter) SubmitFormService(submitDetail model.SubmitForm) error {
	if err := utility.ValidateSubmitForm(submitDetail); err != nil {
		log.Println(err)
		return err
	}
	formValid, formID, err := s.r.FormValidateRepository(submitDetail.RefID)
	if err != nil {
		log.Println("Error validating form:", err)
		return err
	}
	if !formValid {
		return errors.New("form does not match")
	}
	err = s.r.SubmitFormRepository(submitDetail, formID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
