// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package pkgs

import (
	"context"

	"github.com/DataDog/orchestrion/internal/injector/config"
	"github.com/DataDog/orchestrion/internal/jobserver/common"
	"github.com/nats-io/nats.go"
	"golang.org/x/tools/go/packages"
)

const (
	subjectPrefix = "packages."

	resolveSubject = subjectPrefix + "resolve"
	loadSubject    = subjectPrefix + "load"
)

type service struct {
	resolved  common.Cache[ResolveResponse]
	loaded    common.Cache[*packages.Package]
	graph     common.Graph
	serverURL string
}

func Subscribe(ctx context.Context, serverURL string, conn *nats.Conn, stats *common.CacheStats) (config.PackageLoader, error) {
	_ = "STUB: not implemented"
	return *new(config.PackageLoader), nil
}
