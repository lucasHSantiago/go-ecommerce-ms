package gateway

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/internal/util"
	"github.com/lucasHSantiago/go-ecommerce-ms/gateway/proto/auth/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewGateway(ctx context.Context, config util.Config) (http.Handler, error) {
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := gwruntime.NewServeMux(jsonOption())
	err := gen.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, config.AuthAddress, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
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
