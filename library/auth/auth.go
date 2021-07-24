package auth

import (
	"github.com/gogf/gf/database/gdb"
)

var user *gdb.Record

func New(model *gdb.Record) {
	user = model
}

func User() *gdb.Record {
	return user
}
