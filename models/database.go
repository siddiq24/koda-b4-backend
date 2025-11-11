package models

import (
	"github.com/siddiq24/backend-coffee-shop/configs"
)

var Pg = configs.InitPostgres()
var Rdb = configs.Redis()
