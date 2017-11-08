// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XSalesStage represents a row from 'x_showroom.x_sales_stage'.
type XSalesStage struct {
	ID                  uint           `json:"id"`                     // id
	Name                string         `json:"name"`                   // name
	StatusCd            string         `json:"status_cd"`              // status_cd
	PipelineStatusCd    string         `json:"pipeline_status_cd"`     // pipeline_status_cd
	ReportingStatusCd   string         `json:"reporting_status_cd"`    // reporting_status_cd
	DescTxt             string         `json:"desc_txt"`               // desc_txt
	SeqNum              uint           `json:"seq_num"`                // seq_num
	ActiveFlag          string         `json:"active_flag"`            // active_flag
	SalesMethodID       uint           `json:"sales_method_id"`        // sales_method_id
	PipelineSaleStageID sql.NullInt64  `json:"pipeline_sale_stage_id"` // pipeline_sale_stage_id
	CreatedBy           uint           `json:"created_by"`             // created_by
	UpdatedBy           uint           `json:"updated_by"`             // updated_by
	CreatedAt           mysql.NullTime `json:"created_at"`             // created_at
	UpdatedAt           mysql.NullTime `json:"updated_at"`             // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XSalesStage exists in the database.
func (xss *XSalesStage) Exists() bool {
	return xss._exists
}

// Deleted provides information if the XSalesStage has been deleted from the database.
func (xss *XSalesStage) Deleted() bool {
	return xss._deleted
}

// Insert inserts the XSalesStage to the database.
func (xss *XSalesStage) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xss._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_sales_stage (` +
		`name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xss.Name, xss.StatusCd, xss.PipelineStatusCd, xss.ReportingStatusCd, xss.DescTxt, xss.SeqNum, xss.ActiveFlag, xss.SalesMethodID, xss.PipelineSaleStageID, xss.CreatedBy, xss.UpdatedBy, xss.CreatedAt, xss.UpdatedAt)
	res, err := db.Exec(sqlstr, xss.Name, xss.StatusCd, xss.PipelineStatusCd, xss.ReportingStatusCd, xss.DescTxt, xss.SeqNum, xss.ActiveFlag, xss.SalesMethodID, xss.PipelineSaleStageID, xss.CreatedBy, xss.UpdatedBy, xss.CreatedAt, xss.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xss.ID = uint(id)
	xss._exists = true

	return nil
}

// Update updates the XSalesStage in the database.
func (xss *XSalesStage) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xss._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xss._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_sales_stage SET ` +
		`name = ?, status_cd = ?, pipeline_status_cd = ?, reporting_status_cd = ?, desc_txt = ?, seq_num = ?, active_flag = ?, sales_method_id = ?, pipeline_sale_stage_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xss.Name, xss.StatusCd, xss.PipelineStatusCd, xss.ReportingStatusCd, xss.DescTxt, xss.SeqNum, xss.ActiveFlag, xss.SalesMethodID, xss.PipelineSaleStageID, xss.CreatedBy, xss.UpdatedBy, xss.CreatedAt, xss.UpdatedAt, xss.ID)
	_, err = db.Exec(sqlstr, xss.Name, xss.StatusCd, xss.PipelineStatusCd, xss.ReportingStatusCd, xss.DescTxt, xss.SeqNum, xss.ActiveFlag, xss.SalesMethodID, xss.PipelineSaleStageID, xss.CreatedBy, xss.UpdatedBy, xss.CreatedAt, xss.UpdatedAt, xss.ID)
	return err
}

// Save saves the XSalesStage to the database.
func (xss *XSalesStage) Save(db XODB) error {
	if xss.Exists() {
		return xss.Update(db)
	}

	return xss.Insert(db)
}

// Delete deletes the XSalesStage from the database.
func (xss *XSalesStage) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xss._exists {
		return nil
	}

	// if deleted, bail
	if xss._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_sales_stage WHERE id = ?`

	// run query
	XOLog(sqlstr, xss.ID)
	_, err = db.Exec(sqlstr, xss.ID)
	if err != nil {
		return err
	}

	// set deleted
	xss._deleted = true

	return nil
}

// XSalesStagesByCreatedBy retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_created_by_foreign'.
func XSalesStagesByCreatedBy(db XODB, createdBy uint) ([]*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesStage{}
	for q.Next() {
		xss := XSalesStage{
			_exists: true,
		}

		// scan
		err = q.Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xss)
	}

	return res, nil
}

// XSalesStageByID retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_id_pkey'.
func XSalesStageByID(db XODB, id uint) (*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xss := XSalesStage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xss, nil
}

// XSalesStageByNameSalesMethodID retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_name_sales_method_id_unique'.
func XSalesStageByNameSalesMethodID(db XODB, name string, salesMethodID uint) (*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE name = ? AND sales_method_id = ?`

	// run query
	XOLog(sqlstr, name, salesMethodID)
	xss := XSalesStage{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, salesMethodID).Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xss, nil
}

// XSalesStagesByPipelineSaleStageID retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_pipeline_sale_stage_id_foreign'.
func XSalesStagesByPipelineSaleStageID(db XODB, pipelineSaleStageID sql.NullInt64) ([]*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE pipeline_sale_stage_id = ?`

	// run query
	XOLog(sqlstr, pipelineSaleStageID)
	q, err := db.Query(sqlstr, pipelineSaleStageID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesStage{}
	for q.Next() {
		xss := XSalesStage{
			_exists: true,
		}

		// scan
		err = q.Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xss)
	}

	return res, nil
}

// XSalesStagesBySalesMethodID retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_sales_method_id_foreign'.
func XSalesStagesBySalesMethodID(db XODB, salesMethodID uint) ([]*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE sales_method_id = ?`

	// run query
	XOLog(sqlstr, salesMethodID)
	q, err := db.Query(sqlstr, salesMethodID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesStage{}
	for q.Next() {
		xss := XSalesStage{
			_exists: true,
		}

		// scan
		err = q.Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xss)
	}

	return res, nil
}

// XSalesStagesByUpdatedBy retrieves a row from 'x_showroom.x_sales_stage' as a XSalesStage.
//
// Generated from index 'x_sales_stage_updated_by_foreign'.
func XSalesStagesByUpdatedBy(db XODB, updatedBy uint) ([]*XSalesStage, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, status_cd, pipeline_status_cd, reporting_status_cd, desc_txt, seq_num, active_flag, sales_method_id, pipeline_sale_stage_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_sales_stage ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XSalesStage{}
	for q.Next() {
		xss := XSalesStage{
			_exists: true,
		}

		// scan
		err = q.Scan(&xss.ID, &xss.Name, &xss.StatusCd, &xss.PipelineStatusCd, &xss.ReportingStatusCd, &xss.DescTxt, &xss.SeqNum, &xss.ActiveFlag, &xss.SalesMethodID, &xss.PipelineSaleStageID, &xss.CreatedBy, &xss.UpdatedBy, &xss.CreatedAt, &xss.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xss)
	}

	return res, nil
}