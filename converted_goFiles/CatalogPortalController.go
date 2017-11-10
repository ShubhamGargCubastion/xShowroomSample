package main

import (
	"net/http"
	"time"
	"encoding/json"

)

func fetch_product_categories(w http.ResponseWriter, req *http.Request)  {
     db.SingularTable(true)
	var catalogs []XCatalog
	db.Debug().
		Select("id,name").
		Where("id IN (?) AND master_catalog_id IS NOT NULL",
		db.Table("x_catalog_prod").
			Select("DISTINCT catalog_id").Order("seq_num").QueryExpr()).Find(&catalogs)

	json.NewEncoder(w).Encode(catalogs)

}


type XCatalog struct {
	ID              uint           `json:"id,omitempty"`
	Name            string         `json:"name,omitempty"`              // name
	DisplayName     string         `json:"display_name,omitempty"`      // display_name
	SeqNum          int            `json:"seq_num,omitempty"`           // seq_num
	TypeCd          string         `json:"type_cd,omitempty"`           // type_cd
	ParCatalogID    int  		   `json:"par_catalog_id,omitempty"`    // par_catalog_id
	MasterCatalogID int            `json:"master_catalog_id,omitempty"` // master_catalog_id
	ActiveFlag      string         `json:"active_flag,omitempty"`       // active_flag
	AttID           int            `json:"att_id,omitempty"`            // att_id
	DemoID          int            `json:"demo_id,omitempty"`           // demo_id
	OrgID           uint           `json:"org_id,omitempty"`            // org_id
	CreatedBy       uint           `json:"created_by,omitempty"`        // created_by
	UpdatedBy       uint           `json:"updated_by,omitempty"`        // updated_by
	CreatedAt       *time.Time     `json:"created_at,omitempty"`        // created_at
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`        // updated_at
	UserOrgID		*XUser         `gorm:"ForeignKey:OrgID;AssociationForeignKey:OrgID" json:"omitempty"`
	CatProdID		*XCatalogProd  `gorm:"ForeignKey:CatalogID;AssociationForeignKey:ID" json:"omitempty"`

}
type XCatalogProd struct {
	ID        uint           `json:"id,omitempty"`         // id
	SeqNum    uint           `json:"seq_num,omitempty"`    // seq_num
	ProdID    uint           `json:"prod_id,omitempty"`    // prod_id
	CatalogID uint           `json:"catalog_id,omitempty"` // catalog_id
	CreatedBy uint           `json:"created_by,omitempty"` // created_by
	UpdatedBy uint           `json:"updated_by,omitempty"` // updated_by
	CreatedAt *time.Time     `json:"created_at,omitempty"` // created_at
	UpdatedAt *time.Time     `json:"updated_at,omitempty"` // updated_at
}