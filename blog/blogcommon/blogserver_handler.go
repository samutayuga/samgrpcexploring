package blogcommon

import (
	"context"
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"github.com/samutayuga/samgrpcexploring/sandra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	blogKs    = "samdb"
	blogTable = "blog_item"
)

//Server ...
type Server struct {
}

//CreateBlog ...
func (s *Server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	bRaw := req.GetBlog()
	bID := gocql.TimeUUID()
	bRaw.Id = bID.String()
	//DB
	if err := sandra.Csess.Query(`INSERT INTO blog_item (id,author_id,title,content) VALUES(?,?,?,?)`,
		bID, bRaw.GetAuthorId(), bRaw.GetTitle(), bRaw.GetContent()).Exec(); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while inserting a blog with author= %s,title=%s, %v", bRaw.GetAuthorId(), bRaw.GetTitle(), err))
		//log.Fatalf("Error while inserting a blog with author= %s,title=%s, %v", bRaw.GetAuthorId(), bRaw.GetTitle(), err)
	}
	newBlog := blogpb.CreateBlogResponse{Blog: bRaw}

	return &newBlog, nil
}

//FindBlogByID ..
func FindBlogByID(blogID gocql.UUID) (*blogpb.Blog, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var errReading error
	b := blogpb.Blog{}

	if errReading = sandra.Csess.Query(`SELECT author_id,title,content FROM blog_item WHERE id=?`, blogID).Consistency(gocql.One).Scan(
		&b.AuthorId, &b.Title, &b.Content); errReading == nil {

		return &b, nil

	}
	if errReading.Error() == "not found" {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Blog with id=%v is not found", blogID))
	}
	return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while retrieving id=%v.Server Error: %v", blogID, errReading))

}

//ReadBlog ...
func (s *Server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var bID gocql.UUID
	var errParse, errReading error
	var blg *blogpb.Blog
	//b := &blogpb.Blog{}

	if bID, errParse = gocql.ParseUUID(req.GetBlogId()); errParse == nil {

		if blg, errReading = FindBlogByID(bID); errReading == nil {
			blg.Id = req.GetBlogId()
			resp := blogpb.ReadBlogResponse{Blog: blg}
			return &resp, nil
		}
		return nil, errReading
	}
	return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse the ID %s, err=%v", req.GetBlogId(), errParse))

}

//UpdateBlog search if an ID exists in DB
func (s *Server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var anyEror error
	var bID gocql.UUID
	var oldBlg *blogpb.Blog
	log.Printf("Update blog %v", req.GetBlog())
	if bID, anyEror = gocql.ParseUUID(req.GetBlog().GetId()); anyEror == nil {
		log.Printf("Got blog id %v", bID)
		if oldBlg, anyEror = FindBlogByID(bID); anyEror == nil {
			//update
			log.Printf("Found existing blog %v", oldBlg)
			if anyEror = sandra.Csess.Query(`UPDATE blog_item set content=? WHERE id=? and author_id=?`,
				req.GetBlog().GetContent(),
				bID,
				req.GetBlog().GetAuthorId()).Exec(); anyEror == nil {
				log.Printf("Update successfull %s", req.GetBlog().GetContent())
				return &blogpb.UpdateBlogResponse{Blog: req.GetBlog()}, nil
			}
		}
	}
	return nil, anyEror

}

//DeleteBlog ...
func (s *Server) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogRequest) (*blogpb.DeleteBlogResponse, error) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var anyEror error
	var bID gocql.UUID
	var oldBlg *blogpb.Blog
	log.Printf("Delete blog %v", req.GetBlogId())
	if bID, anyEror = gocql.ParseUUID(req.GetBlogId()); anyEror == nil {
		log.Printf("Got blog id %v", bID)
		if oldBlg, anyEror = FindBlogByID(bID); anyEror == nil {
			//update
			log.Printf("Found existing blog %v", oldBlg)
			if anyEror = sandra.Csess.Query(`DELETE FROM blog_item WHERE id=?`,
				bID).Exec(); anyEror == nil {
				log.Printf("Delete successfull %s", req.GetBlogId())
				return &blogpb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
			}
		}
	}
	return nil, anyEror
}

//ListBlog ...
func (s *Server) ListBlog(req *blogpb.ListBlogRequest, stream blogpb.BlogService_ListBlogServer) error {
	iter := sandra.Csess.Query(`SELECT id,author_id,title,content FROM blog_item`).Iter()
	defer iter.Close()
	var blogID gocql.UUID
	blgFromDB := blogpb.Blog{}
	for iter.Scan(&blogID, &blgFromDB.AuthorId, &blgFromDB.Title, &blgFromDB.Content) {
		blgFromDB.Id = blogID.String()
		stream.Send(&blogpb.ListBlogResponse{Blog: &blgFromDB})
	}

	return nil

}

//GetKeySpace ...
func GetKeySpace() string {
	return blogKs
}

//GetBlogTable ...
func GetBlogTable() string {
	return blogTable
}
func init() {
	log.Println("initialize the blog common")
}
