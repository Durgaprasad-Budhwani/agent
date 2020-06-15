package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/pinpt/go-common/v10/fileutil"
	"github.com/pinpt/go-common/v10/log"
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package an integration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		integrationDir := args[0]
		logger := log.NewCommandLogger(cmd)
		defer logger.Close()
		integrationDir, _ = filepath.Abs(integrationDir)
		distDir, _ := cmd.Flags().GetString("dir")
		distDir, _ = filepath.Abs(distDir)
		bundleDir := filepath.Join(distDir, "bundle")
		dataDir := filepath.Join(bundleDir, "data")
		os.MkdirAll(bundleDir, 0700)
		os.MkdirAll(dataDir, 0700)

		buf, err := ioutil.ReadFile(filepath.Join(integrationDir, "integration.yaml"))
		if err != nil {
			log.Fatal(logger, "error loading integration.yaml", "err", err)
		}

		ioutil.WriteFile(filepath.Join(dataDir, "integration.yaml"), buf, 0644)

		dataFn := filepath.Join(bundleDir, "data.zip")
		bundleFn := filepath.Join(distDir, "bundle.zip")

		oss, _ := cmd.Flags().GetStringSlice("os")
		arches, _ := cmd.Flags().GetStringSlice("arch")

		cargs := []string{"build", integrationDir, "--dir", dataDir}
		for _, o := range oss {
			cargs = append(cargs, "--os", o)
		}
		for _, a := range arches {
			cargs = append(cargs, "--arch", a)
		}
		c := exec.Command(os.Args[0], cargs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		if err := c.Run(); err != nil {
			os.Exit(1)
		}

		sha := getBuildCommitForIntegration(integrationDir)

		// write out our version file
		ioutil.WriteFile(filepath.Join(bundleDir, "version.txt"), []byte(sha), 0644)

		// write out the sha sum512 for each file in the zip for integrity checking
		shafilename := filepath.Join(bundleDir, "sha512sum.txt.asc")
		if err := fileutil.ShaFiles(dataDir, shafilename, regexp.MustCompile(".*")); err != nil {
			log.Fatal(logger, "error generating sha sums", "err", err)
		}

		if _, err := fileutil.ZipDir(dataFn, dataDir, regexp.MustCompile(".*")); err != nil {
			log.Fatal(logger, "error building zip file", "err", err)
		}
		if _, err := fileutil.ZipDir(bundleFn, bundleDir, regexp.MustCompile(".(zip|asc|txt)$")); err != nil {
			log.Fatal(logger, "error building zip file", "err", err)
		}
		os.RemoveAll(bundleDir)
		log.Info(logger, "bundle packaged to "+bundleFn)
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.Flags().String("dir", "dist", "the output directory to place the generated file")
	packageCmd.Flags().StringSlice("os", []string{"darwin", "linux"}, "the OS to build for")
	packageCmd.Flags().StringSlice("arch", []string{"amd64"}, "the architecture to build for")
}
