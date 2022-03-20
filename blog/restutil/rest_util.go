package restutil

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	gw "github.com/samutayuga/samgrpcexploring/blog/blogpb"
	"github.com/samutayuga/samgrpcexploring/pg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	BlogClient gw.BlogServiceClient
)

func CreateBlog(writer http.ResponseWriter, request *http.Request) {
	//handle the creation of the blog
	writer.Header().Set("Content-Type", "application/json;charset=UFT-8")

	if reqBody, errReqBody := ioutil.ReadAll(request.Body); errReqBody == nil {
		payload := pg.BlogItem{}

		if errUnm := json.Unmarshal(reqBody, &payload); errUnm == nil {
			bId := uuid.New().String()
			payload.BlogId = bId
			ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
			defer cancel()
			if cresponse, err := BlogClient.CreateBlog(ctx, &gw.CreateBlogRequest{
				Blog: &gw.Blog{Id: payload.BlogId, AuthorId: payload.Author,
					Title: payload.Title, Content: payload.Content}}); err != nil {
				log.Fatalf("Error while creating the blog %v", err)
			} else {
				log.Printf("Blog is created %v", cresponse)
				//build the good response to end user
				if b, mErr := json.Marshal(payload); mErr != nil {
					log.Printf("error while marshalling response to json %v", mErr)
					writer.WriteHeader(http.StatusInternalServerError)
				} else {
					writer.WriteHeader(http.StatusOK)
					writer.Write(b)
				}
			}
		} else {
			log.Printf("error while unmarshall the request payload %v", errUnm)
			writer.WriteHeader(http.StatusInternalServerError)
		}

	} else {
		log.Printf("error while reading the request payload %v", errReqBody)
		writer.WriteHeader(http.StatusBadRequest)
	}
}
func ListBlog(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json;charset=UFT-8")
	//handle the listing
	allBls := make([]pg.BlogItem, 0)
	stream, err := BlogClient.ListBlog(context.Background(), &gw.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling list blog %v", err)
	}
	for {
		res, rcvErr := stream.Recv()
		if rcvErr == io.EOF {
			break
		}
		if rcvErr != nil {
			log.Fatalf("Unexpected error %v\n", rcvErr)
		}
		//good section
		blRs := res.GetBlog()
		log.Printf("receiving %v\n", res.GetBlog())
		allBls = append(allBls, pg.BlogItem{
			BlogId:  blRs.GetId(),
			Author:  blRs.GetAuthorId(),
			Title:   blRs.GetTitle(),
			Content: blRs.GetContent(),
		})
	}
	if b, mErr := json.Marshal(allBls); mErr != nil {
		log.Printf("error while marshalling response to json %v", mErr)
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write(b)
	}
}
func UpdateBlog(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json;charset=UFT-8")
	if reqBody, errReqBody := ioutil.ReadAll(request.Body); errReqBody == nil {
		payload := pg.BlogItem{}

		if errUnm := json.Unmarshal(reqBody, &payload); errUnm == nil {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
			defer cancel()
			//call grpc service
			//need to isolate
			if cresponse, err := BlogClient.UpdateBlog(ctx, &gw.UpdateBlogRequest{
				Blog: &gw.Blog{Id: payload.BlogId, AuthorId: payload.Author,
					Title: payload.Title, Content: payload.Content}}); err != nil {
				log.Printf("Error while updating the blog %v\n", err)
				if st, ok := status.FromError(err); ok {
					if st.Code() == codes.NotFound {
						log.Printf("Blog with id  %s is not found", payload.BlogId)
						writer.WriteHeader(http.StatusNotFound)
					} else {
						writer.WriteHeader(http.StatusInternalServerError)
					}

				} else {
					writer.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				log.Printf("Blog is updated %v", cresponse)
				//build the good response to end user
				if b, mErr := json.Marshal(payload); mErr != nil {
					log.Printf("error while marshalling response to json %v", mErr)
					writer.WriteHeader(http.StatusInternalServerError)
				} else {
					writer.WriteHeader(http.StatusOK)
					writer.Write(b)
				}
			}
		}
	} else {
		log.Printf("error while reading the request payload %v", errReqBody)
		writer.WriteHeader(http.StatusBadRequest)
	}

}
func ProcessASingleBlog(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	if blogId, exists := vars["blogId"]; exists {
		//this is the delete or get by id request
		switch request.Method {
		case "DELETE":
			//delete blog
			resDel, errDel := BlogClient.DeleteBlog(context.Background(), &gw.DeleteBlogRequest{BlogId: blogId})
			if errDel == nil {
				log.Printf("Successfully Delete the blog %v", resDel)
				writer.WriteHeader(http.StatusOK)
			} else {
				log.Printf("Error while Delete the err=%v", errDel)
				if st, ok := status.FromError(errDel); ok {
					if st.Code() == codes.NotFound {
						log.Printf("Blog with id  %s is not found", blogId)
						writer.WriteHeader(http.StatusNotFound)
					} else {
						writer.WriteHeader(http.StatusInternalServerError)
					}

				} else {
					log.Printf("Unexpected error %v", errDel)
					writer.WriteHeader(http.StatusInternalServerError)
				}
			}
		case "GET":
			var successReadErr error
			var b *gw.ReadBlogResponse
			if b, successReadErr = BlogClient.ReadBlog(context.Background(), &gw.ReadBlogRequest{BlogId: blogId}); successReadErr == nil {
				log.Printf("found %v", b)
				existingBl := b.GetBlog()
				bItem := pg.BlogItem{
					BlogId:  existingBl.GetId(),
					Author:  existingBl.GetAuthorId(),
					Title:   existingBl.GetTitle(),
					Content: existingBl.GetContent(),
				}

				if blItemSer, mErr := json.Marshal(bItem); mErr != nil {
					log.Printf("error while marshalling response to json %v", mErr)
					writer.WriteHeader(http.StatusInternalServerError)
				} else {
					writer.WriteHeader(http.StatusOK)
					writer.Write(blItemSer)
				}
			}

			if successReadErr != nil {
				if st, ok := status.FromError(successReadErr); ok {
					if st.Code() == codes.NotFound {
						log.Printf("Blog with id  %s is not found", blogId)
						writer.WriteHeader(http.StatusNotFound)
					} else {
						writer.WriteHeader(http.StatusInternalServerError)
					}

				} else {
					log.Printf("Unexpected error %v", successReadErr)
					writer.WriteHeader(http.StatusInternalServerError)
				}

			} else {
				writer.WriteHeader(http.StatusInternalServerError)
			}

		default:
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
