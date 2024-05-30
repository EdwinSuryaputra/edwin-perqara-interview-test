package handler

import (
	"net/http"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/constant/response"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/generated"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
	"github.com/gin-gonic/gin"
)

// Get Storage Drinks godoc
//
//	@Summary		Get storage drinks
//	@Description	Get storage drinks
//	@Tags			Storage Drinks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.HandlerResponse[[]services.GetDrinksOutput]	"List of drinks"
//	@Failure		400	{object}	response.HandlerErrorResponse							"Bad request - Error code and message"
//	@Router			/api/v1/storage/drinks [GET]
func (s *Server) GetApiV1StorageDrinks(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := s.Service.GetDrinks(ctx)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	err = validateGetStorageDrinksResponse(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.HandlerRespondJsonSuccess(result))
}

// Insert a drink into storage godoc
//
//	@Summary		Insert drink into storage
//	@Description	Insert drink into storage
//	@Tags			Storage Drinks
//	@Accept			json
//	@Produce		json
//	@Param			requestBody	body		generated.PostApiV1StorageDrinkJSONRequestBody			true	"Insert drink request body"
//	@Success		200			{object}	response.HandlerResponse[services.CreateDrinkOutput]	"Inserted drink"
//	@Failure		400			{object}	response.HandlerErrorResponse							"Bad request - Error code and message"
//	@Router			/api/v1/storage/drink [POST]
func (s *Server) PostApiV1StorageDrink(c *gin.Context) {
	params := new(generated.PostApiV1StorageDrinkJSONRequestBody)
	if err := c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}
	if err := validatePostDrinkRequest(params); err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}

	ctx := c.Request.Context()
	result, err := s.Service.InsertDrink(ctx, services.CreateDrinkInput{
		Name:  params.Name,
		Stock: params.Stock,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	err = validatePostDrinkResponse(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.HandlerRespondJsonSuccess(result))
}

// Update a drink on storage godoc
//
//	@Summary		Update a drink on storage
//	@Description	Update a drink on storage
//	@Tags			Storage Drinks
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string													true	"Drink Id"
//	@Param			requestBody	body		generated.PutApiV1StorageDrinkIdJSONRequestBody			true	"Update drink request body"
//	@Success		200			{object}	response.HandlerResponse[services.UpdateDrinkOutput]	"Updated drink"
//	@Failure		400			{object}	response.HandlerErrorResponse							"Bad request - Error code and message"
//	@Router			/api/v1/storage/drink/{id} [PUT]
func (s *Server) PutApiV1StorageDrinkId(c *gin.Context, id generated.Id) {
	params := new(generated.PutApiV1StorageDrinkIdJSONRequestBody)
	if err := c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}
	if err := validatePutDrinkRequest(id, params); err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}

	ctx := c.Request.Context()
	result, err := s.Service.UpdateDrink(ctx, services.UpdateDrinkInput{
		PublicId: id,
		Name:     params.Name,
		Stock:    params.Stock,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	err = validatePutDrinkResponse(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.HandlerRespondJsonSuccess(result))
}

// Delete a drink on storage godoc
//
//	@Summary		Delete a drink on storage
//	@Description	Delete a drink on storage
//	@Tags			Storage Drinks
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string								true	"Drink Id"
//	@Success		200	{object}	response.HandlerResponse[string]	"Deleted drink"
//	@Failure		400	{object}	response.HandlerErrorResponse		"Bad request - Error code and message"
//	@Router			/api/v1/storage/drink/{id} [DELETE]
func (s *Server) DeleteApiV1StorageDrinkId(c *gin.Context, id generated.Id) {
	ctx := c.Request.Context()
	if err := validateDeleteDrinkRequest(id); err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}

	err := s.Service.DeleteDrink(ctx, services.DeleteDrinkInput{
		PublicId: id,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.HandlerRespondJsonSuccess("Success"))
}
