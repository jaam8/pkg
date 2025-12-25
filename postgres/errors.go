package postgres

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	ErrDuplicateKeyCode        = "23505"
	ErrForeignKeyViolationCode = "23503"
)

// IsForeignKeyErr checks if the provided error is a foreign key violation error.
// Optionally checks if the error is related to a specific field constraint.
func IsForeignKeyErr(err error, field ...string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == ErrForeignKeyViolationCode {
			if len(field) != 0 {
				return strings.Contains(pgErr.ConstraintName, field[0])
			}
			return true
		}
	}
	return false
}

// IsUniqueViolationErr determines if the given error is a unique constraint violation.
// Optionally checks if the error is related to a specific field constraint.
func IsUniqueViolationErr(err error, field ...string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == ErrDuplicateKeyCode {
			if len(field) != 0 {
				return strings.Contains(pgErr.ConstraintName, field[0])
			}
			return true
		}
	}

	return false
}
