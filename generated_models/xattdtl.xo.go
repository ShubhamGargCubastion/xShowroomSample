// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XAttDtl represents a row from 'x_showroom.x_att_dtl'.
type XAttDtl struct {
	ID             uint           `json:"id"`              // id
	PlaceHolder    string         `json:"place_holder"`    // place_holder
	VersionNum     int            `json:"version_num"`     // version_num
	FileExtType    sql.NullString `json:"file_ext_type"`   // file_ext_type
	AttachmentName string         `json:"attachment_name"` // attachment_name
	AttID          uint           `json:"att_id"`          // att_id
	CreatedBy      uint           `json:"created_by"`      // created_by
	UpdatedBy      uint           `json:"updated_by"`      // updated_by
	CreatedAt      mysql.NullTime `json:"created_at"`      // created_at
	UpdatedAt      mysql.NullTime `json:"updated_at"`      // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XAttDtl exists in the database.
func (xad *XAttDtl) Exists() bool {
	return xad._exists
}

// Deleted provides information if the XAttDtl has been deleted from the database.
func (xad *XAttDtl) Deleted() bool {
	return xad._deleted
}

// Insert inserts the XAttDtl to the database.
func (xad *XAttDtl) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xad._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_att_dtl (` +
		`place_holder, version_num, file_ext_type, attachment_name, att_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xad.PlaceHolder, xad.VersionNum, xad.FileExtType, xad.AttachmentName, xad.AttID, xad.CreatedBy, xad.UpdatedBy, xad.CreatedAt, xad.UpdatedAt)
	res, err := db.Exec(sqlstr, xad.PlaceHolder, xad.VersionNum, xad.FileExtType, xad.AttachmentName, xad.AttID, xad.CreatedBy, xad.UpdatedBy, xad.CreatedAt, xad.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xad.ID = uint(id)
	xad._exists = true

	return nil
}

// Update updates the XAttDtl in the database.
func (xad *XAttDtl) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xad._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xad._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_att_dtl SET ` +
		`place_holder = ?, version_num = ?, file_ext_type = ?, attachment_name = ?, att_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xad.PlaceHolder, xad.VersionNum, xad.FileExtType, xad.AttachmentName, xad.AttID, xad.CreatedBy, xad.UpdatedBy, xad.CreatedAt, xad.UpdatedAt, xad.ID)
	_, err = db.Exec(sqlstr, xad.PlaceHolder, xad.VersionNum, xad.FileExtType, xad.AttachmentName, xad.AttID, xad.CreatedBy, xad.UpdatedBy, xad.CreatedAt, xad.UpdatedAt, xad.ID)
	return err
}

// Save saves the XAttDtl to the database.
func (xad *XAttDtl) Save(db XODB) error {
	if xad.Exists() {
		return xad.Update(db)
	}

	return xad.Insert(db)
}

// Delete deletes the XAttDtl from the database.
func (xad *XAttDtl) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xad._exists {
		return nil
	}

	// if deleted, bail
	if xad._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_att_dtl WHERE id = ?`

	// run query
	XOLog(sqlstr, xad.ID)
	_, err = db.Exec(sqlstr, xad.ID)
	if err != nil {
		return err
	}

	// set deleted
	xad._deleted = true

	return nil
}

// XAtt returns the XAtt associated with the XAttDtl's AttID (att_id).
//
// Generated from foreign key 'x_att_dtl_att_id_foreign'.
func (xad *XAttDtl) XAtt(db XODB) (*XAtt, error) {
	return XAttByID(db, xad.AttID)
}

// XUserByCreatedBy returns the XUser associated with the XAttDtl's CreatedBy (created_by).
//
// Generated from foreign key 'x_att_dtl_created_by_foreign'.
func (xad *XAttDtl) XUserByCreatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xad.CreatedBy)
}

// XUserByUpdatedBy returns the XUser associated with the XAttDtl's UpdatedBy (updated_by).
//
// Generated from foreign key 'x_att_dtl_updated_by_foreign'.
func (xad *XAttDtl) XUserByUpdatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xad.UpdatedBy)
}

// XAttDtlByAttIDPlaceHolder retrieves a row from 'x_showroom.x_att_dtl' as a XAttDtl.
//
// Generated from index 'x_att_dtl_att_id_place_holder_unique'.
func XAttDtlByAttIDPlaceHolder(db XODB, attID uint, placeHolder string) (*XAttDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, place_holder, version_num, file_ext_type, attachment_name, att_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att_dtl ` +
		`WHERE att_id = ? AND place_holder = ?`

	// run query
	XOLog(sqlstr, attID, placeHolder)
	xad := XAttDtl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, attID, placeHolder).Scan(&xad.ID, &xad.PlaceHolder, &xad.VersionNum, &xad.FileExtType, &xad.AttachmentName, &xad.AttID, &xad.CreatedBy, &xad.UpdatedBy, &xad.CreatedAt, &xad.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xad, nil
}

// XAttDtlsByCreatedBy retrieves a row from 'x_showroom.x_att_dtl' as a XAttDtl.
//
// Generated from index 'x_att_dtl_created_by_foreign'.
func XAttDtlsByCreatedBy(db XODB, createdBy uint) ([]*XAttDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, place_holder, version_num, file_ext_type, attachment_name, att_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att_dtl ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAttDtl{}
	for q.Next() {
		xad := XAttDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xad.ID, &xad.PlaceHolder, &xad.VersionNum, &xad.FileExtType, &xad.AttachmentName, &xad.AttID, &xad.CreatedBy, &xad.UpdatedBy, &xad.CreatedAt, &xad.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xad)
	}

	return res, nil
}

// XAttDtlByID retrieves a row from 'x_showroom.x_att_dtl' as a XAttDtl.
//
// Generated from index 'x_att_dtl_id_pkey'.
func XAttDtlByID(db XODB, id uint) (*XAttDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, place_holder, version_num, file_ext_type, attachment_name, att_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att_dtl ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xad := XAttDtl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xad.ID, &xad.PlaceHolder, &xad.VersionNum, &xad.FileExtType, &xad.AttachmentName, &xad.AttID, &xad.CreatedBy, &xad.UpdatedBy, &xad.CreatedAt, &xad.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xad, nil
}

// XAttDtlsByUpdatedBy retrieves a row from 'x_showroom.x_att_dtl' as a XAttDtl.
//
// Generated from index 'x_att_dtl_updated_by_foreign'.
func XAttDtlsByUpdatedBy(db XODB, updatedBy uint) ([]*XAttDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, place_holder, version_num, file_ext_type, attachment_name, att_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_att_dtl ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAttDtl{}
	for q.Next() {
		xad := XAttDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xad.ID, &xad.PlaceHolder, &xad.VersionNum, &xad.FileExtType, &xad.AttachmentName, &xad.AttID, &xad.CreatedBy, &xad.UpdatedBy, &xad.CreatedAt, &xad.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xad)
	}

	return res, nil
}