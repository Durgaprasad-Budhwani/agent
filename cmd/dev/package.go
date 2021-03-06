package dev

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/pinpt/agent/v4/sdk"
	"github.com/pinpt/go-common/v10/fileutil"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/log"
	pos "github.com/pinpt/go-common/v10/os"
	"github.com/spf13/cobra"
)

// PackageCmd represents the package command
var PackageCmd = &cobra.Command{
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
		appDir := filepath.Join(integrationDir, "app")
		channel, _ := cmd.Flags().GetString("channel")
		os.MkdirAll(bundleDir, 0700)
		os.MkdirAll(dataDir, 0700)

		buf, err := ioutil.ReadFile(filepath.Join(integrationDir, "integration.yaml"))
		if err != nil {
			log.Fatal(logger, "error loading integration.yaml", "err", err)
		}

		descriptor, err := sdk.LoadDescriptor(base64.StdEncoding.EncodeToString(buf), "", "")
		if err != nil {
			log.Fatal(logger, "error loading descriptor", "err", err)
		}
		if err := ioutil.WriteFile(filepath.Join(bundleDir, "integration.json"), []byte(pjson.Stringify(descriptor)), 0644); err != nil {
			log.Fatal(logger, "error writing integration json", "err", err)
		}

		dataFn := filepath.Join(bundleDir, "data.zip")
		uiFn := filepath.Join(bundleDir, "ui.zip")
		bundleFn := filepath.Join(distDir, "bundle.zip")

		oss, _ := cmd.Flags().GetStringArray("os")
		arches, _ := cmd.Flags().GetStringArray("arch")

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
			log.Fatal(logger, "error running command", "command", c.String(), "err", err)
		}

		// build the application
		c = exec.Command("npm", "install", "--loglevel", "error")
		c.Dir = appDir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		if err := c.Run(); err != nil {
			log.Fatal(logger, "error running command", "command", c.String(), "err", err)
		}
		c = exec.Command("npm", "run", "build", "--loglevel", "error")
		c.Dir = appDir
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin
		if err := c.Run(); err != nil {
			log.Fatal(logger, "error running command", "command", c.String(), "err", err)
		}
		if _, err := fileutil.ZipDir(uiFn, filepath.Join(appDir, "build"), regexp.MustCompile(".*")); err != nil {
			log.Fatal(logger, "error building zip file", "err", err)
		}

		sha := getBuildCommitForIntegration(integrationDir)

		// write out our version file
		if err := ioutil.WriteFile(filepath.Join(bundleDir, "version.txt"), []byte(sha), 0644); err != nil {
			log.Fatal(logger, "error writing version file to bundle dir", "err", err)
		}

		// write out dev certificate (if we have one)
		if devCfg, err := loadDevConfig(channel); err == nil {
			if devCfg.Certificate != "" {
				if err := ioutil.WriteFile(filepath.Join(bundleDir, "cert.pem"), []byte(devCfg.Certificate), 0644); err != nil {
					log.Fatal(logger, "error writing developer certificate to bundle dir", "err", err)
				}
			} else {
				log.Debug(logger, "no developer certificate found, not including in bundle", "err", err)
			}
		} else {
			log.Warn(logger, "unable to load developer config, the bundle will not contain your developer certificate", "err", err)
		}

		if _, err := fileutil.ZipDir(dataFn, dataDir, regexp.MustCompile(".*")); err != nil {
			log.Fatal(logger, "error building zip file", "err", err)
		}
		if _, err := fileutil.ZipDir(bundleFn, bundleDir, regexp.MustCompile(".(zip|asc|txt|pem|json)$")); err != nil {
			log.Fatal(logger, "error building zip file", "err", err)
		}
		os.RemoveAll(bundleDir)
		log.Info(logger, "bundle packaged to "+bundleFn)
	},
}

func init() {
	PackageCmd.Flags().String("dir", "dist", "the output directory to place the generated file")
	PackageCmd.Flags().StringArray("os", []string{"darwin", "linux", "windows"}, "the OS to build for")
	PackageCmd.Flags().StringArray("arch", []string{"amd64"}, "the architecture to build for")
	PackageCmd.Flags().String("channel", pos.Getenv("PP_CHANNEL", "stable"), "the channel which can be set")
}
