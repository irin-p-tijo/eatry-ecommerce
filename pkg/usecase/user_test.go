package usecase

import (
	"eatry/pkg/repository/mock"
	"eatry/pkg/utils/models"

	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	cartRepository := mock.NewMockCartRepository(ctrl)

	userUseCase := NewUserUseCase(userRepo, cartRepository)

	testingData := map[string]struct {
		input   int
		stub    func(*mock.MockUserRepository)
		want    []models.Address
		wantErr error
	}{
		"success": {
			input: 1,
			stub: func(ur *mock.MockUserRepository) {
				ur.EXPECT().GetAddresses(1).Return([]models.Address{
					{
						HouseName: "chandhana palli",
						Street:    "viyyor college road",
						City:      "Devagiri",
						District:  "Kozhikkode",
						State:     "kerala",
						Pin:       "673003",
					},
				}, nil)
			},
			want: []models.Address{
				{
					HouseName: "chandhana palli",
					Street:    "viyyor college road",
					City:      "Devagiri",
					District:  "Kozhikkode",
					State:     "kerala",
					Pin:       "673003",
				},
			},
			wantErr: nil,
		},
		"failure": {
			input: 1,
			stub: func(ur *mock.MockUserRepository) {
				ur.EXPECT().GetAddresses(1).Return(nil, errors.New("error in getting addresses"))
			},
			want:    []models.Address{},
			wantErr: errors.New("error in getting addresses"),
		},
	}

	for name, test := range testingData {
		t.Run(name, func(t *testing.T) {
			test.stub(userRepo)
			result, err := userUseCase.GetAllAddress(test.input)
			assert.Equal(t, test.want, result)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
