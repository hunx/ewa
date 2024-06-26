package server

import (
	"net/http"

	guard "gitee.com/wallesoft/ewa/kernel/server"

	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/os/glog"
)

//Server
type Server struct {
	*guard.ServerGuard
	debug  *gtype.Bool  //@deprecated
	logger *glog.Logger //@deprecated
}

const (
	MUX_GROUP = "openplatform" // default config group name
)

func (s *Server) SetMux() {
	//init handler
	s.RegisterMessageType(messageType)
	s.initHandler()
}

//Resolve
func (s *Server) Resolve(msg *guard.Message) bool {

	if msg != nil {
		var t string
		if msg.Contains("InfoType") {
			t = msg.GetString("InfoType")
		} else if msg.Contains("MsgType") {
			t = msg.GetString("MsgType")
		} else {
			s.Response.WriteStatusExit(http.StatusBadRequest, "Invalid message type")
		}
		s.Dispatch(t, msg)
		s.Response.Write(guard.SUCCESS_EMPTY_RESPONSE)
	}
	return true
}

//Should return raw response
func (s *Server) ShouldReturnRawResponse() bool {
	return true
}
