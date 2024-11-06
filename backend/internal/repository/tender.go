package repository

import (
	"avitoTest/backend/internal/entity"
	"avitoTest/backend/pkg/errorsx"
	"context"
	"database/sql"
	"errors"
)

//// NewTenderStorage
//func (s *Storage) NewTenderStorage() error {
//	createType := `
//	CREATE TYPE service_type AS ENUM ('Construction', 'Delivery', 'Manufacture');
//`
//	_, err := s.db.Exec(createType)
//	if err != nil {
//		msgErr := fmt.Errorf("service_type have already been created", err)
//		log.Println(msgErr)
//	}
//
//	createType = `
//	CREATE TYPE tender_status AS ENUM ('Created', 'Published', 'Closed');
//`
//	_, err = s.db.Exec(createType)
//	if err != nil {
//		msgErr := fmt.Errorf("tender_status have already been created", err)
//		log.Println(msgErr)
//	}
//
//	createTableSQL := `
//  	CREATE TABLE IF NOT EXISTS tenders (
//		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//		name VARCHAR(100) NOT NULL,
//		description VARCHAR(500) NOT NULL,
//		serviceType service_type NOT NULL,
//		status tender_status NOT NULL,
//	    organizationId UUID REFERENCES organization(id) ON DELETE CASCADE,
//		creatorUsername VARCHAR(50) REFERENCES employee(username),
//		version INT DEFAULT 1 NOT NULL,
//		createdAt TIMESTAMP NOT NULL
//	 );`
//
//	_, err = s.db.Exec(createTableSQL)
//	if err != nil {
//		msgErr := fmt.Errorf("Error creating table:", err)
//		log.Println(msgErr)
//		return msgErr
//	}
//	return nil
//}

// CreateTender Creating a new tender
func (s *Storage) CreateTender(context context.Context, tender *entity.Tender) (string, error) {
	const op = "repository.tender.CreateTender"

	insertQuery := `
		INSERT INTO tenders (name, description, serviceType, status, organizationId, creatorUsername,version, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	stmt, err := s.db.Prepare(insertQuery)
	defer stmt.Close()
	if err != nil {
		errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	var ID string
	err = stmt.QueryRowContext(context, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationId, tender.CreatorUsername, tender.Version, tender.CreatedAt).Scan(&ID)
	if err != nil {
		errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}
	return ID, nil
}

func (s *Storage) GetTenders(context context.Context, limit, offset int) ([]entity.Tender, error) {
	const op = "repository.tender.GetTenders"

	status := "Published"
	tenders := make([]entity.Tender, limit)

	query := "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders where status=$1 ORDER BY name LIMIT $2 OFFSET $3"
	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	rows, err := stmt.QueryContext(context, status, limit, offset)
	defer rows.Close()
	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}

	for rows.Next() {
		tender := entity.Tender{}
		err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType,
			&tender.Status, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return nil, errorsx.New(errorsx.ErrInternal, "rows.Next error", op, err)
		}
		tenders = append(tenders, tender)
	}

	if err = rows.Err(); err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "rows.Next finished with error", op, err)
	}

	return tenders, nil
}

// GetTendersWithServiceType
func (s *Storage) GetTendersWithServiceType(context context.Context, limit, offset int, ServiceType string) ([]entity.Tender, error) {
	const op = "repository.tender.GetTendersWithServiceType"
	status := "Published"
	query := "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders WHERE status=$1 and serviceType = $2 ORDER BY name LIMIT $3 OFFSET $4"
	tenders := make([]entity.Tender, limit)

	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	rows, err := stmt.QueryContext(context, status, ServiceType, limit, offset)
	defer rows.Close()

	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}

	for rows.Next() {
		tender := entity.Tender{}
		err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType,
			&tender.Status, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return nil, errorsx.New(errorsx.ErrInternal, "rows.Next error", op, err)
		}
		tenders = append(tenders, tender)
	}

	if err = rows.Err(); err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "rows.Next finished with error", op, err)
	}
	return tenders, nil
}

// GetTenders
func (s *Storage) GetUserTenders(context context.Context, limit, offset int, user string) ([]entity.Tender, error) {
	const op = "repository.tender.GetUserTenders"
	query := "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders WHERE  creatorUsername = $1 ORDER BY name LIMIT $2 OFFSET $3"
	tenders := make([]entity.Tender, limit)

	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	rows, err := stmt.QueryContext(context, user, limit, offset)
	defer rows.Close()

	if err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}

	for rows.Next() {
		tender := entity.Tender{}
		err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType,
			&tender.Status, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return nil, errorsx.New(errorsx.ErrNotFound, "rows.Next error", op, err)
		}
		tenders = append(tenders, tender)
	}

	if err = rows.Err(); err != nil {
		return nil, errorsx.New(errorsx.ErrInternal, "rows.Next finished with error", op, err)
	}
	return tenders, nil
}

func (s *Storage) GetTenderStatus(context context.Context, id string) (string, string, error) {
	const op = "repository.tender.GetTenderStatus"

	var status string
	var companId string
	query := "SELECT status, organizationId FROM tenders WHERE id = $1"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	err = stmt.QueryRowContext(context, id).Scan(&status, &companId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", errorsx.New(errorsx.ErrNotFound, "The tender not found", op, err)
		} else {
			return "", "", errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
		}
	}
	return status, companId, nil
}

// GetTender returns one tender by id, or not found error if not found
func (s *Storage) GetTender(context context.Context, tenderId string) (entity.Tender, error) {
	const op = "repository.tender.GetTender"
	var tender entity.Tender

	var query = "SELECT * FROM tenders where id=$1 "

	stmt, err := s.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return tender, errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	err = stmt.QueryRowContext(context, tenderId).Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationId, &tender.CreatorUsername, &tender.Version, &tender.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return tender, errorsx.New(errorsx.ErrNotFound, "The tender not found", op, err)
		}
		return tender, errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}
	return tender, nil
}

// UpdateTenderStatus update Tender Status from tenders DB
func (s *Storage) UpdateTenderStatus(context context.Context, tenderId string, status string) (string, error) {
	const op = "repository.tender.UpdateTenderStatus"

	query := "UPDATE tenders SET status = $1 WHERE id = $2 RETURNING *;"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return "", errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	err = stmt.QueryRowContext(context, status, tenderId).Scan(&status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errorsx.New(errorsx.ErrNotFound, "The tender status not found", op, err)
		}
		return "", errorsx.New(errorsx.ErrInternal, "QueryRow error", op, err)
	}

	return status, nil
}

func (s *Storage) UpdateTender(context context.Context, tender *entity.Tender) error {
	const op = "storage.tender.UpdateTender"
	query := "UPDATE tenders  SET name = $1, description = $2, serviceType = $3 WHERE id = $4 "

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	_, err = stmt.ExecContext(context, tender.Name, tender.Description, tender.ServiceType, tender.Id)
	if err != nil {
		return errorsx.New(errorsx.ErrInternal, "аailed to update status", op, err)
	}
	return nil
}

func (s *Storage) RollBackTender(context context.Context, tenderID string, version int) error {
	const op = "storage.tender.RollBackTender"
	query := "SELECT restore_tender_to_version($1, $2);"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return errorsx.New(errorsx.ErrInternal, "preparation query error", op, err)
	}

	_, err = stmt.ExecContext(context, tenderID, version)
	if err != nil {
		return errorsx.New(errorsx.ErrInternal, "аailed to update status", op, err)
	}
	return nil
}
