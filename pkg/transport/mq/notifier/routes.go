package notifiermq

import (
	mqtp "pms.pkg/transport/mq"
)

var (
	Routes = []mqtp.QueueRoute{"greet"}
)
