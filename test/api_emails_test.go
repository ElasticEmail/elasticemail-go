/*
Elastic Email REST API

Testing EmailsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ElasticEmail

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/elasticemail/elasticemail-go"
)

func Test_ElasticEmail_EmailsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailsAPIService EmailsByMsgidViewGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var msgid string

		resp, httpRes, err := apiClient.EmailsAPI.EmailsByMsgidViewGet(context.Background(), msgid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsAPIService EmailsByTransactionidStatusGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionid string

		resp, httpRes, err := apiClient.EmailsAPI.EmailsByTransactionidStatusGet(context.Background(), transactionid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsAPIService EmailsMergefilePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsAPI.EmailsMergefilePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsAPIService EmailsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsAPI.EmailsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailsAPIService EmailsTransactionalPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EmailsAPI.EmailsTransactionalPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
