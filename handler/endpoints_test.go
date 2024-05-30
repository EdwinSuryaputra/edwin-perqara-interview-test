package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/conf"
	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/constant/response"
	r "github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"
	s "github.com/EdwinSuryaputra/edwin-perqara-interview-test/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetStorageDrinks(t *testing.T) {
	conf.InitTest()

	mockService := s.NewMockServiceInterface(gomock.NewController(t))
	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	endpoint := Server{Service: mockService, Repository: mockRepo}

	t.Run("Success - Get Drinks", func(t *testing.T) {
		r := gin.New()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/storage/drinks", nil)
		w := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(w, r)
		c.Request = req

		mockResult := []s.GetDrinksOutput{
			{
				DrinkPublicId: "bade4ae2-0178-47fc-ad02-37648e882d3c",
				Name:          "Item 1",
				Stock:         1,
			},
			{
				DrinkPublicId: "9a36d04f-61ca-4762-80a6-cef6044180c0",
				Name:          "Item 2",
				Stock:         10,
			},
		}
		mockResultInJson, _ := json.Marshal(response.HandlerRespondJsonSuccess(mockResult))

		mockService.EXPECT().GetDrinks(c.Request.Context()).Return(mockResult, nil).Times(1)
		endpoint.GetApiV1StorageDrinks(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(mockResultInJson), w.Body.String())
	})

	t.Run("Failed - Error occurred during get drinks", func(t *testing.T) {
		r := gin.New()

		req := httptest.NewRequest(http.MethodGet, "/api/v1/storage/drinks", nil)
		w := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(w, r)
		c.Request = req

		mockService.EXPECT().GetDrinks(c.Request.Context()).Return(nil, errors.New("Some error")).Times(1)
		endpoint.GetApiV1StorageDrinks(c)

		mockResult, _ := json.Marshal(response.HandlerResponseJsonError("Some error"))

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		assert.Equal(t, string(mockResult), w.Body.String())
	})
}

func TestPostStorageDrink(t *testing.T) {
	conf.InitTest()

	mockService := s.NewMockServiceInterface(gomock.NewController(t))
	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	endpoint := Server{Service: mockService, Repository: mockRepo}

	t.Run("Success - Insert drink into storage", func(t *testing.T) {
		r := gin.New()

		requestBody := []byte(`{"name":"Item 1","stock": 1}`)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/storage/drink", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(w, r)
		c.Request = req

		mockService.EXPECT().InsertDrink(c.Request.Context(), s.CreateDrinkInput{Name: "Item 1", Stock: 1}).Return(s.CreateDrinkOutput{PublicId: "9a36d04f-61ca-4762-80a6-cef6044180c0", Name: "Item 1", Stock: 1}, nil).Times(1)
		endpoint.PostApiV1StorageDrink(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `{"data":{"id":"9a36d04f-61ca-4762-80a6-cef6044180c0","name":"Item 1","stock":1}}`, w.Body.String())
	})
}

func TestPutStorageDrink(t *testing.T) {
	conf.InitTest()

	mockService := s.NewMockServiceInterface(gomock.NewController(t))
	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	endpoint := Server{Service: mockService, Repository: mockRepo}

	t.Run("Success - Update drink on storage", func(t *testing.T) {
		r := gin.New()

		requestBody := []byte(`{"name":"Item 1","stock": 1}`)
		req := httptest.NewRequest(http.MethodPut, "/api/v1/storage/drink/9a36d04f-61ca-4762-80a6-cef6044180c0", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(w, r)
		c.Request = req

		mockService.EXPECT().UpdateDrink(c.Request.Context(), s.UpdateDrinkInput{PublicId: "9a36d04f-61ca-4762-80a6-cef6044180c0", Name: "Item 1", Stock: 1}).
			Return(s.UpdateDrinkOutput{PublicId: "9a36d04f-61ca-4762-80a6-cef6044180c0", Name: "Item 1", Stock: 1}, nil).Times(1)
		endpoint.PutApiV1StorageDrinkId(c, "9a36d04f-61ca-4762-80a6-cef6044180c0")

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"data":{"id":"9a36d04f-61ca-4762-80a6-cef6044180c0","name":"Item 1","stock":1}}`, w.Body.String())
	})
}

func TestDeleteStorageDrink(t *testing.T) {
	conf.InitTest()

	mockService := s.NewMockServiceInterface(gomock.NewController(t))
	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	endpoint := Server{Service: mockService, Repository: mockRepo}

	t.Run("Success - Delete drink on storage", func(t *testing.T) {
		r := gin.New()

		req := httptest.NewRequest(http.MethodDelete, "/api/v1/storage/drink/9a36d04f-61ca-4762-80a6-cef6044180c0", nil)

		w := httptest.NewRecorder()
		c := gin.CreateTestContextOnly(w, r)
		c.Request = req

		mockService.EXPECT().DeleteDrink(c.Request.Context(), s.DeleteDrinkInput{PublicId: "9a36d04f-61ca-4762-80a6-cef6044180c0"}).
			Return(nil).Times(1)
		endpoint.DeleteApiV1StorageDrinkId(c, "9a36d04f-61ca-4762-80a6-cef6044180c0")

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"data":"Success"}`, w.Body.String())
	})
}
