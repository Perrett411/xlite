// Command quantum-gate runs the quantum consensus function and exits non-zero
// unless consensus is validated. Workflows use it to gate automated changes.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cli/cli/v2/pkg/quantum"
)

func main() {
	load := flag.Float64("load", 1, "current load")
	maxLoad := flag.Float64("max-load", 1, "maximum load")
	flag.Parse()

	q := quantum.New()
	nodes, err := q.AdjustNodes(*load, *maxLoad)
	if err != nil {
		fmt.Fprintln(os.Stderr, "quantum gate:", err)
		os.Exit(1)
	}
	state, err := q.EntangleNodes(nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "quantum gate:", err)
		os.Exit(1)
	}
	if !q.ValidateConsensus(state) {
		fmt.Fprintf(os.Stderr, "quantum gate: consensus failed with %d nodes\n", nodes)
		os.Exit(1)
	}
	fmt.Printf("quantum gate: consensus validated with %d nodes\n", nodes)
}
