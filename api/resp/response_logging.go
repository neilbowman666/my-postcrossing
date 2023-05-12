package resp

import (
	"log"
	"my-postcrossing/util"
)

func ResponseLogging(msg interface{}) string {
	logUUID, _ := util.GetUUID()
	log.Printf("log_uuid-`%v`|msg-`%v`\n", logUUID, msg)
	return logUUID
}
