package cmd

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start batman",
	Long: `Starting batman.
The first argument is the minimum treshhold for battery percentage (e.g. 20).
The second is the maximum (e.g. 95).`,
	Args:    cobra.ExactArgs(2),
	Example: "batman start 20 95",
	RunE:    start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) error {

	// Check if batman is already running
	_, err := os.Stat(Fbatmanfile)
	if err == nil {
		exec.Command("batman", "stop").Run()
	}

	// Start batman loop as a detached process
	go func(args []string) {
		out, err := exec.Command("batman", "loop", args[0], args[1]).Output()
		if err != nil {
			log.Fatalf("Could not stop previous batman: %v: %s", err, out)
		}
	}(args)

	// Give batman loop time to start
	time.Sleep(10 * time.Millisecond)

	return nil
}
