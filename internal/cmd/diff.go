// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package cmd

import (
	"fmt"

	"github.com/DataDog/orchestrion/internal/report"
	"github.com/urfave/cli/v2"
)

var (
	filenameFlag = cli.BoolFlag{
		Name:  "files",
		Usage: "Only show file paths created by orchestrion instead of diff output",
	}

	filterFlag = cli.StringFlag{
		Name:  "filter",
		Usage: "Filter the diff to a regex matched on the package or file paths from the build.",
	}

	packageFlag = cli.BoolFlag{
		Name:  "package",
		Usage: "Print package names instead of printing the diff",
	}

	debugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "Also print synthetic and tracer weaved packages",
	}

	buildFlag = cli.BoolFlag{
		Name:  "build",
		Usage: "Execute a build with -work before generating the diff. All remaining arguments after flags are passed to the build command.",
	}

	noCacheFlag = cli.BoolFlag{
		Name:  "no-cache",
		Usage: "Force a rebuild of all packages when using --build (adds -a flag). Useful for ensuring complete instrumentation coverage.",
	}

	Diff = &cli.Command{
		Name:  "diff",
		Usage: "Generates a diff between a nominal and orchestrion-instrumented build. Use --build to execute a build first, or provide a work directory path obtained from `orchestrion go build -work -a`. This is incompatible with coverage related flags.",
		Args:  true,
		Flags: []cli.Flag{
			&filenameFlag,
			&filterFlag,
			&packageFlag,
			&debugFlag,
			&buildFlag,
			&noCacheFlag,
		},
		Action: func(clictx *cli.Context) error {
			workFolder, err := workFolder(clictx)
			if err != nil {
				return err
			}

			report, err := report.FromWorkDir(clictx.Context, workFolder)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed to read work dir: %s (did you forgot the -work flag during build ?)", err), 1)
			}

			if report.IsEmpty() {
				return cli.Exit("no files to diff (did you forgot the -a flag during build?)", 1)
			}

			if !clictx.Bool(debugFlag.Name) {
				report = report.WithSpecialCasesFilter()
			}

			if filter := clictx.String(filterFlag.Name); filter != "" {
				report, err = report.WithRegexFilter(filter)
				if err != nil {
					return cli.Exit(fmt.Sprintf("failed to filter files: %s", err), 1)
				}
			}

			return outputReport(clictx, report)
		},
	}
)

func workFolder(clictx *cli.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func prepareBuildArgs(args []string, forceRebuild bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func outputReport(clictx *cli.Context, rpt report.Report) error {
	_ = "STUB: not implemented"
	return nil
}

func executeBuildAndCaptureWorkDir(clictx *cli.Context, buildArgs []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func extractWorkDirFromOutput(output string) string { _ = "STUB: not implemented"; return "" }
