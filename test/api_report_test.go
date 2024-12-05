/*
Suger API

Testing ReportAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/sugerio/suger-sdk-go"
	"testing"
)

func Test_openapi_ReportAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ReportAPIService GetRevenueReport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ReportAPI.GetRevenueReport(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportAPIService ListDailyRevenueRecords", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ReportAPI.ListDailyRevenueRecords(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportAPIService ListRevenueRecordDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.ReportAPI.ListRevenueRecordDetails(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportAPIService ListRevenueRecords", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.ReportAPI.ListRevenueRecords(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ReportAPIService ListUsageMeteringDailyRecords", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.ReportAPI.ListUsageMeteringDailyRecords(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
