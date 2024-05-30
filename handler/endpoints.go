package handler

import (
	"net/http"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/constant/response"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/generated"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetStorageDrinks(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := s.Service.GetDrinks(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.HandlerResponseJsonError(err.Error()))
		return
	}

	err = validateGetStorageDrinksResponse(result)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.HandlerResponseJsonError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.HandlerRespondJsonSuccess(result))
}

func (s *Server) PostStorageDrink(c *gin.Context) {
	params := new(generated.PostStorageDrinkJSONRequestBody)
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

func (s *Server) PutStorageDrinkId(c *gin.Context, id generated.Id) {
	params := new(generated.PutStorageDrinkIdJSONRequestBody)
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

func (s *Server) DeleteStorageDrinkId(c *gin.Context, id generated.Id) {
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
