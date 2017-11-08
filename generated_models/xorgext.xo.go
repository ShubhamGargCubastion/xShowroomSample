// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XOrgExt represents a row from 'x_showroom.x_org_ext'.
type XOrgExt struct {
	ID              uint           `json:"id"`                // id
	Name            string         `json:"name"`              // name
	IntegrationID   string         `json:"integration_id"`    // integration_id
	AddressLine1    string         `json:"address_line_1"`    // address_line_1
	AddressLine2    string         `json:"address_line_2"`    // address_line_2
	City            string         `json:"city"`              // city
	State           string         `json:"state"`             // state
	Zipcode         string         `json:"zipcode"`           // zipcode
	Country         string         `json:"country"`           // country
	SalesMethodID   sql.NullInt64  `json:"sales_method_id"`   // sales_method_id
	DeviceLevelAuth string         `json:"device_level_auth"` // device_level_auth
	CreatedBy       uint           `json:"created_by"`        // created_by
	UpdatedBy       uint           `json:"updated_by"`        // updated_by
	CreatedAt       mysql.NullTime `json:"created_at"`        // created_at
	UpdatedAt       mysql.NullTime `json:"updated_at"`        // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XOrgExt exists in the database.
func (xoe *XOrgExt) Exists() bool {
	return xoe._exists
}

// Deleted provides information if the XOrgExt has been deleted from the database.
func (xoe *XOrgExt) Deleted() bool {
	return xoe._deleted
}

// Insert inserts the XOrgExt to the database.
func (xoe *XOrgExt) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xoe._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_org_ext (` +
		`name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xoe.Name, xoe.IntegrationID, xoe.AddressLine1, xoe.AddressLine2, xoe.City, xoe.State, xoe.Zipcode, xoe.Country, xoe.SalesMethodID, xoe.DeviceLevelAuth, xoe.CreatedBy, xoe.UpdatedBy, xoe.CreatedAt, xoe.UpdatedAt)
	res, err := db.Exec(sqlstr, xoe.Name, xoe.IntegrationID, xoe.AddressLine1, xoe.AddressLine2, xoe.City, xoe.State, xoe.Zipcode, xoe.Country, xoe.SalesMethodID, xoe.DeviceLevelAuth, xoe.CreatedBy, xoe.UpdatedBy, xoe.CreatedAt, xoe.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xoe.ID = uint(id)
	xoe._exists = true

	return nil
}

// Update updates the XOrgExt in the database.
func (xoe *XOrgExt) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xoe._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xoe._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_org_ext SET ` +
		`name = ?, integration_id = ?, address_line_1 = ?, address_line_2 = ?, city = ?, state = ?, zipcode = ?, country = ?, sales_method_id = ?, device_level_auth = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xoe.Name, xoe.IntegrationID, xoe.AddressLine1, xoe.AddressLine2, xoe.City, xoe.State, xoe.Zipcode, xoe.Country, xoe.SalesMethodID, xoe.DeviceLevelAuth, xoe.CreatedBy, xoe.UpdatedBy, xoe.CreatedAt, xoe.UpdatedAt, xoe.ID)
	_, err = db.Exec(sqlstr, xoe.Name, xoe.IntegrationID, xoe.AddressLine1, xoe.AddressLine2, xoe.City, xoe.State, xoe.Zipcode, xoe.Country, xoe.SalesMethodID, xoe.DeviceLevelAuth, xoe.CreatedBy, xoe.UpdatedBy, xoe.CreatedAt, xoe.UpdatedAt, xoe.ID)
	return err
}

// Save saves the XOrgExt to the database.
func (xoe *XOrgExt) Save(db XODB) error {
	if xoe.Exists() {
		return xoe.Update(db)
	}

	return xoe.Insert(db)
}

