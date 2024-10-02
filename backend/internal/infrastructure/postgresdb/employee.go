package postgresdb

import (
	"database/sql"
	"errors"
	"fmt"
)

// Domain
type User struct {
	ID      int
	Name    string
	Surname string
}

type UserSnapshot struct {
	ID      int    `db:"id, int64"`
	Name    string `db:"name, varchar"`
	Surname string `db:"surname, varchar"`
}

// Application <- Infrastructure
type CreateUserDTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type CreateUserRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// CreateUserRequest -> CreateUserDTO
// userService.Create(ctx, CreateUserDTO) -> User
// userStorage.Insert(ctx, User)
// {
// User -> UserSnapshot
// }

// CheckUserExists check if the user exists in the employee database. Return TRUE if User EXISTS
func (s *Storage) CheckUserExists(creatorUsername string) (bool, error) {
	const op = "storage.CheckUserExists"
	var Exists bool
	query := `SELECT EXISTS (SELECT 1 FROM employee WHERE username = $1)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	err = stmt.QueryRow(creatorUsername).Scan(&Exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("%s. Error executing query: %v", op, err)
		}
	}
	return Exists, nil
}

// CheckUserExistsByID check if the user exists in the employee database by ID. Return TRUE if User EXISTS
func (s *Storage) CheckUserExistsByID(userID string) (bool, error) {
	const op = "storage.CheckUserExistsByID"
	var Exists bool
	query := `SELECT EXISTS (SELECT 1 FROM employee WHERE id = $1)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	err = stmt.QueryRow(userID).Scan(&Exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("%s. Error executing query: %v", op, err)
		}
	}
	return Exists, nil
}

//// CheckOrganizationIdAndUserIDExists check if the user exists in the employee database
//func (s *Storage) CheckOrganizationIdAndUserIDExists(organizationId, creatorUsername string) (bool, error) {
//	const op = "storage.CheckOrganizationIdAndUserIDExists"
//
//	var exists bool
//	query := `SELECT EXISTS (SELECT 1
//	FROM employee as e join
//	organization_responsible as o
//	on e.id = o.user_id
//	where e.username = $1 and o.organization_id = $2)`
//
//	stmt, err := s.db.Prepare(query)
//	if err != nil {
//		return false, fmt.Errorf("%s. Error preparing statement: %v", op, err)
//	}
//
//	err = stmt.QueryRow(creatorUsername, organizationId).Scan(&exists)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return false, nil
//		} else {
//			return false, fmt.Errorf("%s. Error executing query: %v", op, err)
//		}
//	}
//	return exists, nil
//}

// GetUserIDUsername returns the user id using the username
func (s *Storage) GetUserIDUsername(username string) (string, error) {
	const op = "storage.CheckUserExistsByID"
	var UserdId string
	query := `SELECT id FROM employee WHERE username = $1)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	err = stmt.QueryRow(username).Scan(&UserdId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		} else {
			return "", fmt.Errorf("%s. Error executing query: %v", op, err)
		}
	}
	return UserdId, nil
}
