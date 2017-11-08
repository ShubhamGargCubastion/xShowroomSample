// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XPriCompDef represents a row from 'x_showroom.x_pri_comp_def'.
type XPriCompDef struct {
	ID          uint           `json:"id"`           // id
	Name        string         `json:"name"`         // name
	DisplayName string         `json:"display_name"` // display_name
	TypeCd      string         `json:"type_cd"`      // type_cd
	SeqNum      uint           `json:"seq_num"`      // seq_num
	OrgID       uint           `json:"org_id"`       // org_id
	CreatedBy   uint           `json:"created_by"`   // created_by
	UpdatedBy   uint           `json:"updated_by"`   // updated_by
	CreatedAt   mysql.NullTime `json:"created_at"`   // created_at
	UpdatedAt   mysql.NullTime `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XPriCompDef exists in the database.
func (xpcd *XPriCompDef) Exists() bool {
	return xpcd._exists
}

// Deleted provides information if the XPriCompDef has been deleted from the database.
func (xpcd *XPriCompDef) Deleted() bool {
	return xpcd._deleted
}

// Insert inserts the XPriCompDef to the database.
func (xpcd *XPriCompDef) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpcd._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_pri_comp_def (` +
		`name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpcd.Name, xpcd.DisplayName, xpcd.TypeCd, xpcd.SeqNum, xpcd.OrgID, xpcd.CreatedBy, xpcd.UpdatedBy, xpcd.CreatedAt, xpcd.UpdatedAt)
	res, err := db.Exec(sqlstr, xpcd.Name, xpcd.DisplayName, xpcd.TypeCd, xpcd.SeqNum, xpcd.OrgID, xpcd.CreatedBy, xpcd.UpdatedBy, xpcd.CreatedAt, xpcd.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpcd.ID = uint(id)
	xpcd._exists = true

	return nil
}

// Update updates the XPriCompDef in the database.
func (xpcd *XPriCompDef) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpcd._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpcd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_pri_comp_def SET ` +
		`name = ?, display_name = ?, type_cd = ?, seq_num = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpcd.Name, xpcd.DisplayName, xpcd.TypeCd, xpcd.SeqNum, xpcd.OrgID, xpcd.CreatedBy, xpcd.UpdatedBy, xpcd.CreatedAt, xpcd.UpdatedAt, xpcd.ID)
	_, err = db.Exec(sqlstr, xpcd.Name, xpcd.DisplayName, xpcd.TypeCd, xpcd.SeqNum, xpcd.OrgID, xpcd.CreatedBy, xpcd.UpdatedBy, xpcd.CreatedAt, xpcd.UpdatedAt, xpcd.ID)
	return err
}

// Save saves the XPriCompDef to the database.
func (xpcd *XPriCompDef) Save(db XODB) error {
	if xpcd.Exists() {
		return xpcd.Update(db)
	}

	return xpcd.Insert(db)
}

// Delete deletes the XPriCompDef from the database.
func (xpcd *XPriCompDef) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpcd._exists {
		return nil
	}

	// if deleted, bail
	if xpcd._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_pri_comp_def WHERE id = ?`

	// run query
	XOLog(sqlstr, xpcd.ID)
	_, err = db.Exec(sqlstr, xpcd.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpcd._deleted = true

	return nil
}

// XPriCompDevesByCreatedBy retrieves a row from 'x_showroom.x_pri_comp_def' as a XPriCompDef.
//
// Generated from index 'x_pri_comp_def_created_by_foreign'.
func XPriCompDevesByCreatedBy(db XODB, createdBy uint) ([]*XPriCompDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_pri_comp_def ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XPriCompDef{}
	for q.Next() {
		xpcd := XPriCompDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpcd.ID, &xpcd.Name, &xpcd.DisplayName, &xpcd.TypeCd, &xpcd.SeqNum, &xpcd.OrgID, &xpcd.CreatedBy, &xpcd.UpdatedBy, &xpcd.CreatedAt, &xpcd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpcd)
	}

	return res, nil
}

// XPriCompDefByID retrieves a row from 'x_showroom.x_pri_comp_def' as a XPriCompDef.
//
// Generated from index 'x_pri_comp_def_id_pkey'.
func XPriCompDefByID(db XODB, id uint) (*XPriCompDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_pri_comp_def ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpcd := XPriCompDef{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpcd.ID, &xpcd.Name, &xpcd.DisplayName, &xpcd.TypeCd, &xpcd.SeqNum, &xpcd.OrgID, &xpcd.CreatedBy, &xpcd.UpdatedBy, &xpcd.CreatedAt, &xpcd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpcd, nil
}

// XPriCompDefByNameOrgID retrieves a row from 'x_showroom.x_pri_comp_def' as a XPriCompDef.
//
// Generated from index 'x_pri_comp_def_name_org_id_unique'.
func XPriCompDefByNameOrgID(db XODB, name string, orgID uint) (*XPriCompDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_pri_comp_def ` +
		`WHERE name = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, name, orgID)
	xpcd := XPriCompDef{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, orgID).Scan(&xpcd.ID, &xpcd.Name, &xpcd.DisplayName, &xpcd.TypeCd, &xpcd.SeqNum, &xpcd.OrgID, &xpcd.CreatedBy, &xpcd.UpdatedBy, &xpcd.CreatedAt, &xpcd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpcd, nil
}

// XPriCompDevesByOrgID retrieves a row from 'x_showroom.x_pri_comp_def' as a XPriCompDef.
//
// Generated from index 'x_pri_comp_def_org_id_foreign'.
func XPriCompDevesByOrgID(db XODB, orgID uint) ([]*XPriCompDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_pri_comp_def ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XPriCompDef{}
	for q.Next() {
		xpcd := XPriCompDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpcd.ID, &xpcd.Name, &xpcd.DisplayName, &xpcd.TypeCd, &xpcd.SeqNum, &xpcd.OrgID, &xpcd.CreatedBy, &xpcd.UpdatedBy, &xpcd.CreatedAt, &xpcd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpcd)
	}

	return res, nil
}

// XPriCompDevesByUpdatedBy retrieves a row from 'x_showroom.x_pri_comp_def' as a XPriCompDef.
//
// Generated from index 'x_pri_comp_def_updated_by_foreign'.
func XPriCompDevesByUpdatedBy(db XODB, updatedBy uint) ([]*XPriCompDef, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, seq_num, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_pri_comp_def ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XPriCompDef{}
	for q.Next() {
		xpcd := XPriCompDef{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpcd.ID, &xpcd.Name, &xpcd.DisplayName, &xpcd.TypeCd, &xpcd.SeqNum, &xpcd.OrgID, &xpcd.CreatedBy, &xpcd.UpdatedBy, &xpcd.CreatedAt, &xpcd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpcd)
	}

	return res, nil
}