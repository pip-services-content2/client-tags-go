package build

import (
	clients1 "github.com/pip-services-content2/client-tags-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type TagsClientFactory struct {
	*cbuild.Factory
}

func NewTagsClientFactory() *TagsClientFactory {
	c := &TagsClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-tags", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-tags", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-tags", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewTagsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewTagsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewTagsCommandableHttpClientV1)

	return c
}
