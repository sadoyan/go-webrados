package wrados

import "C"
import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"time"
)

type Radcon struct {
	Connection *rados.Conn
	Poolnames  map[string]bool
}

var Rconnect = &Radcon{
	Connection: nil,
	Poolnames:  map[string]bool{},
}

func RadoConnect() {
	conn, err := rados.NewConn()
	if err != nil {
		fmt.Println("Error when invoke a new connection:", err)
	}
	err = conn.ReadDefaultConfigFile()
	if err != nil {
		fmt.Println("Error when read default config file:", err)
	}
	err = conn.Connect()
	if err != nil {
		fmt.Println("Error when connect:", err)
	}
	fmt.Println("Connected to Ceph cluster")
	Rconnect.Connection = conn
}

func ListPools() {
	for {
		if Rconnect.Connection == nil {
			RadoConnect()
		}
		pools, _ := Rconnect.Connection.ListPools()
		for p := range pools {
			o := pools[p]
			Rconnect.Poolnames[o] = true
		}
		time.Sleep(10 * time.Second)
	}
}

//func PutData(pool string, name string, input []byte) {
//	ioctx, _ := Rconnect.Connection.OpenIOContext(pool)
//	_ = ioctx.Write(name, input, 0)
//}
//
//func GetData(pool string, name string) []byte {
//	if _, ok := Rconnect.Poolnames[pool]; ok {
//		ioctx, e := Rconnect.Connection.OpenIOContext(pool)
//		if e != nil {
//			fmt.Println(e)
//		}
//		xo, _ := ioctx.Stat(name)
//		bytesOut := make([]byte, xo.Size)
//		out, _ := ioctx.Read(name, bytesOut, 0)
//		fmt.Println(out, pool, name, xo.Size)
//		return bytesOut
//	} else {
//		fmt.Println("Pool " + pool + " does not exists")
//		return nil
//	}
//}
