package staticpageservice

import (
	"context"
	"fmt"
	"metadatasvc/gen/go/staticpagepb"
)

type StaticPageServiceServer struct {
	staticpagepb.UnimplementedStaticPageServiceServer
}

func NewStaticPageServiceServer() *StaticPageServiceServer {
	return &StaticPageServiceServer{}
}

func (s *StaticPageServiceServer) Create(ctx context.Context, req *staticpagepb.CreateRequest) (*staticpagepb.CreateResponse, error) {
	fmt.Println("____________Create a new static page______________")
	return &staticpagepb.CreateResponse{}, nil
}
