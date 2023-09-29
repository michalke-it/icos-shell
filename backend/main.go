/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	shellbackend "shellbackend/go"

	"github.com/spf13/viper"
)

func main() {

	cfgFile := flag.String("config", "config.yml", "config file")
	flag.Parse()
	cfgFileString := *cfgFile

	viper.SetConfigFile(cfgFileString)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Config:", viper.ConfigFileUsed())
	} else {
		fmt.Fprintln(os.Stderr, "Config file not found!")
	}

	log.Printf("Backend service starting...")

	ControllerApiService := shellbackend.NewControllerApiService()
	ControllerApiController := shellbackend.NewControllerApiController(ControllerApiService)

	DefaultApiService := shellbackend.NewDefaultApiService()
	DefaultApiController := shellbackend.NewDefaultApiController(DefaultApiService)

	DeploymentApiService := shellbackend.NewDeploymentApiService()
	DeploymentApiController := shellbackend.NewDeploymentApiController(DeploymentApiService)

	ResourceApiService := shellbackend.NewResourceApiService()
	ResourceApiController := shellbackend.NewResourceApiController(ResourceApiService)

	UserApiService := shellbackend.NewUserApiService()
	UserApiController := shellbackend.NewUserApiController(UserApiService)

	router := shellbackend.NewRouter(ControllerApiController, DefaultApiController, DeploymentApiController, ResourceApiController, UserApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
