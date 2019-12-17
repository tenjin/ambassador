// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/udp_listener_config.proto

package envoy_api_v3alpha_listener

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on UdpListenerConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UdpListenerConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UdpListenerName

	switch m.ConfigType.(type) {

	case *UdpListenerConfig_Config:

		{
			tmp := m.GetConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return UdpListenerConfigValidationError{
						field:  "Config",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *UdpListenerConfig_TypedConfig:

		{
			tmp := m.GetTypedConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return UdpListenerConfigValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// UdpListenerConfigValidationError is the validation error returned by
// UdpListenerConfig.Validate if the designated constraints aren't met.
type UdpListenerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpListenerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpListenerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpListenerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpListenerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpListenerConfigValidationError) ErrorName() string {
	return "UdpListenerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e UdpListenerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpListenerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpListenerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpListenerConfigValidationError{}
