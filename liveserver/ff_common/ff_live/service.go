package ff_live

import (
	"time"
)

var deviceList = make(map[string]*LiveService)

func GetDevice(deviceId string) (*LiveService, error) {
	if d, ok := deviceList[deviceId]; !ok {
		LiveService, err := NewLiveService(deviceId)
		if err != nil {
			return nil, err
		}
		return LiveService, nil
	} else {
		return d, nil
	}
}

type LiveService struct {
	DeviceId     string
	VideoChanOne chan []byte
	VideoChanTwo chan []byte
	Clients      map[*Client]bool
	Register     chan *Client
	Unregister   chan *Client
}

func NewLiveService(deviceId string) (*LiveService, error) {
	deviceList[deviceId] = &LiveService{
		DeviceId:     deviceId,
		VideoChanOne: make(chan []byte, 40*1024),
		VideoChanTwo: make(chan []byte, 40*1024),
		Clients:      make(map[*Client]bool),
		Register:     make(chan *Client),
		Unregister:   make(chan *Client),
	}
	go deviceList[deviceId].Run()
	return deviceList[deviceId], nil
}

func (p *LiveService) Broadcast(data []byte, typeId int) {
	if len(p.Clients) > 0 {
		switch typeId {
		case 1:
			p.VideoChanOne <- data
		case 2:
			p.VideoChanTwo <- data
		}
	}
}

func (p *LiveService) Run() {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
		case client := <-p.Unregister:
			if _, ok := p.Clients[client]; ok {
				delete(p.Clients, client)
				close(client.SendQueueOne)
				close(client.SendQueueTwo)
			}
		case message := <-p.VideoChanOne:
			for client := range p.Clients {
				select {
				case client.SendQueueOne <- message:
				default:
				}
			}
		case message := <-p.VideoChanTwo:
			for client := range p.Clients {
				select {
				case client.SendQueueTwo <- message:
				default:
				}
			}
		case <-ticker.C:
		}
	}
}
