// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdRptGroup represents a row from 'x_showroom.x_prod_rpt_group'.
type XProdRptGroup struct {
	ID          uint           `json:"id"`           // id
	Name        string         `json:"name"`         // name
	DisplayName string         `json:"display_name"` // display_name
	OrgID       uint           `json:"org_id"`       // org_id
	CreatedBy   uint           `json:"created_by"`   // created_by
	UpdatedBy   uint           `json:"updated_by"`   // updated_by
	CreatedAt   mysql.NullTime `json:"created_at"`   // created_at
	UpdatedAt   mysql.NullTime `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdRptGroup exists in the database.
func (xprg *XProdRptGroup) Exists() bool {
	return xprg._exists
}

// Deleted provides information if the XProdRptGroup has been deleted from the database.
func (xprg *XProdRptGroup) Deleted() bool {
	return xprg._deleted
}

// Insert inserts the XProdRptGroup to the database.
func (xprg *XProdRptGroup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xprg._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_rpt_group (` +
		`name, display_name, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xprg.Name, xprg.DisplayName, xprg.OrgID, xprg.CreatedBy, xprg.UpdatedBy, xprg.CreatedAt, xprg.UpdatedAt)
	res, err := db.Exec(sqlstr, xprg.Name, xprg.DisplayName, xprg.OrgID, xprg.CreatedBy, xprg.UpdatedBy, xprg.CreatedAt, xprg.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xprg.ID = uint(id)
	xprg._exists = true

	return nil
}

// Update updates the XProdRptGroup in the database.
func (xprg *XProdRptGroup) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xprg._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xprg._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_rpt_group SET ` +
		`name = ?, display_name = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xprg.Name, xprg.DisplayName, xprg.OrgID, xprg.CreatedBy, xprg.UpdatedBy, xprg.CreatedAt, xprg.UpdatedAt, xprg.ID)
	_, err = db.Exec(sqlstr, xprg.Name, xprg.DisplayName, xprg.OrgID, xprg.CreatedBy, xprg.UpdatedBy, xprg.CreatedAt, xprg.UpdatedAt, xprg.ID)
	return err
}

// Save saves the XProdRptGroup to the database.
func (xprg *XProdRptGroup) Save(db XODB) error {
	if xprg.Exists() {
		return xprg.Update(db)
	}

	return xprg.Insert(db)
}

// Delete deletes the XProdRptGroup from the database.
func (xprg *XProdRptGroup) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xprg._exists {
		return nil
	}

	// if deleted, bail
	if xprg._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_rpt_group WHERE id = ?`

	// run query
	XOLog(sqlstr, xprg.ID)
	_, err = db.Exec(sqlstr, xprg.ID)
	if err != nil {
		return err
	}

	// set deleted
	xprg._deleted = true

	return nil
}

// XProdRptGroupsByCreatedBy retrieves a row from 'x_showroom.x_prod_rpt_group' as a XProdRptGroup.
//
// Generated from index 'x_prod_rpt_group_created_by_foreign'.
func XProdRptGroupsByCreatedBy(db XODB, createdBy uint) ([]*XProdRptGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rpt_group ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRptGroup{}
	for q.Next() {
		xprg := XProdRptGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprg.ID, &xprg.Name, &xprg.DisplayName, &xprg.OrgID, &xprg.CreatedBy, &xprg.UpdatedBy, &xprg.CreatedAt, &xprg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprg)
	}

	return res, nil
}

// XProdRptGroupByID retrieves a row from 'x_showroom.x_prod_rpt_group' as a XProdRptGroup.
//
// Generated from index 'x_prod_rpt_group_id_pkey'.
func XProdRptGroupByID(db XODB, id uint) (*XProdRptGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rpt_group ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xprg := XProdRptGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xprg.ID, &xprg.Name, &xprg.DisplayName, &xprg.OrgID, &xprg.CreatedBy, &xprg.UpdatedBy, &xprg.CreatedAt, &xprg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xprg, nil
}

// XProdRptGroupsByOrgID retrieves a row from 'x_showroom.x_prod_rpt_group' as a XProdRptGroup.
//
// Generated from index 'x_prod_rpt_group_org_id_foreign'.
func XProdRptGroupsByOrgID(db XODB, orgID uint) ([]*XProdRptGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rpt_group ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRptGroup{}
	for q.Next() {
		xprg := XProdRptGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprg.ID, &xprg.Name, &xprg.DisplayName, &xprg.OrgID, &xprg.CreatedBy, &xprg.UpdatedBy, &xprg.CreatedAt, &xprg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprg)
	}

	return res, nil
}

// XProdRptGroupsByUpdatedBy retrieves a row from 'x_showroom.x_prod_rpt_group' as a XProdRptGroup.
//
// Generated from index 'x_prod_rpt_group_updated_by_foreign'.
func XProdRptGroupsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdRptGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rpt_group ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRptGroup{}
	for q.Next() {
		xprg := XProdRptGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprg.ID, &xprg.Name, &xprg.DisplayName, &xprg.OrgID, &xprg.CreatedBy, &xprg.UpdatedBy, &xprg.CreatedAt, &xprg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprg)
	}

	return res, nil
}

// XProdRptGroupByNameOrgID retrieves a row from 'x_showroom.x_prod_rpt_group' as a XProdRptGroup.
//
// Generated from index 'x_unique_key_1'.
func XProdRptGroupByNameOrgID(db XODB, name string, orgID uint) (*XProdRptGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rpt_group ` +
		`WHERE name = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, name, orgID)
	xprg := XProdRptGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, orgID).Scan(&xprg.ID, &xprg.Name, &xprg.DisplayName, &xprg.OrgID, &xprg.CreatedBy, &xprg.UpdatedBy, &xprg.CreatedAt, &xprg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xprg, nil
}
