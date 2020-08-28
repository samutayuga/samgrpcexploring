package main

import (
	"context"
	"io"
	"log"

	"github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("This is client ")
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Error while dialing %v", err)
	}
	defer conn.Close()
	c := blogpb.NewBlogServiceClient(conn)
	cresponse, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{AuthorId: "Jack",
			Title: "This is a blog", Content: "This is content"}})
	if err != nil {
		log.Fatalf("Error while creating the blog %v", err)
	}
	log.Printf("Blog is created %v", cresponse)

	if _, errRead := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "blogID"}); errRead != nil {
		log.Printf("Error expected as the blogid is invalid %v", errRead)
	}

	//993651b1-e813-11ea-9935-54270a1c270d
	if _, readNotFound := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "993651b1-e813-11ea-9935-54270a1c270f"}); readNotFound != nil {
		log.Printf("error %v", readNotFound)
	}
	var successReadErr error
	var b *blogpb.ReadBlogResponse
	if b, successReadErr = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: cresponse.GetBlog().GetId()}); successReadErr == nil {
		log.Printf("found %v", b)
	}
	if successReadErr != nil {
		log.Printf("Unexpected error %v", successReadErr)
	}
	//update blog
	//newBlg := cresponse.GetBlog()
	//newBlg.Content = "This is updated by client"
	b.Blog.Content = "Update content...."
	res, errUpdate := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: b.Blog})
	if errUpdate == nil {
		log.Printf("Successfully updated the blog %v", res)
	} else {
		log.Printf("Error while updating the blog %v,err=%v", res, errUpdate)
	}
	//delete blog
	resDel, errDel := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: b.GetBlog().GetId()})
	if errDel == nil {
		log.Printf("Successfully Delete the blog %v", resDel)
	} else {
		log.Printf("Error while Delete the err=%v", errDel)
	}
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling list blog %v", err)
	}
	for {
		res, rcvErr := stream.Recv()
		if rcvErr == io.EOF {
			break
		}
		if rcvErr != nil {
			log.Fatalf("Unexpected error %v", rcvErr)
		}
		//good section
		log.Printf("receiving %v", res.GetBlog())
	}
}
