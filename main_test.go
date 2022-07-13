// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// https://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or https://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

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
