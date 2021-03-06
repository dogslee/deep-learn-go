// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/coin/v1/coin.proto

package v1

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on AddCoinRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddCoinRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Num

	return nil
}

// AddCoinRequestValidationError is the validation error returned by
// AddCoinRequest.Validate if the designated constraints aren't met.
type AddCoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCoinRequestValidationError) ErrorName() string { return "AddCoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddCoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCoinRequestValidationError{}

// Validate checks the field values on AddCoinReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AddCoinReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	return nil
}

// AddCoinReplyValidationError is the validation error returned by
// AddCoinReply.Validate if the designated constraints aren't met.
type AddCoinReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCoinReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCoinReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCoinReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCoinReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCoinReplyValidationError) ErrorName() string { return "AddCoinReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddCoinReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCoinReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCoinReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCoinReplyValidationError{}

// Validate checks the field values on ReduceCoinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ReduceCoinRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Num

	return nil
}

// ReduceCoinRequestValidationError is the validation error returned by
// ReduceCoinRequest.Validate if the designated constraints aren't met.
type ReduceCoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReduceCoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReduceCoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReduceCoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReduceCoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReduceCoinRequestValidationError) ErrorName() string {
	return "ReduceCoinRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReduceCoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReduceCoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReduceCoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReduceCoinRequestValidationError{}

// Validate checks the field values on ReduceCoinReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ReduceCoinReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	return nil
}

// ReduceCoinReplyValidationError is the validation error returned by
// ReduceCoinReply.Validate if the designated constraints aren't met.
type ReduceCoinReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReduceCoinReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReduceCoinReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReduceCoinReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReduceCoinReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReduceCoinReplyValidationError) ErrorName() string { return "ReduceCoinReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReduceCoinReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReduceCoinReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReduceCoinReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReduceCoinReplyValidationError{}

// Validate checks the field values on ShowCoinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ShowCoinRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ShowCoinRequestValidationError is the validation error returned by
// ShowCoinRequest.Validate if the designated constraints aren't met.
type ShowCoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShowCoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShowCoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShowCoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShowCoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShowCoinRequestValidationError) ErrorName() string { return "ShowCoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e ShowCoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShowCoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShowCoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShowCoinRequestValidationError{}

// Validate checks the field values on ShowCoinReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ShowCoinReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Amount

	// no validation rules for Message

	return nil
}

// ShowCoinReplyValidationError is the validation error returned by
// ShowCoinReply.Validate if the designated constraints aren't met.
type ShowCoinReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShowCoinReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShowCoinReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShowCoinReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShowCoinReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShowCoinReplyValidationError) ErrorName() string { return "ShowCoinReplyValidationError" }

// Error satisfies the builtin error interface
func (e ShowCoinReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShowCoinReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShowCoinReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShowCoinReplyValidationError{}
