package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

var (
	input  = flag.String("input", "perimeter/beam/kinglear.txt", "File(s) to read")
	output = flag.String("output", "", "Output file (required)")
)

func init() {
	register.DoFn3x0(&extractFn{})
	register.Function2x1(formatFn)
	register.Emitter1[string]()
}

var (
	wordRE          = regexp.MustCompile(`[a-zA-Z]+('[a-z]')?`)
	empty           = beam.NewCounter("extract", "emptyLines")
	smallWordLength = flag.Int("small_word_length", 9, "length of small words (default: 9)")
	smallWords      = beam.NewCounter("extract", "smallWords")
	lineLen         = beam.NewDistribution("extract", "lineLenDistro")
)

// extractFn is a structural DoFn that emits the words in a given line and keeps
// a count for small words. Its ProcessElement function will be invoked on each
// element in the input PCollection
type extractFn struct {
	SmallWordLength int `json:"smallWordLength"`
}

func (f *extractFn) ProcessElement(ctx context.Context, line string, emit func(string)) {
	lineLen.Update(ctx, int64(len(line)))
	if len(strings.TrimSpace(line)) == 0 {
		empty.Inc(ctx, 1)
	}
	for _, word := range wordRE.FindAllString(line, -1) {
		if len(word) < f.SmallWordLength {
			smallWords.Inc(ctx, 1)
		}
		emit(word)
	}
}

// formatFn is a functional DoFn that formats a word and its count as a string
func formatFn(w string, c int) string {
	return fmt.Sprintf("%s: %v", w, c)
}

func CountWords(s beam.Scope, lines beam.PCollection) beam.PCollection {
	s = s.Scope("CountWords")
	//Convert lines of text into individual words
	col := beam.ParDo(s, &extractFn{SmallWordLength: *smallWordLength}, lines)
	return stats.Count(s, col)
}

func main() {
	flag.Parse()
	beam.Init()
	if *output == "" {
		log.Fatal("No output provided")
	}
	p := beam.NewPipeline()
	s := p.Root()

	lines := textio.Read(s, *input)
	counted := CountWords(s, lines)
	formatted := beam.ParDo(s, formatFn, counted)
	textio.Write(s, *output, formatted)
	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

//////////////////////////////////////////////////////////////////
// Output
/////////////////////////////////////////////////////////////////

/* time=2026-02-27T10:38:16.026+07:00 level=INFO msg="Serving JobManagement" endpoint=localhost:42539
time=2026-02-27T10:38:16.027+07:00 level=INFO msg="starting Loopback server at 127.0.0.1:42975"
time=2026-02-27T10:38:16.040+07:00 level=INFO msg="Prepared job with id: job-001 and staging token: job-001"
time=2026-02-27T10:38:16.042+07:00 level=INFO msg="Staged binary artifact with token: job-001"
time=2026-02-27T10:38:16.042+07:00 level=INFO msg="Submitted job: job-001"
time=2026-02-27T10:38:16.043+07:00 level=INFO msg="starting job-001[go-job-1-1772163496029120347]"
time=2026-02-27T10:38:16.043+07:00 level=INFO msg="running job-001[go-job-1-1772163496029120347]"
time=2026-02-27T10:38:16.043+07:00 level=INFO msg="Job[job-001] state: RUNNING"
time=2026-02-27T10:38:16.049+07:00 level=INFO msg="[SDK] Reading from perimeter/beam/kinglear.txt"
time=2026-02-27T10:38:16.167+07:00 level=INFO msg="[SDK] Writing to counts"
time=2026-02-27T10:38:16.168+07:00 level=INFO msg="pipeline done!" job=job-001[go-job-1-1772163496029120347]
time=2026-02-27T10:38:16.168+07:00 level=INFO msg="stopping worker job-001[go-job-1-1772163496029120347]_go"
time=2026-02-27T10:38:16.168+07:00 level=INFO msg="[SDK] pipeline completed job-001[go-job-1-1772163496029120347]"
time=2026-02-27T10:38:16.168+07:00 level=INFO msg="[SDK] terminating job-001[go-job-1-1772163496029120347]" */
