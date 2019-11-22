// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	configuration "github.com/kyma-incubator/compass/components/provisioner/internal/hydroform/configuration"
	gqlschema "github.com/kyma-incubator/compass/components/provisioner/pkg/gqlschema"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/provisioner/internal/model"
)

// BuilderFactory is an autogenerated mock type for the BuilderFactory type
type BuilderFactory struct {
	mock.Mock
}

// NewDeprovisioningBuilder provides a mock function with given fields: runtimeConfig
func (_m *BuilderFactory) NewDeprovisioningBuilder(runtimeConfig model.RuntimeConfig) configuration.Builder {
	ret := _m.Called(runtimeConfig)

	var r0 configuration.Builder
	if rf, ok := ret.Get(0).(func(model.RuntimeConfig) configuration.Builder); ok {
		r0 = rf(runtimeConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(configuration.Builder)
		}
	}

	return r0
}

// NewProvisioningBuilder provides a mock function with given fields: provisionInput
func (_m *BuilderFactory) NewProvisioningBuilder(provisionInput gqlschema.ProvisionRuntimeInput) configuration.Builder {
	ret := _m.Called(provisionInput)

	var r0 configuration.Builder
	if rf, ok := ret.Get(0).(func(gqlschema.ProvisionRuntimeInput) configuration.Builder); ok {
		r0 = rf(provisionInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(configuration.Builder)
		}
	}

	return r0
}