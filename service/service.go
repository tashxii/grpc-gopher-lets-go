package service

import (
	"errors"
	"fmt"
	"grpc-gopher-lets-go/message"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/context"
)

const (
	maxLife          int64 = 5
	maxSize          int64 = 100
	bombLife         int64 = 25
	bombSize         int64 = 30
	bombFireRange    int64 = 20
	bombSpeed        int64 = 20
	bombDmg          int64 = 1
	missileLife      int64 = 40
	missileSize      int64 = 50
	missileFireRnage int64 = 36
	missileSpeed     int64 = 32
	missileDmg       int64 = 4
	dmgSize          int64 = 12
)

// MessageServer model
type MessageServer struct {
	lock        *sync.Mutex
	connections map[string]*Connection
}

// Player model
type Player struct {
	id, name, charge string
	x, y, life, size int64
}

// Bullet model
type Bullet struct {
	id                             string
	x, y, life, maxlife, direction int64
	damage, speed, firerange, size int64
	fire, special                  bool
}

// Connection model
type Connection struct {
	id         string
	player     *Player
	bomb       *Bullet
	missile    *Bullet
	queue      chan string
	disconnect chan struct{}
}

// NewMessageServer returns new instance
func NewMessageServer() *MessageServer {
	return &MessageServer{
		lock:        &sync.Mutex{},
		connections: make(map[string]*Connection),
	}
}

// Connect method
func (s *MessageServer) Connect(ctx context.Context, req *message.ConnectRequest) (res *message.Empty, err error) {
	res = &message.Empty{}
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.connections[req.Id]; ok {
		delete(s.connections, req.Id)
	}
	player := &Player{
		id:   req.Id,
		name: req.Name,
		life: maxLife,
		size: maxSize,
		x:    req.X,
		y:    req.Y,
	}
	bomb := &Bullet{
		id:      req.Id,
		life:    bombLife,
		size:    bombSize,
		damage:  bombDmg,
		maxlife: bombLife,
		speed:   bombSpeed,
		fire:    false,
	}
	missile := &Bullet{
		id:      req.Id,
		life:    missileLife,
		size:    missileSize,
		damage:  missileDmg,
		maxlife: missileLife,
		speed:   missileSpeed,
		special: true,
		fire:    false,
	}

	//Show new player to other players
	for _, connection := range s.connections {
		msg := fmt.Sprintf("show %s %s %d %d %d %d %s",
			player.id,
			player.name,
			player.life,
			player.x,
			player.y,
			player.size,
			player.charge)
		connection.queue <- msg
	}

	s.connections[req.Id] = &Connection{
		id:         req.Id,
		player:     player,
		bomb:       bomb,
		missile:    missile,
		queue:      make(chan string, 1000),
		disconnect: make(chan struct{}),
	}
	fmt.Printf("Connect is called. ID [%s] \n", req.Id)
	return
}

// Disconnect method
func (s *MessageServer) Disconnect(ctx context.Context, req *message.DisconnectRequest) (res *message.Empty, err error) {
	fmt.Printf("Disonnect is called. ID [%s] \n", req.Id)
	res = &message.Empty{}
	s.lock.Lock()
	defer s.lock.Unlock()
	if connection, ok := s.connections[req.Id]; ok {
		connection.disconnect <- struct{}{}
		delete(s.connections, req.Id)
	}
	return
}

//SendMessageOthers method
func (s *MessageServer) SendMessageOthers(ctx context.Context, req *message.SendMessageRequest) (res *message.Empty, err error) {
	res = &message.Empty{}
	sender, ok := s.connections[req.FromId]
	if !ok {
		msg := fmt.Sprintf("FromId:%s does not exists", req.FromId)
		return nil, errors.New(msg)
	}
	handleMessageOthers(sender, s.connections, req)
	return
}

// SendMessageAll method
func (s *MessageServer) SendMessageAll(ctx context.Context, req *message.SendMessageRequest) (res *message.Empty, err error) {
	res = &message.Empty{}
	handleMessageAll(s, req)
	return
}

// ReceiveMessage method
func (s *MessageServer) ReceiveMessage(req *message.ReceiveMessageRequest, stream message.MessageService_ReceiveMessageServer) (err error) {
	fmt.Println("ReceiveMessage is called")
	connection, ok := s.connections[req.FromId]
	if !ok {
		msg := fmt.Sprintf("FromId:%s does not exist", req.FromId)
		return errors.New(msg)
	}
loop:
	for {
		select {
		case msg := <-connection.queue:
			params := strings.Split(msg, " ")
			stream.Send(&message.ReceiveMessageResponse{
				FromId: connection.id,
				Type:   params[0],
				Params: params[1:],
			})
		case <-connection.disconnect:
			break loop
		}
	}
	return nil
}

