// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XProdDemoGroup represents a row from 'x_showroom.x_prod_demo_group'.
type XProdDemoGroup struct {
	ID             uint           `json:"id"`                // id
	ActionCode     sql.NullString `json:"action_code"`       // action_code
	Name           sql.NullString `json:"name"`              // name
	DisplayName    sql.NullString `json:"display_name"`      // display_name
	TypeCd         sql.NullString `json:"type_cd"`           // type_cd
	ParID          sql.NullInt64  `json:"par_id"`            // par_id
	AttID          sql.NullInt64  `json:"att_id"`            // att_id
	ProdID         uint           `json:"prod_id"`           // prod_id
	DemoGroupDefID sql.NullInt64  `json:"demo_group_def_id"` // demo_group_def_id
	SeqNum         sql.NullInt64  `json:"seq_num"`           // seq_num
	DemoID         sql.NullInt64  `json:"demo_id"`           // demo_id
	DemoGroupID    sql.NullInt64  `json:"demo_group_id"`     // demo_group_id
	CreatedBy      uint           `json:"created_by"`        // created_by
	UpdatedBy      uint           `json:"updated_by"`        // updated_by
	CreatedAt      mysql.NullTime `json:"created_at"`        // created_at
	UpdatedAt      mysql.NullTime `json:"updated_at"`        // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XProdDemoGroup exists in the database.
func (xpdg *XProdDemoGroup) Exists() bool {
	return xpdg._exists
}

// Deleted provides information if the XProdDemoGroup has been deleted from the database.
func (xpdg *XProdDemoGroup) Deleted() bool {
	return xpdg._deleted
}

// Insert inserts the XProdDemoGroup to the database.
func (xpdg *XProdDemoGroup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xpdg._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_prod_demo_group (` +
		`action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xpdg.ActionCode, xpdg.Name, xpdg.DisplayName, xpdg.TypeCd, xpdg.ParID, xpdg.AttID, xpdg.ProdID, xpdg.DemoGroupDefID, xpdg.SeqNum, xpdg.DemoID, xpdg.DemoGroupID, xpdg.CreatedBy, xpdg.UpdatedBy, xpdg.CreatedAt, xpdg.UpdatedAt)
	res, err := db.Exec(sqlstr, xpdg.ActionCode, xpdg.Name, xpdg.DisplayName, xpdg.TypeCd, xpdg.ParID, xpdg.AttID, xpdg.ProdID, xpdg.DemoGroupDefID, xpdg.SeqNum, xpdg.DemoID, xpdg.DemoGroupID, xpdg.CreatedBy, xpdg.UpdatedBy, xpdg.CreatedAt, xpdg.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xpdg.ID = uint(id)
	xpdg._exists = true

	return nil
}

// Update updates the XProdDemoGroup in the database.
func (xpdg *XProdDemoGroup) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpdg._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xpdg._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_prod_demo_group SET ` +
		`action_code = ?, name = ?, display_name = ?, type_cd = ?, par_id = ?, att_id = ?, prod_id = ?, demo_group_def_id = ?, seq_num = ?, demo_id = ?, demo_group_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xpdg.ActionCode, xpdg.Name, xpdg.DisplayName, xpdg.TypeCd, xpdg.ParID, xpdg.AttID, xpdg.ProdID, xpdg.DemoGroupDefID, xpdg.SeqNum, xpdg.DemoID, xpdg.DemoGroupID, xpdg.CreatedBy, xpdg.UpdatedBy, xpdg.CreatedAt, xpdg.UpdatedAt, xpdg.ID)
	_, err = db.Exec(sqlstr, xpdg.ActionCode, xpdg.Name, xpdg.DisplayName, xpdg.TypeCd, xpdg.ParID, xpdg.AttID, xpdg.ProdID, xpdg.DemoGroupDefID, xpdg.SeqNum, xpdg.DemoID, xpdg.DemoGroupID, xpdg.CreatedBy, xpdg.UpdatedBy, xpdg.CreatedAt, xpdg.UpdatedAt, xpdg.ID)
	return err
}

