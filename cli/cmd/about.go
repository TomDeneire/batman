package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:     "about",
	Short:   "About batman",
	Long:    `Version and build time information about the batman executable.`,
	Args:    cobra.NoArgs,
	Example: `batman about`,
	RunE:    about,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func about(cmd *cobra.Command, args []string) error {
	msg := map[string]string{"BuildTime": BuildTime, "BuildHost": BuildHost, "BuildWith": GoVersion}
	host, e := os.Hostname()

	if e == nil {
		msg["uname"] = host
	}
	user, err := user.Current()
	if err == nil {
		msg["user.name"] = user.Name
		msg["user.username"] = user.Username
	}
	b, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))

	return nil
}
