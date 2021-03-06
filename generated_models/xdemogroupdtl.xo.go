// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XDemoGroupDtl represents a row from 'x_showroom.x_demo_group_dtl'.
type XDemoGroupDtl struct {
	ID                uint            `json:"id"`                    // id
	Name              string          `json:"name"`                  // name
	DisplayName       sql.NullString  `json:"display_name"`          // display_name
	TypeCd            string          `json:"type_cd"`               // type_cd
	Value             sql.NullString  `json:"value"`                 // value
	Height            sql.NullString  `json:"height"`                // height
	Width             sql.NullString  `json:"width"`                 // width
	BgColor           sql.NullString  `json:"bg_color"`              // bg_color
	Alpha             sql.NullFloat64 `json:"alpha"`                 // alpha
	StyleCd           sql.NullString  `json:"style_cd"`              // style_cd
	SeqNum            uint            `json:"seq_num"`               // seq_num
	AttID             sql.NullInt64   `json:"att_id"`                // att_id
	ThumbnailAttID    sql.NullInt64   `json:"thumbnail_att_id"`      // thumbnail_att_id
	AttTypeCd         sql.NullString  `json:"att_type_cd"`           // att_type_cd
	SeqGrpStyle       sql.NullString  `json:"seq_grp_style"`         // seq_grp_style
	SeqGrpNum         int             `json:"seq_grp_num"`           // seq_grp_num
	DescTextLine1     sql.NullString  `json:"desc_text_line_1"`      // desc_text_line_1
	DescTextLine2     sql.NullString  `json:"desc_text_line_2"`      // desc_text_line_2
	DescTextLine3     sql.NullString  `json:"desc_text_line_3"`      // desc_text_line_3
	DemoGroupID       sql.NullInt64   `json:"demo_group_id"`         // demo_group_id
	CreatedBy         uint            `json:"created_by"`            // created_by
	UpdatedBy         uint            `json:"updated_by"`            // updated_by
	CreatedAt         mysql.NullTime  `json:"created_at"`            // created_at
	UpdatedAt         mysql.NullTime  `json:"updated_at"`            // updated_at
	DemoGroupDefDtlID sql.NullInt64   `json:"demo_group_def_dtl_id"` // demo_group_def_dtl_id
	ProdFeatureGrpID  sql.NullInt64   `json:"prod_feature_grp_id"`   // prod_feature_grp_id
	ContainerID       sql.NullInt64   `json:"container_id"`          // container_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XDemoGroupDtl exists in the database.
func (xdgd *XDemoGroupDtl) Exists() bool {
	return xdgd._exists
}

// Deleted provides information if the XDemoGroupDtl has been deleted from the database.
func (xdgd *XDemoGroupDtl) Deleted() bool {
	return xdgd._deleted
}

// Insert inserts the XDemoGroupDtl to the database.
func (xdgd *XDemoGroupDtl) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xdgd._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_demo_group_dtl (` +
		`name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.Value, xdgd.Height, xdgd.Width, xdgd.BgColor, xdgd.Alpha, xdgd.StyleCd, xdgd.SeqNum, xdgd.AttID, xdgd.ThumbnailAttID, xdgd.AttTypeCd, xdgd.SeqGrpStyle, xdgd.SeqGrpNum, xdgd.DescTextLine1, xdgd.DescTextLine2, xdgd.DescTextLine3, xdgd.DemoGroupID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.DemoGroupDefDtlID, xdgd.ProdFeatureGrpID, xdgd.ContainerID)
	res, err := db.Exec(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.Value, xdgd.Height, xdgd.Width, xdgd.BgColor, xdgd.Alpha, xdgd.StyleCd, xdgd.SeqNum, xdgd.AttID, xdgd.ThumbnailAttID, xdgd.AttTypeCd, xdgd.SeqGrpStyle, xdgd.SeqGrpNum, xdgd.DescTextLine1, xdgd.DescTextLine2, xdgd.DescTextLine3, xdgd.DemoGroupID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.DemoGroupDefDtlID, xdgd.ProdFeatureGrpID, xdgd.ContainerID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xdgd.ID = uint(id)
	xdgd._exists = true

	return nil
}

// Update updates the XDemoGroupDtl in the database.
func (xdgd *XDemoGroupDtl) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xdgd._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xdgd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_demo_group_dtl SET ` +
		`name = ?, display_name = ?, type_cd = ?, value = ?, height = ?, width = ?, bg_color = ?, alpha = ?, style_cd = ?, seq_num = ?, att_id = ?, thumbnail_att_id = ?, att_type_cd = ?, seq_grp_style = ?, seq_grp_num = ?, desc_text_line_1 = ?, desc_text_line_2 = ?, desc_text_line_3 = ?, demo_group_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?, demo_group_def_dtl_id = ?, prod_feature_grp_id = ?, container_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.Value, xdgd.Height, xdgd.Width, xdgd.BgColor, xdgd.Alpha, xdgd.StyleCd, xdgd.SeqNum, xdgd.AttID, xdgd.ThumbnailAttID, xdgd.AttTypeCd, xdgd.SeqGrpStyle, xdgd.SeqGrpNum, xdgd.DescTextLine1, xdgd.DescTextLine2, xdgd.DescTextLine3, xdgd.DemoGroupID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.DemoGroupDefDtlID, xdgd.ProdFeatureGrpID, xdgd.ContainerID, xdgd.ID)
	_, err = db.Exec(sqlstr, xdgd.Name, xdgd.DisplayName, xdgd.TypeCd, xdgd.Value, xdgd.Height, xdgd.Width, xdgd.BgColor, xdgd.Alpha, xdgd.StyleCd, xdgd.SeqNum, xdgd.AttID, xdgd.ThumbnailAttID, xdgd.AttTypeCd, xdgd.SeqGrpStyle, xdgd.SeqGrpNum, xdgd.DescTextLine1, xdgd.DescTextLine2, xdgd.DescTextLine3, xdgd.DemoGroupID, xdgd.CreatedBy, xdgd.UpdatedBy, xdgd.CreatedAt, xdgd.UpdatedAt, xdgd.DemoGroupDefDtlID, xdgd.ProdFeatureGrpID, xdgd.ContainerID, xdgd.ID)
	return err
}

