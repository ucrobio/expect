package expect

import (
	"fmt"
	"testing"

	. "spec.ucrob.io"
)

func TestExpectFunIntegrationSuite(t *testing.T) {
	err := Var[error](nil)

	Run(
		Define(
			"fun integration suite",

			Describe(
				"BeError()",

				Describe(
					"positive case",

					Let(err, func() error { return fmt.Errorf("err") }),

					Inline(func() error { return Error(err.Value()).To(BeError()) }),

					Describe(
						"fail case",

						Let(err, func() error { return nil }),

						Inline(func() error { return Error(Error(err.Value()).To(BeError())).To(BeError()) }),
					),
				),

				Describe(
					"positive case",

					Let(err, func() error { return nil }),

					Inline(func() error { return Error(err.Value()).NotTo(BeError()) }),

					Describe(
						"fail case",

						Let(err, func() error { return fmt.Errorf("err") }),

						Inline(func() error { return Error(Error(err.Value()).NotTo(BeError())).To(BeError()) }),
					),
				),
			),
		),

		func(err error) { t.Fatal(err) },
	)
}
