package endpoint_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"unicomer-test/internal/endpoint"
	holiday "unicomer-test/internal/holiday/application"
	"unicomer-test/internal/holiday/application/find"
	"unicomer-test/internal/holiday/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_HolidayEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type test struct {
		name           string
		holidayApp     *holiday.App
		mockResponse   []domain.Holiday
		mockError      error
		queryParams    string
		expectedStatus int
		expectedBody   string
	}

	mockHolidays := []domain.Holiday{
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
	}

	cases := []test{
		{
			name: "valid request - JSON response",
			holidayApp: holiday.NewApp(holiday.Queries{
				FindHoliday: createMockHandler(mockHolidays, nil),
			}),
			queryParams:    "",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"holidays":[{"date":"2024-05-01T00:00:00Z","title":"Día Nacional del Trabajo","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2024-05-21T00:00:00Z","title":"Día de las Glorias Navales","type":"Civil","inalienable":false,"extra":"Civil"}]}`,
		},
		{
			name: "valid request - XML response",
			holidayApp: holiday.NewApp(holiday.Queries{
				FindHoliday: createMockHandler(mockHolidays, nil),
			}),
			queryParams:    "",
			expectedStatus: http.StatusOK,
			expectedBody:   `<Holidays><holidays><holiday><Date>2024-05-01T00:00:00Z</Date><Title>Día Nacional del Trabajo</Title><Type>Civil</Type><Inalienable>true</Inalienable><Extra>Civil e Irrenunciable</Extra></holiday><holiday><Date>2024-05-21T00:00:00Z</Date><Title>Día de las Glorias Navales</Title><Type>Civil</Type><Inalienable>false</Inalienable><Extra>Civil</Extra></holiday></holidays></Holidays>`,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			ep := endpoint.NewEndpoint(tt.holidayApp)
			router.GET("/holidays", func(c *gin.Context) {
				err := ep.Invoke(c)
				if err != nil {
					t.Error(err.Error())
				}
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/holidays"+tt.queryParams, nil)
			if tt.name == "valid request - XML response" {
				req.Header.Set("Content-Type", "application/xml")
			} else {
				router.ServeHTTP(w, req)

				assert.Equal(t, tt.expectedStatus, w.Code)
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}

		})
	}
}

func createMockHandler(holidays []domain.Holiday, err error) find.HolidayFindQueryHandler {
	mockHandler := new(HolidayFindQueryHandler)
	mockHandler.On("Handle", mock.Anything, mock.Anything).Return(holidays, err)
	return mockHandler
}

func mustParseTime(value string) time.Time {
	const layout = time.RFC3339
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

type HolidayFindQueryHandler struct {
	mock.Mock
}

func (m *HolidayFindQueryHandler) Handle(ctx context.Context, query find.HolidayFindQuery) ([]domain.Holiday, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]domain.Holiday), args.Error(1)
}
