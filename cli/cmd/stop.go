package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop batman",
	Long:    `Stopping batman.`,
	Args:    cobra.NoArgs,
	Example: "batman stop",
	RunE:    stop,
}

func init() {
	rootCmd.AddCommand(stopCmd)
}

func stop(cmd *cobra.Command, args []string) error {

	removeCache := func() {
		err := os.Remove(Fbatmanfile)
		if err != nil {
			log.Fatalf("Could not remove batman cache file: %v", err)
		}
	}

	data, err := os.ReadFile(Fbatmanfile)
	if err != nil {
		log.Fatalf("Could not read batman cache file: %v", err)
	}

	pid := string(data)

	if pid == "" {
		removeCache()
		log.Fatalf("No valid batman pid")
	}

	id, err := strconv.Atoi(pid)
	if err != nil {
		removeCache()
		log.Fatalf("No valid batman pid: %v", err)
	}

	err = syscall.Kill(id, 2)
	if err != nil {
		// if errors.Is(err, syscall.ErrNoSuchProcess) {
		removeCache()
		// }
		log.Fatalf("Could not kill batman pid (%s): %v", pid, err)
	}

	fmt.Println("Succesfully terminated batman!")
	removeCache()

	return nil
}
