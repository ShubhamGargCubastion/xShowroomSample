// Package generated_models contains the types for schema 'x_showroom'.
package generated_models

// Code generated by xo. DO NOT EDIT.

import (
	"time"
)

// PasswordReset represents a row from 'x_showroom.password_resets'.
type PasswordReset struct {
	Email     string    `json:"email"`      // email
	Token     string    `json:"token"`      // token
	CreatedAt time.Time `json:"created_at"` // created_at
}

// PasswordResetsByEmail retrieves a row from 'x_showroom.password_resets' as a PasswordReset.
//
// Generated from index 'password_resets_email_index'.
func PasswordResetsByEmail(db XODB, email string) ([]*PasswordReset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`email, token, created_at ` +
		`FROM x_showroom.password_resets ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	q, err := db.Query(sqlstr, email)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PasswordReset{}
	for q.Next() {
		pr := PasswordReset{}

		// scan
		err = q.Scan(&pr.Email, &pr.Token, &pr.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &pr)
	}

	return res, nil
}

// PasswordResetsByToken retrieves a row from 'x_showroom.password_resets' as a PasswordReset.
//
// Generated from index 'password_resets_token_index'.
func PasswordResetsByToken(db XODB, token string) ([]*PasswordReset, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`email, token, created_at ` +
		`FROM x_showroom.password_resets ` +
		`WHERE token = ?`

	// run query
	XOLog(sqlstr, token)
	q, err := db.Query(sqlstr, token)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*PasswordReset{}
	for q.Next() {
		pr := PasswordReset{}

		// scan
		err = q.Scan(&pr.Email, &pr.Token, &pr.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &pr)
	}

	return res, nil
}