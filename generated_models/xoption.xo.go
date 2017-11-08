// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XOption represents a row from 'x_showroom.x_option'.
type XOption struct {
	ID                 uint           `json:"id"`                    // id
	Name               string         `json:"name"`                  // name
	DisplayName        string         `json:"display_name"`          // display_name
	TypeCd             string         `json:"type_cd"`               // type_cd
	AttID              sql.NullInt64  `json:"att_id"`                // att_id
	TaxSlabID          sql.NullInt64  `json:"tax_slab_id"`           // tax_slab_id
	OrgID              uint           `json:"org_id"`                // org_id
	CreatedBy          uint           `json:"created_by"`            // created_by
	UpdatedBy          uint           `json:"updated_by"`            // updated_by
	CreatedAt          mysql.NullTime `json:"created_at"`            // created_at
	UpdatedAt          mysql.NullTime `json:"updated_at"`            // updated_at
	PricingCompGroupID sql.NullInt64  `json:"pricing_comp_group_id"` // pricing_comp_group_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XOption exists in the database.
func (xo *XOption) Exists() bool {
	return xo._exists
}

// Deleted provides information if the XOption has been deleted from the database.
func (xo *XOption) Deleted() bool {
	return xo._deleted
}

// Insert inserts the XOption to the database.
func (xo *XOption) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xo._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_option (` +
		`name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xo.Name, xo.DisplayName, xo.TypeCd, xo.AttID, xo.TaxSlabID, xo.OrgID, xo.CreatedBy, xo.UpdatedBy, xo.CreatedAt, xo.UpdatedAt, xo.PricingCompGroupID)
	res, err := db.Exec(sqlstr, xo.Name, xo.DisplayName, xo.TypeCd, xo.AttID, xo.TaxSlabID, xo.OrgID, xo.CreatedBy, xo.UpdatedBy, xo.CreatedAt, xo.UpdatedAt, xo.PricingCompGroupID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xo.ID = uint(id)
	xo._exists = true

	return nil
}

// Update updates the XOption in the database.
func (xo *XOption) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xo._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xo._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_option SET ` +
		`name = ?, display_name = ?, type_cd = ?, att_id = ?, tax_slab_id = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?, pricing_comp_group_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xo.Name, xo.DisplayName, xo.TypeCd, xo.AttID, xo.TaxSlabID, xo.OrgID, xo.CreatedBy, xo.UpdatedBy, xo.CreatedAt, xo.UpdatedAt, xo.PricingCompGroupID, xo.ID)
	_, err = db.Exec(sqlstr, xo.Name, xo.DisplayName, xo.TypeCd, xo.AttID, xo.TaxSlabID, xo.OrgID, xo.CreatedBy, xo.UpdatedBy, xo.CreatedAt, xo.UpdatedAt, xo.PricingCompGroupID, xo.ID)
	return err
}

// Save saves the XOption to the database.
func (xo *XOption) Save(db XODB) error {
	if xo.Exists() {
		return xo.Update(db)
	}

	return xo.Insert(db)
}

// Delete deletes the XOption from the database.
func (xo *XOption) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xo._exists {
		return nil
	}

	// if deleted, bail
	if xo._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_option WHERE id = ?`

	// run query
	XOLog(sqlstr, xo.ID)
	_, err = db.Exec(sqlstr, xo.ID)
	if err != nil {
		return err
	}

	// set deleted
	xo._deleted = true

	return nil
}

// XOptionsByAttID retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_att_id_foreign'.
func XOptionsByAttID(db XODB, attID sql.NullInt64) ([]*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE att_id = ?`

	// run query
	XOLog(sqlstr, attID)
	q, err := db.Query(sqlstr, attID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOption{}
	for q.Next() {
		xo := XOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xo)
	}

	return res, nil
}

// XOptionsByCreatedBy retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_created_by_foreign'.
func XOptionsByCreatedBy(db XODB, createdBy uint) ([]*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOption{}
	for q.Next() {
		xo := XOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xo)
	}

	return res, nil
}

// XOptionByID retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_id_pkey'.
func XOptionByID(db XODB, id uint) (*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xo := XOption{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
	if err != nil {
		return nil, err
	}

	return &xo, nil
}

// XOptionByNameOrgID retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_name_org_id_unique'.
func XOptionByNameOrgID(db XODB, name string, orgID uint) (*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE name = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, name, orgID)
	xo := XOption{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, orgID).Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
	if err != nil {
		return nil, err
	}

	return &xo, nil
}

// XOptionsByOrgID retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_org_id_foreign'.
func XOptionsByOrgID(db XODB, orgID uint) ([]*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOption{}
	for q.Next() {
		xo := XOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xo)
	}

	return res, nil
}

// XOptionsByTaxSlabID retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_tax_slab_id_foreign'.
func XOptionsByTaxSlabID(db XODB, taxSlabID sql.NullInt64) ([]*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE tax_slab_id = ?`

	// run query
	XOLog(sqlstr, taxSlabID)
	q, err := db.Query(sqlstr, taxSlabID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOption{}
	for q.Next() {
		xo := XOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xo)
	}

	return res, nil
}

// XOptionsByUpdatedBy retrieves a row from 'x_showroom.x_option' as a XOption.
//
// Generated from index 'x_option_updated_by_foreign'.
func XOptionsByUpdatedBy(db XODB, updatedBy uint) ([]*XOption, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, att_id, tax_slab_id, org_id, created_by, updated_by, created_at, updated_at, pricing_comp_group_id ` +
		`FROM x_showroom.x_option ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XOption{}
	for q.Next() {
		xo := XOption{
			_exists: true,
		}

		// scan
		err = q.Scan(&xo.ID, &xo.Name, &xo.DisplayName, &xo.TypeCd, &xo.AttID, &xo.TaxSlabID, &xo.OrgID, &xo.CreatedBy, &xo.UpdatedBy, &xo.CreatedAt, &xo.UpdatedAt, &xo.PricingCompGroupID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xo)
	}

	return res, nil
}