package main

import (
	_ "database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm"
	"net/http"
	"encoding/json"
)

type XCatalog struct {
	ID              uint           `json:"id,omitempty"`                // id
	Name            string         `json:"name,omitempty"`              // name
	DisplayName     string         `json:"display_name,omitempty"`      // display_name
	SeqNum          int            `json:"seq_num,omitempty"`           // seq_num
	TypeCd          string         `json:"type_cd,omitempty"`           // type_cd
	ParCatalogID    int            `json:"par_catalog_id,omitempty"`    // par_catalog_id
	MasterCatalogID int            `json:"master_catalog_id,omitempty"` // master_catalog_id
	ActiveFlag      string         `json:"active_flag,omitempty"`       // active_flag
	AttID           int            `json:"att_id,omitempty"`            // att_id
	DemoID          int            `json:"demo_id,omitempty"`           // demo_id
	OrgID           uint           `json:"org_id,omitempty"`            // org_id
	CreatedBy       uint           `json:"created_by,omitempty"`        // created_by
	UpdatedBy       uint           `json:"updated_by,omitempty"`        // updated_by
	CreatedAt       *time.Time      `json:"created_at,omitempty"`        // created_at
	UpdatedAt       *time.Time      `json:"updated_at,omitempty"`        // updated_at
	//MasterCatalog   *XCatalog      `gorm:"ForeignKey:ID;AssociationForeignKey:MasterCatalogID" json:"mastercatalog,omitempty"`
	//ParentOrg       XOrgExt       `gorm:"ForeignKey:ID;AssociationForeignKey:OrgID"           json:"parentorg,omitempty"`
	//Products        []XCatalogProd                                                            `json:"products,omitempty"`

}

type XOrgExt struct {
	ID              uint           `json:"id,omitempty"`                // id
	Name            string         `json:"name,omitempty"`              // name
	IntegrationID   string         `json:"integration_id,omitempty"`    // integration_id
	AddressLine1    string         `json:"address_line_1,omitempty"`    // address_line_1
	AddressLine2    string         `json:"address_line_2,omitempty"`    // address_line_2
	City            string         `json:"city,omitempty"`              // city
	State           string         `json:"state,omitempty"`             // state
	Zipcode         string         `json:"zipcode,omitempty"`           // zipcode
	Country         string         `json:"country,omitempty"`           // country
	SalesMethodID   int            `json:"sales_method_id,omitempty"`   // sales_method_id
	DeviceLevelAuth string         `json:"device_level_auth,omitempty"` // device_level_auth
	CreatedBy       uint           `json:"created_by,omitempty"`        // created_by
	UpdatedBy       uint           `json:"updated_by,omitempty"`        // updated_by
	CreatedAt       *time.Time     `json:"created_at,omitempty"`        // created_at
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`        // updated_at

}

type XCatalogProd struct {
	ID        uint           `json:"id,omitempty"`         // id
	SeqNum    uint           `json:"seq_num,omitempty"`    // seq_num
	ProdID    uint           `json:"prod_id,omitempty"`    // prod_id
	CatalogID uint           `json:"catalog_id,omitempty"` // catalog_id
	CreatedBy uint           `json:"created_by,omitempty"` // created_by
	UpdatedBy uint           `json:"updated_by,omitempty"` // updated_by
	CreatedAt *time.Time      `json:"created_at,omitempty"` // created_at
	UpdatedAt *time.Time      `json:"updated_at,omitempty"` // updated_at
	//XCatalogID uint
}


func fetch_product_categories(w http.ResponseWriter, req *http.Request){
	db.SingularTable(true)

	var catalog []XCatalog
	//db.Debug().Preload("MasterCatalog").Where("type_cd=?","Category").Find(&catalog)
	db.Debug().
		Select("id,name").
		Where("id IN (?) AND master_catalog_id IS NOT NULL",
		db.Model(XCatalogProd{}).
			Select("DISTINCT catalog_id").
			Order("seq_num").
			QueryExpr()).
		Where("master_catalog_id IN (?)",
		db.Model(XCatalog{}).
			Select("id").
			QueryExpr()).
		Find(&catalog)

		json.NewEncoder(w).Encode(catalog)


}
