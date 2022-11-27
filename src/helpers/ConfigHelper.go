package helpers

import (
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"os"
)
func CertData(path string ) []byte{
	f,err:=os.Open(path )
	goft.Error(err)
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	goft.Error(err)
	//这里去掉了上节课的decode 。不需要的
	return b
}