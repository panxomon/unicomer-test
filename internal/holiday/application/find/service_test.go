package find_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
	"unicomer-test/internal/holiday/application/find"
	"unicomer-test/internal/holiday/domain"
	"unicomer-test/tests/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_Service(t *testing.T) {
	type test struct {
		name             string
		holidaysResponse *domain.HolidayResponse
		repositoryError  error
		expectedResult   []domain.Holiday
		expectedError    error
	}

	cases := []test{
		{
			name: "when service is ok",
			holidaysResponse: &domain.HolidayResponse{
				Status: "OK",
				Data: []domain.Holiday{
					{
						Date:        mustParseTime("2024-05-01T00:00:00Z"),
						Title:       "Día Nacional del Trabajo",
						Type:        "Civil",
						Inalienable: true,
						Extra:       "Civil e Irrenunciable",
					},
					{
						Date:        mustParseTime("2024-05-21T00:00:00Z"),
						Title:       "Día de las Glorias Navales",
						Type:        "Civil",
						Inalienable: false,
						Extra:       "Civil",
					},
					{
						Date:        mustParseTime("2024-06-09T00:00:00Z"),
						Title:       "Elecciones Primarias Alcaldes y Gobernadores",
						Type:        "Civil",
						Inalienable: false,
						Extra:       "Civil",
					},
					{
						Date:        mustParseTime("2024-06-20T00:00:00Z"),
						Title:       "Día Nacional de los Pueblos Indígenas",
						Type:        "Civil",
						Inalienable: false,
						Extra:       "Civil",
					},
				},
			},
			repositoryError: nil,
			expectedResult: []domain.Holiday{
				{
					Date:        mustParseTime("2024-05-01T00:00:00Z"),
					Title:       "Día Nacional del Trabajo",
					Type:        "Civil",
					Inalienable: true,
					Extra:       "Civil e Irrenunciable",
				},
				{
					Date:        mustParseTime("2024-05-21T00:00:00Z"),
					Title:       "Día de las Glorias Navales",
					Type:        "Civil",
					Inalienable: false,
					Extra:       "Civil",
				},
				{
					Date:        mustParseTime("2024-06-09T00:00:00Z"),
					Title:       "Elecciones Primarias Alcaldes y Gobernadores",
					Type:        "Civil",
					Inalienable: false,
					Extra:       "Civil",
				},
				{
					Date:        mustParseTime("2024-06-20T00:00:00Z"),
					Title:       "Día Nacional de los Pueblos Indígenas",
					Type:        "Civil",
					Inalienable: false,
					Extra:       "Civil",
				},
			},
			expectedError: nil,
		},
		{
			name:             "when repository returns error",
			holidaysResponse: nil,
			repositoryError:  errors.New("repository error"),
			expectedResult:   nil,
			expectedError:    errors.New("repository error"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			repositoryMock := new(mocks.HolidayRepository)
			repositoryMock.On("Retrieve", mock.Anything).Return(tt.holidaysResponse, tt.repositoryError)

			svc := find.NewHolidayFinder(repositoryMock)
			response, err := svc.Execute(context.Background())

			if tt.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedResult, response)
			}

			fmt.Println(response)
		})
	}
}

func mustParseTime(value string) time.Time {
	const layout = time.RFC3339
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}
