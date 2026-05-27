// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package buildid

import (
	"context"
	"sync"

	"github.com/DataDog/orchestrion/internal/injector/config"
	"github.com/DataDog/orchestrion/internal/jobserver/common"
	"github.com/nats-io/nats.go"
)

const (
	subjectPrefix = "buildid."

	versionSubject = subjectPrefix + "versionSuffix"
)

type service struct {
	packageLoader   config.PackageLoader
	stats           *common.CacheStats
	resolvedVersion VersionSuffixResponse
	mu              sync.Mutex
}

func Subscribe(ctx context.Context, conn *nats.Conn, pkgLoader config.PackageLoader, stats *common.CacheStats) error {
	_ = "STUB: not implemented"
	return nil
}
