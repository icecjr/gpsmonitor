package prepareKnow

//没成功
import (
	"fmt"
	"time"
)

func main1() {

	var firstName, lastName string

	fmt.Scanln(&firstName, &lastName)

	time.AfterFunc(5*time.Second, func() {
		return
	})
	for {

		fmt.Scanln(&firstName, &lastName)

		fmt.Printf("read str succ, ret:%s\n", firstName)
		time.AfterFunc(5*time.Second, func() {
			return
		})
		time.Sleep(1 * time.Second)

	}

}
