// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package aspect

import (
	"context"

	"github.com/DataDog/orchestrion/internal/toolexec/proxy"
)

func (w Weaver) OnLink(ctx context.Context, cmd *proxy.LinkCommand) (err error) {
	_ = "STUB: not implemented"
	return nil
}
