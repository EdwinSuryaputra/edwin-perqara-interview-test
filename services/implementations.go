package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"
	"github.com/google/uuid"
)

func (s *Service) GetDrinks(ctx context.Context) (output []GetDrinksOutput, err error) {
	drinks, err := s.repository.GetDrinks(ctx)
	if err != nil {
		return nil, err
	}

	for _, drink := range drinks {
		output = append(output, GetDrinksOutput{DrinkPublicId: drink.DrinkPublicId, Name: drink.Name, Stock: drink.Stock})
	}

	return output, nil
}

func (s *Service) InsertDrink(ctx context.Context, input CreateDrinkInput) (output CreateDrinkOutput, err error) {
	publicId := uuid.NewString()

	err = s.repository.InsertDrink(ctx, repository.InsertDrinkInput{
		DrinkPublicId: publicId,
		Name:          input.Name,
		Stock:         input.Stock,
		CreatedAt:     time.Now(),
		CreatedBy:     "dummy",
	})
	if err != nil {
		return output, err
	}

	return CreateDrinkOutput{PublicId: publicId}, nil
}

func (s *Service) UpdateDrink(ctx context.Context, input UpdateDrinkInput) (output UpdateDrinkOutput, err error) {
	drink, err := s.repository.GetDrinkByPublicId(ctx, input.PublicId)
	if err != nil {
		if err == sql.ErrNoRows {
			return output, errors.New("drink is not exist")
		}

		return output, err
	}

	err = s.repository.UpdateDrink(ctx, repository.UpdateDrinkInput{
		DrinkId:   drink.Id,
		Name:      input.Name,
		Stock:     input.Stock,
		UpdatedAt: time.Now(),
		UpdatedBy: "dummy",
	})
	if err != nil {
		return output, err
	}

	return UpdateDrinkOutput{PublicId: drink.DrinkPublicId, Name: input.Name, Stock: input.Stock}, nil
}

func (s *Service) DeleteDrink(ctx context.Context, input DeleteDrinkInput) error {
	drink, err := s.repository.GetDrinkByPublicId(ctx, input.PublicId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("drink is not exist")
		}

		return err
	}

	err = s.repository.DeleteDrink(ctx, repository.DeleteDrinkInput{
		DrinkId:   drink.Id,
		DeletedAt: time.Now(),
		DeletedBy: "dummy",
	})
	if err != nil {
		return err
	}

	return nil
}
