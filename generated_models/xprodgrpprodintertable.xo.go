// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdGrpProdIntertable represents a row from 'x_showroom.x_prod_grp_prod_intertables'.
type XProdGrpProdIntertable struct {
	ID               uint           `json:"id"`                 // id
	ProdRelprodgrpID uint           `json:"prod_relprodgrp_id"` // prod_relprodgrp_id
	ProdID           uint           `json:"prod_id"`            // prod_id
	SeqNum           uint           `json:"seq_num"`            // seq_num
	CreatedBy        uint           `json:"created_by"`         // created_by
	UpdatedBy        uint           `json:"updated_by"`         // updated_by
	CreatedAt        mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt        mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdGrpProdIntertable exists in the database.
func (xpgpi *XProdGrpProdIntertable) Exists() bool {
	return xpgpi._exists
}

// Deleted provides information if the XProdGrpProdIntertable has been deleted from the database.
func (xpgpi *XProdGrpProdIntertable) Deleted() bool {
	return xpgpi._deleted
}

// Insert inserts the XProdGrpProdIntertable to the database.
func (xpgpi *XProdGrpProdIntertable) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpgpi._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_grp_prod_intertables (` +
		`prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpgpi.ProdRelprodgrpID, xpgpi.ProdID, xpgpi.SeqNum, xpgpi.CreatedBy, xpgpi.UpdatedBy, xpgpi.CreatedAt, xpgpi.UpdatedAt)
	res, err := db.Exec(sqlstr, xpgpi.ProdRelprodgrpID, xpgpi.ProdID, xpgpi.SeqNum, xpgpi.CreatedBy, xpgpi.UpdatedBy, xpgpi.CreatedAt, xpgpi.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpgpi.ID = uint(id)
	xpgpi._exists = true

	return nil
}

// Update updates the XProdGrpProdIntertable in the database.
func (xpgpi *XProdGrpProdIntertable) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpgpi._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpgpi._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_grp_prod_intertables SET ` +
		`prod_relprodgrp_id = ?, prod_id = ?, seq_num = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpgpi.ProdRelprodgrpID, xpgpi.ProdID, xpgpi.SeqNum, xpgpi.CreatedBy, xpgpi.UpdatedBy, xpgpi.CreatedAt, xpgpi.UpdatedAt, xpgpi.ID)
	_, err = db.Exec(sqlstr, xpgpi.ProdRelprodgrpID, xpgpi.ProdID, xpgpi.SeqNum, xpgpi.CreatedBy, xpgpi.UpdatedBy, xpgpi.CreatedAt, xpgpi.UpdatedAt, xpgpi.ID)
	return err
}

// Save saves the XProdGrpProdIntertable to the database.
func (xpgpi *XProdGrpProdIntertable) Save(db XODB) error {
	if xpgpi.Exists() {
		return xpgpi.Update(db)
	}

	return xpgpi.Insert(db)
}

