// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XEmailTemplate represents a row from 'x_showroom.x_email_template'.
type XEmailTemplate struct {
	ID           uint           `json:"id"`            // id
	TemplateName string         `json:"template_name"` // template_name
	TemplateCode string         `json:"template_code"` // template_code
	ActiveFlag   string         `json:"active_flag"`   // active_flag
	OrgID        uint           `json:"org_id"`        // org_id
	CreatedBy    uint           `json:"created_by"`    // created_by
	UpdatedBy    uint           `json:"updated_by"`    // updated_by
	CreatedAt    mysql.NullTime `json:"created_at"`    // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`    // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XEmailTemplate exists in the database.
func (xet *XEmailTemplate) Exists() bool {
	return xet._exists
}

// Deleted provides information if the XEmailTemplate has been deleted from the database.
func (xet *XEmailTemplate) Deleted() bool {
	return xet._deleted
}

// Insert inserts the XEmailTemplate to the database.
func (xet *XEmailTemplate) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xet._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_email_template (` +
		`template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xet.TemplateName, xet.TemplateCode, xet.ActiveFlag, xet.OrgID, xet.CreatedBy, xet.UpdatedBy, xet.CreatedAt, xet.UpdatedAt)
	res, err := db.Exec(sqlstr, xet.TemplateName, xet.TemplateCode, xet.ActiveFlag, xet.OrgID, xet.CreatedBy, xet.UpdatedBy, xet.CreatedAt, xet.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xet.ID = uint(id)
	xet._exists = true

	return nil
}

// Update updates the XEmailTemplate in the database.
func (xet *XEmailTemplate) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xet._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xet._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_email_template SET ` +
		`template_name = ?, template_code = ?, active_flag = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xet.TemplateName, xet.TemplateCode, xet.ActiveFlag, xet.OrgID, xet.CreatedBy, xet.UpdatedBy, xet.CreatedAt, xet.UpdatedAt, xet.ID)
	_, err = db.Exec(sqlstr, xet.TemplateName, xet.TemplateCode, xet.ActiveFlag, xet.OrgID, xet.CreatedBy, xet.UpdatedBy, xet.CreatedAt, xet.UpdatedAt, xet.ID)
	return err
}

// Save saves the XEmailTemplate to the database.
func (xet *XEmailTemplate) Save(db XODB) error {
	if xet.Exists() {
		return xet.Update(db)
	}

	return xet.Insert(db)
}

// Delete deletes the XEmailTemplate from the database.
func (xet *XEmailTemplate) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xet._exists {
		return nil
	}

	// if deleted, bail
	if xet._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_email_template WHERE id = ?`

	// run query
	XOLog(sqlstr, xet.ID)
	_, err = db.Exec(sqlstr, xet.ID)
	if err != nil {
		return err
	}

	// set deleted
	xet._deleted = true

	return nil
}

// XEmailTemplatesByCreatedBy retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_email_template_created_by_foreign'.
func XEmailTemplatesByCreatedBy(db XODB, createdBy uint) ([]*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XEmailTemplate{}
	for q.Next() {
		xet := XEmailTemplate{
			_exists: true,
		}

		// scan
		err = q.Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xet)
	}

	return res, nil
}

// XEmailTemplateByID retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_email_template_id_pkey'.
func XEmailTemplateByID(db XODB, id uint) (*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xet := XEmailTemplate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xet, nil
}

// XEmailTemplatesByOrgID retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_email_template_org_id_foreign'.
func XEmailTemplatesByOrgID(db XODB, orgID uint) ([]*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XEmailTemplate{}
	for q.Next() {
		xet := XEmailTemplate{
			_exists: true,
		}

		// scan
		err = q.Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xet)
	}

	return res, nil
}

// XEmailTemplatesByUpdatedBy retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_email_template_updated_by_foreign'.
func XEmailTemplatesByUpdatedBy(db XODB, updatedBy uint) ([]*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XEmailTemplate{}
	for q.Next() {
		xet := XEmailTemplate{
			_exists: true,
		}

		// scan
		err = q.Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xet)
	}

	return res, nil
}

// XEmailTemplateByTemplateNameOrgID retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_unique_key_1'.
func XEmailTemplateByTemplateNameOrgID(db XODB, templateName string, orgID uint) (*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE template_name = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, templateName, orgID)
	xet := XEmailTemplate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, templateName, orgID).Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xet, nil
}

// XEmailTemplateByTemplateCodeOrgID retrieves a row from 'x_showroom.x_email_template' as a XEmailTemplate.
//
// Generated from index 'x_unique_key_2'.
func XEmailTemplateByTemplateCodeOrgID(db XODB, templateCode string, orgID uint) (*XEmailTemplate, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, template_name, template_code, active_flag, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_email_template ` +
		`WHERE template_code = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, templateCode, orgID)
	xet := XEmailTemplate{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, templateCode, orgID).Scan(&xet.ID, &xet.TemplateName, &xet.TemplateCode, &xet.ActiveFlag, &xet.OrgID, &xet.CreatedBy, &xet.UpdatedBy, &xet.CreatedAt, &xet.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xet, nil
}
