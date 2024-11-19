// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	buildv1 "github.com/openshift/api/build/v1"
	corev1 "k8s.io/api/core/v1"
)

// BuildConfigSpecApplyConfiguration represents a declarative configuration of the BuildConfigSpec type for use
// with apply.
type BuildConfigSpecApplyConfiguration struct {
	Triggers                     []BuildTriggerPolicyApplyConfiguration `json:"triggers,omitempty"`
	RunPolicy                    *buildv1.BuildRunPolicy                `json:"runPolicy,omitempty"`
	CommonSpecApplyConfiguration `json:",inline"`
	SuccessfulBuildsHistoryLimit *int32 `json:"successfulBuildsHistoryLimit,omitempty"`
	FailedBuildsHistoryLimit     *int32 `json:"failedBuildsHistoryLimit,omitempty"`
}

// BuildConfigSpecApplyConfiguration constructs a declarative configuration of the BuildConfigSpec type for use with
// apply.
func BuildConfigSpec() *BuildConfigSpecApplyConfiguration {
	return &BuildConfigSpecApplyConfiguration{}
}

// WithTriggers adds the given value to the Triggers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Triggers field.
func (b *BuildConfigSpecApplyConfiguration) WithTriggers(values ...*BuildTriggerPolicyApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTriggers")
		}
		b.Triggers = append(b.Triggers, *values[i])
	}
	return b
}

// WithRunPolicy sets the RunPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunPolicy field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithRunPolicy(value buildv1.BuildRunPolicy) *BuildConfigSpecApplyConfiguration {
	b.RunPolicy = &value
	return b
}

// WithServiceAccount sets the ServiceAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccount field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithServiceAccount(value string) *BuildConfigSpecApplyConfiguration {
	b.ServiceAccount = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithSource(value *BuildSourceApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	b.Source = value
	return b
}

// WithRevision sets the Revision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Revision field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithRevision(value *SourceRevisionApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	b.Revision = value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithStrategy(value *BuildStrategyApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	b.Strategy = value
	return b
}

// WithOutput sets the Output field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Output field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithOutput(value *BuildOutputApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	b.Output = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *BuildConfigSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithPostCommit sets the PostCommit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PostCommit field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithPostCommit(value *BuildPostCommitSpecApplyConfiguration) *BuildConfigSpecApplyConfiguration {
	b.PostCommit = value
	return b
}

// WithCompletionDeadlineSeconds sets the CompletionDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletionDeadlineSeconds field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithCompletionDeadlineSeconds(value int64) *BuildConfigSpecApplyConfiguration {
	b.CompletionDeadlineSeconds = &value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithNodeSelector(value buildv1.OptionalNodeSelector) *BuildConfigSpecApplyConfiguration {
	b.NodeSelector = &value
	return b
}

// WithMountTrustedCA sets the MountTrustedCA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountTrustedCA field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithMountTrustedCA(value bool) *BuildConfigSpecApplyConfiguration {
	b.MountTrustedCA = &value
	return b
}

// WithSuccessfulBuildsHistoryLimit sets the SuccessfulBuildsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuccessfulBuildsHistoryLimit field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithSuccessfulBuildsHistoryLimit(value int32) *BuildConfigSpecApplyConfiguration {
	b.SuccessfulBuildsHistoryLimit = &value
	return b
}

// WithFailedBuildsHistoryLimit sets the FailedBuildsHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailedBuildsHistoryLimit field is set to the value of the last call.
func (b *BuildConfigSpecApplyConfiguration) WithFailedBuildsHistoryLimit(value int32) *BuildConfigSpecApplyConfiguration {
	b.FailedBuildsHistoryLimit = &value
	return b
}
