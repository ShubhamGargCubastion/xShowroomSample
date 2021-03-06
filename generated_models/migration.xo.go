// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// Migration represents a row from 'x_showroom.migrations'.
type Migration struct {
	ID        uint   `json:"id"`        // id
	Migration string `json:"migration"` // migration
	Batch     int    `json:"batch"`     // batch

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Migration exists in the database.
func (m *Migration) Exists() bool {
	return m._exists
}

// Deleted provides information if the Migration has been deleted from the database.
func (m *Migration) Deleted() bool {
	return m._deleted
}

// Insert inserts the Migration to the database.
func (m *Migration) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO x_showroom.migrations (` +
		`migration, batch` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, m.Migration, m.Batch)
	res, err := db.Exec(sqlstr, m.Migration, m.Batch)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	m.ID = uint(id)
	m._exists = true

	return nil
}

// Update updates the Migration in the database.
func (m *Migration) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if m._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE x_showroom.migrations SET ` +
		`migration = ?, batch = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, m.Migration, m.Batch, m.ID)
	_, err = db.Exec(sqlstr, m.Migration, m.Batch, m.ID)
	return err
}

// Save saves the Migration to the database.
func (m *Migration) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Delete deletes the Migration from the database.
func (m *Migration) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return nil
	}

	// if deleted, bail
	if m._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM x_showroom.migrations WHERE id = ?`

	// run query
	XOLog(sqlstr, m.ID)
	_, err = db.Exec(sqlstr, m.ID)
	if err != nil {
		return err
	}

	// set deleted
	m._deleted = true

	return nil
}

// MigrationByID retrieves a row from 'x_showroom.migrations' as a Migration.
//
// Generated from index 'migrations_id_pkey'.
func MigrationByID(db XODB, id uint) (*Migration, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, migration, batch ` +
		`FROM x_showroom.migrations ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	m := Migration{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&m.ID, &m.Migration, &m.Batch)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
