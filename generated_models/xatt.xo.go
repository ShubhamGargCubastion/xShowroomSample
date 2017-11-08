// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XAtt represents a row from 'x_showroom.x_att'.
type XAtt struct {
	ID        uint           `json:"id"`         // id
	Name      string         `json:"name"`       // name
	FolderID  sql.NullInt64  `json:"folder_id"`  // folder_id
	OrgID     uint           `json:"org_id"`     // org_id
	CreatedBy uint           `json:"created_by"` // created_by
	UpdatedBy uint           `json:"updated_by"` // updated_by
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XAtt exists in the database.
func (xa *XAtt) Exists() bool {
	return xa._exists
}

// Deleted provides information if the XAtt has been deleted from the database.
func (xa *XAtt) Deleted() bool {
	return xa._deleted
}

// Insert inserts the XAtt to the database.
func (xa *XAtt) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xa._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_att (` +
		`name, folder_id, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xa.Name, xa.FolderID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt)
	res, err := db.Exec(sqlstr, xa.Name, xa.FolderID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xa.ID = uint(id)
	xa._exists = true

	return nil
}

// Update updates the XAtt in the database.
func (xa *XAtt) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xa._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_att SET ` +
		`name = ?, folder_id = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xa.Name, xa.FolderID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt, xa.ID)
	_, err = db.Exec(sqlstr, xa.Name, xa.FolderID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt, xa.ID)
	return err
}

// Save saves the XAtt to the database.
func (xa *XAtt) Save(db XODB) error {
	if xa.Exists() {
		return xa.Update(db)
	}

	return xa.Insert(db)
}

// Delete deletes the XAtt from the database.
func (xa *XAtt) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xa._exists {
		return nil
	}

	// if deleted, bail
	if xa._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_att WHERE id = ?`

	// run query
	XOLog(sqlstr, xa.ID)
	_, err = db.Exec(sqlstr, xa.ID)
	if err != nil {
		return err
	}

	// set deleted
	xa._deleted = true

	return nil
}

// XUserByCreatedBy returns the XUser associated with the XAtt's CreatedBy (created_by).
//
// Generated from foreign key 'x_att_created_by_foreign'.
func (xa *XAtt) XUserByCreatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xa.CreatedBy)
}

// XAssetFolder returns the XAssetFolder associated with the XAtt's FolderID (folder_id).
//
// Generated from foreign key 'x_att_folder_id_foreign'.
func (xa *XAtt) XAssetFolder(db XODB) (*XAssetFolder, error) {
	return XAssetFolderByID(db, uint(xa.FolderID.Int64))
}

// XOrgExt returns the XOrgExt associated with the XAtt's OrgID (org_id).
//
// Generated from foreign key 'x_att_org_id_foreign'.
func (xa *XAtt) XOrgExt(db XODB) (*XOrgExt, error) {
	return XOrgExtByID(db, xa.OrgID)
}

// XUserByUpdatedBy returns the XUser associated with the XAtt's UpdatedBy (updated_by).
//
// Generated from foreign key 'x_att_updated_by_foreign'.
func (xa *XAtt) XUserByUpdatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xa.UpdatedBy)
}

// XAttsByCreatedBy retrieves a row from 'x_showroom.x_att' as a XAtt.
//
// Generated from index 'x_att_created_by_foreign'.
func XAttsByCreatedBy(db XODB, createdBy uint) ([]*XAtt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, folder_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAtt{}
	for q.Next() {
		xa := XAtt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.FolderID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}

// XAttsByFolderID retrieves a row from 'x_showroom.x_att' as a XAtt.
//
// Generated from index 'x_att_folder_id_foreign'.
func XAttsByFolderID(db XODB, folderID sql.NullInt64) ([]*XAtt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, folder_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att ` +
		`WHERE folder_id = ?`

	// run query
	XOLog(sqlstr, folderID)
	q, err := db.Query(sqlstr, folderID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAtt{}
	for q.Next() {
		xa := XAtt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.FolderID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}

// XAttByID retrieves a row from 'x_showroom.x_att' as a XAtt.
//
// Generated from index 'x_att_id_pkey'.
func XAttByID(db XODB, id uint) (*XAtt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, folder_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xa := XAtt{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xa.ID, &xa.Name, &xa.FolderID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xa, nil
}

// XAttsByOrgID retrieves a row from 'x_showroom.x_att' as a XAtt.
//
// Generated from index 'x_att_org_id_foreign'.
func XAttsByOrgID(db XODB, orgID uint) ([]*XAtt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, folder_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAtt{}
	for q.Next() {
		xa := XAtt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.FolderID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}

// XAttsByUpdatedBy retrieves a row from 'x_showroom.x_att' as a XAtt.
//
// Generated from index 'x_att_updated_by_foreign'.
func XAttsByUpdatedBy(db XODB, updatedBy uint) ([]*XAtt, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, folder_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAtt{}
	for q.Next() {
		xa := XAtt{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.FolderID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}