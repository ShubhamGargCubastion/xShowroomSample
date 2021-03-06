// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdFeatureGrp represents a row from 'x_showroom.x_prod_feature_grp'.
type XProdFeatureGrp struct {
	ID          uint           `json:"id"`           // id
	Name        string         `json:"name"`         // name
	DisplayName string         `json:"display_name"` // display_name
	SeqNum      uint           `json:"seq_num"`      // seq_num
	ParID       uint           `json:"par_id"`       // par_id
	CreatedBy   uint           `json:"created_by"`   // created_by
	UpdatedBy   uint           `json:"updated_by"`   // updated_by
	CreatedAt   mysql.NullTime `json:"created_at"`   // created_at
	UpdatedAt   mysql.NullTime `json:"updated_at"`   // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdFeatureGrp exists in the database.
func (xpfg *XProdFeatureGrp) Exists() bool {
	return xpfg._exists
}

// Deleted provides information if the XProdFeatureGrp has been deleted from the database.
func (xpfg *XProdFeatureGrp) Deleted() bool {
	return xpfg._deleted
}

// Insert inserts the XProdFeatureGrp to the database.
func (xpfg *XProdFeatureGrp) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpfg._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_feature_grp (` +
		`name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpfg.Name, xpfg.DisplayName, xpfg.SeqNum, xpfg.ParID, xpfg.CreatedBy, xpfg.UpdatedBy, xpfg.CreatedAt, xpfg.UpdatedAt)
	res, err := db.Exec(sqlstr, xpfg.Name, xpfg.DisplayName, xpfg.SeqNum, xpfg.ParID, xpfg.CreatedBy, xpfg.UpdatedBy, xpfg.CreatedAt, xpfg.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpfg.ID = uint(id)
	xpfg._exists = true

	return nil
}

// Update updates the XProdFeatureGrp in the database.
func (xpfg *XProdFeatureGrp) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpfg._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpfg._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_feature_grp SET ` +
		`name = ?, display_name = ?, seq_num = ?, par_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpfg.Name, xpfg.DisplayName, xpfg.SeqNum, xpfg.ParID, xpfg.CreatedBy, xpfg.UpdatedBy, xpfg.CreatedAt, xpfg.UpdatedAt, xpfg.ID)
	_, err = db.Exec(sqlstr, xpfg.Name, xpfg.DisplayName, xpfg.SeqNum, xpfg.ParID, xpfg.CreatedBy, xpfg.UpdatedBy, xpfg.CreatedAt, xpfg.UpdatedAt, xpfg.ID)
	return err
}

// Save saves the XProdFeatureGrp to the database.
func (xpfg *XProdFeatureGrp) Save(db XODB) error {
	if xpfg.Exists() {
		return xpfg.Update(db)
	}

	return xpfg.Insert(db)
}

// Delete deletes the XProdFeatureGrp from the database.
func (xpfg *XProdFeatureGrp) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpfg._exists {
		return nil
	}

	// if deleted, bail
	if xpfg._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_feature_grp WHERE id = ?`

	// run query
	XOLog(sqlstr, xpfg.ID)
	_, err = db.Exec(sqlstr, xpfg.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpfg._deleted = true

	return nil
}

// XProdFeatureGrpsByCreatedBy retrieves a row from 'x_showroom.x_prod_feature_grp' as a XProdFeatureGrp.
//
// Generated from index 'x_prod_feature_grp_created_by_foreign'.
func XProdFeatureGrpsByCreatedBy(db XODB, createdBy uint) ([]*XProdFeatureGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_grp ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureGrp{}
	for q.Next() {
		xpfg := XProdFeatureGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfg.ID, &xpfg.Name, &xpfg.DisplayName, &xpfg.SeqNum, &xpfg.ParID, &xpfg.CreatedBy, &xpfg.UpdatedBy, &xpfg.CreatedAt, &xpfg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfg)
	}

	return res, nil
}

// XProdFeatureGrpByID retrieves a row from 'x_showroom.x_prod_feature_grp' as a XProdFeatureGrp.
//
// Generated from index 'x_prod_feature_grp_id_pkey'.
func XProdFeatureGrpByID(db XODB, id uint) (*XProdFeatureGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_grp ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpfg := XProdFeatureGrp{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpfg.ID, &xpfg.Name, &xpfg.DisplayName, &xpfg.SeqNum, &xpfg.ParID, &xpfg.CreatedBy, &xpfg.UpdatedBy, &xpfg.CreatedAt, &xpfg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpfg, nil
}

// XProdFeatureGrpsByParID retrieves a row from 'x_showroom.x_prod_feature_grp' as a XProdFeatureGrp.
//
// Generated from index 'x_prod_feature_grp_par_id_foreign'.
func XProdFeatureGrpsByParID(db XODB, parID uint) ([]*XProdFeatureGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_grp ` +
		`WHERE par_id = ?`

	// run query
	XOLog(sqlstr, parID)
	q, err := db.Query(sqlstr, parID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureGrp{}
	for q.Next() {
		xpfg := XProdFeatureGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfg.ID, &xpfg.Name, &xpfg.DisplayName, &xpfg.SeqNum, &xpfg.ParID, &xpfg.CreatedBy, &xpfg.UpdatedBy, &xpfg.CreatedAt, &xpfg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfg)
	}

	return res, nil
}

// XProdFeatureGrpsByUpdatedBy retrieves a row from 'x_showroom.x_prod_feature_grp' as a XProdFeatureGrp.
//
// Generated from index 'x_prod_feature_grp_updated_by_foreign'.
func XProdFeatureGrpsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdFeatureGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_grp ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdFeatureGrp{}
	for q.Next() {
		xpfg := XProdFeatureGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpfg.ID, &xpfg.Name, &xpfg.DisplayName, &xpfg.SeqNum, &xpfg.ParID, &xpfg.CreatedBy, &xpfg.UpdatedBy, &xpfg.CreatedAt, &xpfg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpfg)
	}

	return res, nil
}

// XProdFeatureGrpByNameParID retrieves a row from 'x_showroom.x_prod_feature_grp' as a XProdFeatureGrp.
//
// Generated from index 'x_unique_key_1'.
func XProdFeatureGrpByNameParID(db XODB, name string, parID uint) (*XProdFeatureGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, seq_num, par_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_feature_grp ` +
		`WHERE name = ? AND par_id = ?`

	// run query
	XOLog(sqlstr, name, parID)
	xpfg := XProdFeatureGrp{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, parID).Scan(&xpfg.ID, &xpfg.Name, &xpfg.DisplayName, &xpfg.SeqNum, &xpfg.ParID, &xpfg.CreatedBy, &xpfg.UpdatedBy, &xpfg.CreatedAt, &xpfg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpfg, nil
}
