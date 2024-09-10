package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/charmbracelet/log"
	"github.com/numtide/godoc/pkg/markdown"
	"github.com/numtide/godoc/pkg/parse"
	"github.com/spf13/cobra"
)

var (
	clean  bool
	outDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godoc [source directory]",
	Short: "Custom Go doc generation",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		if outDir == "" {
			return fmt.Errorf("output directory must be specified")
		}

		if clean {
			log.Infof("cleaning output directory: %s", outDir)
			if err := os.RemoveAll(outDir); err != nil {
				return fmt.Errorf("failed to clean output directory: %w", err)
			}
		}

		sourceDir := args[0]
		log.Infof("processing Go files in: %s", sourceDir)

		eg := errgroup.Group{}
		eg.SetLimit(runtime.NumCPU())

		if err := filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(path, ".go") {
				// skip
				return nil
			}

			eg.Go(func() error {
				log.Infof("processing file: %s", path)
				data, err := parse.File(path)
				if err != nil {
					return fmt.Errorf("failed to parse file: %w", err)
				}

				return markdown.Write(outDir, data)
			})
			return nil
		}); err != nil {
			return fmt.Errorf("failed to walk source directory: %w", err)
		}

		// wait for processing to complete
		return eg.Wait()
	},
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&outDir, "out", "o", "", "output directory")
	pf.BoolVarP(&clean, "clean", "c", false, "clean output directory before writing")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
