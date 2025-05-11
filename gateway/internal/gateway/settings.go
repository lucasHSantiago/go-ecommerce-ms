package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

type GatewaySettings struct {
	Services []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
		Acl  []struct {
			Http  string   `json:"http"`
			Route string   `json:"route"`
			Roles []string `json:"roles"`
		} `json:"acl"`
	} `json:"services"`
}

func (gs *GatewaySettings) NeedAuth(r *http.Request) bool {
	for _, service := range gs.Services {
		for _, authRoute := range service.Acl {
			if authRoute.Route != "" && strings.Contains(strings.ToLower(r.URL.Path), strings.ToLower(authRoute.Route)) && r.Method == strings.ToUpper(authRoute.Http) {
				return true
			}
		}
	}
	return false
}

func (gs *GatewaySettings) HasPermission(userRole string) bool {
	for _, service := range gs.Services {
		if service.Acl == nil || len(service.Acl) == 0 {
			return true
		}

		for _, acl := range service.Acl {
			if acl.Roles == nil || len(acl.Roles) == 0 {
				return true
			}

			if slices.Contains(acl.Roles, userRole) {
				return true
			}
		}
	}

	return false
}

func readSetting() (*GatewaySettings, error) {
	file, err := os.Open("settings.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open setting.json: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings.json: %w", err)
	}

	gatewaySettings := &GatewaySettings{}
	err = json.Unmarshal(byteValue, gatewaySettings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal settings.json: %w", err)
	}

	return gatewaySettings, nil
}
