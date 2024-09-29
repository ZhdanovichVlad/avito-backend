package postgresdb

import (
	"avitoTest/backend/domain/tender"
	"avitoTest/backend/infrastructure/storageerror"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func (s *Storage) NewTenderStorage() error {

	createType := `
	CREATE TYPE service_type AS ENUM ('Construction', 'Delivery', 'Manufacture');
`
	_, err := s.db.Exec(createType)
	if err != nil {
		msgErr := fmt.Errorf("service_type have already been created", err)
		log.Println(msgErr)
	}

	createType = `
	CREATE TYPE tender_status AS ENUM ('Created', 'Published', 'Closed');
`
	_, err = s.db.Exec(createType)
	if err != nil {
		msgErr := fmt.Errorf("tender_status have already been created", err)
		log.Println(msgErr)
	}

	createTableSQL := `
  	CREATE TABLE IF NOT EXISTS tenders (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(100) NOT NULL,
		description VARCHAR(500) NOT NULL,
		serviceType service_type NOT NULL,
		status tender_status NOT NULL,
	    organizationId UUID REFERENCES organization(id) ON DELETE CASCADE,
		creatorUsername VARCHAR(50) REFERENCES employee(username),
		version INT DEFAULT 1 NOT NULL,
		createdAt TIMESTAMP NOT NULL
	 );`

	_, err = s.db.Exec(createTableSQL)
	if err != nil {
		msgErr := fmt.Errorf("Error creating table:", err)
		log.Println(msgErr)
		return msgErr
	}

	return nil
}

// CreateTender Creating a new tender
func (s *Storage) CreateTender(tender *tender.Tender) (string, error) {
	const op = "storage.CreateTender"

	insertQuery := `
		INSERT INTO tenders (name, description, serviceType, status, organizationId, creatorUsername,version, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	stmt, err := s.db.Prepare(insertQuery)
	defer stmt.Close()
	if err != nil {
		return "", fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	var ID string
	err = stmt.QueryRow(tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationId, tender.CreatorUsername, tender.Version, tender.CreatedAt).Scan(&ID)
	if err != nil {
		return "", fmt.Errorf("%s. Error executing query: %v", op, err)
	}

	return ID, nil
}

func (s *Storage) GetTenders(tenders *[]tender.Tender, limit, offset int, searchInfo string, serchingType int) error {
	const op = "storage.GetTenders"
	status := "Published"
	var query string
	switch serchingType {
	case 0:
		query = "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders where status=$1 ORDER BY name LIMIT $2 OFFSET $3"
	case 1:
		query = "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders WHERE status=$1 and serviceType = $2 ORDER BY name LIMIT $3 OFFSET $4"
	case 2:
		query = "SELECT id, name, description, serviceType, status, version, createdAt FROM tenders WHERE  creatorUsername = $1 ORDER BY name LIMIT $2 OFFSET $3"
	default:
		return fmt.Errorf("unknown serchingType: %d", serchingType)
	}
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}
	defer stmt.Close()

	var rows *sql.Rows
	switch serchingType {
	case 0:
		rows, err = stmt.Query(status, limit, offset)
	case 1:
		rows, err = stmt.Query(status, searchInfo, limit, offset)
	case 2:
		rows, err = stmt.Query(searchInfo, limit, offset)
	}
	if err != nil {
		return fmt.Errorf("%s. Error executing query: %v", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		tender := tender.Tender{}
		err = rows.Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType,
			&tender.Status, &tender.Version, &tender.CreatedAt)
		if err != nil {
			return fmt.Errorf("%s. failed scan from database: %v", op, err)
		}
		*tenders = append(*tenders, tender)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("%s. rows.Next() contains errors: %v", op, err)
	}

	return nil
}

// GetFullTender returns one tender(name, description, serviceType, version) by id
func (s *Storage) GetFullTender(tender *tender.Tender, tenderId string) error {
	const op = "storage.GetFullTender"

	var query = "SELECT * FROM tenders where id=$1 "

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s. Error preparing statement: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(tenderId).Scan(&tender.Id, &tender.Name, &tender.Description, &tender.ServiceType, &tender.Status, &tender.OrganizationId, &tender.CreatorUsername, &tender.Version, &tender.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storageerror.ErrTenderNotFound
		}
		return fmt.Errorf("error executing query: %w", err)
	}
	return nil
}

// UpdateTenderStatus update Tender Status from tenders DB
func (s *Storage) UpdateTenderStatus(tenderId string, status string) (err error) {
	const op = "storage.UpdateTenderStatus"

	query := "UPDATE tenders SET status = $1 WHERE id = $2"

	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	_, err = stmt.Exec(status, tenderId)
	if err != nil {
		return fmt.Errorf("%s. Error executing query: %v", op, err)
	}

	return nil
}
