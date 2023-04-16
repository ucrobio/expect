package expect

import (
	"testing"

	. "spec.ucrob.io"
)

func TestExpectFunIntegrationSuite(t *testing.T) {
	Run(
		Define(
			"fun integration suite",

			Describe(
				"error matchers",

				Describe("BeError() matcher"),
				Describe("ContainsError(...) matcher"),
				Describe("IsError(...) matcher"),
			),
		),

		func(err error) { t.Fatal(err) },
	)
}
