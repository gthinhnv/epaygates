package staticpageservice

import (
	"context"
	"metadatasvc/gen/go/staticpagepb"
	"metadatasvc/internal/bootstrap"
	"shared/models/staticpagemodel"
	"shared/pkg/utils/dbutil"
	"shared/pkg/utils/grpcutil"

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
	page := req.GetPage()

	var pageModel staticpagemodel.StaticPage
	if err := dbutil.MapStruct(page, &pageModel); err != nil {
		return nil, status.Error(codes.InvalidArgument, "map struct issue")
	}

	if err := bootstrap.Validate.Struct(pageModel); err != nil {
		return nil, grpcutil.BuildValidationError(err)
	}

	id, err := bootstrap.Repos.StaticPageRepo.Create(ctx, &pageModel)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create static page: %v", err)
	}

	return &staticpagepb.CreateResponse{
		Id: id,
	}, nil
}

func (s *StaticPageServiceServer) Update(ctx context.Context, req *staticpagepb.UpdateRequest) (*staticpagepb.UpdateResponse, error) {
	if len(req.Fields) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no fields need to update")
	}

	page := req.GetPage()

	pageDB, err := bootstrap.Repos.StaticPageRepo.GetByID(ctx, page.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "not found page")
	}

	for _, f := range req.Fields {
		switch f {
		case "title":
			pageDB.Title = req.Page.Title
		}
	}

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

	var pageProto staticpagepb.StaticPage
	if err := dbutil.MapStruct(page, &pageProto); err != nil {
		return nil, status.Error(codes.Internal, "failed to map struct")
	}

	return &staticpagepb.GetResponse{
		Page: &pageProto,
	}, nil
}

func (s *StaticPageServiceServer) List(ctx context.Context, req *staticpagepb.ListRequest) (*staticpagepb.ListResponse, error) {
	pages, err := bootstrap.Repos.StaticPageRepo.List(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get static page: %v", err)
	}

	pageProtos := make([]*staticpagepb.StaticPage, len(pages))
	for i, page := range pages {
		var pageProto staticpagepb.StaticPage

		if err := dbutil.MapStruct(page, &pageProto); err != nil {
			return nil, err
		}

		pageProtos[i] = &pageProto
	}

	return &staticpagepb.ListResponse{
		Pages: pageProtos,
	}, nil
}
