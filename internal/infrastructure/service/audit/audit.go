package audit

import (
	"log"
)

// simulator of Audit Service
func AuditMsg(data string) {
	log.Printf("Data received in audit service\n")

	//saving data
	//audit.save(data)
}
