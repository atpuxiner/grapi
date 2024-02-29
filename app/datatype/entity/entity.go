// 每个功能对应一个子文件夹，便于后期扩展（ent作为前缀）

package entity

import (
	"database/sql"
	"time"
)

type DefaultEntity struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type DefaultEntityWithoutDlt struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PageInfo struct {
	Page     int
	PageSize int
}

func (r PageInfo) Limit() int {
	return r.PageSize
}

func (r PageInfo) Offset() int {
	return (r.Page - 1) * r.PageSize
}
