package pool

import (
	"io"
	"log"
	"math/rand"
	"sync/atomic"
	"time"
)

type DbConnection struct {
	ID int32
}

//DbConnection 实现Close接口
func (conn *DbConnection) Close() error {
	log.Println("close connection", conn.ID)
	return nil
}

var idCounter int32

//工厂方法创建数据库连接
func CreateDbConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new connection", id)
	return &DbConnection{id}, nil
}

func PerformQueries(i int, p *Pool) {
	//获取连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//释放资源
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	log.Printf("Query ID:%d Connect ID: %d", i, conn.(*DbConnection).ID)
}
