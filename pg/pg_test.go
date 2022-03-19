package pg

import (
	"github.com/google/uuid"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	assert := assert2.New(t)
	newId := uuid.New().String()
	author := "sam123"
	newBl := BlogItem{newId,
		"sam123", "This is a new era of agriculture", "Testing"}
	assert.Equalf("OK", PingDb(), "should be OK if successful")
	assert.Equalf(1, InsertBlog(&BlogItem{newId,
		author, "This is a new era of agriculture", "Testing"}),
		"should return 1")
	assert.Equalf(&newBl, SelectBlogById(newId), "should return a  record")
	//update the content to be something else
	newContent := "This is about agriculture in the sky"
	assert.Equalf(1, UpdateBlogWithAffectedRecords(newId, author, newContent), "Should successfully update the content")
	//retrieve it back
	newBl.Content = newContent
	assert.Equalf(&newBl, SelectBlogById(newId), "The value should be the same")
	//use another method to update
	newContent = "The agriculture is now better than ever..."
	assert.Equalf(1, UpdateBlogWithReturn(newId, author, newContent), "Should successfully update the content")
	//retrieve it back
	newBl.Content = newContent
	assert.Equalf(&newBl, SelectBlogById(newId), "The value should be the same")

	assert.Equalf(1, DeleteSingleBlog(newId),
		"should return 1")
	assert.Equalf(&BlogItem{}, SelectBlogById(newId), "should return  no record")
}

func TestListing(t *testing.T) {
	//delete all
	assert := assert2.New(t)
	allBls := SelectAll()
	for _, aBl := range allBls {
		assert.Equalf(1, DeleteSingleBlog(aBl.BlogId), "OK")
	}

	newBl1 := BlogItem{uuid.New().String(),
		"sam123", "This is a new era of agriculture1", "Testing1"}
	newBl2 := BlogItem{uuid.New().String(),
		"sam123", "This is a new era of agriculture2", "Testing2"}
	newBl3 := BlogItem{uuid.New().String(),
		"sam123", "This is a new era of agriculture3", "Testing3"}

	assert.Equalf(1, InsertBlog(&newBl1), "OK")
	assert.Equalf(1, InsertBlog(&newBl2), "OK")
	assert.Equalf(1, InsertBlog(&newBl3), "OK")

	allBlNews := SelectAll()
	assert.Equal(3, len(allBlNews))

}
