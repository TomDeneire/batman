package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	util "../../lib/util"
	"github.com/spf13/cobra"
)

var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Loop batman",
	Long: `Battery check loop for batman.
The first argument is the minimum treshhold for battery percentage (e.g. 20).
The second is the maximum (e.g. 95).
Should not be used directly! Use batman start instead!`,
	Args:    cobra.ExactArgs(2),
	Example: "batman loop 20 95",
	RunE:    loop,
}

func init() {
	rootCmd.AddCommand(loopCmd)
}

func loop(cmd *cobra.Command, args []string) error {

	// Parameters
	max, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Maximum should be an integer: %v", err)
	}
	min, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Minimum should be an integer: %v", err)
	}

	// Cache batman process id
	pid := os.Getpid()
	err = os.WriteFile(Fbatmanfile, []byte(fmt.Sprint(pid)), 0660)
	if err != nil {
		log.Fatalf("Could not write to batman cachefile: %v", err)
	}

	// Battery check loop
	for {
		time.Sleep(time.Second * 60)

		state, percentage, err := util.BatteryInfo()
		if err != nil {
			log.Fatalf("Could not get battery status: %v", err)
		}

		message := ""

		if state == "Full" {
			message = "Stop charging your laptop!"
		}
		if state == "Charging" && percentage > max {
			message = "Stop charging your laptop!"
		}
		if state == "Discharging" && percentage < min {
			message = "Start charging your laptop!"
		}

		if message == "" {
			continue
		}

		util.Display(state, strconv.Itoa(percentage), message, false)

	}

	return nil
}
