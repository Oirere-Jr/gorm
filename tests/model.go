package tests

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	gorm.Model
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *int
	Company   Company
	ManagerID uint
	Manager   *User
	Team      []User     `foreignkey:ManagerID`
	Friends   []*User    `gorm:"many2many:user_friends"`
	Languages []Language `gorm:"many2many:user_speaks"`
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	OwnerID   string
	OwnerType string
}

type Company struct {
	ID   uint
	Name string
}

type Language struct {
	Code string `gorm:primarykey`
	Name string
}
