package global

import "deliverwater/config"
import "gorm.io/gorm"

var Config config.Config
var DB *gorm.DB
