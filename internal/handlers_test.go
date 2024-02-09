package internal_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlgorithmChopda/Website-Checker.git/internal"
	"github.com/AlgorithmChopda/Website-Checker.git/mocks"
	"github.com/stretchr/testify/mock"
)

func TestReadWebsiteHandler(t *testing.T) {
	websitesSvc := mocks.NewWebsiteService(t)
	addWebsitesHandler := internal.ReadWebsiteHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.WebsiteService)
		expectedStatusCode int
	}{
		{
			name: "Add website",
			input: `
                {
                    "data": ["hi.com"]
                }
            `,
			setup: func(mockSvc *mocks.WebsiteService) {
				mockSvc.On("ReadWebsite", mock.Anything).Return().Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error for invalid json",
			input: `
                {
                    "data": ["]
                }
            `,
			setup:              func(mockSvc *mocks.WebsiteService) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "empty website",
			input: `
                {
                    "data": [""]
                }
            `,
			setup: func(mockSvc *mocks.WebsiteService) {
				mockSvc.On("ReadWebsite", mock.Anything).Return().Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", "/website", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}

}

func TestGetWebsiteStatus8(t *testing.T) {
	websitesSvc := mocks.NewWebsiteService(t)
	addWebsitesHandler := internal.GetWebsiteStatus(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.WebsiteService)
		expectedStatusCode int
	}{
		{
			name:  "url present",
			input: fmt.Sprintf("?url=%s", "abc.com"),
			setup: func(mockSvc *mocks.WebsiteService) {
				mockSvc.On("GetWebsiteStatus", mock.Anything).Return(internal.Status{
					IsActive:    true,
					LastFetched: "now",
				}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "url not found",
			input: fmt.Sprintf("?url=%s", "abc.com"),
			setup: func(mockSvc *mocks.WebsiteService) {
				mockSvc.On("GetWebsiteStatus", mock.Anything).Return(internal.Status{
					IsActive:    true,
					LastFetched: "now",
				}, errors.New("error occured")).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "all url status",
			input: "", //fmt.Sprintf("?url=%s", "abc.com"),
			setup: func(mockSvc *mocks.WebsiteService) {
				mockSvc.On("GetAllWebsiteStatus").Return(internal.Status{
					IsActive:    true,
					LastFetched: "now",
				}).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", fmt.Sprintf("/website/status%s", test.input), bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}

}
