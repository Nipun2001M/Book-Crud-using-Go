package channel

import "time"

var LogChannel = make(chan Logging, 100)

type Logging struct {
	Time   time.Time
	Method string
	Data   string
}

func AddToChannel( Methodapi string, Dataapi string) {
	logmsg := Logging{
		Time:   time.Now(),
		Method: Methodapi,
		Data:   Dataapi,
	}
	LogChannel <- logmsg

}
