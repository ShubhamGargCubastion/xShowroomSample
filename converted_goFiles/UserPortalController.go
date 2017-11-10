package main

import (
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

func fetch_users(w http.ResponseWriter, req *http.Request)  {

          db.SingularTable(true)
	var position []XPostn

	db.Debug().Preload("ParentPosition").Preload("Creator", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, first_name,last_name")
	}).Preload("Updater",func(db *gorm.DB) *gorm.DB {
		return db.Select("id, first_name,last_name")
	}).Find(&position)

	json.NewEncoder(w).Encode(position)



}
type XPostn struct {
	ID          uint           `json:"id,omitempty"`           // id
	Name        string         `gorm:"varchar(191)" json:"name,omitempty"`         // name
	DisplayName string         `gorm:"varchar(191)" json:"display_name,omitempty"` // display_name
	TypeCd      string         `gorm:"varchar(191)" json:"type_cd,omitempty"`      // type_cd
	ParPostnID  int            `json:"par_postn_id,omitempty"` // par_postn_id
	OrgID       uint           `json:"org_id,omitempty"`       // org_id
	CreatedBy   uint           `json:"created_by,omitempty"`   // created_by
	UpdatedBy   uint           `json:"updated_by,omitempty"`   // updated_by
	CreatedAt   *time.Time      `json:"created_at,omitempty"`   // created_at
	UpdatedAt   *time.Time 	   `json:"updated_at,omitempty"`   // updated_at
	ParentPosition *XPostn     `gorm:"ForeignKey:ID;AssociationForeignKey:ParPostnID" json:"parent_position,omitempty"`
	Creator *XUser			   `gorm:"ForeignKey:ID;AssociationForeignKey:CreatedBy" json:"creator,omitempty"`
	Updater *XUser			   `gorm:"ForeignKey:ID;AssociationForeignKey:UpdatedBy" json:"updater,omitempty"`
}

type XUser struct {
	ID             uint          	`json:"id,omitempty"`              // id
	FirstName      string         	`json:"first_name,omitempty"`      // first_name
	LastName       string         	`json:"last_name,omitempty"`       // last_name
	Email          string         	`json:"email,omitempty"`           // email
	PhoneNumber    string         	`gorm:"varchar(191)" json:"phone_number,omitempty"`    // phone_number
	AddressLine1   string 			`gorm:"varchar(191)" json:"address_line_1,omitempty"`  // address_line_1
	AddressLine2   string 			`gorm:"varchar(191)" json:"address_line_2,omitempty"`  // address_line_2
	City           string 			`gorm:"varchar(191)" json:"city,omitempty"`            // city
	State          string 			`gorm:"varchar(191)" json:"state,omitempty"`           // state
	Zipcode        string 			`gorm:"varchar(191)" json:"zipcode,omitempty"`         // zipcode
	Country        string 			`gorm:"varchar(191)" json:"country,omitempty"`         // country
	ActiveFlag     string         	`json:"active_flag,omitempty"`     // active_flag
	UserType       string         	`json:"user_type,omitempty"`       // user_type
	JoiningDt      *time.Time      	`json:"joining_dt,omitempty"`      // joining_dt
	AttachmentName string       	`gorm:"varchar(191)" json:"attachment_name,omitempty"` // attachment_name
	Password       string         	`json:"password,omitempty"`        // password
	RememberToken  string 			`gorm:"varchar(100)" json:"remember_token,omitempty"`  // remember_token
	PostnID        int  			`json:"postn_id,omitempty"`        // postn_id
	OrgID          int  			`json:"org_id,omitempty"`          // org_id
	CreatedBy      int  			`json:"created_by,omitempty"`      // created_by
	UpdatedBy      int  			`json:"updated_by,omitempty"`      // updated_by
	CreatedAt      *time.Time		`json:"created_at,omitempty"`      // created_at
	UpdatedAt      *time.Time      	`json:"updated_at,omitempty"`      // updated_at

}