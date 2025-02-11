package redis

import "fmt"

func RemoveFromRedis(timelimit int64) {
	var BookKeys = []string{}
	var cursor uint64
	for {
		key, nextCursor, err := Client.Scan(Ctx, cursor, "Book:*", 0).Result()
		if err != nil {
			fmt.Println("error in getting keys", err)
		}
		BookKeys = append(BookKeys, key...)
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	for _,key:=range BookKeys{
		idleTime:=GetIdleTime(key)
		if idleTime>timelimit{
			err := Client.Del(Ctx, key).Err()
			if err!=nil{
				fmt.Println("error in deleting ",err)
			}
			fmt.Println("deleted data of book ->",key) 

		}

	}
	

}