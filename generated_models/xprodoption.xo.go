// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdOption represents a row from 'x_showroom.x_prod_option'.
type XProdOption struct {
	ID        uint           `json:"id"`         // id
	ProdID    uint           `json:"prod_id"`    // prod_id
	OptionID  uint           `json:"option_id"`  // option_id
	SeqNum    uint           `json:"seq_num"`    // seq_num
	CreatedBy uint           `json:"created_by"` // created_by
	UpdatedBy uint           `json:"updated_by"` // updated_by
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdOption exists in the database.
func (xpo *XProdOption) Exists() bool {
	return xpo._exists
}

// Deleted provides information if the XProdOption has been deleted from the database.
func (xpo *XProdOption) Deleted() bool {
	return xpo._deleted
}

// Insert inserts the XProdOption to the database.
func (xpo *XProdOption) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpo._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_option (` +
		`prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpo.ProdID, xpo.OptionID, xpo.SeqNum, xpo.CreatedBy, xpo.UpdatedBy, xpo.CreatedAt, xpo.UpdatedAt)
	res, err := db.Exec(sqlstr, xpo.ProdID, xpo.OptionID, xpo.SeqNum, xpo.CreatedBy, xpo.UpdatedBy, xpo.CreatedAt, xpo.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpo.ID = uint(id)
	xpo._exists = true

	return nil
}

// Update updates the XProdOption in the database.
func (xpo *XProdOption) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpo._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpo._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_option SET ` +
		`prod_id = ?, option_id = ?, seq_num = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpo.ProdID, xpo.OptionID, xpo.SeqNum, xpo.CreatedBy, xpo.UpdatedBy, xpo.CreatedAt, xpo.UpdatedAt, xpo.ID)
	_, err = db.Exec(sqlstr, xpo.ProdID, xpo.OptionID, xpo.SeqNum, xpo.CreatedBy, xpo.UpdatedBy, xpo.CreatedAt, xpo.UpdatedAt, xpo.ID)
	return err
}

// Save saves the XProdOption to the database.
func (xpo *XProdOption) Save(db XODB) error {
	if xpo.Exists() {
		return xpo.Update(db)
	}

	return xpo.Insert(db)
}

// Delete deletes the XProdOption from the database.
func (xpo *XProdOption) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpo._exists {
		return nil
	}

	// if deleted, bail
	if xpo._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_option WHERE id = ?`

	// run query
	XOLog(sqlstr, xpo.ID)
	_, err = db.Exec(sqlstr, xpo.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpo._deleted = true

	return nil
}

// XProdOptionsByCreatedBy retrieves a row from 'x_showroom.x_prod_option' as a XProdOption.
//
// Generated from index 'x_prod_option_created_by_foreign'.
func XProdOptionsByCreatedBy(db XODB, createdBy uint) ([]*XProdOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_option ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdOption{}
	for q.Next() {
		xpo := XProdOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpo.ID, &xpo.ProdID, &xpo.OptionID, &xpo.SeqNum, &xpo.CreatedBy, &xpo.UpdatedBy, &xpo.CreatedAt, &xpo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpo)
	}

	return res, nil
}

// XProdOptionByID retrieves a row from 'x_showroom.x_prod_option' as a XProdOption.
//
// Generated from index 'x_prod_option_id_pkey'.
func XProdOptionByID(db XODB, id uint) (*XProdOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_option ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpo := XProdOption{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpo.ID, &xpo.ProdID, &xpo.OptionID, &xpo.SeqNum, &xpo.CreatedBy, &xpo.UpdatedBy, &xpo.CreatedAt, &xpo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpo, nil
}

// XProdOptionsByOptionID retrieves a row from 'x_showroom.x_prod_option' as a XProdOption.
//
// Generated from index 'x_prod_option_option_id_foreign'.
func XProdOptionsByOptionID(db XODB, optionID uint) ([]*XProdOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_option ` +
		`WHERE option_id = ?`

	// run query
	XOLog(sqlstr, optionID)
	q, err := db.Query(sqlstr, optionID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdOption{}
	for q.Next() {
		xpo := XProdOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpo.ID, &xpo.ProdID, &xpo.OptionID, &xpo.SeqNum, &xpo.CreatedBy, &xpo.UpdatedBy, &xpo.CreatedAt, &xpo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpo)
	}

	return res, nil
}

// XProdOptionByProdIDOptionID retrieves a row from 'x_showroom.x_prod_option' as a XProdOption.
//
// Generated from index 'x_prod_option_prod_id_option_id_unique'.
func XProdOptionByProdIDOptionID(db XODB, prodID uint, optionID uint) (*XProdOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_option ` +
		`WHERE prod_id = ? AND option_id = ?`

	// run query
	XOLog(sqlstr, prodID, optionID)
	xpo := XProdOption{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, prodID, optionID).Scan(&xpo.ID, &xpo.ProdID, &xpo.OptionID, &xpo.SeqNum, &xpo.CreatedBy, &xpo.UpdatedBy, &xpo.CreatedAt, &xpo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpo, nil
}

// XProdOptionsByUpdatedBy retrieves a row from 'x_showroom.x_prod_option' as a XProdOption.
//
// Generated from index 'x_prod_option_updated_by_foreign'.
func XProdOptionsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, option_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_option ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdOption{}
	for q.Next() {
		xpo := XProdOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpo.ID, &xpo.ProdID, &xpo.OptionID, &xpo.SeqNum, &xpo.CreatedBy, &xpo.UpdatedBy, &xpo.CreatedAt, &xpo.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpo)
	}

	return res, nil
}
