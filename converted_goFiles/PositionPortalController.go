package main

import (
	"net/http"
	_ "database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"github.com/jinzhu/gorm"
)

type XPostn struct {
	ID          uint           `json:"id,omitempty"`           // id
	Name        string         `json:"name,omitempty"`         // name
	DisplayName string         `json:"display_name,omitempty"` // display_name
	TypeCd      string         `json:"type_cd,omitempty"`      // type_cd
	ParPostnID  int            `json:"par_postn_id,omitempty"` // par_postn_id
	OrgID       uint           `json:"org_id,omitempty"`       // org_id
	CreatedBy   uint           `json:"created_by,omitempty"`   // created_by
	UpdatedBy   uint           `json:"updated_by,omitempty"`   // updated_by
	CreatedAt   *time.Time      `json:"created_at,omitempty"`   // created_at
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`   // updated_at
	Creater       *XUser          `gorm:"ForeignKey:ID;AssociationForeignKey:CreatedBy" json:"creater,omitempty"`      //pos belongs to user
	Updater       *XUser         `gorm:"ForeignKey:ID;AssociationForeignKey:UpdatedBy" json:"updater,omitempty"`      //pos belongs to user
    Pos           *XPostn         `gorm:"ForeignKey:ID;AssociationForeignKey:ParPostnID" json:"parentposition,omitempty"`
	}


type XUser struct {
	ID             uint           `json:"id,omitempty"`              // id
	FirstName      string         `json:"first_name,omitempty"`      // first_name
	LastName       string         `json:"last_name,omitempty"`       // last_name
	Email          string         `json:"email,omitempty"`           // email
	PhoneNumber    string         `json:"phone_number,omitempty"`    // phone_number
	AddressLine1   string         `json:"address_line_1,omitempty"`  // address_line_1
	AddressLine2   string         `json:"address_line_2,omitempty"`  // address_line_2
	City           string         `json:"city,omitempty"`            // city
	State          string         `json:"state,omitempty"`           // state
	Zipcode        string         `json:"zipcode,omitempty"`         // zipcode
	Country        string         `json:"country,omitempty"`         // country
	ActiveFlag     string         `json:"active_flag,omitempty"`     // active_flag
	UserType       string         `json:"user_type,omitempty"`       // user_type
	JoiningDt      *time.Time      `json:"joining_dt,omitempty"`      // joining_dt
	AttachmentName string         `json:"attachment_name,omitempty"` // attachment_name
	Password       string         `json:"password,omitempty"`        // password
	RememberToken  string         `json:"remember_token,omitempty"`  // remember_token
	PostnID        int            `json:"postn_id,omitempty"`        // postn_id
	OrgID          int            `json:"org_id,omitempty"`          // org_id
	CreatedBy      int            `json:"created_by,omitempty"`      // created_by
	UpdatedBy      int            `json:"updated_by,omitempty"`      // updated_by
	CreatedAt      *time.Time      `json:"created_at,omitempty"`      // created_at
	UpdatedAt      *time.Time      `json:"updated_at,omitempty"`      // updated_at

}

func fetch_position(w http.ResponseWriter, req *http.Request){
	db.SingularTable(true)

	var pos []XPostn
	db.Debug().Preload("Pos").Preload("Creater", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, first_name,last_name").Where("id=?",1)
	}).Preload("Updater",func(db *gorm.DB) *gorm.DB {
		return db.Select("id, first_name,last_name")
	}).Where("id=?",2).Or("id=?",3).Find(&pos)

		json.NewEncoder(w).Encode(pos)

}

