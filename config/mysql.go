package config

import "time"

type Mysql struct {
	Host                      string        `yaml:"host"`
	Port                      int64         `yaml:"port"`
	User                      string        `yaml:"user"`
	Password                  string        `yaml:"password"`
	Dbname                    string        `yaml:"dbname"`
	CharSet                   string        `yaml:"charSet"`
	ParseTime                 bool          `yaml:"parseTime"`
	TimeZone                  string        `yaml:"loc"`
	Gorm                      Gorm          `yaml:"gorm"`
	SlowSql                   time.Duration `yaml:"slowSql"`
	LogLevel                  string        `yaml:"logLevel"`
	IgnoreRecordNotFoundError bool          `yaml:"ignoreRecordNotFoundError"`
}
type Gorm struct {
	SkipDefaultTx   bool   `yaml:"skipDefaultTx"`
	TablePrefix     string `yaml:"tablePrefix"`
	SingularTable   bool   `yaml:"singularTable"`
	CoverLogger     bool   `yaml:"coverLogger"`
	PrepareStmt     bool   `yaml:"prepareStmt"`
	CloseForeignKey bool   `yaml:"disableForeignKeyConstraintWhenMigrating"`
}