func handleMessageOthers(sender *Connection, connections map[string]*Connection, req *message.SendMessageRequest) {
	var msg string
	switch req.Type {
	case "show":
		//["show", e.pageX, e.pageY, charge]
		msg = showPlayer(sender.player, req.Params)
	case "fire-bomb":
		//["fire-xxx", e.pageX, e.pageY, direction]
		msg = fireBullet(sender.bomb, req.Params)
	case "fire-missile":
		//["fire-xxx", e.pageX, e.pageY, direction]
		msg = fireBullet(sender.missile, req.Params)
	}
	for id, receiver := range connections {
		if id != sender.id {
			receiver.queue <- msg
		}
	}
}

var previousTime int64
var currentTime int64

func handleMessageAll(s *MessageServer, req *message.SendMessageRequest) {
	if req.Type == "refresh" {
		// currentTime := time.Now().UnixNano() / int64(time.Nanosecond)
		// if currentTime-previousTime >= 20 {
		// 	previousTime = currentTime
		for _, sender := range s.connections {
			// player := sender.player
			missile := sender.missile
			bomb := sender.bomb
			moveBullet(missile, s.connections)
			moveBullet(bomb, s.connections)
			judgeHitBullet(sender.player, missile, s.connections)
			judgeHitBullet(sender.player, bomb, s.connections)
		}
		// }
	}
}

func showPlayer(player *Player, params []string) string {
	newX, _ := strconv.Atoi(params[0])
	player.x = int64(newX)
	newY, _ := strconv.Atoi(params[1])
	player.y = int64(newY)
	player.charge = params[2]
	msg := fmt.Sprintf("show %s %s %d %d %d %d %s",
		player.id,
		player.name,
		player.life,
		player.x,
		player.y,
		player.size,
		player.charge)
	return msg
}

func fireBullet(bullet *Bullet, params []string) string {
	bullet.fire = true
	bullet.life = bullet.maxlife
	newX, _ := strconv.Atoi(params[0])
	bullet.x = int64(newX)
	newY, _ := strconv.Atoi(params[1])
	bullet.y = int64(newY)
	direction, _ := strconv.Atoi(params[2])
	bullet.direction = int64(direction)
	msg := fmt.Sprintf("bullet %s %d %d %d %d %t",
		bullet.id,
		bullet.x,
		bullet.y,
		bullet.direction,
		bullet.size,
		bullet.special)
	return msg
}

func moveBullet(bullet *Bullet, receivers map[string]*Connection) {
	if bullet.fire == false {
		return
	}
	bullet.life = bullet.life - 1
	if bullet.life <= 0 {
		bullet.fire = false
		msg := fmt.Sprintf("miss %s %t", bullet.id, bullet.special)
		for _, receiver := range receivers {
			receiver.queue <- msg
		}
		return
	}
	var dx, dy int64
	switch bullet.direction {
	case 0:
		dy = bullet.speed
	case 1:
		dx = bullet.speed
	case 2:
		dy = -bullet.speed
	case 3:
		dx = -bullet.speed
	}
	bullet.x += dx
	bullet.y += dy
	msg := fmt.Sprintf("bullet %s %d %d %d %t", bullet.id, bullet.x, bullet.y, bullet.direction, bullet.special)
	for _, receiver := range receivers {
		receiver.queue <- msg
	}
}

func judgeHitBullet(player *Player, bullet *Bullet, receivers map[string]*Connection) {
	if bullet.fire == false || player.life <= 0 || bullet.life >= bullet.firerange {
		return
	}
	for _, receiver := range receivers {
		player := receiver.player
		if player.x-player.size/2 <= bullet.x-bullet.size/2 &&
			bullet.x+bullet.size/2 <= player.x+player.size/2 &&
			player.y-player.size/2 <= bullet.y-bullet.size/2 &&
			bullet.y+bullet.size/2 <= player.y+player.size/2 {
			player.life = player.life - bullet.damage
			player.size = player.size - dmgSize*bullet.damage
			bullet.life = 0
			bullet.fire = false
			var msg string
			if player.life <= 0 {
				msg = fmt.Sprintf("dead %s %s %t", player.id, bullet.id, bullet.special)
			} else {
				msg = fmt.Sprintf("hit %s %s %d %d %t", player.id, bullet.id, player.life, player.size, bullet.special)
			}
			for _, receiver := range receivers {
				receiver.queue <- msg
			}
		}
	}
}
