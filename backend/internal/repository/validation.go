package repository

import (
	"avitoTest/backend/pkg/errorsx"
	"context"
	"database/sql"
	"errors"
)

// ValidateResponsibleEmployee method checks whether the employee is responsible in the organization
func (s *Storage) ValidateResponsibleEmployee(context context.Context, organizationId, creatorUsername string) (bool, error) {
	const op = "repository.ValidateResponsibleEmployee"

	var exists bool
	query := `SELECT EXISTS (SELECT 1
	FROM employee as e join
	organization_responsible as o
	on e.id = o.user_id
	where e.username = $1 and o.organization_id = $2)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, errorsx.ErrScanRows
	}

	err = stmt.QueryRowContext(context, creatorUsername, organizationId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, errorsx.ErrScanRows
		}
	}
	return exists, nil
}