// Save saves the XDemoGroupDtl to the database.
func (xdgd *XDemoGroupDtl) Save(db XODB) error {
	if xdgd.Exists() {
		return xdgd.Update(db)
	}

	return xdgd.Insert(db)
}

// Delete deletes the XDemoGroupDtl from the database.
func (xdgd *XDemoGroupDtl) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xdgd._exists {
		return nil
	}

	// if deleted, bail
	if xdgd._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_demo_group_dtl WHERE id = ?`

	// run query
	XOLog(sqlstr, xdgd.ID)
	_, err = db.Exec(sqlstr, xdgd.ID)
	if err != nil {
		return err
	}

	// set deleted
	xdgd._deleted = true

	return nil
}

// XDemoGroupDtlsByContainerID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_container_id_foreign'.
func XDemoGroupDtlsByContainerID(db XODB, containerID sql.NullInt64) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE container_id = ?`

	// run query
	XOLog(sqlstr, containerID)
	q, err := db.Query(sqlstr, containerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDtlsByCreatedBy retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_created_by_foreign'.
func XDemoGroupDtlsByCreatedBy(db XODB, createdBy uint) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDtlsByDemoGroupDefDtlID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_demo_group_def_dtl_id_foreign'.
func XDemoGroupDtlsByDemoGroupDefDtlID(db XODB, demoGroupDefDtlID sql.NullInt64) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE demo_group_def_dtl_id = ?`

	// run query
	XOLog(sqlstr, demoGroupDefDtlID)
	q, err := db.Query(sqlstr, demoGroupDefDtlID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDtlsByDemoGroupID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_demo_group_id_foreign'.
func XDemoGroupDtlsByDemoGroupID(db XODB, demoGroupID sql.NullInt64) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE demo_group_id = ?`

	// run query
	XOLog(sqlstr, demoGroupID)
	q, err := db.Query(sqlstr, demoGroupID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDtlByID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_id_pkey'.
func XDemoGroupDtlByID(db XODB, id uint) (*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xdgd := XDemoGroupDtl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
	if err != nil {
		return nil, err
	}

	return &xdgd, nil
}

// XDemoGroupDtlByNameTypeCdDemoGroupID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_name_type_cd_demo_group_id_unique'.
func XDemoGroupDtlByNameTypeCdDemoGroupID(db XODB, name string, typeCd string, demoGroupID sql.NullInt64) (*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE name = ? AND type_cd = ? AND demo_group_id = ?`

	// run query
	XOLog(sqlstr, name, typeCd, demoGroupID)
	xdgd := XDemoGroupDtl{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name, typeCd, demoGroupID).Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
	if err != nil {
		return nil, err
	}

	return &xdgd, nil
}

// XDemoGroupDtlsByProdFeatureGrpID retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_prod_feature_grp_id_foreign'.
func XDemoGroupDtlsByProdFeatureGrpID(db XODB, prodFeatureGrpID sql.NullInt64) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE prod_feature_grp_id = ?`

	// run query
	XOLog(sqlstr, prodFeatureGrpID)
	q, err := db.Query(sqlstr, prodFeatureGrpID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}

// XDemoGroupDtlsByUpdatedBy retrieves a row from 'x_showroom.x_demo_group_dtl' as a XDemoGroupDtl.
//
// Generated from index 'x_demo_group_dtl_updated_by_foreign'.
func XDemoGroupDtlsByUpdatedBy(db XODB, updatedBy uint) ([]*XDemoGroupDtl, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, display_name, type_cd, value, height, width, bg_color, alpha, style_cd, seq_num, att_id, thumbnail_att_id, att_type_cd, seq_grp_style, seq_grp_num, desc_text_line_1, desc_text_line_2, desc_text_line_3, demo_group_id, created_by, updated_by, created_at, updated_at, demo_group_def_dtl_id, prod_feature_grp_id, container_id ` +
		`FROM x_showroom.x_demo_group_dtl ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XDemoGroupDtl{}
	for q.Next() {
		xdgd := XDemoGroupDtl{
			_exists: true,
		}

		// scan
		err = q.Scan(&xdgd.ID, &xdgd.Name, &xdgd.DisplayName, &xdgd.TypeCd, &xdgd.Value, &xdgd.Height, &xdgd.Width, &xdgd.BgColor, &xdgd.Alpha, &xdgd.StyleCd, &xdgd.SeqNum, &xdgd.AttID, &xdgd.ThumbnailAttID, &xdgd.AttTypeCd, &xdgd.SeqGrpStyle, &xdgd.SeqGrpNum, &xdgd.DescTextLine1, &xdgd.DescTextLine2, &xdgd.DescTextLine3, &xdgd.DemoGroupID, &xdgd.CreatedBy, &xdgd.UpdatedBy, &xdgd.CreatedAt, &xdgd.UpdatedAt, &xdgd.DemoGroupDefDtlID, &xdgd.ProdFeatureGrpID, &xdgd.ContainerID)
		if err != nil {
			return nil, err
		}

		res = append(res, &xdgd)
	}

	return res, nil
}
