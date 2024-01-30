// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package config

const (
	defaultRefreshRate  = 2
	defaultMaxConnRetry = 5
)

// UI tracks ui specific configs.
type UI struct {
	// EnableMouse toggles mouse support.
	EnableMouse bool `json:"enableMouse" yaml:"enableMouse"`

	// Headless toggles top header display.
	Headless bool `json:"headless" yaml:"headless"`

	// LogoLess toggles k9s logo.
	Logoless bool `json:"logoless" yaml:"logoless"`

	// Crumbsless toggles nav crumb display.
	Crumbsless bool `json:"crumbsless" yaml:"crumbsless"`

	// Reactive toggles reactive ui changes.
	Reactive bool `json:"reactive" yaml:"reactive"`

	// NoIcons toggles icons display.
	NoIcons bool `json:"noIcons" yaml:"noIcons"`

	// Skin reference the general k9s skin name.
	// Can be overridden per context.
	Skin string `json:"skin" yaml:"skin,omitempty"`

	// FullScreenLView sets default full screen for live_view.
	FullScreenLView bool `json:"fullScreenLView" yaml:"fullScreenLView"`
}
