package openuem_utils

import (
	"log"
	"time"

	"golang.org/x/sys/windows/svc"
)

type OpenUEMWindowsService struct {
	ServiceStart func()
	ServiceStop  func()
}

func NewOpenUEMWindowsService() *OpenUEMWindowsService {
	return &OpenUEMWindowsService{}
}

func (s *OpenUEMWindowsService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}

	s.ServiceStart()

	// service control manager loop
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				log.Println("[INFO]: service has received the stop or shutdown command")
				s.ServiceStop()
				break loop
			default:
				log.Println("[WARN]: unexpected control request")
				return true, 1
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return true, 0
}
