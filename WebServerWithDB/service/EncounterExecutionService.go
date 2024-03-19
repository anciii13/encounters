package service

import (
	"database-example/model"
	"database-example/repo"
	"time"
)

type EncounterExecutionService struct {
	EncounterExecutionRepo *repo.EncounterExecutionRepository
}

func (service *EncounterExecutionService) GetExecutionByUser(userID int) (*model.EncounterExecution, error) {
	encounter, err := service.EncounterExecutionRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	if encounter.IsCompleted {
		return nil, nil
	}
	return &encounter, nil
}

func (service *EncounterExecutionService) CompleteEncounter(userID int) (*model.EncounterExecution, error) {
	encounter, err := service.EncounterExecutionRepo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}

	encounter.CompletionTime = time.Now()
	encounter.IsCompleted = true

	// Neki update XP za korisnika

	err = service.EncounterExecutionRepo.Update(&encounter)
	if err != nil {
		return nil, err
	}
	return &encounter, nil
}

func (service *EncounterExecutionService) CreateEncounter(encounter *model.EncounterExecution) error {
	err := service.EncounterExecutionRepo.Create(encounter)
	if err != nil {
		return err
	}
	return nil
}

func (service *EncounterExecutionService) UpdateEncounter(id int, encounter *model.EncounterExecution) error {
	encounter.ID = id
	err := service.EncounterExecutionRepo.Update(encounter)
	if err != nil {
		return err
	}
	return nil
}
