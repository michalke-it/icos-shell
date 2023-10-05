/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package shellbackend

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type TimerData struct {
	timer      *time.Timer
	controller Controller
}

var timers = make(map[string]TimerData)

func DeleteController(key string) {
	fmt.Printf("Timeout, deleting controller: '%s'\n", key)
	delete(timers, key)
}

func AddController(controller Controller) bool {
	key := controller.Address
	val, exists := timers[key]
	duration := time.Second * time.Duration(viper.GetInt("lighthouse.controller_timeout"))
	if exists {
		val.timer.Reset(duration)
		fmt.Printf("Timer reset for controller: '%s'\n", controller.Address)
		return true
	} else {
		timer := time.NewTimer(duration)
		timers[key] = TimerData{
			timer:      timer,
			controller: controller,
		}
		fmt.Printf("Controller added: '%s'\n", controller.Address)
		// Start a goroutine to wait for the timer to expire
		go func() {
			<-timer.C
			DeleteController(key)
		}()
		return false
	}
}

func GetControllersList() []Controller {
	var controllers []Controller
	for _, value := range timers {
		controllers = append(controllers, value.controller)
	}
	return controllers
}

// ControllerAPIService is a service that implements the logic for the ControllerAPIServicer
// This service should implement the business logic for every endpoint for the ControllerAPI API.
// Include any external packages or services that will be required by this service.
type ControllerAPIService struct {
}

// NewControllerAPIService creates a default api service
func NewControllerAPIService() ControllerAPIServicer {
	return &ControllerAPIService{}
}

// AddController - Adds a new controller
func (s *ControllerAPIService) AddController(ctx context.Context, controller Controller, apiKey string) (ImplResponse, error) {
	// TO-DO keycloak
	// if (strings.Compare(username, viper.GetString("lighthouse.username")) == 0) && (strings.Compare(password, viper.GetString("lighthouse.password")) == 0) {
	// 	exists := AddController(controller)
	// 	if exists {
	// 		return Response(202, "Controller already exists, timer has been reset"), nil
	// 	} else {
	// 		return Response(201, "New controller correctly added"), nil
	// 	}
	// } else {
	// 	return Response(405, nil), errors.New("wrong user or password")
	// }
	return Response(http.StatusNotImplemented, nil), errors.New("AddController method not implemented")
}

// GetControllers - Returns a list of controllers
func (s *ControllerAPIService) GetControllers(ctx context.Context) (ImplResponse, error) {
	var controllers = GetControllersList()
	if controllers == nil {
		return Response(204, nil), nil
	} else {
		return Response(200, controllers), nil
	}
}
