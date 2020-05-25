package heartbeat

import (
	"os"
	"time"

	"github.com/gt449060397/objectStorage/src/lib/rabbitmq"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}

}
