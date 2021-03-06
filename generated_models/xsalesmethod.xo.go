// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XSalesMethod represents a row from 'x_showroom.x_sales_method'.
type XSalesMethod struct {
	ID        uint           `json:"id"`         // id
	Name      string         `json:"name"`       // name
	DescTxt   string         `json:"desc_txt"`   // desc_txt
	OrgID     sql.NullInt64  `json:"org_id"`     // org_id
	CreatedBy uint           `json:"created_by"` // created_by
	UpdatedBy uint           `json:"updated_by"` // updated_by
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XSalesMethod exists in the database.
func (xsm *XSalesMethod) Exists() bool {
	return xsm._exists
}

// Deleted provides information if the XSalesMethod has been deleted from the database.
func (xsm *XSalesMethod) Deleted() bool {
	return xsm._deleted
}

// Insert inserts the XSalesMethod to the database.
func (xsm *XSalesMethod) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xsm._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_sales_method (` +
		`name, desc_txt, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xsm.Name, xsm.DescTxt, xsm.OrgID, xsm.CreatedBy, xsm.UpdatedBy, xsm.CreatedAt, xsm.UpdatedAt)
	res, err := db.Exec(sqlstr, xsm.Name, xsm.DescTxt, xsm.OrgID, xsm.CreatedBy, xsm.UpdatedBy, xsm.CreatedAt, xsm.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xsm.ID = uint(id)
	xsm._exists = true

	return nil
}

// Update updates the XSalesMethod in the database.
func (xsm *XSalesMethod) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xsm._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xsm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_sales_method SET ` +
		`name = ?, desc_txt = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xsm.Name, xsm.DescTxt, xsm.OrgID, xsm.CreatedBy, xsm.UpdatedBy, xsm.CreatedAt, xsm.UpdatedAt, xsm.ID)
	_, err = db.Exec(sqlstr, xsm.Name, xsm.DescTxt, xsm.OrgID, xsm.CreatedBy, xsm.UpdatedBy, xsm.CreatedAt, xsm.UpdatedAt, xsm.ID)
	return err
}

// Save saves the XSalesMethod to the database.
func (xsm *XSalesMethod) Save(db XODB) error {
	if xsm.Exists() {
		return xsm.Update(db)
	}

	return xsm.Insert(db)
}

// Delete deletes the XSalesMethod from the database.
func (xsm *XSalesMethod) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xsm._exists {
		return nil
	}

	// if deleted, bail
	if xsm._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_sales_method WHERE id = ?`

	// run query
	XOLog(sqlstr, xsm.ID)
	_, err = db.Exec(sqlstr, xsm.ID)
	if err != nil {
		return err
	}

	// set deleted
	xsm._deleted = true

	return nil
}

// XSalesMethodsByCreatedBy retrieves a row from 'x_showroom.x_sales_method' as a XSalesMethod.
//
// Generated from index 'x_sales_method_created_by_foreign'.
func XSalesMethodsByCreatedBy(db XODB, createdBy uint) ([]*XSalesMethod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, desc_txt, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_method ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesMethod{}
	for q.Next() {
		xsm := XSalesMethod{
			_exists: true,
		}

		// scan
		err = q.Scan(&xsm.ID, &xsm.Name, &xsm.DescTxt, &xsm.OrgID, &xsm.CreatedBy, &xsm.UpdatedBy, &xsm.CreatedAt, &xsm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xsm)
	}

	return res, nil
}

// XSalesMethodByID retrieves a row from 'x_showroom.x_sales_method' as a XSalesMethod.
//
// Generated from index 'x_sales_method_id_pkey'.
func XSalesMethodByID(db XODB, id uint) (*XSalesMethod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, desc_txt, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_method ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xsm := XSalesMethod{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xsm.ID, &xsm.Name, &xsm.DescTxt, &xsm.OrgID, &xsm.CreatedBy, &xsm.UpdatedBy, &xsm.CreatedAt, &xsm.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xsm, nil
}

// XSalesMethodByNameOrgID retrieves a row from 'x_showroom.x_sales_method' as a XSalesMethod.
//
// Generated from index 'x_sales_method_name_org_id_unique'.
func XSalesMethodByNameOrgID(db XODB, name string, orgID sql.NullInt64) (*XSalesMethod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, desc_txt, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_method ` +
		`WHERE name = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, name, orgID)
	xsm := XSalesMethod{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, orgID).Scan(&xsm.ID, &xsm.Name, &xsm.DescTxt, &xsm.OrgID, &xsm.CreatedBy, &xsm.UpdatedBy, &xsm.CreatedAt, &xsm.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xsm, nil
}

// XSalesMethodsByOrgID retrieves a row from 'x_showroom.x_sales_method' as a XSalesMethod.
//
// Generated from index 'x_sales_method_org_id_foreign'.
func XSalesMethodsByOrgID(db XODB, orgID sql.NullInt64) ([]*XSalesMethod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, desc_txt, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_method ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesMethod{}
	for q.Next() {
		xsm := XSalesMethod{
			_exists: true,
		}

		// scan
		err = q.Scan(&xsm.ID, &xsm.Name, &xsm.DescTxt, &xsm.OrgID, &xsm.CreatedBy, &xsm.UpdatedBy, &xsm.CreatedAt, &xsm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xsm)
	}

	return res, nil
}

// XSalesMethodsByUpdatedBy retrieves a row from 'x_showroom.x_sales_method' as a XSalesMethod.
//
// Generated from index 'x_sales_method_updated_by_foreign'.
func XSalesMethodsByUpdatedBy(db XODB, updatedBy uint) ([]*XSalesMethod, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, desc_txt, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_method ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesMethod{}
	for q.Next() {
		xsm := XSalesMethod{
			_exists: true,
		}

		// scan
		err = q.Scan(&xsm.ID, &xsm.Name, &xsm.DescTxt, &xsm.OrgID, &xsm.CreatedBy, &xsm.UpdatedBy, &xsm.CreatedAt, &xsm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xsm)
	}

	return res, nil
}
