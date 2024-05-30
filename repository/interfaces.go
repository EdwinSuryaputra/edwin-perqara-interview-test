package repository

import "context"

type RepositoryInterface interface {
	GetDrinks(ctx context.Context) (output []GetDrinkOutput, err error)

	GetDrinkByPublicId(ctx context.Context, publicId string) (output GetDrinkOutput, err error)

	InsertDrink(ctx context.Context, input InsertDrinkInput) error

	UpdateDrink(ctx context.Context, input UpdateDrinkInput) error

	DeleteDrink(ctx context.Context, input DeleteDrinkInput) error
}