// Save saves the XProdDemoGroup to the database.
func (xpdg *XProdDemoGroup) Save(db XODB) error {
	if xpdg.Exists() {
		return xpdg.Update(db)
	}

	return xpdg.Insert(db)
}

// Delete deletes the XProdDemoGroup from the database.
func (xpdg *XProdDemoGroup) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xpdg._exists {
		return nil
	}

	// if deleted, bail
	if xpdg._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_prod_demo_group WHERE id = ?`

	// run query
	XOLog(sqlstr, xpdg.ID)
	_, err = db.Exec(sqlstr, xpdg.ID)
	if err != nil {
		return err
	}

	// set deleted
	xpdg._deleted = true

	return nil
}

// XProdDemoGroupsByAttID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_att_id_foreign'.
func XProdDemoGroupsByAttID(db XODB, attID sql.NullInt64) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE att_id = ?`

	// run query
	XOLog(sqlstr, attID)
	q, err := db.Query(sqlstr, attID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}

// XProdDemoGroupsByCreatedBy retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_created_by_foreign'.
func XProdDemoGroupsByCreatedBy(db XODB, createdBy uint) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}

// XProdDemoGroupsByDemoGroupDefID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_demo_group_def_id_foreign'.
func XProdDemoGroupsByDemoGroupDefID(db XODB, demoGroupDefID sql.NullInt64) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE demo_group_def_id = ?`

	// run query
	XOLog(sqlstr, demoGroupDefID)
	q, err := db.Query(sqlstr, demoGroupDefID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}

// XProdDemoGroupByID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_id_pkey'.
func XProdDemoGroupByID(db XODB, id uint) (*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xpdg := XProdDemoGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpdg, nil
}

// XProdDemoGroupByNameTypeCdDemoGroupDefIDProdID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_name_type_cd_demo_group_def_id_prod_id_unique'.
func XProdDemoGroupByNameTypeCdDemoGroupDefIDProdID(db XODB, name sql.NullString, typeCd sql.NullString, demoGroupDefID sql.NullInt64, prodID uint) (*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE name = ? AND type_cd = ? AND demo_group_def_id = ? AND prod_id = ?`

	// run query
	XOLog(sqlstr, name, typeCd, demoGroupDefID, prodID)
	xpdg := XProdDemoGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, typeCd, demoGroupDefID, prodID).Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xpdg, nil
}

// XProdDemoGroupsByParID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_par_id_foreign'.
func XProdDemoGroupsByParID(db XODB, parID sql.NullInt64) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE par_id = ?`

	// run query
	XOLog(sqlstr, parID)
	q, err := db.Query(sqlstr, parID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}

// XProdDemoGroupsByProdID retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_prod_id_foreign'.
func XProdDemoGroupsByProdID(db XODB, prodID uint) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE prod_id = ?`

	// run query
	XOLog(sqlstr, prodID)
	q, err := db.Query(sqlstr, prodID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}

// XProdDemoGroupsByUpdatedBy retrieves a row from 'x_showroom.x_prod_demo_group' as a XProdDemoGroup.
//
// Generated from index 'x_prod_demo_group_updated_by_foreign'.
func XProdDemoGroupsByUpdatedBy(db XODB, updatedBy uint) ([]*XProdDemoGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_code, name, display_name, type_cd, par_id, att_id, prod_id, demo_group_def_id, seq_num, demo_id, demo_group_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_prod_demo_group ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XProdDemoGroup{}
	for q.Next() {
		xpdg := XProdDemoGroup{
			_exists: true,
		}

		// scan
		err = q.Scan(&xpdg.ID, &xpdg.ActionCode, &xpdg.Name, &xpdg.DisplayName, &xpdg.TypeCd, &xpdg.ParID, &xpdg.AttID, &xpdg.ProdID, &xpdg.DemoGroupDefID, &xpdg.SeqNum, &xpdg.DemoID, &xpdg.DemoGroupID, &xpdg.CreatedBy, &xpdg.UpdatedBy, &xpdg.CreatedAt, &xpdg.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xpdg)
	}

	return res, nil
}