// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pin

import (
	"context"
	"io"
)

const envVarCheckedGoMod = "DD_ORCHESTRION_IS_GOMOD_VERSION"

// AutoPinOrchestrion automatically runs [PinOrchestrion] if the necessary
// requirements are not already met. It prints messages to `stderr` to inform
// the user about what is going on.
func AutoPinOrchestrion(ctx context.Context, stdout io.Writer, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// A parent process (or ourselves earlier) has already done the check

// Make sure we don't do this again

// There is already a required version, but we're not running that one!

// We're good to go
