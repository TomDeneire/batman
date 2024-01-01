package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"

	util "tomdeneire.be/batman/lib/util"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Short:   "Check batman",
	Long:    `Checking batman.`,
	Args:    cobra.NoArgs,
	Example: "batman check",
	RunE:    check,
}

var Fraw bool

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().BoolVar(&Fraw, "raw", false, "Raw output")
}

func check(cmd *cobra.Command, args []string) error {

	state, percentage, err := util.BatteryInfo()
	if err != nil {
		log.Fatalf("Could not get battery status: %v", err)
	}

	util.Display(state, strconv.Itoa(percentage), "", Fraw)

	return nil
}
