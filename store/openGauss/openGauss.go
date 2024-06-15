package openGauss

import (
	"fmt"
	og "github.com/stitchcula/OpenGauss"
	"gorm.io/gorm"
	"os"
	"strconv"
)

type (
	Orm struct {
		Host     string
		Port     string
		User     string
		Pass     string
		Database string

		Conf *gorm.Config
		*gorm.DB
	}
	Option func(r *Orm)
)

func (r *Orm) GetOrm() *gorm.DB {
	return r.DB
}

func (r *Orm) OrmConnectionUpdate(conf OrmConf) *Orm {
	orm, err := NewMysqlOrm(conf)
	if err != nil {
		return r
	}
	return orm
}

func MustNewMysqlOrm(conf OrmConf, opts ...Option) *Orm {
	orm, err := NewMysqlOrm(conf, opts...)
	if err != nil {
		os.Exit(1)
	}
	return orm
}

func NewMysqlOrm(conf OrmConf, opts ...Option) (*Orm, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	opts = append([]Option{WithAddr(conf.Host, conf.Port)}, opts...)
	opts = append([]Option{WithAuth(conf.User, conf.Pass)}, opts...)
	opts = append([]Option{WithDBName(conf.Database)}, opts...)

	return newOrm(opts...)
}

func WithGormConf(conf *gorm.Config) Option {
	return func(r *Orm) {
		r.Conf = conf
	}
}

func WithAddr(host, port string) Option {
	return func(r *Orm) {
		r.Host = host
		r.Port = port
	}
}

func WithAuth(user, pass string) Option {
	return func(r *Orm) {
		r.Pass = pass
		r.User = user
	}
}

func WithDBName(db string) Option {
	return func(r *Orm) {
		r.Database = db
	}
}

func newOrm(opts ...Option) (*Orm, error) {
	m := &Orm{}
	for _, opt := range opts {
		opt(m)
	}
	conf := m.Conf
	if conf == nil {
		conf = &gorm.Config{}
	}
	//TODO 这里做opengauss的连接
	num64, err := strconv.ParseUint(m.Port, 10, 16)
	port := uint16(num64)
	config := og.Config{
		Host:     m.Host,
		Port:     port,
		Database: m.Database,
		User:     m.User,
		Password: m.Pass,
	}
	db, err := gorm.Open(og.New(config), conf)
	if err != nil {
		fmt.Sprintf("open opengauss error: %v", err)
		return nil, fmt.Errorf("open opengauss error: %v", err)
	}
	m.DB = db
	return m, err
}

//config := og.Config{
//Host:     "192.168.xx.xx",
//Port:     5432,
//Database: "db_tpcc",
//User:     "user_persistence",
//Password: "1234@abc",
//RuntimeParams: map[string]string{
//"search_path": "my_schema",
//},
//}
//db, err := gorm.Open(og.New(config), &gorm.Config{})
