package model

import "gorm.io/gorm"

type Address struct {
	*gorm.Model
	Line1   string `json:"line_1"`
	Line2   string `json:"line_2,omitempty"`
	Line3   string `json:"line_3,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	UserID  string
}
