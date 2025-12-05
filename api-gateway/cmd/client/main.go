package main

import (
	"apigateway/internal/bootstrap"
	"context"
	"fmt"
	"log"
	"metadatasvc/gen/go/commonpb"
	"metadatasvc/gen/go/staticpagepb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := bootstrap.Init(); err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", bootstrap.Config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	staticPageClient := staticpagepb.NewStaticPageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := staticPageClient.Create(ctx, &staticpagepb.CreateRequest{
		Page: &staticpagepb.StaticPage{
			Title:     "Test Static Page 1",
			Slug:      "test-static-page 1",
			Content:   "<h1>This is a test static page</h1>",
			PageType:  commonpb.PageType_PAGE_TYPE_HEADER,
			SortOrder: 1,
			Seo: &commonpb.SEO{
				MetaTitle:    "seo meta title",
				MetaDesc:     "seo meta desc",
				MetaKeywords: []string{"kw1"},
			},
			Status: commonpb.Status_STATUS_ACTIVE,
		},
	})
	fmt.Println("errrrrrrrr", err)
	fmt.Println("resp", resp)
}
