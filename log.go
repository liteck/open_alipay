package open_alipay

import (
	"log"
)

func init() {
	log.SetPrefix("ALIPAY: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

