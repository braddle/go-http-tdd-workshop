package pact_test

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/types"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/suite"
)

type PactProviderSuite struct {
	suite.Suite
	pact dsl.Pact
}

func TestPactProviderSuite(t *testing.T) {
	p := new(PactProviderSuite)
	p.pact = dsl.Pact{
		Consumer: "Book-Consumer",
		Provider: "Book-Provider",
		Host:     "localhost",
	}

	suite.Run(t, p)
}

func (s *PactProviderSuite) TestPacts() {
	s.T().Skip("PACT FUN. issues with response body matching!")
	res, _ := s.pact.VerifyProvider(
		s.T(),
		types.VerifyRequest{
			ProviderBaseURL:            "http://http:8080",
			PactURLs:                   []string{"http://pact_broker:9292/pacts/provider/Book-Provider/consumer/Book-Consumer/latest/dev"},
			FailIfNoPactsFound:         true,
			PublishVerificationResults: true,    // TODO should come from an ENV var
			ProviderVersion:            "1.0.0", // TODO should come from an env var
			Verbose:                    true,
		},
	)

	fmt.Printf("SUMMARY: %s\n", res.SummaryLine)
	for _, exp := range res.Examples {
		if exp.Status == "failed" {
			fmt.Printf("DESCRIPTION: %s\n", exp.Description)
			fmt.Printf("FULL DESCRIPTION: %s\n", exp.FullDescription)
			fmt.Printf("PENDING MESSAGE: %v\n", exp.PendingMessage)
			fmt.Printf("STATUS: %s\n", exp.Status)
			fmt.Printf("EXEPCTION CLASS: %s\n", exp.Exception.Class)
			fmt.Printf("EXEPCTION MESSAGE: %s\n", exp.Exception.Message)
			//fmt.Printf("EXEPCTION MESSAGE: %s\n", exp.Exception.)

			fmt.Print("_________________________________________\n")
			fmt.Print("_________________________________________\n\n\n\n\n\n\n\n")
		}

	}
}
