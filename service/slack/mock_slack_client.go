// Code generated by mockery v2.14.0. DO NOT EDIT.

package slack

import (
	context "context"

	slack_goslack "github.com/slack-go/slack"
	mock "github.com/stretchr/testify/mock"
)

// mockSlackClient is an autogenerated mock type for the slackClient type
type mockSlackClient struct {
	mock.Mock
}

// PostMessageContext provides a mock function with given fields: ctx, channelID, options
func (_m *mockSlackClient) PostMessageContext(ctx context.Context, channelID string, options ...slack_goslack.MsgOption) (string, string, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, channelID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, ...slack_goslack.MsgOption) string); ok {
		r0 = rf(ctx, channelID, options...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, ...slack_goslack.MsgOption) string); ok {
		r1 = rf(ctx, channelID, options...)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, ...slack_goslack.MsgOption) error); ok {
		r2 = rf(ctx, channelID, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTnewMockSlackClient interface {
	mock.TestingT
	Cleanup(func())
}

// newMockSlackClient creates a new instance of mockSlackClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockSlackClient(t mockConstructorTestingTnewMockSlackClient) *mockSlackClient {
	mock := &mockSlackClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
