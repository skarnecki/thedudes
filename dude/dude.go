package dude

import (
	"time"
)


// Model Struct
type Dude struct {
	Id  int64 `xorm:"autoincr pk"`
	Name string `xorm:"" json:"name" binding:"required"`
	CreatedAt time.Time `xorm:"created"`
}
