/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package shellbackend




type Controller struct {

	// Name of the controller
	Name string `json:"name"`

	// IP address of the controller
	Address string `json:"address"`
}

// AssertControllerRequired checks if the required fields are not zero-ed
func AssertControllerRequired(obj Controller) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"address": obj.Address,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertControllerConstraints checks if the values respects the defined constraints
func AssertControllerConstraints(obj Controller) error {
	return nil
}
