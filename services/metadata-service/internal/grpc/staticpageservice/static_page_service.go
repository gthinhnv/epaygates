package staticpageservice

import (
	"context"
	"fmt"
	"metadatasvc/gen/go/staticpagepb"
	"metadatasvc/internal/bootstrap"

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
		if ve, ok := err.(*protovalidate.ValidationError); ok {
			for _, violation := range ve.Violations {
				fmt.Println("fieldName", protovalidate.FieldPathString(violation.Proto.GetField()))
				fmt.Println("message", violation.Proto.GetMessage())
				fmt.Println("ruleId", violation.Proto.GetRuleId())
			}
		}
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %v", err)
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
