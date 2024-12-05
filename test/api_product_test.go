/*
Suger API

Testing ProductAPIService

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

func Test_openapi_ProductAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	header := apiClient.GetConfig().DefaultHeader
	header["Authorization"] = "Key b277c95e5e92ff7a8e96e74baf6ee2fb080db3e6507977c0067791abc1f52da4220e866e2081117a1721788aa2e9dc6fe009f2a699f17a7bba23973af6954db4"

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

		var orgId = "w43Vc6UfM"
		var productId = "5BXI1hZRl"

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

	t.Run("Test ProductAPIService PublishProduct", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.PublishProduct(context.Background(), orgId, productId).Execute()

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

	t.Run("Test ProductAPIService UpdateProductFulfillmentUrl", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var productId string

		resp, httpRes, err := apiClient.ProductAPI.UpdateProductFulfillmentUrl(context.Background(), orgId, productId).Execute()

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
