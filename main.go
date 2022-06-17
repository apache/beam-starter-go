package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	beamLog "github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

var (
	input_text = flag.String("input-text", "default input text", "Input text to print.")
)

func init() {
	// DoFns should be registered with Beam to be available in distributed runners.
	beam.RegisterFunction(toTitleCase)
	beam.RegisterFunction(logAndEmit)
}

// A simple function would take the element as an argument and return a single value.
func toTitleCase(element string) string {
	return strings.Title(element)
}

// You can also access the Context and "emit" zero or more values like FlatMap.
func logAndEmit(ctx context.Context, element string, emit func(string)) {
	beamLog.Infoln(ctx, element)
	emit(element)
}

func MyPipeline(scope beam.Scope, input_text string) beam.PCollection {
	elements := beam.Create(scope, "hello", "world!", input_text)
	elements = beam.ParDo(scope, toTitleCase, elements)
	return beam.ParDo(scope, logAndEmit, elements)
}

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	pipeline, scope := beam.NewPipelineWithRoot()
	MyPipeline(scope, *input_text)

	// Run the pipeline. You can specify your runner with the --runner flag.
	if err := beamx.Run(ctx, pipeline); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
