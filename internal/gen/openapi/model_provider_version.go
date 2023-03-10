/*
 * Alexandria
 *
 * Alexandria
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProviderVersion struct {
	Versions []ProviderVersionVersionsInner `json:"versions,omitempty"`
}

// AssertProviderVersionRequired checks if the required fields are not zero-ed
func AssertProviderVersionRequired(obj ProviderVersion) error {
	for _, el := range obj.Versions {
		if err := AssertProviderVersionVersionsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseProviderVersionRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ProviderVersion (e.g. [][]ProviderVersion), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseProviderVersionRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aProviderVersion, ok := obj.(ProviderVersion)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertProviderVersionRequired(aProviderVersion)
	})
}