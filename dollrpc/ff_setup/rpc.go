package ff_setup

import (
	"net/rpc/jsonrpc"
	"net"
	"os"
	"strings"
	"github.com/sirupsen/logrus"
)

func SetupRpc(host string, allowIp string) error {
	//服务注册
	RegisterService()

	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Errorf("Accept data fail. Error : %s ", err.Error())
			continue
		}
		visitIp := conn.RemoteAddr().String()
		if checkAuth(allowIp, visitIp) {
			go jsonrpc.ServeConn(conn)
		}else{
			logrus.Errorf("RemoteIp %s not allow ip.", visitIp)
			conn.Close()
		}
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		logrus.Fatalf("[Rpc] Fatal error : %s ", err.Error())
		os.Exit(1)
	}
}

func checkAuth(allowIp string, visitIp string) bool {
	allowList := strings.Split(allowIp, ",")
	addr := strings.Split(visitIp, ":")
	authorized := false
	for _, v := range allowList {
		if v == addr[0] {
			authorized = true
			break
		}
	}
	if authorized == true {
		return true
	} else {
		return false
	}
}
