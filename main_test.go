package main

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
)

func TestPipeline(t *testing.T) {
	beam.Init()

	pipeline, scope := beam.NewPipelineWithRoot()

	// Note that the order of the elements doesn't matter.
	elements := myPipeline(scope, "Test")
	passert.Equals(scope, elements, "Test", "Hello", "World!")
	ptest.RunAndValidate(t, pipeline)
}
