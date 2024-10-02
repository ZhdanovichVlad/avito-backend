package postgresdb

import "fmt"

// CheckOrganizationExists check if the Organizatio exists in the Organizatio database. Return TRUE if User EXISTS
func (s *Storage) OrganizationExists(organizationid string) (bool, error) {
	const op = "storage.CheckCompanyExists"
	var Exists bool
	query := `SELECT EXISTS (SELECT 1 FROM organization WHERE id = $1)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	err = stmt.QueryRow(organizationid).Scan(&Exists)
	if err != nil {
		return false, nil
	}
	return Exists, nil
}
