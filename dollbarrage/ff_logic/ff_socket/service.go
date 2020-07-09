package ff_socket

var ServiceList = make(map[int64]*Service)

type Service struct {
	 //注册client
	 Clients map[*Client]bool

	 //消息通道channel
	 BroadCast chan []byte

	 //注册来自Client的请求
	 Register chan *Client

	 //移除来自Client的请求
	 UnRegister chan *Client
}

func NewService(SId int64) *Service {
	if s, ok := ServiceList[SId]; !ok {
		ServiceList[SId] = &Service{
			Clients:make(map[*Client]bool),
			BroadCast:make(chan []byte),
			Register:make(chan *Client),
			UnRegister:make(chan *Client),
		}
		go ServiceList[SId].Run()
		return ServiceList[SId]
	}else{
		return s
	}
}

func (s *Service) Run(){
	for {
		select {
		case client := <- s.Register:
		s.Clients[client] = true
		case client := <- s.UnRegister:
			if _, ok := s.Clients[client]; ok {
				delete(s.Clients, client)
				close(client.Send)
			}
		case message := <- s.BroadCast:
			for client := range s.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(s.Clients, client)
				}
			}
		}
	}
}