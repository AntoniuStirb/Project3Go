package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHowMuchTillPayday(t *testing.T) {
	tests := []struct {
		name         string
		payDay       string
		expectedCode int
	}{
		{
			name:         "valid pay_day parameter",
			payDay:       "15",
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid pay_day parameter",
			payDay:       "32",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid pay_day parameter",
			payDay:       "0",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid pay_day parameter",
			payDay:       "-5",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid pay_day parameter",
			payDay:       "abcdefg/",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "missing pay_day parameter",
			payDay:       "",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a new request with the test parameters.
			req, err := http.NewRequest("GET", fmt.Sprintf("/?pay_day=%s", testCase.payDay), nil)
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}

			// Set the query parameters for the request.
			//q := req.URL.Query()
			//q.Add("pay_day", tt.payDay)
			rec := httptest.NewRecorder()

			// Call the handler function with the test request and response.
			handler := http.HandlerFunc(HowMuchTillPayday)
			handler.ServeHTTP(rec, req)

			// Check the response status code.
			if rec.Code != testCase.expectedCode {
				t.Errorf("unexpected status code: got %d, want %d", rec.Code, testCase.expectedCode)
			}
		})
	}
}

func TestPayDayListDates(t *testing.T) {

	testCases := []struct {
		name           string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid URL",
			url:            "/till-salary/payday/12/list-dates/asdasd",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid URL\n",
		},
		{
			name:           "Invalid URL",
			url:            "/till-salary/payday//list-dates/asdasd",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid URL\n",
		},
		{
			name:           "Invalid Pay Day",
			url:            "/till-salary/payday/40/list-dates",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid pay day\n",
		},
		{
			name:           "Successful Request",
			url:            "/till-salary/payday/15/list-dates",
			expectedStatus: http.StatusOK,
			expectedBody: `{"pay_days":["15-03-2023","15-04-2023","15-05-2023","15-06-2023","15-07-2023","15-08-2023",
							"15-09-2023","15-10-2023","15-11-2023","15-12-2023"]}`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, testCase.url, nil)
			if err != nil {
				t.Errorf("failed to create request: %v", err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(PayDayListDates)
			handler.ServeHTTP(rec, req)

			if rec.Code != testCase.expectedStatus {
				t.Errorf("unexpected status code: got %d, want %d", rec.Code, testCase.expectedStatus)
			}

			if rec.Body.String() != testCase.expectedBody {
				t.Errorf("Expected response body %q, but got %q", testCase.expectedBody, rec.Body.String())
			}
		})
	}
}
