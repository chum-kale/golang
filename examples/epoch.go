//getting the unix epoch for precise tam
//The Unix epoch is the point in time at midnight UTC on Thursday, January 1, 1970
//Unix time increases by one for each non-leap second that passes.

package main

import (
	"fmt"
	"time"
)

// we can get unix time in seconds, milliseconds or nanoseconds, respectively.
func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	//convert secs to current time since the epoch
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
