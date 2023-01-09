package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-tags-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type tagsCommandableHttpClientV1Test struct {
	client  *version1.TagsCommandableHttpClientV1
	fixture *TagsClientFixtureV1
}

func newTagsCommandableHttpClientV1Test() *tagsCommandableHttpClientV1Test {
	return &tagsCommandableHttpClientV1Test{}
}

func (c *tagsCommandableHttpClientV1Test) setup(t *testing.T) *TagsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewTagsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewTagsClientFixtureV1(c.client)

	return c.fixture
}

func (c *tagsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newTagsCommandableHttpClientV1Test()

	fixture := c.setup(t)
	fixture.TestGetAndSetTags(t)
	c.teardown(t)

	fixture = c.setup(t)
	fixture.TestRecordTags(t)
	c.teardown(t)
}
