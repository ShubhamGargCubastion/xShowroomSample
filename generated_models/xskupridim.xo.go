// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XSkuPriDim represents a row from 'x_showroom.x_sku_pri_dim'.
type XSkuPriDim struct {
	ID                    uint           `json:"id"`                      // id
	SkuID                 uint           `json:"sku_id"`                  // sku_id
	PricingDimensionName  string         `json:"pricing_dimension_name"`  // pricing_dimension_name
	PricingDimensionValue string         `json:"pricing_dimension_value"` // pricing_dimension_value
	CreatedBy             uint           `json:"created_by"`              // created_by
	UpdatedBy             uint           `json:"updated_by"`              // updated_by
	CreatedAt             mysql.NullTime `json:"created_at"`              // created_at
	UpdatedAt             mysql.NullTime `json:"updated_at"`              // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XSkuPriDim exists in the database.
func (xspd *XSkuPriDim) Exists() bool {
	return xspd._exists
}

// Deleted provides information if the XSkuPriDim has been deleted from the database.
func (xspd *XSkuPriDim) Deleted() bool {
	return xspd._deleted
}

// Insert inserts the XSkuPriDim to the database.
func (xspd *XSkuPriDim) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xspd._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_sku_pri_dim (` +
		`sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xspd.SkuID, xspd.PricingDimensionName, xspd.PricingDimensionValue, xspd.CreatedBy, xspd.UpdatedBy, xspd.CreatedAt, xspd.UpdatedAt)
	res, err := db.Exec(sqlstr, xspd.SkuID, xspd.PricingDimensionName, xspd.PricingDimensionValue, xspd.CreatedBy, xspd.UpdatedBy, xspd.CreatedAt, xspd.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xspd.ID = uint(id)
	xspd._exists = true

	return nil
}

// Update updates the XSkuPriDim in the database.
func (xspd *XSkuPriDim) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xspd._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xspd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_sku_pri_dim SET ` +
		`sku_id = ?, pricing_dimension_name = ?, pricing_dimension_value = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xspd.SkuID, xspd.PricingDimensionName, xspd.PricingDimensionValue, xspd.CreatedBy, xspd.UpdatedBy, xspd.CreatedAt, xspd.UpdatedAt, xspd.ID)
	_, err = db.Exec(sqlstr, xspd.SkuID, xspd.PricingDimensionName, xspd.PricingDimensionValue, xspd.CreatedBy, xspd.UpdatedBy, xspd.CreatedAt, xspd.UpdatedAt, xspd.ID)
	return err
}

// Save saves the XSkuPriDim to the database.
func (xspd *XSkuPriDim) Save(db XODB) error {
	if xspd.Exists() {
		return xspd.Update(db)
	}

	return xspd.Insert(db)
}

// Delete deletes the XSkuPriDim from the database.
func (xspd *XSkuPriDim) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xspd._exists {
		return nil
	}

	// if deleted, bail
	if xspd._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_sku_pri_dim WHERE id = ?`

	// run query
	XOLog(sqlstr, xspd.ID)
	_, err = db.Exec(sqlstr, xspd.ID)
	if err != nil {
		return err
	}

	// set deleted
	xspd._deleted = true

	return nil
}

// XSkuPriDimsByCreatedBy retrieves a row from 'x_showroom.x_sku_pri_dim' as a XSkuPriDim.
//
// Generated from index 'x_sku_pri_dim_created_by_foreign'.
func XSkuPriDimsByCreatedBy(db XODB, createdBy uint) ([]*XSkuPriDim, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sku_pri_dim ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSkuPriDim{}
	for q.Next() {
		xspd := XSkuPriDim{
			_exists: true,
		}

		// scan
		err = q.Scan(&xspd.ID, &xspd.SkuID, &xspd.PricingDimensionName, &xspd.PricingDimensionValue, &xspd.CreatedBy, &xspd.UpdatedBy, &xspd.CreatedAt, &xspd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xspd)
	}

	return res, nil
}

// XSkuPriDimByID retrieves a row from 'x_showroom.x_sku_pri_dim' as a XSkuPriDim.
//
// Generated from index 'x_sku_pri_dim_id_pkey'.
func XSkuPriDimByID(db XODB, id uint) (*XSkuPriDim, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sku_pri_dim ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xspd := XSkuPriDim{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xspd.ID, &xspd.SkuID, &xspd.PricingDimensionName, &xspd.PricingDimensionValue, &xspd.CreatedBy, &xspd.UpdatedBy, &xspd.CreatedAt, &xspd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xspd, nil
}

// XSkuPriDimByPricingDimensionValueSkuID retrieves a row from 'x_showroom.x_sku_pri_dim' as a XSkuPriDim.
//
// Generated from index 'x_sku_pri_dim_pricing_dimension_value_sku_id_unique'.
func XSkuPriDimByPricingDimensionValueSkuID(db XODB, pricingDimensionValue string, skuID uint) (*XSkuPriDim, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sku_pri_dim ` +
		`WHERE pricing_dimension_value = ? AND sku_id = ?`

	// run query
	XOLog(sqlstr, pricingDimensionValue, skuID)
	xspd := XSkuPriDim{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, pricingDimensionValue, skuID).Scan(&xspd.ID, &xspd.SkuID, &xspd.PricingDimensionName, &xspd.PricingDimensionValue, &xspd.CreatedBy, &xspd.UpdatedBy, &xspd.CreatedAt, &xspd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xspd, nil
}

// XSkuPriDimsBySkuID retrieves a row from 'x_showroom.x_sku_pri_dim' as a XSkuPriDim.
//
// Generated from index 'x_sku_pri_dim_sku_id_foreign'.
func XSkuPriDimsBySkuID(db XODB, skuID uint) ([]*XSkuPriDim, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sku_pri_dim ` +
		`WHERE sku_id = ?`

	// run query
	XOLog(sqlstr, skuID)
	q, err := db.Query(sqlstr, skuID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSkuPriDim{}
	for q.Next() {
		xspd := XSkuPriDim{
			_exists: true,
		}

		// scan
		err = q.Scan(&xspd.ID, &xspd.SkuID, &xspd.PricingDimensionName, &xspd.PricingDimensionValue, &xspd.CreatedBy, &xspd.UpdatedBy, &xspd.CreatedAt, &xspd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xspd)
	}

	return res, nil
}

// XSkuPriDimsByUpdatedBy retrieves a row from 'x_showroom.x_sku_pri_dim' as a XSkuPriDim.
//
// Generated from index 'x_sku_pri_dim_updated_by_foreign'.
func XSkuPriDimsByUpdatedBy(db XODB, updatedBy uint) ([]*XSkuPriDim, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, sku_id, pricing_dimension_name, pricing_dimension_value, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sku_pri_dim ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSkuPriDim{}
	for q.Next() {
		xspd := XSkuPriDim{
			_exists: true,
		}

		// scan
		err = q.Scan(&xspd.ID, &xspd.SkuID, &xspd.PricingDimensionName, &xspd.PricingDimensionValue, &xspd.CreatedBy, &xspd.UpdatedBy, &xspd.CreatedAt, &xspd.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xspd)
	}

	return res, nil
}