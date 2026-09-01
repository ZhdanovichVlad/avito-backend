package repository

import (
	"avitoTest/backend/pkg/errorsx"
	"context"
	"database/sql"
	"errors"
)

// ValidateResponsibleEmployee method checks whether the employee is responsible in the organization
func (s *Storage) ValidateResponsibleEmployee(context context.Context, organizationId, creatorUsername string) error {
	const op = "repository.validation.ValidateResponsibleEmployee"

	var exists bool
	query := `SELECT EXISTS (SELECT 1
	FROM employee as e join
	organization_responsible as o
	on e.id = o.user_id
	where e.username = $1 and o.organization_id = $2)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	err = stmt.QueryRowContext(context, creatorUsername, organizationId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorsx.New(errorsx.ErrForbidden, "An employee of the organization is not responsible for the organization", op, err)
		} else {
			return errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)

		}
	}
	return nil
}

// CheckUserExists
func (s *Storage) CheckUserExists(context context.Context, creatorUsername string) error {
	const op = "repository.validation.CheckUserExists"
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM employee WHERE username = $1)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	err = stmt.QueryRowContext(context, creatorUsername).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorsx.New(errorsx.ErrBadRequest, "The employee doesn't exist", op, err)
		} else {
			return errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
		}
	}
	return nil
}
