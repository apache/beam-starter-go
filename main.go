package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

var (
	input_text = flag.String("input-text", "Default input text", "Input text to print.")
)

func init() {
	// DoFns should be registered with Beam.
	beam.RegisterFunction(printAndEmit)
}

func printAndEmit(element string, emit func(string)) {
	fmt.Println(element)
	emit(element)
}

func my_pipeline(scope beam.Scope, input_text string) beam.PCollection {
	elements := beam.Create(scope, "Hello", "World!", input_text)
	return beam.ParDo(scope, printAndEmit, elements)
}

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	pipeline, scope := beam.NewPipelineWithRoot()
	my_pipeline(scope, *input_text)

	// Run the pipeline. You can specify your runner with the --runner flag.
	if err := beamx.Run(ctx, pipeline); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
