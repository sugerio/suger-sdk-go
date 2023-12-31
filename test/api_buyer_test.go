/*
Suger API

Testing BuyerAPIService

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

func Test_openapi_BuyerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BuyerAPIService GetBuyer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var buyerId string

		resp, httpRes, err := apiClient.BuyerAPI.GetBuyer(context.Background(), orgId, buyerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BuyerAPIService ListBuyersByContact", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var contactId string

		resp, httpRes, err := apiClient.BuyerAPI.ListBuyersByContact(context.Background(), orgId, contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BuyerAPIService ListBuyersByOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.BuyerAPI.ListBuyersByOrganization(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BuyerAPIService ListBuyersByPartner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.BuyerAPI.ListBuyersByPartner(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BuyerAPIService UpdateBuyer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var buyerId string

		resp, httpRes, err := apiClient.BuyerAPI.UpdateBuyer(context.Background(), orgId, buyerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
