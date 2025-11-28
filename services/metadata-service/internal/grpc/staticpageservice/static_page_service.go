package staticpageservice

import (
	"context"
	"fmt"
	"metadatasvc/gen/go/staticpagepb"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StaticPageServiceServer struct {
	staticpagepb.UnimplementedStaticPageServiceServer
}

func NewStaticPageServiceServer() *StaticPageServiceServer {
	return &StaticPageServiceServer{}
}

func (s *StaticPageServiceServer) Create(ctx context.Context, req *staticpagepb.CreateRequest) (*staticpagepb.CreateResponse, error) {
	fmt.Println("____________Create a new static page______________")
	if err := protovalidate.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
	}
	return &staticpagepb.CreateResponse{}, nil
}
