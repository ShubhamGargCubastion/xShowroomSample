// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdFeatureItem represents a row from 'x_showroom.x_prod_feature_item'.
type XProdFeatureItem struct {
	ID        uint           `json:"id"`         // id
	Name      string         `json:"name"`       // name
	Value     string         `json:"value"`      // value
	SeqNum    uint           `json:"seq_num"`    // seq_num
	ParID     uint           `json:"par_id"`     // par_id
	CreatedBy uint           `json:"created_by"` // created_by
	UpdatedBy uint           `json:"updated_by"` // updated_by
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdFeatureItem exists in the database.
func (xpfi *XProdFeatureItem) Exists() bool {
	return xpfi._exists
}

// Deleted provides information if the XProdFeatureItem has been deleted from the database.
func (xpfi *XProdFeatureItem) Deleted() bool {
	return xpfi._deleted
}

// Insert inserts the XProdFeatureItem to the database.
func (xpfi *XProdFeatureItem) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpfi._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_feature_item (` +
		`name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpfi.Name, xpfi.Value, xpfi.SeqNum, xpfi.ParID, xpfi.CreatedBy, xpfi.UpdatedBy, xpfi.CreatedAt, xpfi.UpdatedAt)
	res, err := db.Exec(sqlstr, xpfi.Name, xpfi.Value, xpfi.SeqNum, xpfi.ParID, xpfi.CreatedBy, xpfi.UpdatedBy, xpfi.CreatedAt, xpfi.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpfi.ID = uint(id)
	xpfi._exists = true

	return nil
}

// Update updates the XProdFeatureItem in the database.
func (xpfi *XProdFeatureItem) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpfi._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpfi._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_feature_item SET ` +
		`name = ?, value = ?, seq_num = ?, par_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpfi.Name, xpfi.Value, xpfi.SeqNum, xpfi.ParID, xpfi.CreatedBy, xpfi.UpdatedBy, xpfi.CreatedAt, xpfi.UpdatedAt, xpfi.ID)
	_, err = db.Exec(sqlstr, xpfi.Name, xpfi.Value, xpfi.SeqNum, xpfi.ParID, xpfi.CreatedBy, xpfi.UpdatedBy, xpfi.CreatedAt, xpfi.UpdatedAt, xpfi.ID)
	return err
}

// Save saves the XProdFeatureItem to the database.
func (xpfi *XProdFeatureItem) Save(db XODB) error {
	if xpfi.Exists() {
		return xpfi.Update(db)
	}

	return xpfi.Insert(db)
}

// Delete deletes the XProdFeatureItem from the database.
func (xpfi *XProdFeatureItem) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpfi._exists {
		return nil
	}

	// if deleted, bail
	if xpfi._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_feature_item WHERE id = ?`

	// run query
	XOLog(sqlstr, xpfi.ID)
	_, err = db.Exec(sqlstr, xpfi.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpfi._deleted = true

	return nil
}

// XProdFeatureItemsByCreatedBy retrieves a row from 'x_showroom.x_prod_feature_item' as a XProdFeatureItem.
//
// Generated from index 'x_prod_feature_item_created_by_foreign'.
func XProdFeatureItemsByCreatedBy(db XODB, createdBy uint) ([]*XProdFeatureItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_item ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureItem{}
	for q.Next() {
		xpfi := XProdFeatureItem{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfi.ID, &xpfi.Name, &xpfi.Value, &xpfi.SeqNum, &xpfi.ParID, &xpfi.CreatedBy, &xpfi.UpdatedBy, &xpfi.CreatedAt, &xpfi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfi)
	}

	return res, nil
}

// XProdFeatureItemByID retrieves a row from 'x_showroom.x_prod_feature_item' as a XProdFeatureItem.
//
// Generated from index 'x_prod_feature_item_id_pkey'.
func XProdFeatureItemByID(db XODB, id uint) (*XProdFeatureItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_item ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpfi := XProdFeatureItem{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpfi.ID, &xpfi.Name, &xpfi.Value, &xpfi.SeqNum, &xpfi.ParID, &xpfi.CreatedBy, &xpfi.UpdatedBy, &xpfi.CreatedAt, &xpfi.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpfi, nil
}

// XProdFeatureItemsByParID retrieves a row from 'x_showroom.x_prod_feature_item' as a XProdFeatureItem.
//
// Generated from index 'x_prod_feature_item_par_id_foreign'.
func XProdFeatureItemsByParID(db XODB, parID uint) ([]*XProdFeatureItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_item ` +
		`WHERE par_id = ?`

	// run query
	XOLog(sqlstr, parID)
	q, err := db.Query(sqlstr, parID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureItem{}
	for q.Next() {
		xpfi := XProdFeatureItem{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfi.ID, &xpfi.Name, &xpfi.Value, &xpfi.SeqNum, &xpfi.ParID, &xpfi.CreatedBy, &xpfi.UpdatedBy, &xpfi.CreatedAt, &xpfi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfi)
	}

	return res, nil
}

// XProdFeatureItemsByUpdatedBy retrieves a row from 'x_showroom.x_prod_feature_item' as a XProdFeatureItem.
//
// Generated from index 'x_prod_feature_item_updated_by_foreign'.
func XProdFeatureItemsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdFeatureItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_item ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureItem{}
	for q.Next() {
		xpfi := XProdFeatureItem{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfi.ID, &xpfi.Name, &xpfi.Value, &xpfi.SeqNum, &xpfi.ParID, &xpfi.CreatedBy, &xpfi.UpdatedBy, &xpfi.CreatedAt, &xpfi.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfi)
	}

	return res, nil
}

// XProdFeatureItemByNameParID retrieves a row from 'x_showroom.x_prod_feature_item' as a XProdFeatureItem.
//
// Generated from index 'x_unique_key_1'.
func XProdFeatureItemByNameParID(db XODB, name string, parID uint) (*XProdFeatureItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, value, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_item ` +
		`WHERE name = ? AND par_id = ?`

	// run query
	XOLog(sqlstr, name, parID)
	xpfi := XProdFeatureItem{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, parID).Scan(&xpfi.ID, &xpfi.Name, &xpfi.Value, &xpfi.SeqNum, &xpfi.ParID, &xpfi.CreatedBy, &xpfi.UpdatedBy, &xpfi.CreatedAt, &xpfi.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpfi, nil
}
