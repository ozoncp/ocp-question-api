// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-question-api/ocp-question-api.proto

package ocp_question_api

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

// Validate checks the field values on Question with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Question) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Text

	return nil
}

// QuestionValidationError is the validation error returned by
// Question.Validate if the designated constraints aren't met.
type QuestionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionValidationError) ErrorName() string { return "QuestionValidationError" }

// Error satisfies the builtin error interface
func (e QuestionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionValidationError{}

// Validate checks the field values on MultiCreateQuestionsV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateQuestionsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetQuestions()) < 1 {
		return MultiCreateQuestionsV1RequestValidationError{
			field:  "Questions",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetQuestions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateQuestionsV1RequestValidationError{
					field:  fmt.Sprintf("Questions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateQuestionsV1RequestValidationError is the validation error
// returned by MultiCreateQuestionsV1Request.Validate if the designated
// constraints aren't met.
type MultiCreateQuestionsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateQuestionsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateQuestionsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateQuestionsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateQuestionsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateQuestionsV1RequestValidationError) ErrorName() string {
	return "MultiCreateQuestionsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateQuestionsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateQuestionsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateQuestionsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateQuestionsV1RequestValidationError{}

// Validate checks the field values on MultiCreateQuestionsV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateQuestionsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiCreateQuestionsV1ResponseValidationError is the validation error
// returned by MultiCreateQuestionsV1Response.Validate if the designated
// constraints aren't met.
type MultiCreateQuestionsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateQuestionsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateQuestionsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateQuestionsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateQuestionsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateQuestionsV1ResponseValidationError) ErrorName() string {
	return "MultiCreateQuestionsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateQuestionsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateQuestionsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateQuestionsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateQuestionsV1ResponseValidationError{}

// Validate checks the field values on CreateQuestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateQuestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateQuestionV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetText()) < 3 {
		return CreateQuestionV1RequestValidationError{
			field:  "Text",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// CreateQuestionV1RequestValidationError is the validation error returned by
// CreateQuestionV1Request.Validate if the designated constraints aren't met.
type CreateQuestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateQuestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateQuestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateQuestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateQuestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateQuestionV1RequestValidationError) ErrorName() string {
	return "CreateQuestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateQuestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateQuestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateQuestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateQuestionV1RequestValidationError{}

// Validate checks the field values on CreateQuestionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateQuestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for QuestionId

	return nil
}

// CreateQuestionV1ResponseValidationError is the validation error returned by
// CreateQuestionV1Response.Validate if the designated constraints aren't met.
type CreateQuestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateQuestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateQuestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateQuestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateQuestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateQuestionV1ResponseValidationError) ErrorName() string {
	return "CreateQuestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateQuestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateQuestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateQuestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateQuestionV1ResponseValidationError{}

// Validate checks the field values on DescribeQuestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeQuestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuestionId() <= 0 {
		return DescribeQuestionV1RequestValidationError{
			field:  "QuestionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeQuestionV1RequestValidationError is the validation error returned by
// DescribeQuestionV1Request.Validate if the designated constraints aren't met.
type DescribeQuestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeQuestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeQuestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeQuestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeQuestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeQuestionV1RequestValidationError) ErrorName() string {
	return "DescribeQuestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeQuestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeQuestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeQuestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeQuestionV1RequestValidationError{}

// Validate checks the field values on DescribeQuestionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeQuestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetQuestion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeQuestionV1ResponseValidationError{
				field:  "Question",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeQuestionV1ResponseValidationError is the validation error returned
// by DescribeQuestionV1Response.Validate if the designated constraints aren't met.
type DescribeQuestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeQuestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeQuestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeQuestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeQuestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeQuestionV1ResponseValidationError) ErrorName() string {
	return "DescribeQuestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeQuestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeQuestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeQuestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeQuestionV1ResponseValidationError{}

// Validate checks the field values on ListQuestionsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListQuestionsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() < 0 {
		return ListQuestionsV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetOffset() < 0 {
		return ListQuestionsV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListQuestionsV1RequestValidationError is the validation error returned by
// ListQuestionsV1Request.Validate if the designated constraints aren't met.
type ListQuestionsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQuestionsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQuestionsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQuestionsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQuestionsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQuestionsV1RequestValidationError) ErrorName() string {
	return "ListQuestionsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListQuestionsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQuestionsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQuestionsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQuestionsV1RequestValidationError{}

// Validate checks the field values on ListQuestionsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListQuestionsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetQuestions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListQuestionsV1ResponseValidationError{
					field:  fmt.Sprintf("Questions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListQuestionsV1ResponseValidationError is the validation error returned by
// ListQuestionsV1Response.Validate if the designated constraints aren't met.
type ListQuestionsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListQuestionsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListQuestionsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListQuestionsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListQuestionsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListQuestionsV1ResponseValidationError) ErrorName() string {
	return "ListQuestionsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListQuestionsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListQuestionsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListQuestionsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListQuestionsV1ResponseValidationError{}

// Validate checks the field values on UpdateQuestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateQuestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuestionId() <= 0 {
		return UpdateQuestionV1RequestValidationError{
			field:  "QuestionId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return UpdateQuestionV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetText()) < 3 {
		return UpdateQuestionV1RequestValidationError{
			field:  "Text",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// UpdateQuestionV1RequestValidationError is the validation error returned by
// UpdateQuestionV1Request.Validate if the designated constraints aren't met.
type UpdateQuestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateQuestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateQuestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateQuestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateQuestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateQuestionV1RequestValidationError) ErrorName() string {
	return "UpdateQuestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateQuestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateQuestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateQuestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateQuestionV1RequestValidationError{}

// Validate checks the field values on UpdateQuestionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateQuestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// UpdateQuestionV1ResponseValidationError is the validation error returned by
// UpdateQuestionV1Response.Validate if the designated constraints aren't met.
type UpdateQuestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateQuestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateQuestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateQuestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateQuestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateQuestionV1ResponseValidationError) ErrorName() string {
	return "UpdateQuestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateQuestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateQuestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateQuestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateQuestionV1ResponseValidationError{}

// Validate checks the field values on RemoveQuestionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveQuestionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetQuestionId() <= 0 {
		return RemoveQuestionV1RequestValidationError{
			field:  "QuestionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveQuestionV1RequestValidationError is the validation error returned by
// RemoveQuestionV1Request.Validate if the designated constraints aren't met.
type RemoveQuestionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveQuestionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveQuestionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveQuestionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveQuestionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveQuestionV1RequestValidationError) ErrorName() string {
	return "RemoveQuestionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveQuestionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveQuestionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveQuestionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveQuestionV1RequestValidationError{}

// Validate checks the field values on RemoveQuestionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveQuestionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// RemoveQuestionV1ResponseValidationError is the validation error returned by
// RemoveQuestionV1Response.Validate if the designated constraints aren't met.
type RemoveQuestionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveQuestionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveQuestionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveQuestionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveQuestionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveQuestionV1ResponseValidationError) ErrorName() string {
	return "RemoveQuestionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveQuestionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveQuestionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveQuestionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveQuestionV1ResponseValidationError{}
