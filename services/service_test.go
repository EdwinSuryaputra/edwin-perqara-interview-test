package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"github.com/EdwinSuryaputra/edwin-perqara-interview-test/conf"
	r "github.com/EdwinSuryaputra/edwin-perqara-interview-test/repository"
)

func TestGetDrinks(t *testing.T) {
	conf.InitTest()

	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	service := Service{mockRepo}

	type param struct {
		ctx context.Context
	}

	type scenario struct {
		name     string
		param    param
		mock     func(param param)
		expected []GetDrinksOutput
		isError  bool
		err      error
	}

	ctx := context.Background()

	for _, s := range []scenario{
		{
			name: "Success - Get drinks list",
			param: param{
				ctx: ctx,
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinks(ctx).Return(
					[]r.GetDrinkOutput{{Id: 1, DrinkPublicId: "test", Name: "Item 1", Stock: 1}}, nil).Times(1)
			},
			expected: []GetDrinksOutput{{DrinkPublicId: "test", Name: "Item 1", Stock: 1}},
			isError:  false,
			err:      nil,
		},
		{
			name: "Failed - Error while get drinks",
			param: param{
				ctx: ctx,
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinks(ctx).Return(
					[]r.GetDrinkOutput{}, errors.New("Some error")).Times(1)
			},
			expected: []GetDrinksOutput{},
			isError:  true,
			err:      errors.New("Some error"),
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mock != nil {
				s.mock(s.param)
			}

			result, err := service.GetDrinks(s.param.ctx)
			if s.isError {
				assert.Equal(t, s.isError, err != nil)
				assert.Equal(t, s.err, err)
				return
			}

			assert.NotEmpty(t, result)
			for idx, r := range result {
				assert.Equal(t, s.expected[idx].DrinkPublicId, r.DrinkPublicId)
				assert.Equal(t, s.expected[idx].Name, r.Name)
				assert.Equal(t, s.expected[idx].Stock, r.Stock)
			}
		})
	}
}

