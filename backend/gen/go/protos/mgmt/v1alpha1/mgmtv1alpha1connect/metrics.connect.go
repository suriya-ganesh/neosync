// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/metrics.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// MetricsServiceName is the fully-qualified name of the MetricsService service.
	MetricsServiceName = "mgmt.v1alpha1.MetricsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MetricsServiceGetDailyMetricCountProcedure is the fully-qualified name of the MetricsService's
	// GetDailyMetricCount RPC.
	MetricsServiceGetDailyMetricCountProcedure = "/mgmt.v1alpha1.MetricsService/GetDailyMetricCount"
	// MetricsServiceGetMetricCountProcedure is the fully-qualified name of the MetricsService's
	// GetMetricCount RPC.
	MetricsServiceGetMetricCountProcedure = "/mgmt.v1alpha1.MetricsService/GetMetricCount"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	metricsServiceServiceDescriptor                   = v1alpha1.File_mgmt_v1alpha1_metrics_proto.Services().ByName("MetricsService")
	metricsServiceGetDailyMetricCountMethodDescriptor = metricsServiceServiceDescriptor.Methods().ByName("GetDailyMetricCount")
	metricsServiceGetMetricCountMethodDescriptor      = metricsServiceServiceDescriptor.Methods().ByName("GetMetricCount")
)

// MetricsServiceClient is a client for the mgmt.v1alpha1.MetricsService service.
type MetricsServiceClient interface {
	// Retrieve a timed range of records
	GetDailyMetricCount(context.Context, *connect.Request[v1alpha1.GetDailyMetricCountRequest]) (*connect.Response[v1alpha1.GetDailyMetricCountResponse], error)
	// For the given metric and time range, returns the total count found
	GetMetricCount(context.Context, *connect.Request[v1alpha1.GetMetricCountRequest]) (*connect.Response[v1alpha1.GetMetricCountResponse], error)
}

// NewMetricsServiceClient constructs a client for the mgmt.v1alpha1.MetricsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMetricsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MetricsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &metricsServiceClient{
		getDailyMetricCount: connect.NewClient[v1alpha1.GetDailyMetricCountRequest, v1alpha1.GetDailyMetricCountResponse](
			httpClient,
			baseURL+MetricsServiceGetDailyMetricCountProcedure,
			connect.WithSchema(metricsServiceGetDailyMetricCountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getMetricCount: connect.NewClient[v1alpha1.GetMetricCountRequest, v1alpha1.GetMetricCountResponse](
			httpClient,
			baseURL+MetricsServiceGetMetricCountProcedure,
			connect.WithSchema(metricsServiceGetMetricCountMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// metricsServiceClient implements MetricsServiceClient.
type metricsServiceClient struct {
	getDailyMetricCount *connect.Client[v1alpha1.GetDailyMetricCountRequest, v1alpha1.GetDailyMetricCountResponse]
	getMetricCount      *connect.Client[v1alpha1.GetMetricCountRequest, v1alpha1.GetMetricCountResponse]
}

// GetDailyMetricCount calls mgmt.v1alpha1.MetricsService.GetDailyMetricCount.
func (c *metricsServiceClient) GetDailyMetricCount(ctx context.Context, req *connect.Request[v1alpha1.GetDailyMetricCountRequest]) (*connect.Response[v1alpha1.GetDailyMetricCountResponse], error) {
	return c.getDailyMetricCount.CallUnary(ctx, req)
}

// GetMetricCount calls mgmt.v1alpha1.MetricsService.GetMetricCount.
func (c *metricsServiceClient) GetMetricCount(ctx context.Context, req *connect.Request[v1alpha1.GetMetricCountRequest]) (*connect.Response[v1alpha1.GetMetricCountResponse], error) {
	return c.getMetricCount.CallUnary(ctx, req)
}

// MetricsServiceHandler is an implementation of the mgmt.v1alpha1.MetricsService service.
type MetricsServiceHandler interface {
	// Retrieve a timed range of records
	GetDailyMetricCount(context.Context, *connect.Request[v1alpha1.GetDailyMetricCountRequest]) (*connect.Response[v1alpha1.GetDailyMetricCountResponse], error)
	// For the given metric and time range, returns the total count found
	GetMetricCount(context.Context, *connect.Request[v1alpha1.GetMetricCountRequest]) (*connect.Response[v1alpha1.GetMetricCountResponse], error)
}

// NewMetricsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMetricsServiceHandler(svc MetricsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	metricsServiceGetDailyMetricCountHandler := connect.NewUnaryHandler(
		MetricsServiceGetDailyMetricCountProcedure,
		svc.GetDailyMetricCount,
		connect.WithSchema(metricsServiceGetDailyMetricCountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	metricsServiceGetMetricCountHandler := connect.NewUnaryHandler(
		MetricsServiceGetMetricCountProcedure,
		svc.GetMetricCount,
		connect.WithSchema(metricsServiceGetMetricCountMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/mgmt.v1alpha1.MetricsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MetricsServiceGetDailyMetricCountProcedure:
			metricsServiceGetDailyMetricCountHandler.ServeHTTP(w, r)
		case MetricsServiceGetMetricCountProcedure:
			metricsServiceGetMetricCountHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMetricsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMetricsServiceHandler struct{}

func (UnimplementedMetricsServiceHandler) GetDailyMetricCount(context.Context, *connect.Request[v1alpha1.GetDailyMetricCountRequest]) (*connect.Response[v1alpha1.GetDailyMetricCountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.MetricsService.GetDailyMetricCount is not implemented"))
}

func (UnimplementedMetricsServiceHandler) GetMetricCount(context.Context, *connect.Request[v1alpha1.GetMetricCountRequest]) (*connect.Response[v1alpha1.GetMetricCountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.MetricsService.GetMetricCount is not implemented"))
}
