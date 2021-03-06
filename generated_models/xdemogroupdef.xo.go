// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XDemoGroupDef represents a row from 'x_showroom.x_demo_group_def'.
type XDemoGroupDef struct {
	ID          uint           `json:"id"`           // id
	Name        string         `json:"name"`         // name
	DisplayName string         `json:"display_name"` // display_name
	TypeCd      string         `json:"type_cd"`      // type_cd
	SeqNum      uint           `json:"seq_num"`      // seq_num
	AttID       sql.NullInt64  `json:"att_id"`       // att_id
	OrgID       uint           `json:"org_id"`       // org_id
	CreatedBy   uint           `json:"created_by"`   // created_by
	UpdatedBy   uint           `json:"updated_by"`   // updated_by
	CreatedAt   mysql.NullTime `json:"created_at"`   // created_at
	UpdatedAt   mysql.NullTime `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XDemoGroupDef exists in the database.
func (xdgd *XDemoGroupDef) Exists() bool {
	return xdgd._exists
}

// Deleted provides information if the XDemoGroupDef has been deleted from the database.
func (xdgd *XDemoGroupDef) Deleted() bool {
	return xdgd._deleted
}

// Insert inserts the XDemoGroupDef to the database.
func (xdgd *XDemoGroupDef) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xdgd._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_demo_group_def (` +
		`name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.SeqNum, xdgd.AttID, xdgd.OrgID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt)
	res, err := db.Exec(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.SeqNum, xdgd.AttID, xdgd.OrgID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xdgd.ID = uint(id)
	xdgd._exists = true

	return nil
}

// Update updates the XDemoGroupDef in the database.
func (xdgd *XDemoGroupDef) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xdgd._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xdgd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_demo_group_def SET ` +
		`name = ?, display_name = ?, type_cd = ?, seq_num = ?, att_id = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.SeqNum, xdgd.AttID, xdgd.OrgID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.ID)
	_, err = db.Exec(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.SeqNum, xdgd.AttID, xdgd.OrgID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.ID)
	return err
}

// Save saves the XDemoGroupDef to the database.
func (xdgd *XDemoGroupDef) Save(db XODB) error {
	if xdgd.Exists() {
		return xdgd.Update(db)
	}

	return xdgd.Insert(db)
}

// Delete deletes the XDemoGroupDef from the database.
func (xdgd *XDemoGroupDef) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xdgd._exists {
		return nil
	}

	// if deleted, bail
	if xdgd._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_demo_group_def WHERE id = ?`

	// run query
	XOLog(sqlstr, xdgd.ID)
	_, err = db.Exec(sqlstr, xdgd.ID)
	if err != nil {
		return err
	}

	// set deleted
	xdgd._deleted = true

	return nil
}

// XDemoGroupDevesByAttID retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_att_id_foreign'.
func XDemoGroupDevesByAttID(db XODB, attID sql.NullInt64) ([]*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE att_id = ?`

	// run query
	XOLog(sqlstr, attID)
	q, err := db.Query(sqlstr, attID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDef{}
	for q.Next() {
		xdgd := XDemoGroupDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDevesByCreatedBy retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_created_by_foreign'.
func XDemoGroupDevesByCreatedBy(db XODB, createdBy uint) ([]*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDef{}
	for q.Next() {
		xdgd := XDemoGroupDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDefByID retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_id_pkey'.
func XDemoGroupDefByID(db XODB, id uint) (*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xdgd := XDemoGroupDef{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xdgd, nil
}

// XDemoGroupDefByNameTypeCdOrgID retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_name_type_cd_org_id_unique'.
func XDemoGroupDefByNameTypeCdOrgID(db XODB, name string, typeCd string, orgID uint) (*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE name = ? AND type_cd = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, name, typeCd, orgID)
	xdgd := XDemoGroupDef{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, typeCd, orgID).Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xdgd, nil
}

// XDemoGroupDevesByOrgID retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_org_id_foreign'.
func XDemoGroupDevesByOrgID(db XODB, orgID uint) ([]*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDef{}
	for q.Next() {
		xdgd := XDemoGroupDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDevesByUpdatedBy retrieves a row from 'x_showroom.x_demo_group_def' as a XDemoGroupDef.
//
// Generated from index 'x_demo_group_def_updated_by_foreign'.
func XDemoGroupDevesByUpdatedBy(db XODB, updatedBy uint) ([]*XDemoGroupDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, att_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_demo_group_def ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDef{}
	for q.Next() {
		xdgd := XDemoGroupDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.OrgID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}
