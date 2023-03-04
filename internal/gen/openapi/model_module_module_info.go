/*
 * Alexandria
 *
 * Alexandria
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ModuleModuleInfo struct {
	Valid bool `json:"valid,omitempty"`

	ErrorCount int32 `json:"error_count,omitempty"`

	WarningCount int32 `json:"warning_count,omitempty"`
}

// AssertModuleModuleInfoRequired checks if the required fields are not zero-ed
func AssertModuleModuleInfoRequired(obj ModuleModuleInfo) error {
	return nil
}

// AssertRecurseModuleModuleInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ModuleModuleInfo (e.g. [][]ModuleModuleInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseModuleModuleInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aModuleModuleInfo, ok := obj.(ModuleModuleInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertModuleModuleInfoRequired(aModuleModuleInfo)
	})
}
