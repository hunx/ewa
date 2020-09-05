package server_test

import (
	"testing"

	"gitee.com/wallesoft/ewa/kernel/encryptor"
	"github.com/gogf/gf/encoding/gxml"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/test/gtest"
	"github.com/gogf/gf/util/gconv"
)

var decrypted = []byte(`<xml><AppId><![CDATA[wx3832e3725afda621]]></AppId><CreateTime>1598631541</CreateTime><InfoType><![CDATA[component_verify_ticket]]></InfoType><ComponentVerifyTicket><![CDATA[ticket@@@fzNlzI6afoWrdNO9MkEk2T6M4sQ59CWqdSHezqWuV1vkJGQcv9Dr16dBAXTNWSXNrQbAsw_aOWEwqDxbPNFWvQ]]></ComponentVerifyTicket></xml>`)
var bodyRaw = []byte(`<xml><AppId><![CDATA[wx3832e3725afda621]]></AppId><Encrypt><![CDATA[lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA==]]></Encrypt></xml>`)

func Test_parseXmlMessage(t *testing.T) {
	data := []byte(`<xml><AppId><![CDATA[wx3832e3725afda621]]></AppId><Encrypt><![CDATA[lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA==]]></Encrypt></xml>`)
	if m, err := gxml.DecodeWithoutRoot(data); err != nil {
		panic(err)
	} else {
		gtest.C(t, func(t *gtest.T) {
			if val, ok := m["AppId"]; ok {
				t.Assert(gconv.String(val), "wx3832e3725afda621")
			}
			if encrypted, ok := m["Encrypt"]; ok {
				g.Dump(encrypted)
			}
		})

	}

}

func Test_DecryptMessage(t *testing.T) {
	a := []string{"kB90oxaqQK7Aaj7qXEQVXN2Q21fbWXKO", "1598631541", "199454771", "lT3mY6ueqntYmQV4bl/sj+Zjp/jy1v7eI4NWVqmHKsLTZbe8we6BwfedsWixX5fV5IMljzoNCoK7xt1S4Bld6RCmAipTFDgqudmAkCz8Jjw7S0JVm3zXn/hQgjVnKtE1PId52kFSOyoaRYjQ+bL6mGsPvnDBQnTkX8tl8BiSY9PCRQSurr2P++dX4hazreE7UCzV6wbFJKIpi5F36jtvyzWcbkRS0s/Fix9/qu0IPEg6aW6E91E/OAGE7v4nMa5nU9Fvh/KJF24TKThNoyJvqP8UFhBnmtakGmMM2ZUItXX+pfoLX7pk3SC4yG8KQu5HPUSpbr3oZri0v2gRxf7sW8zqdB5lj4rQxUUcx6GL5xd9Q4kb/RSkN49yzLQhYTW6WzLBLI0sVkVjUfxi9kP/q2cDP9+YuEeuH4Rdn/eiXiEiaWH54EcoocDjF058/cdO5cP4LXzGlNyM15B3fZXPxA=="}
	gtest.C(t, func(t *gtest.T) {
		t.Assert(encryptor.Signature(a), "9096dc4f1ee13911603d3e436f5b0dfd22b088dd")
	})
}
func Test_GetMessage(t *testing.T) {

}