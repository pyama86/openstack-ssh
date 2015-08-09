package openstack_ssh

import (
	"log"
	"log/syslog"
)

func ERR(message error) {
	logger, err := syslog.New(syslog.LOG_ERR|syslog.LOG_USER, "openstack-ssh")
	if err != nil {
		panic(err)
	}
	log.SetOutput(logger)
	log.Println(message)
}
