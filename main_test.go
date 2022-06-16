package main

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
)

func TestPipeline(t *testing.T) {
	// Note that the order of the elements doesn't matter.
	result := my_pipeline("Test")
	passert.Equals(result.Scope, result.PCollection, "Test", "Hello", "World!")
	ptest.RunAndValidate(t, result.Pipeline)
}
