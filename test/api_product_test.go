/*
Suger API

Testing ProductAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ProductAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductAPIService CreateOrUpdateDraftProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ProductAPI.CreateOrUpdateDraftProduct(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService CreateProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ProductAPI.CreateProduct(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService DeleteProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.DeleteProduct(context.Background(), orgId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService GetProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		orgId = "EH3R7hwfM"
		productId = "eWeOhF0tp"
		resp, httpRes, err := apiClient.ProductAPI.GetProduct(context.Background(), orgId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ListProductMeteringDimensions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.ListProductMeteringDimensions(context.Background(), orgId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ListProductsByOrganization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.ProductAPI.ListProductsByOrganization(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService ListProductsByPartner", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var partner string

		resp, httpRes, err := apiClient.ProductAPI.ListProductsByPartner(context.Background(), orgId, partner).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService UpdateProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.UpdateProduct(context.Background(), orgId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductAPIService UpdateProductMetaInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.UpdateProductMetaInfo(context.Background(), orgId, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
