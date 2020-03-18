package consumer_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/braddle/go-http-template/consumer"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"

	"github.com/stretchr/testify/suite"
)

type ConsumerSuite struct {
	suite.Suite
	pact dsl.Pact
}

func TestConsumerSuite(t *testing.T) {
	c := new(ConsumerSuite)
	c.pact = dsl.Pact{
		Consumer: "Book-Consumer",
		Provider: "Book-Provider",
		Host:     "localhost",
		PactDir:  "../pacts",
	}

	suite.Run(t, c)

	c.pact.Teardown()

	p := dsl.Publisher{}
	err := p.Publish(types.PublishRequest{
		PactURLs:        []string{"../pacts/book-consumer-book-provider.json"},
		PactBroker:      "http://pact_broker:9292",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"master", "dev"},
	})

	if err != nil {
		t.Error(err)
	}
}

func (s *ConsumerSuite) TestHealthCheck() {

	var test = func() (err error) {
		host := fmt.Sprintf("http://localhost:%d", s.pact.Server.Port)
		c := consumer.NewClient(host)
		h, err := c.HeathCheck()

		s.Equal("OK", h.Status)
		s.Empty(h.Errors)

		return err
	}

	s.pact.AddInteraction().
		Given("The Service is healthy").
		UponReceiving("A request for health check").
		WithRequest(dsl.Request{
			Method:  http.MethodGet,
			Path:    dsl.String("/healthcheck"),
			Headers: dsl.MapMatcher{"Accept": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  http.StatusOK,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    `{"status": "OK", "errors": []}`,
		})

	s.NoError(s.pact.Verify(test))
}
