package services

type GetDrinksOutput struct {
	DrinkPublicId string
	Name          string
	Stock         int
}

type CreateDrinkInput struct {
	Name  string
	Stock int
}
type CreateDrinkOutput struct {
	PublicId string `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
}

type UpdateDrinkInput struct {
	PublicId string
	Name     string
	Stock    int
}
type UpdateDrinkOutput struct {
	PublicId string `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
}

type DeleteDrinkInput struct {
	PublicId string
}
