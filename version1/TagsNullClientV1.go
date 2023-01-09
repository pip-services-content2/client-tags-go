package version1

import "context"

type TagsNullClientV1 struct {
}

func NewTagsNullClientV1() *TagsNullClientV1 {
	return &TagsNullClientV1{}
}

func (c *TagsNullClientV1) GetTags(ctx context.Context, correlationId string, partyId string) (*PartyTagsV1, error) {
	return nil, nil
}

func (c *TagsNullClientV1) SetTags(ctx context.Context, correlationId string, partyTags *PartyTagsV1) (*PartyTagsV1, error) {
	return partyTags, nil
}

func (c *TagsNullClientV1) RecordTags(ctx context.Context, correlationId string, partyId string, tags []string) (*PartyTagsV1, error) {
	return nil, nil
}
