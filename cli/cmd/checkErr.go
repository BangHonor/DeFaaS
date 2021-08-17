package cmd

import (
	"fmt"
	"os"
)

// CheckErr prints the msg with the prefix 'Error:' and exits with error code 1. If the msg is nil, it does nothing.
// https://github.com/spf13/cobra/blob/8380ddd3132bdf8fd77731725b550c181dda0aa8/cobra.go#L211
func CheckErr(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}
