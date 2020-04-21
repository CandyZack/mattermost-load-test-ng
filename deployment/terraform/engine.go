// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package terraform

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/mattermost/mattermost-server/v5/mlog"
)

// Cleanup is called at the end of each command to clean temporary files
func (t *Terraform) Cleanup() {
	if t.dir != "" {
		os.RemoveAll(t.dir)
	}
}

// runCommand runs terraform with the args supplied. If dst is not nil, it writes the output there.
// Otherwise, it logs the output to console.
func (t *Terraform) runCommand(dst io.Writer, args ...string) error {
	terraformBin := "terraform"
	if _, err := exec.LookPath(terraformBin); err != nil {
		return fmt.Errorf("terraform not installed. Please install terraform. (https://www.terraform.io/downloads.html): %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cmdExecTimeoutMinutes*time.Minute)
	defer cancel()

	mlog.Debug("Running terraform command", mlog.String("args", fmt.Sprintf("%v", args)))
	cmd := exec.CommandContext(ctx, terraformBin, args...)

	// If dst is set, that means we want to capture the output.
	// We write a simple case to handle that using CombinedOutput.
	if dst != nil {
		out, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		_, err = dst.Write(out)
		return err
	}

	// From here, we want to stream the output concurrently from stderr and stdout
	// to mlog.
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			mlog.Info(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		mlog.Info(scanner.Text())
	}
	// No need to check for scanner.Error as cmd.Wait() already does that.
	wg.Wait()

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func checkTerraformVersion() error {
	if out, err := exec.Command("terraform", "version").Output(); err != nil {
		return fmt.Errorf("could not run %q command: %w", "terraform version", err)
	} else if re := regexp.MustCompile(`v\d+.?\d+`); !re.Match(out) {
		return fmt.Errorf("could not parse terraform command output: %s", out)
	} else if version, err := strconv.ParseFloat(string(re.FindAll(out, 1)[0][1:]), 64); err != nil {
		return fmt.Errorf("could not parse terraform command output: %w", err)
	} else if version < 0.12 {
		return fmt.Errorf("at least terraform v0.12 is required, you have %.2f", version)
	} else if version > 0.12 {
		mlog.Warn(fmt.Sprintf("This tool tested with terraform v0.12, you have v%.2f. Proceed with your own risk.", version))
	}

	return nil
}
