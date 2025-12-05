package staticpageservice

import (
	"context"
	"metadatasvc/gen/go/staticpagepb"
	"metadatasvc/internal/bootstrap"
	"shared/pkg/utils/grpcutil"

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
	if err := protovalidate.Validate(req); err != nil {
		return nil, grpcutil.BuildValidationError(err)
	}

	id, err := bootstrap.Repos.StaticPageRepo.Create(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create static page: %v", err)
	}

	return &staticpagepb.CreateResponse{
		Id: id,
	}, nil
}

func (s *StaticPageServiceServer) Update(ctx context.Context, req *staticpagepb.UpdateRequest) (*staticpagepb.UpdateResponse, error) {
	return &staticpagepb.UpdateResponse{}, nil
}

func (s *StaticPageServiceServer) Delete(ctx context.Context, req *staticpagepb.DeleteRequest) (*staticpagepb.DeleteResponse, error) {
	return &staticpagepb.DeleteResponse{}, nil
}

func (s *StaticPageServiceServer) Get(ctx context.Context, req *staticpagepb.GetRequest) (*staticpagepb.GetResponse, error) {
	page, err := bootstrap.Repos.StaticPageRepo.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get static page: %v", err)
	}

	return &staticpagepb.GetResponse{
		Page: page,
	}, nil
}

func (s *StaticPageServiceServer) List(ctx context.Context, req *staticpagepb.ListRequest) (*staticpagepb.ListResponse, error) {
	pages, err := bootstrap.Repos.StaticPageRepo.List(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get static page: %v", err)
	}

	return &staticpagepb.ListResponse{
		Pages: pages,
	}, nil
}
