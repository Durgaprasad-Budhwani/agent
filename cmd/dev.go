package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/pinpt/go-common/v10/fileutil"
	"github.com/pinpt/go-common/v10/log"
	pos "github.com/pinpt/go-common/v10/os"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "run an integration in development mode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		integrationDir := args[0]
		_logger := log.NewCommandLogger(cmd)
		defer _logger.Close()
		integrationDir, _ = filepath.Abs(integrationDir)
		integration := strings.Replace(filepath.Base(integrationDir), "agent.next.", "", -1)
		fp := filepath.Join(integrationDir, "integration.go")
		if !fileutil.FileExists(fp) {
			log.Fatal(_logger, "couldn't find the integration at "+fp)
		}
		distDir := filepath.Join(os.TempDir(), "agent.next")
		os.MkdirAll(distDir, 0700)
		integrationFile := filepath.Join(distDir, integration)

		// build our integration
		c := exec.Command(os.Args[0], "build", "--dir", distDir, integrationDir)
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
		if err := c.Run(); err != nil {
			os.Exit(1)
		}

		channel, _ := cmd.Flags().GetString("channel")
		dir, _ := cmd.Flags().GetString("dir")

		devargs := []string{"--dev", "--dir", dir, "--channel", channel, "--log-level", "debug"}

		config, _ := cmd.Flags().GetStringSlice("config")
		for _, str := range config {
			devargs = append(devargs, "--set", str)
		}

		c = exec.Command(integrationFile, devargs...)

		pos.OnExit(func(_ int) {
			if c != nil {
				syscall.Kill(-c.Process.Pid, syscall.SIGINT)
				c = nil
			}
		})

		c.Stderr = os.Stderr
		c.Stdout = os.Stdout
		c.Dir = distDir
		c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		if err := c.Start(); err != nil {
			os.Exit(1)
		}
		c.Wait()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
	devCmd.Flags().StringSlice("config", []string{}, "a config key/value pair such as a=b")
	devCmd.Flags().String("dir", "", "the directory to output pipe contents")
	devCmd.Flags().String("channel", "", "the channel which can be set")
	devCmd.Flags().MarkHidden("channel")
}
