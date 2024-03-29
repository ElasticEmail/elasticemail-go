/*
Elastic Email REST API

Testing SuppressionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ElasticEmail

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ElasticEmail_SuppressionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SuppressionsApiService SuppressionsBouncesGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsBouncesGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsBouncesImportPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsBouncesImportPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsBouncesPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsBouncesPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsByEmailDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var email string

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsByEmailDelete(context.Background(), email).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsByEmailGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var email string

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsByEmailGet(context.Background(), email).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsComplaintsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsComplaintsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsComplaintsImportPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsComplaintsImportPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsComplaintsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsComplaintsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsUnsubscribesGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsUnsubscribesGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsUnsubscribesImportPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsUnsubscribesImportPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SuppressionsApiService SuppressionsUnsubscribesPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SuppressionsApi.SuppressionsUnsubscribesPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
