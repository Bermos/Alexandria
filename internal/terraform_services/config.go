package terraform_services

import "fmt"

var (
	enabledServices = map[string]string{}
)

func init() {
	EnableService("modules.v1", "/modules")
}

func GetEnabledServices() map[string]string {
	return enabledServices
}

func EnableService(name, endpoint string) {
	enabledServices[name] = endpoint
}

func DisableService(name string) error {
	_, ok := enabledServices[name]
	if ok {
		delete(enabledServices, name)
	} else {
		return fmt.Errorf("could not remove key %s, it does not exist in the map", name)
	}

	return nil
}
