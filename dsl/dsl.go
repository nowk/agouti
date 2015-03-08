// Package dsl is DEPRECATED. Use "github.com/sclevine/agouti" instead.
package dsl

import (
	"fmt"
	"os"

	"github.com/onsi/ginkgo"
)

func init() {
	if os.Getenv("SUPPRESS_DEPRECATION_NOTICE") == "true" {
		return
	}
	fmt.Fprintln(os.Stderr, `****************
NOTICE: The "github.com/sclevine/agouti/dsl" package is deprecated and may
eventually perish. Importing "github.com/sclevine/agouti" and dot-importing
"github.com/sclevine/agouti/matchers" is the prefered way of writing
integration or acceptance tests when using Agouti with Ginkgo/Gomega.
****************`)
}

// Background is equivalent to Ginkgo BeforeEach.
func Background(body interface{}, timeout ...float64) bool {
	return ginkgo.BeforeEach(body, timeout...)
}

// Feature is equivalent to Ginkgo Describe.
func Feature(text string, body func()) bool {
	return ginkgo.Describe(text, body)
}

// FFeature is equilavent to Ginkgo FDescribe (Focused Describe).
func FFeature(text string, body func()) bool {
	return ginkgo.FDescribe(text, body)
}

// PFeature is equilavent to Ginkgo PDescribe (Pending Describe).
func PFeature(text string, body func()) bool {
	return ginkgo.PDescribe(text, body)
}

// XFeature is equilavent to Ginkgo XDescribe (Pending Describe).
func XFeature(text string, body func()) bool {
	return ginkgo.XDescribe(text, body)
}

// Scenario is equivalent to Ginkgo It.
func Scenario(description string, body func(), timeout ...float64) bool {
	return ginkgo.It(description, body, timeout...)
}

// FScenario is equivalent to Ginkgo FIt (Focused It).
func FScenario(description string, body func(), timeout ...float64) bool {
	return ginkgo.FIt(description, body, timeout...)
}

// PScenario is equivalent to Ginkgo PIt (Pending It).
func PScenario(description string, ignored ...interface{}) bool {
	return ginkgo.PIt(description, ignored...)
}

// XScenario is equivalent to Ginkgo XIt (Pending It).
func XScenario(description string, ignored ...interface{}) bool {
	return ginkgo.XIt(description, ignored...)
}

// Step is equivalent to Ginkgo By.
func Step(text string, callbacks ...func()) {
	ginkgo.By(text, callbacks...)
}
