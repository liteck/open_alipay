package open_alipay

import (
	"io"
	"io/ioutil"
	"log"
)

var (
	Trace *log.Logger // 记录所有日志
)

func init() {
	Trace = log.New(ioutil.Discard,
		"OPEN_ALIPAY: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func ConfigLog(prefix string, w io.Writer) {
	if prefix != "" {
		Trace.SetPrefix(prefix)
	}
	if w != nil {
		Trace.SetOutput(w)
	}
}
