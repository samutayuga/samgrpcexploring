package blogcommon

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"github.com/samutayuga/samgrpcexploring/pg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// PgBlogServer PgBlogServcer ...
type PgBlogServer struct {
}

//
//func FindBlogById(blogId string) (*blogpb.Blog, error) {
//
//}
func (s *PgBlogServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {

	bRaw := req.GetBlog()
	log.Printf("request has payload %v\n", bRaw)
	bID := uuid.New().String()
	countRec := pg.InsertBlog(&pg.BlogItem{BlogId: bID,
		Author: bRaw.AuthorId,
		Title:  bRaw.Title, Content: bRaw.Content})
	if countRec == 1 {
		return &blogpb.CreateBlogResponse{Blog: bRaw}, nil
	}
	return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while inserting a blog with author= %s,title=%s", bRaw.GetAuthorId(), bRaw.GetTitle()))
}

//ReadBlog ....
func (s *PgBlogServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	if bl := pg.SelectBlogById(req.BlogId); bl == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Blog with id=%v is not found", req.BlogId))
	} else {
		return &blogpb.ReadBlogResponse{Blog: &blogpb.Blog{Id: bl.BlogId, AuthorId: bl.Author, Title: bl.Title, Content: bl.Content}}, nil
	}

}

//
func (s *PgBlogServer) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	oldBlog := req.GetBlog()

	if count := pg.UpdateBlogWithAffectedRecords(oldBlog.Id, oldBlog.AuthorId, oldBlog.Content); count == 1 {
		return &blogpb.UpdateBlogResponse{Blog: req.GetBlog()}, nil
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Can not update blog %s", req.Blog.Id))
}

//
func (s *PgBlogServer) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	if count := pg.DeleteSingleBlog(req.BlogId); count == 1 {
		return &blogpb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot delete record for id %s, because it is not found", req.GetBlogId()))
}

//
func (s *PgBlogServer) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	for _, ab := range pg.SelectAll() {
		stream.Send(&blogpb.ListBlogResponse{Blog: &blogpb.Blog{Id: ab.BlogId, AuthorId: ab.Author, Title: ab.Title, Content: ab.Content}})
	}

	return nil
}
