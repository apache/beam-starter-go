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

type Result = struct {
	Pipeline    *beam.Pipeline
	Scope       beam.Scope
	PCollection beam.PCollection
}

func my_pipeline(input_text string) Result {
	beam.Init()

	pipeline, scope := beam.NewPipelineWithRoot()

	elements := beam.Create(scope, "Hello", "World!", input_text)
	elements = beam.ParDo(scope, func(elem string, emit func(string)) { fmt.Println(elem); emit(elem) }, elements)

	return Result{pipeline, scope, elements}
}

func main() {
	flag.Parse()

	ctx := context.Background()
	result := my_pipeline(*input_text)

	// Run the pipeline. You can specify your runner with the --runner flag.
	if err := beamx.Run(ctx, result.Pipeline); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