// Delete deletes the XOrgExt from the database.
func (xoe *XOrgExt) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xoe._exists {
		return nil
	}

	// if deleted, bail
	if xoe._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_org_ext WHERE id = ?`

	// run query
	XOLog(sqlstr, xoe.ID)
	_, err = db.Exec(sqlstr, xoe.ID)
	if err != nil {
		return err
	}

	// set deleted
	xoe._deleted = true

	return nil
}

// XOrgExtByIntegrationID retrieves a row from 'x_showroom.x_org_ext' as a XOrgExt.
//
// Generated from index 'USER_KEY_1'.
func XOrgExtByIntegrationID(db XODB, integrationID string) (*XOrgExt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_org_ext ` +
		`WHERE integration_id = ?`

	// run query
	XOLog(sqlstr, integrationID)
	xoe := XOrgExt{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, integrationID).Scan(&xoe.ID, &xoe.Name, &xoe.IntegrationID, &xoe.AddressLine1, &xoe.AddressLine2, &xoe.City, &xoe.State, &xoe.Zipcode, &xoe.Country, &xoe.SalesMethodID, &xoe.DeviceLevelAuth, &xoe.CreatedBy, &xoe.UpdatedBy, &xoe.CreatedAt, &xoe.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xoe, nil
}

// XOrgExtsByCreatedBy retrieves a row from 'x_showroom.x_org_ext' as a XOrgExt.
//
// Generated from index 'x_org_ext_created_by_foreign'.
func XOrgExtsByCreatedBy(db XODB, createdBy uint) ([]*XOrgExt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_org_ext ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOrgExt{}
	for q.Next() {
		xoe := XOrgExt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xoe.ID, &xoe.Name, &xoe.IntegrationID, &xoe.AddressLine1, &xoe.AddressLine2, &xoe.City, &xoe.State, &xoe.Zipcode, &xoe.Country, &xoe.SalesMethodID, &xoe.DeviceLevelAuth, &xoe.CreatedBy, &xoe.UpdatedBy, &xoe.CreatedAt, &xoe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xoe)
	}

	return res, nil
}

// XOrgExtByID retrieves a row from 'x_showroom.x_org_ext' as a XOrgExt.
//
// Generated from index 'x_org_ext_id_pkey'.
func XOrgExtByID(db XODB, id uint) (*XOrgExt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_org_ext ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xoe := XOrgExt{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xoe.ID, &xoe.Name, &xoe.IntegrationID, &xoe.AddressLine1, &xoe.AddressLine2, &xoe.City, &xoe.State, &xoe.Zipcode, &xoe.Country, &xoe.SalesMethodID, &xoe.DeviceLevelAuth, &xoe.CreatedBy, &xoe.UpdatedBy, &xoe.CreatedAt, &xoe.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xoe, nil
}

// XOrgExtsBySalesMethodID retrieves a row from 'x_showroom.x_org_ext' as a XOrgExt.
//
// Generated from index 'x_org_ext_sales_method_id_foreign'.
func XOrgExtsBySalesMethodID(db XODB, salesMethodID sql.NullInt64) ([]*XOrgExt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_org_ext ` +
		`WHERE sales_method_id = ?`

	// run query
	XOLog(sqlstr, salesMethodID)
	q, err := db.Query(sqlstr, salesMethodID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOrgExt{}
	for q.Next() {
		xoe := XOrgExt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xoe.ID, &xoe.Name, &xoe.IntegrationID, &xoe.AddressLine1, &xoe.AddressLine2, &xoe.City, &xoe.State, &xoe.Zipcode, &xoe.Country, &xoe.SalesMethodID, &xoe.DeviceLevelAuth, &xoe.CreatedBy, &xoe.UpdatedBy, &xoe.CreatedAt, &xoe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xoe)
	}

	return res, nil
}

// XOrgExtsByUpdatedBy retrieves a row from 'x_showroom.x_org_ext' as a XOrgExt.
//
// Generated from index 'x_org_ext_updated_by_foreign'.
func XOrgExtsByUpdatedBy(db XODB, updatedBy uint) ([]*XOrgExt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, integration_id, address_line_1, address_line_2, city, state, zipcode, country, sales_method_id, device_level_auth, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_org_ext ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOrgExt{}
	for q.Next() {
		xoe := XOrgExt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xoe.ID, &xoe.Name, &xoe.IntegrationID, &xoe.AddressLine1, &xoe.AddressLine2, &xoe.City, &xoe.State, &xoe.Zipcode, &xoe.Country, &xoe.SalesMethodID, &xoe.DeviceLevelAuth, &xoe.CreatedBy, &xoe.UpdatedBy, &xoe.CreatedAt, &xoe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xoe)
	}

	return res, nil
}