package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/proto/auth/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
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

func NewGateway(ctx context.Context) (http.Handler, *GatewaySettings, error) {
	gatewaySettings, err := readSetting()
	if err != nil {
		return nil, nil, err
	}

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := gwruntime.NewServeMux(jsonOption())
	for _, service := range gatewaySettings.Services {
		log.Info().Str("name", service.Name).Msg("registering service")
		err := gen.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, service.Url, dialOpts)
		if err != nil {
			log.Error().Err(err).Str("name", service.Name).Msg("cannot register service")
			return nil, nil, err
		}
	}

	return mux, gatewaySettings, nil
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

func jsonOption() gwruntime.ServeMuxOption {
	return gwruntime.WithMarshalerOption(gwruntime.MIMEWildcard, &gwruntime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
}
