package redis

import "time"

func PeriodicFunction() {
	ticker := time.NewTicker(10 *time.Second)

	for{
		select{
		case <-ticker.C:
			RemoveFromRedis(5)
		}
	}

}