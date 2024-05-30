package repository

import "time"

type GetDrinkOutput struct {
	Id            int
	DrinkPublicId string
	Name          string
	Stock         int
}

type InsertDrinkInput struct {
	DrinkPublicId string
	Name          string
	Stock         int
	CreatedAt     time.Time
	CreatedBy     string
}

type UpdateDrinkInput struct {
	DrinkId   int
	Name      string
	Stock     int
	UpdatedAt time.Time
	UpdatedBy string
}

type DeleteDrinkInput struct {
	DrinkId   int
	DeletedAt time.Time
	DeletedBy string
}