func TestInsertDrink(t *testing.T) {
	conf.InitTest()

	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	service := Service{mockRepo}

	type param struct {
		ctx   context.Context
		input CreateDrinkInput
	}

	type scenario struct {
		name     string
		param    param
		mock     func(param param)
		expected CreateDrinkOutput
		isError  bool
		err      error
	}

	ctx := context.Background()

	for _, s := range []scenario{
		{
			name: "Success - Insert drink",
			param: param{
				ctx: ctx,
				input: CreateDrinkInput{
					Name:  "Item 1",
					Stock: 10,
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().InsertDrink(ctx, gomock.Any()).Return(nil).Times(1)
			},
			expected: CreateDrinkOutput{PublicId: "test", Name: "Item 1", Stock: 10},
			isError:  false,
			err:      nil,
		},
		{
			name: "Failed - Error while insert drink",
			param: param{
				ctx: ctx,
				input: CreateDrinkInput{
					Name:  "Item 2",
					Stock: 20,
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().InsertDrink(ctx, gomock.Any()).Return(errors.New("Some error")).Times(1)
			},
			expected: CreateDrinkOutput{},
			isError:  true,
			err:      errors.New("Some error"),
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mock != nil {
				s.mock(s.param)
			}

			result, err := service.InsertDrink(s.param.ctx, s.param.input)
			if s.isError {
				assert.Equal(t, s.isError, err != nil)
				assert.Equal(t, s.err, err)
				return
			}

			assert.Equal(t, s.expected.Name, result.Name)
			assert.Equal(t, s.expected.Stock, result.Stock)
		})
	}
}

func TestUpdateDrink(t *testing.T) {
	conf.InitTest()

	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	service := Service{mockRepo}

	type param struct {
		ctx   context.Context
		input UpdateDrinkInput
	}

	type scenario struct {
		name     string
		param    param
		mock     func(param param)
		expected UpdateDrinkOutput
		isError  bool
		err      error
	}

	ctx := context.Background()

	for _, s := range []scenario{
		{
			name: "Success - Update drink",
			param: param{
				ctx: ctx,
				input: UpdateDrinkInput{
					PublicId: "test",
					Name:     "Item 1",
					Stock:    10,
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test").
					Return(r.GetDrinkOutput{Id: 1, DrinkPublicId: "test", Name: "Item 0", Stock: 5}, nil).Times(1)

				mockRepo.EXPECT().UpdateDrink(ctx, gomock.Any()).Return(nil).Times(1)
			},
			expected: UpdateDrinkOutput{PublicId: "test", Name: "Item 1", Stock: 10},
			isError:  false,
			err:      nil,
		},
		{
			name: "Failed - Error while get drink",
			param: param{
				ctx: ctx,
				input: UpdateDrinkInput{
					PublicId: "test_2",
					Name:     "Item 2",
					Stock:    20,
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test_2").
					Return(r.GetDrinkOutput{}, errors.New("Error while get drink")).Times(1)
			},
			expected: UpdateDrinkOutput{},
			isError:  true,
			err:      errors.New("Error while get drink"),
		},
		{
			name: "Failed - Error while update drink",
			param: param{
				ctx: ctx,
				input: UpdateDrinkInput{
					PublicId: "test_3",
					Name:     "Item 3",
					Stock:    50,
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test_3").
					Return(r.GetDrinkOutput{Id: 3, DrinkPublicId: "test_3", Name: "Item 999", Stock: 1}, nil).Times(1)

				mockRepo.EXPECT().UpdateDrink(ctx, gomock.Any()).Return(errors.New("Error while update drink")).Times(1)
			},
			expected: UpdateDrinkOutput{},
			isError:  true,
			err:      errors.New("Error while update drink"),
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mock != nil {
				s.mock(s.param)
			}

			result, err := service.UpdateDrink(s.param.ctx, s.param.input)
			if s.isError {
				assert.Equal(t, s.isError, err != nil)
				assert.Equal(t, s.err, err)
				return
			}

			assert.Equal(t, s.expected.PublicId, result.PublicId)
			assert.Equal(t, s.expected.Name, result.Name)
			assert.Equal(t, s.expected.Stock, result.Stock)
		})
	}
}

func TestDeleteDrink(t *testing.T) {
	conf.InitTest()

	mockRepo := r.NewMockRepositoryInterface(gomock.NewController(t))
	service := Service{mockRepo}

	type param struct {
		ctx   context.Context
		input DeleteDrinkInput
	}

	type scenario struct {
		name    string
		param   param
		mock    func(param param)
		isError bool
		err     error
	}

	ctx := context.Background()

	for _, s := range []scenario{
		{
			name: "Success - Delete drink",
			param: param{
				ctx: ctx,
				input: DeleteDrinkInput{
					PublicId: "test",
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test").
					Return(r.GetDrinkOutput{Id: 1, DrinkPublicId: "test", Name: "Item 0", Stock: 5}, nil).Times(1)

				mockRepo.EXPECT().DeleteDrink(ctx, gomock.Any()).Return(nil).Times(1)
			},
			isError: false,
			err:     nil,
		},
		{
			name: "Failed - Error while get drink",
			param: param{
				ctx: ctx,
				input: DeleteDrinkInput{
					PublicId: "test_2",
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test_2").
					Return(r.GetDrinkOutput{}, errors.New("Error while get drink")).Times(1)
			},
			isError: true,
			err:     errors.New("Error while get drink"),
		},
		{
			name: "Failed - Error while delete drink",
			param: param{
				ctx: ctx,
				input: DeleteDrinkInput{
					PublicId: "test_3",
				},
			},
			mock: func(param param) {
				mockRepo.EXPECT().GetDrinkByPublicId(ctx, "test_3").
					Return(r.GetDrinkOutput{Id: 3, DrinkPublicId: "test_3", Name: "Item 999", Stock: 1}, nil).Times(1)

				mockRepo.EXPECT().DeleteDrink(ctx, gomock.Any()).Return(errors.New("Error while delete drink")).Times(1)
			},
			isError: true,
			err:     errors.New("Error while delete drink"),
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.mock != nil {
				s.mock(s.param)
			}

			err := service.DeleteDrink(s.param.ctx, s.param.input)
			if s.isError {
				assert.Equal(t, s.isError, err != nil)
				assert.Equal(t, s.err, err)
				return
			}
		})
	}
}
