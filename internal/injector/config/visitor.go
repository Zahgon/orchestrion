// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2023-present Datadog, Inc.

package config

import (
	"github.com/DataDog/orchestrion/internal/injector/aspect"
)

type (
	Visitor = func(cfg File, pkgPath string) error

	File interface {
		Name() string
		Description() string
		Caveats() string
		Icon() string

		OwnAspects() []*aspect.Aspect
	}
)

// Visit calls the visitor for each configuration found in the specified root
// [Config].
func Visit(cfg Config, visitor Visitor) error { _ = "STUB: not implemented"; return nil }

func (c *configYML) Name() string { _ = "STUB: not implemented"; return "" }

func (c *configYML) Description() string { _ = "STUB: not implemented"; return "" }

func (c *configYML) Caveats() string { _ = "STUB: not implemented"; return "" }

func (c *configYML) Icon() string { _ = "STUB: not implemented"; return "" }

func (c *configYML) OwnAspects() []*aspect.Aspect { _ = "STUB: not implemented"; return nil }
