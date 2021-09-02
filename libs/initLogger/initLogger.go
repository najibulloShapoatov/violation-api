package initLogger

import (
	"log"
	"os"
	"time"
)

//InitLog func
func InitLog(Logfile string) {

	f, err := os.OpenFile(Logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("#======================================================================================#")
	log.Println("#======================================================================================#")
	log.Println("#-------- \t\t\t REST Started    \t\t" + time.Now().Format("2006-01-02 15:04:05") + "     \t\t---#")
	log.Println("#======================================================================================#")
	log.Println("#======================================================================================#")

}
