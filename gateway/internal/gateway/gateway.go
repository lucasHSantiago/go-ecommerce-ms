package gateway

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/proto/auth/gen"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

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
