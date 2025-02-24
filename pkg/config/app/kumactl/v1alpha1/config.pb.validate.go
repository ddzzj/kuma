// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/config/app/kumactl/v1alpha1/config.proto

package v1alpha1

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Configuration with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Configuration) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetControlPlanes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigurationValidationError{
					field:  fmt.Sprintf("ControlPlanes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetContexts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigurationValidationError{
					field:  fmt.Sprintf("Contexts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CurrentContext

	return nil
}

// ConfigurationValidationError is the validation error returned by
// Configuration.Validate if the designated constraints aren't met.
type ConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationValidationError) ErrorName() string { return "ConfigurationValidationError" }

// Error satisfies the builtin error interface
func (e ConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationValidationError{}

// Validate checks the field values on ControlPlane with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ControlPlane) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return ControlPlaneValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetCoordinates() == nil {
		return ControlPlaneValidationError{
			field:  "Coordinates",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetCoordinates()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ControlPlaneValidationError{
				field:  "Coordinates",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ControlPlaneValidationError is the validation error returned by
// ControlPlane.Validate if the designated constraints aren't met.
type ControlPlaneValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ControlPlaneValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ControlPlaneValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ControlPlaneValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ControlPlaneValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ControlPlaneValidationError) ErrorName() string { return "ControlPlaneValidationError" }

// Error satisfies the builtin error interface
func (e ControlPlaneValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sControlPlane.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ControlPlaneValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ControlPlaneValidationError{}

// Validate checks the field values on ControlPlaneCoordinates with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ControlPlaneCoordinates) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetApiServer() == nil {
		return ControlPlaneCoordinatesValidationError{
			field:  "ApiServer",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetApiServer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ControlPlaneCoordinatesValidationError{
				field:  "ApiServer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ControlPlaneCoordinatesValidationError is the validation error returned by
// ControlPlaneCoordinates.Validate if the designated constraints aren't met.
type ControlPlaneCoordinatesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ControlPlaneCoordinatesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ControlPlaneCoordinatesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ControlPlaneCoordinatesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ControlPlaneCoordinatesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ControlPlaneCoordinatesValidationError) ErrorName() string {
	return "ControlPlaneCoordinatesValidationError"
}

// Error satisfies the builtin error interface
func (e ControlPlaneCoordinatesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sControlPlaneCoordinates.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ControlPlaneCoordinatesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ControlPlaneCoordinatesValidationError{}

// Validate checks the field values on Context with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Context) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return ContextValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetControlPlane()) < 1 {
		return ContextValidationError{
			field:  "ControlPlane",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetDefaults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContextValidationError{
				field:  "Defaults",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ContextValidationError{
				field:  "Credentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ContextValidationError is the validation error returned by Context.Validate
// if the designated constraints aren't met.
type ContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContextValidationError) ErrorName() string { return "ContextValidationError" }

// Error satisfies the builtin error interface
func (e ContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContextValidationError{}

// Validate checks the field values on ControlPlaneCoordinates_ApiServer with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ControlPlaneCoordinates_ApiServer) Validate() error {
	if m == nil {
		return nil
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		return ControlPlaneCoordinates_ApiServerValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return ControlPlaneCoordinates_ApiServerValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
	}

	return nil
}

// ControlPlaneCoordinates_ApiServerValidationError is the validation error
// returned by ControlPlaneCoordinates_ApiServer.Validate if the designated
// constraints aren't met.
type ControlPlaneCoordinates_ApiServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ControlPlaneCoordinates_ApiServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ControlPlaneCoordinates_ApiServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ControlPlaneCoordinates_ApiServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ControlPlaneCoordinates_ApiServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ControlPlaneCoordinates_ApiServerValidationError) ErrorName() string {
	return "ControlPlaneCoordinates_ApiServerValidationError"
}

// Error satisfies the builtin error interface
func (e ControlPlaneCoordinates_ApiServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sControlPlaneCoordinates_ApiServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ControlPlaneCoordinates_ApiServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ControlPlaneCoordinates_ApiServerValidationError{}

// Validate checks the field values on Context_Defaults with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Context_Defaults) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mesh

	return nil
}

// Context_DefaultsValidationError is the validation error returned by
// Context_Defaults.Validate if the designated constraints aren't met.
type Context_DefaultsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Context_DefaultsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Context_DefaultsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Context_DefaultsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Context_DefaultsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Context_DefaultsValidationError) ErrorName() string { return "Context_DefaultsValidationError" }

// Error satisfies the builtin error interface
func (e Context_DefaultsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContext_Defaults.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Context_DefaultsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Context_DefaultsValidationError{}

// Validate checks the field values on Context_Credentials with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Context_Credentials) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetDataplaneTokenApi()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Context_CredentialsValidationError{
				field:  "DataplaneTokenApi",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Context_CredentialsValidationError is the validation error returned by
// Context_Credentials.Validate if the designated constraints aren't met.
type Context_CredentialsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Context_CredentialsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Context_CredentialsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Context_CredentialsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Context_CredentialsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Context_CredentialsValidationError) ErrorName() string {
	return "Context_CredentialsValidationError"
}

// Error satisfies the builtin error interface
func (e Context_CredentialsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContext_Credentials.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Context_CredentialsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Context_CredentialsValidationError{}

// Validate checks the field values on Context_DataplaneTokenApiCredentials
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *Context_DataplaneTokenApiCredentials) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClientCert

	// no validation rules for ClientKey

	return nil
}

// Context_DataplaneTokenApiCredentialsValidationError is the validation error
// returned by Context_DataplaneTokenApiCredentials.Validate if the designated
// constraints aren't met.
type Context_DataplaneTokenApiCredentialsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Context_DataplaneTokenApiCredentialsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Context_DataplaneTokenApiCredentialsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Context_DataplaneTokenApiCredentialsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Context_DataplaneTokenApiCredentialsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Context_DataplaneTokenApiCredentialsValidationError) ErrorName() string {
	return "Context_DataplaneTokenApiCredentialsValidationError"
}

// Error satisfies the builtin error interface
func (e Context_DataplaneTokenApiCredentialsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContext_DataplaneTokenApiCredentials.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Context_DataplaneTokenApiCredentialsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Context_DataplaneTokenApiCredentialsValidationError{}
