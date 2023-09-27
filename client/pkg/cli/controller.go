package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/tubskns/icos-shell/client/pkg/openapi"
)

func AddController(name string, address string) (result string) {
	controller := *openapi.NewController(name, address)
	resp, err := openapi.Client.ControllerApi.AddController(context.Background()).Username(viper.GetString("username")).Password(viper.GetString("password")).Controller(controller).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", resp.Body)
		return "Error when adding a controller"
	} else {
		if resp.StatusCode == 201 {
			return "Controller successfully added!"
		} else if resp.StatusCode == 202 {
			return "Controller already exists"
		} else {
			return "Wrong status code received"
		}
	}
}

func GetController() (result string) {
	controllers, resp, err := openapi.Client.ControllerApi.GetControllers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return "Error when retrieving controllers"
	} else {
		if resp.StatusCode == 200 {
			for _, element := range controllers {
				fmt.Println(element)
			}
			return "Controller list retrieved"
		} else if resp.StatusCode == 204 {
			return "No controllers found"
		} else {
			return "Wrong status code received"
		}
	}
}
