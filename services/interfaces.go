package services

import "context"

type ServiceInterface interface {
	GetDrinks(ctx context.Context) (output []GetDrinksOutput, err error)

	InsertDrink(ctx context.Context, input CreateDrinkInput) (output CreateDrinkOutput, err error)

	UpdateDrink(ctx context.Context, input UpdateDrinkInput) (output UpdateDrinkOutput, err error)

	DeleteDrink(ctx context.Context, input DeleteDrinkInput) error
}
