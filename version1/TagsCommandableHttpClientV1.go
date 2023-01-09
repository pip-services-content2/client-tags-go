package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type TagsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewTagsCommandableHttpClientV1() *TagsCommandableHttpClientV1 {
	return &TagsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/tags"),
	}
}

func (c *TagsCommandableHttpClientV1) GetTags(ctx context.Context, correlationId string, partyId string) (*PartyTagsV1, error) {
	res, err := c.CallCommand(ctx, "get_tags", correlationId, data.NewAnyValueMapFromTuples(
		"party_id", partyId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*PartyTagsV1](res, correlationId)
}

func (c *TagsCommandableHttpClientV1) SetTags(ctx context.Context, correlationId string, partyTags *PartyTagsV1) (*PartyTagsV1, error) {
	res, err := c.CallCommand(ctx, "set_tags", correlationId, data.NewAnyValueMapFromTuples(
		"party_tags", partyTags,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*PartyTagsV1](res, correlationId)
}

func (c *TagsCommandableHttpClientV1) RecordTags(ctx context.Context, correlationId string, partyId string, tags []string) (*PartyTagsV1, error) {
	res, err := c.CallCommand(ctx, "record_tags", correlationId, data.NewAnyValueMapFromTuples(
		"party_id", partyId,
		"tags", tags,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*PartyTagsV1](res, correlationId)
}
