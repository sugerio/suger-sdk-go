/*
Suger API

Testing EntitlementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package suger

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/sugerio/suger-sdk-go"
	"testing"
)

func Test_suger_EntitlementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EntitlementAPIService AddEntitlementCredit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.AddEntitlementCredit(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ApplyAddonToEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.ApplyAddonToEntitlement(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ApproveEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.ApproveEntitlement(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService CancelEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.CancelEntitlement(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService CreateEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.EntitlementAPI.CreateEntitlement(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService DeleteEntitlementTerm", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string
		var entitlementTermId string

		resp, httpRes, err := apiClient.EntitlementAPI.DeleteEntitlementTerm(context.Background(), orgId, entitlementId, entitlementTermId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService DivideEntitlementCommit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.DivideEntitlementCommit(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService GetEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.GetEntitlement(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService GetEntitlementTerm", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string
		var entitlementTermId string

		resp, httpRes, err := apiClient.EntitlementAPI.GetEntitlementTerm(context.Background(), orgId, entitlementId, entitlementTermId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlementTerms", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlementTerms(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlements(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ScheduleEntitlementCancellation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.ScheduleEntitlementCancellation(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService UnscheduleEntitlementCancellation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.UnscheduleEntitlementCancellation(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService UpdateEntitlementMetaInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.UpdateEntitlementMetaInfo(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService UpdateEntitlementName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.UpdateEntitlementName(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService UpdateEntitlementSeat", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.UpdateEntitlementSeat(context.Background(), orgId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
