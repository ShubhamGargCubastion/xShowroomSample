// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// XAccount represents a row from 'x_showroom.x_account'.
type XAccount struct {
	ID               uint           `json:"id"`                 // id
	Name             string         `json:"name"`               // name
	Email            string         `json:"email"`              // email
	PhoneNumber      string         `json:"phone_number"`       // phone_number
	AddressLine1     sql.NullString `json:"address_line_1"`     // address_line_1
	AddressLine2     sql.NullString `json:"address_line_2"`     // address_line_2
	City             sql.NullString `json:"city"`               // city
	State            sql.NullString `json:"state"`              // state
	Zipcode          sql.NullString `json:"zipcode"`            // zipcode
	Country          sql.NullString `json:"country"`            // country
	ProfilePhotoName sql.NullString `json:"profile_photo_name"` // profile_photo_name
	GenderCd         sql.NullString `json:"gender_cd"`          // gender_cd
	OccupationCd     sql.NullString `json:"occupation_cd"`      // occupation_cd
	SalaryLevel      sql.NullString `json:"salary_level"`       // salary_level
	TwitterID        sql.NullString `json:"twitter_id"`         // twitter_id
	FacebookID       sql.NullString `json:"facebook_id"`        // facebook_id
	LinkedinID       sql.NullString `json:"linkedin_id"`        // linkedin_id
	OrgID            sql.NullInt64  `json:"org_id"`             // org_id
	CreatedBy        uint           `json:"created_by"`         // created_by
	UpdatedBy        uint           `json:"updated_by"`         // updated_by
	CreatedAt        mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt        mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the XAccount exists in the database.
func (xa *XAccount) Exists() bool {
	return xa._exists
}

// Deleted provides information if the XAccount has been deleted from the database.
func (xa *XAccount) Deleted() bool {
	return xa._deleted
}

// Insert inserts the XAccount to the database.
func (xa *XAccount) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if xa._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.x_account (` +
		`name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, xa.Name, xa.Email, xa.PhoneNumber, xa.AddressLine1, xa.AddressLine2, xa.City, xa.State, xa.Zipcode, xa.Country, xa.ProfilePhotoName, xa.GenderCd, xa.OccupationCd, xa.SalaryLevel, xa.TwitterID, xa.FacebookID, xa.LinkedinID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt)
	res, err := db.Exec(sqlstr, xa.Name, xa.Email, xa.PhoneNumber, xa.AddressLine1, xa.AddressLine2, xa.City, xa.State, xa.Zipcode, xa.Country, xa.ProfilePhotoName, xa.GenderCd, xa.OccupationCd, xa.SalaryLevel, xa.TwitterID, xa.FacebookID, xa.LinkedinID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	xa.ID = uint(id)
	xa._exists = true

	return nil
}

// Update updates the XAccount in the database.
func (xa *XAccount) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xa._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if xa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.x_account SET ` +
		`name = ?, email = ?, phone_number = ?, address_line_1 = ?, address_line_2 = ?, city = ?, state = ?, zipcode = ?, country = ?, profile_photo_name = ?, gender_cd = ?, occupation_cd = ?, salary_level = ?, twitter_id = ?, facebook_id = ?, linkedin_id = ?, org_id = ?, created_by = ?, updated_by = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, xa.Name, xa.Email, xa.PhoneNumber, xa.AddressLine1, xa.AddressLine2, xa.City, xa.State, xa.Zipcode, xa.Country, xa.ProfilePhotoName, xa.GenderCd, xa.OccupationCd, xa.SalaryLevel, xa.TwitterID, xa.FacebookID, xa.LinkedinID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt, xa.ID)
	_, err = db.Exec(sqlstr, xa.Name, xa.Email, xa.PhoneNumber, xa.AddressLine1, xa.AddressLine2, xa.City, xa.State, xa.Zipcode, xa.Country, xa.ProfilePhotoName, xa.GenderCd, xa.OccupationCd, xa.SalaryLevel, xa.TwitterID, xa.FacebookID, xa.LinkedinID, xa.OrgID, xa.CreatedBy, xa.UpdatedBy, xa.CreatedAt, xa.UpdatedAt, xa.ID)
	return err
}

// Save saves the XAccount to the database.
func (xa *XAccount) Save(db XODB) error {
	if xa.Exists() {
		return xa.Update(db)
	}

	return xa.Insert(db)
}

