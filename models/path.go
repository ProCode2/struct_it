package models

import "time"

type Path struct {
	ID          int64    `bun:"id,pk,autoincrement"`
	Title       string   `bun:"title,notnull"`
	Description string   `bun:"description,notnull"`
	Tags        []string `bun:"tags,array"`

	Levels []*Level `bun:"rel:has-many,join:id=path_id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Level struct {
	ID      int64 `bun:"id,pk,autoincrement"`
	LevelNo int   `bun:"level_no,notnull"`

	Bits []*Bit `bun:"rel:has-many,join:id=level_id"`
	Path *Path  `bun:"rel:belongs-to,join:path_id=id"`

	PathId int64 `bun:",notnull"`
}

type Bit struct {
	ID          int64  `bun:"id,pk,autoincrement"`
	Link        string `bun:"link,notnull"`
	Description string `bun:"description,notnull"`

	Level *Level `bun:"rel:belongs-to,join:level_id=id"`

	LevelId int64 `bun:",notnull"`
}
