// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdRelProdGrp represents a row from 'x_showroom.x_prod_rel_prod_grps'.
type XProdRelProdGrp struct {
	ID                  uint           `json:"id"`                      // id
	ProdID              uint           `json:"prod_id"`                 // prod_id
	RelatedProdGrpDefID uint           `json:"related_prod_grp_def_id"` // related_prod_grp_def_id
	SeqNum              uint           `json:"seq_num"`                 // seq_num
	CreatedBy           uint           `json:"created_by"`              // created_by
	UpdatedBy           uint           `json:"updated_by"`              // updated_by
	CreatedAt           mysql.NullTime `json:"created_at"`              // created_at
	UpdatedAt           mysql.NullTime `json:"updated_at"`              // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdRelProdGrp exists in the database.
func (xprpg *XProdRelProdGrp) Exists() bool {
	return xprpg._exists
}

// Deleted provides information if the XProdRelProdGrp has been deleted from the database.
func (xprpg *XProdRelProdGrp) Deleted() bool {
	return xprpg._deleted
}

// Insert inserts the XProdRelProdGrp to the database.
func (xprpg *XProdRelProdGrp) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xprpg._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_rel_prod_grps (` +
		`prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xprpg.ProdID, xprpg.RelatedProdGrpDefID, xprpg.SeqNum, xprpg.CreatedBy, xprpg.UpdatedBy, xprpg.CreatedAt, xprpg.UpdatedAt)
	res, err := db.Exec(sqlstr, xprpg.ProdID, xprpg.RelatedProdGrpDefID, xprpg.SeqNum, xprpg.CreatedBy, xprpg.UpdatedBy, xprpg.CreatedAt, xprpg.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xprpg.ID = uint(id)
	xprpg._exists = true

	return nil
}

// Update updates the XProdRelProdGrp in the database.
func (xprpg *XProdRelProdGrp) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xprpg._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xprpg._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_rel_prod_grps SET ` +
		`prod_id = ?, related_prod_grp_def_id = ?, seq_num = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xprpg.ProdID, xprpg.RelatedProdGrpDefID, xprpg.SeqNum, xprpg.CreatedBy, xprpg.UpdatedBy, xprpg.CreatedAt, xprpg.UpdatedAt, xprpg.ID)
	_, err = db.Exec(sqlstr, xprpg.ProdID, xprpg.RelatedProdGrpDefID, xprpg.SeqNum, xprpg.CreatedBy, xprpg.UpdatedBy, xprpg.CreatedAt, xprpg.UpdatedAt, xprpg.ID)
	return err
}

// Save saves the XProdRelProdGrp to the database.
func (xprpg *XProdRelProdGrp) Save(db XODB) error {
	if xprpg.Exists() {
		return xprpg.Update(db)
	}

	return xprpg.Insert(db)
}

// Delete deletes the XProdRelProdGrp from the database.
func (xprpg *XProdRelProdGrp) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xprpg._exists {
		return nil
	}

	// if deleted, bail
	if xprpg._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_rel_prod_grps WHERE id = ?`

	// run query
	XOLog(sqlstr, xprpg.ID)
	_, err = db.Exec(sqlstr, xprpg.ID)
	if err != nil {
		return err
	}

	// set deleted
	xprpg._deleted = true

	return nil
}

// XProdRelProdGrpsByCreatedBy retrieves a row from 'x_showroom.x_prod_rel_prod_grps' as a XProdRelProdGrp.
//
// Generated from index 'x_prod_rel_prod_grps_created_by_foreign'.
func XProdRelProdGrpsByCreatedBy(db XODB, createdBy uint) ([]*XProdRelProdGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rel_prod_grps ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRelProdGrp{}
	for q.Next() {
		xprpg := XProdRelProdGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprpg.ID, &xprpg.ProdID, &xprpg.RelatedProdGrpDefID, &xprpg.SeqNum, &xprpg.CreatedBy, &xprpg.UpdatedBy, &xprpg.CreatedAt, &xprpg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprpg)
	}

	return res, nil
}

// XProdRelProdGrpByID retrieves a row from 'x_showroom.x_prod_rel_prod_grps' as a XProdRelProdGrp.
//
// Generated from index 'x_prod_rel_prod_grps_id_pkey'.
func XProdRelProdGrpByID(db XODB, id uint) (*XProdRelProdGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rel_prod_grps ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xprpg := XProdRelProdGrp{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xprpg.ID, &xprpg.ProdID, &xprpg.RelatedProdGrpDefID, &xprpg.SeqNum, &xprpg.CreatedBy, &xprpg.UpdatedBy, &xprpg.CreatedAt, &xprpg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xprpg, nil
}

// XProdRelProdGrpsByProdID retrieves a row from 'x_showroom.x_prod_rel_prod_grps' as a XProdRelProdGrp.
//
// Generated from index 'x_prod_rel_prod_grps_prod_id_foreign'.
func XProdRelProdGrpsByProdID(db XODB, prodID uint) ([]*XProdRelProdGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rel_prod_grps ` +
		`WHERE prod_id = ?`

	// run query
	XOLog(sqlstr, prodID)
	q, err := db.Query(sqlstr, prodID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRelProdGrp{}
	for q.Next() {
		xprpg := XProdRelProdGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprpg.ID, &xprpg.ProdID, &xprpg.RelatedProdGrpDefID, &xprpg.SeqNum, &xprpg.CreatedBy, &xprpg.UpdatedBy, &xprpg.CreatedAt, &xprpg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprpg)
	}

	return res, nil
}

// XProdRelProdGrpByRelatedProdGrpDefIDProdID retrieves a row from 'x_showroom.x_prod_rel_prod_grps' as a XProdRelProdGrp.
//
// Generated from index 'x_prod_rel_prod_grps_related_prod_grp_def_id_prod_id_unique'.
func XProdRelProdGrpByRelatedProdGrpDefIDProdID(db XODB, relatedProdGrpDefID uint, prodID uint) (*XProdRelProdGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rel_prod_grps ` +
		`WHERE related_prod_grp_def_id = ? AND prod_id = ?`

	// run query
	XOLog(sqlstr, relatedProdGrpDefID, prodID)
	xprpg := XProdRelProdGrp{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, relatedProdGrpDefID, prodID).Scan(&xprpg.ID, &xprpg.ProdID, &xprpg.RelatedProdGrpDefID, &xprpg.SeqNum, &xprpg.CreatedBy, &xprpg.UpdatedBy, &xprpg.CreatedAt, &xprpg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xprpg, nil
}

// XProdRelProdGrpsByUpdatedBy retrieves a row from 'x_showroom.x_prod_rel_prod_grps' as a XProdRelProdGrp.
//
// Generated from index 'x_prod_rel_prod_grps_updated_by_foreign'.
func XProdRelProdGrpsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdRelProdGrp, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, prod_id, related_prod_grp_def_id, seq_num, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_rel_prod_grps ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdRelProdGrp{}
	for q.Next() {
		xprpg := XProdRelProdGrp{
			_exists: true,
		}

		// scan
		err = q.Scan(&xprpg.ID, &xprpg.ProdID, &xprpg.RelatedProdGrpDefID, &xprpg.SeqNum, &xprpg.CreatedBy, &xprpg.UpdatedBy, &xprpg.CreatedAt, &xprpg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xprpg)
	}

	return res, nil
}
