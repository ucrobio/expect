package expect

import (
	"fmt"
	"testing"

	. "spec.ucrob.io"
)

func TestExpectFunIntegrationSuite(t *testing.T) {
	firstNumber := Var(0)
	secondNumber := Var(0)

	err := Var[error](nil)

	Run(
		Define(
			"fun integration suite",

			Describe(
				"BeEq()",

				Let(firstNumber, func() int { return 2 }),
				Let(secondNumber, func() int { return 5 }),

				Describe(
					"positive case",

					Inline(func() error { return Value(firstNumber.Value() + secondNumber.Value()).To(BeEq(7)) }),

					Describe(
						"fail case",

						Inline(func() error {
							return Error(Value(firstNumber.Value() + secondNumber.Value()).To(BeEq(0))).To(BeError())
						}),
					),
				),

				Describe(
					"negative case",

					Inline(func() error { return Value(firstNumber.Value() + secondNumber.Value()).NotTo(BeEq(5)) }),

					Describe(
						"fail case",

						Inline(func() error {
							return Error(Value(firstNumber.Value() + secondNumber.Value()).NotTo(BeEq(7))).To(BeError())
						}),
					),
				),
			),

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
					"negative case",

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