// Delete deletes the XAccount from the database.
func (xa *XAccount) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !xa._exists {
		return nil
	}

	// if deleted, bail
	if xa._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.x_account WHERE id = ?`

	// run query
	XOLog(sqlstr, xa.ID)
	_, err = db.Exec(sqlstr, xa.ID)
	if err != nil {
		return err
	}

	// set deleted
	xa._deleted = true

	return nil
}

// XUserByCreatedBy returns the XUser associated with the XAccount's CreatedBy (created_by).
//
// Generated from foreign key 'x_account_created_by_foreign'.
func (xa *XAccount) XUserByCreatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xa.CreatedBy)
}

// XOrgExt returns the XOrgExt associated with the XAccount's OrgID (org_id).
//
// Generated from foreign key 'x_account_org_id_foreign'.
func (xa *XAccount) XOrgExt(db XODB) (*XOrgExt, error) {
	return XOrgExtByID(db, uint(xa.OrgID.Int64))
}

// XUserByUpdatedBy returns the XUser associated with the XAccount's UpdatedBy (updated_by).
//
// Generated from foreign key 'x_account_updated_by_foreign'.
func (xa *XAccount) XUserByUpdatedBy(db XODB) (*XUser, error) {
	return XUserByID(db, xa.UpdatedBy)
}

// XAccountsByCreatedBy retrieves a row from 'x_showroom.x_account' as a XAccount.
//
// Generated from index 'x_account_created_by_foreign'.
func XAccountsByCreatedBy(db XODB, createdBy uint) ([]*XAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_account ` +
		`WHERE created_by = ?`

	// run query
	XOLog(sqlstr, createdBy)
	q, err := db.Query(sqlstr, createdBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAccount{}
	for q.Next() {
		xa := XAccount{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.Email, &xa.PhoneNumber, &xa.AddressLine1, &xa.AddressLine2, &xa.City, &xa.State, &xa.Zipcode, &xa.Country, &xa.ProfilePhotoName, &xa.GenderCd, &xa.OccupationCd, &xa.SalaryLevel, &xa.TwitterID, &xa.FacebookID, &xa.LinkedinID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}

// XAccountByEmailPhoneNumberOrgID retrieves a row from 'x_showroom.x_account' as a XAccount.
//
// Generated from index 'x_account_email_phone_number_org_id_unique'.
func XAccountByEmailPhoneNumberOrgID(db XODB, email string, phoneNumber string, orgID sql.NullInt64) (*XAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_account ` +
		`WHERE email = ? AND phone_number = ? AND org_id = ?`

	// run query
	XOLog(sqlstr, email, phoneNumber, orgID)
	xa := XAccount{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email, phoneNumber, orgID).Scan(&xa.ID, &xa.Name, &xa.Email, &xa.PhoneNumber, &xa.AddressLine1, &xa.AddressLine2, &xa.City, &xa.State, &xa.Zipcode, &xa.Country, &xa.ProfilePhotoName, &xa.GenderCd, &xa.OccupationCd, &xa.SalaryLevel, &xa.TwitterID, &xa.FacebookID, &xa.LinkedinID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xa, nil
}

// XAccountByID retrieves a row from 'x_showroom.x_account' as a XAccount.
//
// Generated from index 'x_account_id_pkey'.
func XAccountByID(db XODB, id uint) (*XAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_account ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	xa := XAccount{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&xa.ID, &xa.Name, &xa.Email, &xa.PhoneNumber, &xa.AddressLine1, &xa.AddressLine2, &xa.City, &xa.State, &xa.Zipcode, &xa.Country, &xa.ProfilePhotoName, &xa.GenderCd, &xa.OccupationCd, &xa.SalaryLevel, &xa.TwitterID, &xa.FacebookID, &xa.LinkedinID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &xa, nil
}

// XAccountsByOrgID retrieves a row from 'x_showroom.x_account' as a XAccount.
//
// Generated from index 'x_account_org_id_foreign'.
func XAccountsByOrgID(db XODB, orgID sql.NullInt64) ([]*XAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_account ` +
		`WHERE org_id = ?`

	// run query
	XOLog(sqlstr, orgID)
	q, err := db.Query(sqlstr, orgID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAccount{}
	for q.Next() {
		xa := XAccount{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.Email, &xa.PhoneNumber, &xa.AddressLine1, &xa.AddressLine2, &xa.City, &xa.State, &xa.Zipcode, &xa.Country, &xa.ProfilePhotoName, &xa.GenderCd, &xa.OccupationCd, &xa.SalaryLevel, &xa.TwitterID, &xa.FacebookID, &xa.LinkedinID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}

// XAccountsByUpdatedBy retrieves a row from 'x_showroom.x_account' as a XAccount.
//
// Generated from index 'x_account_updated_by_foreign'.
func XAccountsByUpdatedBy(db XODB, updatedBy uint) ([]*XAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, email, phone_number, address_line_1, address_line_2, city, state, zipcode, country, profile_photo_name, gender_cd, occupation_cd, salary_level, twitter_id, facebook_id, linkedin_id, org_id, created_by, updated_by, created_at, updated_at ` +
		`FROM x_showroom.x_account ` +
		`WHERE updated_by = ?`

	// run query
	XOLog(sqlstr, updatedBy)
	q, err := db.Query(sqlstr, updatedBy)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*XAccount{}
	for q.Next() {
		xa := XAccount{
			_exists: true,
		}

		// scan
		err = q.Scan(&xa.ID, &xa.Name, &xa.Email, &xa.PhoneNumber, &xa.AddressLine1, &xa.AddressLine2, &xa.City, &xa.State, &xa.Zipcode, &xa.Country, &xa.ProfilePhotoName, &xa.GenderCd, &xa.OccupationCd, &xa.SalaryLevel, &xa.TwitterID, &xa.FacebookID, &xa.LinkedinID, &xa.OrgID, &xa.CreatedBy, &xa.UpdatedBy, &xa.CreatedAt, &xa.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &xa)
	}

	return res, nil
}