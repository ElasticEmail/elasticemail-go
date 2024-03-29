/*
Elastic Email REST API

Testing CampaignsApiService

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

func Test_ElasticEmail_CampaignsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CampaignsApiService CampaignsByNameDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.CampaignsApi.CampaignsByNameDelete(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CampaignsApiService CampaignsByNameGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.CampaignsApi.CampaignsByNameGet(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CampaignsApiService CampaignsByNamePut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.CampaignsApi.CampaignsByNamePut(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CampaignsApiService CampaignsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CampaignsApi.CampaignsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CampaignsApiService CampaignsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.CampaignsApi.CampaignsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
