package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-tags-go/version1"
	"github.com/stretchr/testify/assert"
)

type TagsClientFixtureV1 struct {
	Client version1.ITagsClientV1
	TAGS   *version1.PartyTagsV1
}

func NewTagsClientFixtureV1(client version1.ITagsClientV1) *TagsClientFixtureV1 {
	return &TagsClientFixtureV1{
		Client: client,
		TAGS: version1.NewPartyTagsV1("1", []*version1.TagRecordV1{
			version1.NewTagRecordV1("tag1", 10),
			version1.NewTagRecordV1("Tag 2", 2),
			version1.NewTagRecordV1("TAG3", 4),
		}),
	}
}

func (c *TagsClientFixtureV1) TestGetAndSetTags(t *testing.T) {
	// Update party tags
	partyTags, err := c.Client.SetTags(context.Background(), "", c.TAGS)
	assert.Nil(t, err)

	assert.Len(t, partyTags.Tags, 3)

	// Read and check party tags
	partyTags, err = c.Client.GetTags(context.Background(), "123", "1")
	assert.Nil(t, err)

	assert.Len(t, partyTags.Tags, 3)
}

func (c *TagsClientFixtureV1) TestRecordTags(t *testing.T) {
	// Record tags first time
	partyTags, err := c.Client.RecordTags(context.Background(), "", "1", []string{"tag1", "tag 2", "tag_3"})
	assert.Nil(t, err)

	assert.Len(t, partyTags.Tags, 3)

	// Record tags second time
	partyTags, err = c.Client.RecordTags(context.Background(), "", "1", []string{"TAG2", "tag3", "tag__4"})
	assert.Nil(t, err)

	assert.Len(t, partyTags.Tags, 4)

	// Get tags
	partyTags, err = c.Client.GetTags(context.Background(), "123", "1")
	assert.Nil(t, err)

	assert.Len(t, partyTags.Tags, 4)
}
