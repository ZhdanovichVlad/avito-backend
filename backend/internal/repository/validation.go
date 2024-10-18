package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

// ValidateResponsibleEmployee.
func (s *Storage) ValidateResponsibleEmployee(organizationId, creatorUsername string) (bool, error) {
	const op = "repository.ValidateResponsibleEmployee"

	var exists bool
	query := `SELECT EXISTS (SELECT 1
	FROM employee as e join
	organization_responsible as o
	on e.id = o.user_id
	where e.username = $1 and o.organization_id = $2)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	err = stmt.QueryRow(creatorUsername, organizationId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("%s. Error executing query: %v", op, err)
		}
	}
	return exists, nil
}
