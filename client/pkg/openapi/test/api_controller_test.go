/*
ICOS Shell

Testing ControllerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tubskns/icos-shell/client/pkg/openapi"
)

func Test_openapi_ControllerApiService(t *testing.T) {

	openapi.Init("lighthouse.icos-project.eu:8080")
	username := "admin"
	password := "Iki946D56!!J@gSHpuonoUyH1uB*^"

	t.Run("Test ControllerApiService AddController", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		// httpRes, err := openapi.Client.ControllerApi.AddController(context.Background()).Execute()
		controller := *openapi.NewController("name_test", "address_test")
		httpRes, err := openapi.Client.ControllerApi.AddController(context.Background()).Username(username).Password(password).Controller(controller).Execute()

		require.Nil(t, err)
		assert.Equal(t, 201, httpRes.StatusCode)

	})

	t.Run("Test ControllerApiService GetControllers", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		resp, httpRes, err := openapi.Client.ControllerApi.GetControllers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
