/*
Suger API

Testing EntitlementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/sugerio/suger-sdk-go"
)

func Test_openapi_EntitlementAPIService(t *testing.T) {

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

	t.Run("Test EntitlementAPIService ApproveEntitlement", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementAPI.ApproveEntitlement(context.Background(), orgId, entitlementId).Execute()

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

	t.Run("Test EntitlementAPIService ListEntitlements", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlements(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlementsByBuyer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var buyerId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlementsByBuyer(context.Background(), orgId, buyerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlementsByOffer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var offerId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlementsByOffer(context.Background(), orgId, offerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlementsByPartner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlementsByPartner(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementAPIService ListEntitlementsByProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.EntitlementAPI.ListEntitlementsByProduct(context.Background(), orgId, productId).Execute()

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

}
