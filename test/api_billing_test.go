/*
Suger API

Testing BillingAPIService

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

func Test_suger_BillingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BillingAPIService CreateAddon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.BillingAPI.CreateAddon(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService CreateRefund", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var buyerId string
		var paymentTransactionId string

		resp, httpRes, err := apiClient.BillingAPI.CreateRefund(context.Background(), orgId, buyerId, paymentTransactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService DeleteAddon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var addonId string

		resp, httpRes, err := apiClient.BillingAPI.DeleteAddon(context.Background(), orgId, addonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService GetAddon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var addonId string

		resp, httpRes, err := apiClient.BillingAPI.GetAddon(context.Background(), orgId, addonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService GetInvoiceV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.GetInvoiceV2(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService IssueInvoiceV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.IssueInvoiceV2(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService ListAddons", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.BillingAPI.ListAddons(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService ListInvoices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.BillingAPI.ListInvoices(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService ListPaymentTransactions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.BillingAPI.ListPaymentTransactions(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService ListRefundOfPaymentTransaction", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var buyerId string
		var paymentTransactionId string

		resp, httpRes, err := apiClient.BillingAPI.ListRefundOfPaymentTransaction(context.Background(), orgId, buyerId, paymentTransactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService PayInvoiceV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.PayInvoiceV2(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService PreviewInvoiceEmail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.PreviewInvoiceEmail(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService UpdateAddon", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var addonId string

		resp, httpRes, err := apiClient.BillingAPI.UpdateAddon(context.Background(), orgId, addonId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService UpdateInvoiceInfoV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.UpdateInvoiceInfoV2(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BillingAPIService VoidInvoiceV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string
		var invoiceId string

		resp, httpRes, err := apiClient.BillingAPI.VoidInvoiceV2(context.Background(), orgId, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
