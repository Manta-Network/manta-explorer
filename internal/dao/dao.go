package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/itering/subscan/configs"
	"github.com/jinzhu/gorm"
)

var (
	DaemonAction = []string{"substrate"}
)

// dao
type Dao struct {
	db    *gorm.DB
	redis *redis.Pool
}

// New new a dao and return.
func New() (dao *Dao, storage *DbStorage) {
	var err error
	var err1 error
	var dc configs.MysqlConf
	var rc configs.RedisConf
	dc.MergeConf()
	rc.MergeConf()
	db, err := newDb(dc)
	dao = &Dao{
		db:    db,
		redis: redis.NewPool(rc.Config, redis.DialDatabase(rc.DbName)),
	}
	if err != nil {
		log.Error("newDb ERROR", err)
	}
	err1 = dao.Migration()
	storage = &DbStorage{db: db}
	if err1 != nil {
		log.Error("Migration ERROR", err)
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	var err1 error
	var err2 error
	if d.redis != nil {
		err1 = d.redis.Close()
	}
	if err1 != nil {
		log.Info("Redis Close ERROR", err1)
	}
	err2 = d.db.Close()
	if err2 != nil {
		log.Info("DB Close ERROR", err2)
	}
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(ctx); err != nil {
		return
	}
	// gorm auto ping
	return
}
