package handler

import (
	"errors"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/generated"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
)

func validateGetStorageDrinksResponse(result []services.GetDrinksOutput) error {
	for _, res := range result {
		if res.DrinkPublicId == "" {
			return errors.New("public id must not empty")
		}
		if len(res.DrinkPublicId) != 36 {
			return errors.New("invalid uuid length")
		}

		if res.Name == "" {
			return errors.New("name must not empty")
		}
		if len(res.Name) < 1 && len(res.Name) > 50 {
			return errors.New("invalid name length")
		}

		if res.Stock < 0 {
			return errors.New("invalid stock")
		}
	}

	return nil
}

func validatePostDrinkRequest(params *generated.PostStorageDrinkJSONRequestBody) error {
	if params.Name == "" {
		return errors.New("name is required")
	}
	if len(params.Name) < 1 {
		return errors.New("the minimum length of name is 1")
	}
	if len(params.Name) > 50 {
		return errors.New("the maximum length of name is 50")
	}

	if params.Stock < 1 {
		return errors.New("stock must not empty")
	}

	return nil
}

func validatePostDrinkResponse(result services.CreateDrinkOutput) error {
	if result.PublicId == "" {
		return errors.New("public id must not empty")
	}
	if len(result.PublicId) != 36 {
		return errors.New("invalid uuid length")
	}

	if result.Name == "" {
		return errors.New("name must not empty")
	}
	if len(result.Name) < 1 && len(result.Name) > 50 {
		return errors.New("invalid name length")
	}

	if result.Stock < 0 {
		return errors.New("invalid stock")
	}

	return nil
}

func validatePutDrinkRequest(publicId string, params *generated.UpdateSpec) error {
	if publicId == "" {
		return errors.New("public id is required")
	}

	if len(params.Name) < 1 {
		return errors.New("the minimum length of name is 1")
	}
	if len(params.Name) > 50 {
		return errors.New("the maximum length of name is 50")
	}

	if params.Stock < 1 {
		return errors.New("stock must not empty")
	}

	return nil
}

func validatePutDrinkResponse(result services.UpdateDrinkOutput) error {
	if result.PublicId == "" {
		return errors.New("public id must not empty")
	}
	if len(result.PublicId) != 36 {
		return errors.New("invalid uuid length")
	}

	if result.Name == "" {
		return errors.New("name must not empty")
	}
	if len(result.Name) < 1 && len(result.Name) > 50 {
		return errors.New("invalid name length")
	}

	if result.Stock < 0 {
		return errors.New("invalid stock")
	}

	return nil
}

func validateDeleteDrinkRequest(publicId string) error {
	if publicId == "" {
		return errors.New("public id is required")
	}

	return nil
}
