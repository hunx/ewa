package main

import (
	// "github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"

	"gitee.com/wallesoft/ewa/kernel/cache"
	"gitee.com/wallesoft/ewa/kernel/cache/adapter"
)

func main() {
	// // v := gvar.New("test a cache")
	// t := gvar.New(gtime.Timestamp())
	// // con := t.Bytes()
	// // gfile.PutBytes("/tmp/test.tmp", append(con, v.Bytes()...))
	// // g.Dump(len(t.Bytes()))
	// v := gfile.GetBytes("/tmp/test.tmp")
	// g.Dump(gvar.New(v[0:8]).Int64())
	// g.Dump(gvar.New(v[8:]).String())
	// var a float64 = 168.23
	// g.Dump(gvar.New(a).Int64())
	// // g.Dump(len(gvar.New(int64{0}).Bytes()))
	// var m = map[interface{}]interface{}{
	// 	"testindex": "testvalue",
	// 	"test":      32,
	// }
	// gfile.PutBytes("/tmp/sts.tmp", append(t.Bytes(), gvar.New(m).Bytes()...))
	// g.Dump(gvar.New(gfile.GetBytes("/tmp/sts.tmp")[8:]).Map)
	// cache := gcache.New()
	// adapter := adapter.New("/tmp/cache/")
	// cache.SetAdapter(adapter)
	// cache.Set("test.cache.one", "this istest", time.Second*1)
	// time.Sleep(time.Second * 2)
	// g.Dump(cache.Get("test.cache.one"))
	a := cache.Get()
	a.Set("test.cache.one", "this istest", 0)
	b := cache.Get()
	b.SetAdapter(adapter.New("/tmp/cache/", ".ewa"))
	a.Set("test.default.one", "this is test", 0)
	// b.Set("test.default.one", "this is test", 0)
	g.Dump(gfile.TempDir())
	g.Dump(a.Get("test.cache.one"))
	g.Dump(b.Get("test.default.one"))
}
