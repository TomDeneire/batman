package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// BuildTime defined by compilation
var BuildTime = ""

// GoVersion defined by compilation
var GoVersion = ""

// BuildHost defined by compilation
var BuildHost = ""

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(buildTime string, goVersion string, buildHost string, args []string) {
	BuildTime = buildTime
	BuildHost = buildHost
	GoVersion = goVersion
	rootCmd.Execute()
}

var rootCmd = &cobra.Command{
	Use:           "batman",
	Short:         "batman - CLI application to monitor battery charging",
	SilenceUsage:  true,
	SilenceErrors: true,
	Long:          `batman is a CLI application to monitor battery charging`,
}

var Fbatmanfile string

func init() {
	cachedir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Could not find user cachedir: %v", err)
	}

	batmandir := filepath.Join(cachedir, "batman")
	err = os.MkdirAll(batmandir, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("Could not create batman cachedir: %v", err)
	}

	Fbatmanfile = filepath.Join(batmandir, "batman_pid")
}
