package redis

import "fmt"

func GetIdleTime(key string) int64{
	

	idleTime, err := Client.Do(Ctx, "OBJECT", "IDLETIME", key).Int64()
	if err != nil {
		fmt.Println("Error getting last access time:", err)
	} 
	return idleTime
}