// Delete deletes the XProdGrpProdIntertable from the database.
func (xpgpi *XProdGrpProdIntertable) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpgpi._exists {
		return nil
	}

	// if deleted, bail
	if xpgpi._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_grp_prod_intertables WHERE id = ?`

	// run query
	XOLog(sqlstr, xpgpi.ID)
	_, err = db.Exec(sqlstr, xpgpi.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpgpi._deleted = true

	return nil
}

// XProdGrpProdIntertablesByCreatedBy retrieves a row from 'x_showroom.x_prod_grp_prod_intertables' as a XProdGrpProdIntertable.
//
// Generated from index 'x_prod_grp_prod_intertables_created_by_foreign'.
func XProdGrpProdIntertablesByCreatedBy(db XODB, createdBy uint) ([]*XProdGrpProdIntertable, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_grp_prod_intertables ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdGrpProdIntertable{}
	for q.Next() {
		xpgpi := XProdGrpProdIntertable{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpgpi.ID, &xpgpi.ProdRelprodgrpID, &xpgpi.ProdID, &xpgpi.SeqNum, &xpgpi.CreatedBy, &xpgpi.UpdatedBy, &xpgpi.CreatedAt, &xpgpi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpgpi)
	}

	return res, nil
}

// XProdGrpProdIntertableByID retrieves a row from 'x_showroom.x_prod_grp_prod_intertables' as a XProdGrpProdIntertable.
//
// Generated from index 'x_prod_grp_prod_intertables_id_pkey'.
func XProdGrpProdIntertableByID(db XODB, id uint) (*XProdGrpProdIntertable, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_grp_prod_intertables ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpgpi := XProdGrpProdIntertable{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpgpi.ID, &xpgpi.ProdRelprodgrpID, &xpgpi.ProdID, &xpgpi.SeqNum, &xpgpi.CreatedBy, &xpgpi.UpdatedBy, &xpgpi.CreatedAt, &xpgpi.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpgpi, nil
}

// XProdGrpProdIntertablesByProdID retrieves a row from 'x_showroom.x_prod_grp_prod_intertables' as a XProdGrpProdIntertable.
//
// Generated from index 'x_prod_grp_prod_intertables_prod_id_foreign'.
func XProdGrpProdIntertablesByProdID(db XODB, prodID uint) ([]*XProdGrpProdIntertable, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_grp_prod_intertables ` +
		`WHERE prod_id = ?`

	// run query
	XOLog(sqlstr, prodID)
	q, err := db.Query(sqlstr, prodID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdGrpProdIntertable{}
	for q.Next() {
		xpgpi := XProdGrpProdIntertable{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpgpi.ID, &xpgpi.ProdRelprodgrpID, &xpgpi.ProdID, &xpgpi.SeqNum, &xpgpi.CreatedBy, &xpgpi.UpdatedBy, &xpgpi.CreatedAt, &xpgpi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpgpi)
	}

	return res, nil
}

// XProdGrpProdIntertablesByProdRelprodgrpID retrieves a row from 'x_showroom.x_prod_grp_prod_intertables' as a XProdGrpProdIntertable.
//
// Generated from index 'x_prod_grp_prod_intertables_prod_relprodgrp_id_foreign'.
func XProdGrpProdIntertablesByProdRelprodgrpID(db XODB, prodRelprodgrpID uint) ([]*XProdGrpProdIntertable, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_grp_prod_intertables ` +
		`WHERE prod_relprodgrp_id = ?`

	// run query
	XOLog(sqlstr, prodRelprodgrpID)
	q, err := db.Query(sqlstr, prodRelprodgrpID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdGrpProdIntertable{}
	for q.Next() {
		xpgpi := XProdGrpProdIntertable{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpgpi.ID, &xpgpi.ProdRelprodgrpID, &xpgpi.ProdID, &xpgpi.SeqNum, &xpgpi.CreatedBy, &xpgpi.UpdatedBy, &xpgpi.CreatedAt, &xpgpi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpgpi)
	}

	return res, nil
}

// XProdGrpProdIntertablesByUpdatedBy retrieves a row from 'x_showroom.x_prod_grp_prod_intertables' as a XProdGrpProdIntertable.
//
// Generated from index 'x_prod_grp_prod_intertables_updated_by_foreign'.
func XProdGrpProdIntertablesByUpdatedBy(db XODB, updatedBy uint) ([]*XProdGrpProdIntertable, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_relprodgrp_id, prod_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_grp_prod_intertables ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdGrpProdIntertable{}
	for q.Next() {
		xpgpi := XProdGrpProdIntertable{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpgpi.ID, &xpgpi.ProdRelprodgrpID, &xpgpi.ProdID, &xpgpi.SeqNum, &xpgpi.CreatedBy, &xpgpi.UpdatedBy, &xpgpi.CreatedAt, &xpgpi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpgpi)
	}

	return res, nil
}