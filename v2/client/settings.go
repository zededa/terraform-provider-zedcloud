package client

import "time"

// Settings carries provider-level options that resources need at run time.
// It is hand-written (not generated) and set by resources.ProviderConfigure.
// See docs/design/ue-177-high-parallelism-robustness.md.
type Settings struct {
	// AsyncWrites makes zedcloud_application_instance create, update and
	// delete use the job-based /v2/apps/instances routes and poll the job,
	// instead of waiting on the synchronous /v1 call (UE-177 §4.6).
	AsyncWrites bool

	// JobPollIntervalMin and JobPollIntervalMax bound the interval between
	// job status polls. The interval starts at the minimum and doubles up to
	// the maximum.
	JobPollIntervalMin time.Duration
	JobPollIntervalMax time.Duration
}

// DefaultSettings returns the settings used when the provider block does not
// override them.
func DefaultSettings() *Settings {
	return &Settings{
		AsyncWrites:        true,
		JobPollIntervalMin: 2 * time.Second,
		JobPollIntervalMax: 10 * time.Second,
	}
}
