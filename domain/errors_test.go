// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package domain

import (
	"database/sql"

	dqlite "github.com/canonical/go-dqlite/driver"
	jc "github.com/juju/testing/checkers"
	"github.com/mattn/go-sqlite3"
	gc "gopkg.in/check.v1"

	interrors "github.com/juju/juju/internal/errors"
)

type asError struct {
	Message string
}

func (a asError) Error() string {
	return a.Message
}

type errorsSuite struct{}

var _ = gc.Suite(&errorsSuite{})

// TestCoerceForNilError checks that if you pass a nil error to CoerceError you
// get back a nil error.
func (e *errorsSuite) TestCoerceForNilError(c *gc.C) {
	err := CoerceError(nil)
	c.Check(err, jc.ErrorIsNil)
}

// TestMaskErrorIsHidesSqlErrors is testing that if we construct a maskError
// with with an error chain that contains either sqlite or dqlite errors calls
// to [errors.Is] will return false and mask the errors presence.
func (e *errorsSuite) TestMaskErrorIsHidesSqlErrors(c *gc.C) {
	tests := []struct {
		Name  string
		Error error
	}{
		{
			Name: "Test sqlite3 errors are hidden from Is()",
			Error: sqlite3.Error{
				Code:         sqlite3.ErrAbort,
				ExtendedCode: sqlite3.ErrBusyRecovery,
			},
		},
		{
			Name: "Test dqlite errors are hidden from Is()",
			Error: dqlite.Error{
				Code:    dqlite.ErrBusy,
				Message: "something went wrong",
			},
		},
		{
			Name:  "Test sql.ErrNoRows errors are hidden from Is()",
			Error: sql.ErrNoRows,
		},
		{
			Name:  "Test sql.ErrTxDone errors are hidden from Is()",
			Error: sql.ErrTxDone,
		},
		{
			Name:  "Test sql.ErrConnDone errors are hidden from Is()",
			Error: sql.ErrConnDone,
		},
	}

	for _, test := range tests {
		err := maskError{interrors.Errorf("%q %w", test.Name, test.Error)}
		c.Check(interrors.Is(err, test.Error), jc.IsFalse, gc.Commentf(test.Name))
	}
}

func (e *errorsSuite) TestErrorMessagePreserved(c *gc.C) {
	tests := []struct {
		Error    error
		Expected string
	}{
		{
			Error:    interrors.Errorf("wrap orig error: %w", sql.ErrNoRows),
			Expected: "wrap orig error: sql: no rows in result set",
		},
		{
			Error:    interrors.Errorf("wrap orig error: %w%w", sql.ErrNoRows, dqlite.Error{Code: dqlite.ErrBusy}),
			Expected: "wrap orig error: sql: no rows in result set",
		},
		{
			Error:    interrors.Errorf("wrap orig error: %w - %w", sql.ErrNoRows, interrors.Errorf("nested error")),
			Expected: "wrap orig error: sql: no rows in result set - nested error",
		},
	}
	for _, test := range tests {
		err := CoerceError(test.Error)
		c.Check(err.Error(), gc.Equals, test.Expected)
	}
}

// TestMaskErrorIsNoHide is here to check that if maskError contains non sql
// errors within its chain that it doesn't attempt to hide their existence.
func (e *errorsSuite) TestMaskErrorIsNoHide(c *gc.C) {
	origError := interrors.New("test error")
	err := interrors.Errorf("wrap orig error: %w", origError)
	maskErr := maskError{err}
	c.Check(interrors.Is(maskErr, origError), jc.IsTrue)

	sqlErr := sqlite3.Error{
		Code:         sqlite3.ErrAbort,
		ExtendedCode: sqlite3.ErrBusyRecovery,
	}

	err = interrors.Errorf("double wrap %w %w", sqlErr, origError)
	maskErr = maskError{err}
	c.Check(interrors.Is(maskErr, origError), jc.IsTrue)
}

// TestMaskErrorAsNoHide is here to check that if maskError contains non sql
// errors within its chain that it doesn't attempt to hide their existence.
func (e *errorsSuite) TestMaskErrorAsNoHide(c *gc.C) {
	origError := asError{"ipv6 rocks"}
	err := interrors.Errorf("wrap orig error: %w", origError)
	maskErr := maskError{err}

	var rval asError
	c.Check(interrors.As(maskErr, &rval), jc.IsTrue)

	sqlErr := sqlite3.Error{
		Code:         sqlite3.ErrAbort,
		ExtendedCode: sqlite3.ErrBusyRecovery,
	}

	err = interrors.Errorf("double wrap %w %w", sqlErr, origError)
	maskErr = maskError{err}
	c.Check(interrors.As(maskErr, &rval), jc.IsTrue)
}

// TestMaskErrorAsHidesSqlLiteErrors is here to assert that if we try and
// extract a sqlite error from a [maskError] that we get back false even though
// it does exist.
func (e *errorsSuite) TestMaskErrorAsHidesSqlLiteErrors(c *gc.C) {
	var rval sqlite3.Error
	err := maskError{sqlite3.Error{
		Code:         sqlite3.ErrAbort,
		ExtendedCode: sqlite3.ErrBusyRecovery,
	}}

	c.Check(interrors.As(err, &rval), jc.IsFalse)
}

// TestMaskErrorAsHidesSqlLiteErrors is here to assert that if we try and
// extract a dqlite error from a [maskError] that we get back false even though
// it does exist.
func (e *errorsSuite) TestMaskErrorAsHidesDQLiteErrors(c *gc.C) {
	var rval dqlite.Error
	err := maskError{dqlite.Error{
		Code:    dqlite.ErrBusy,
		Message: "something went wrong",
	}}

	c.Check(interrors.As(err, &rval), jc.IsFalse)
}
