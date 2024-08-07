package discovery

import (
  "fmt"
  "net"
  "github.com/hashicorp/serf/serf"
)

type Membership struct {
  Config
  handler Handler
  serf *serf.Serf
  events chan serf.Event
}

type Config struct {
  NodeName       string
  BindAddr       string
  Tags           map[string]string
  SerfJoinAddrs  []string
}

func New(config Config, handler Handler) (*Membership, error) {
  c := &Membership{
    Config: config,
    handler: handler,
  }
  if err := c.setupSerf(); err != nil {
    return nil, err
  }
  return c, nil
}

func (m *Membership) setupSerf() error {
  addr, err := net.ResolveTCPAddr("tcp", m.BindAddr)
  if err != nil {
    return err
  }
  config := serf.DefaultConfig()
  config.Init()
  config.MemberlistConfig.BindAddr = addr.IP.String()
  config.MemberlistConfig.BindPort = addr.Port
  m.events = make(chan serf.Event)
  config.EventCh = m.events
  config.Tags = m.Tags
  config.NodeName = m.Config.NodeName

  m.serf, err = serf.Create(config)
  if err != nil {
    return err
  }

  if m.SerfJoinAddrs != nil {
    _, err = m.serf.Join(m.SerfJoinAddrs, true)
    if err != nil {
      return err
    }
  }

  go m.runHandler()
  return nil
}

type Handler interface {
  Join(name, addr string) error
  Leave(name string) error
  /* TODO: derive these commands to separate cmd driver */
  PrintLeader() error
  PrintConfig() error
  PrintMyAddr() error
}

func (m *Membership) runHandler() {
  for e := range m.events {
    switch e.EventType() {
    case serf.EventMemberJoin:
      for _, member := range e.(serf.MemberEvent).Members {
        if m.isLocal(member) {
          continue
        }
        m.handleJoin(member)
      }
    case serf.EventMemberLeave, serf.EventMemberFailed:
      for _, member := range e.(serf.MemberEvent).Members {
        if m.isLocal(member) {
          return
        }
        m.handleLeave(member)
      }
    case serf.EventUser:
      switch e.(serf.UserEvent).Name {
      case "leader":
        m.handler.PrintLeader()
      case "config":
        m.handler.PrintConfig()
      case "me":
        m.handler.PrintMyAddr()
      }
    default:
      fmt.Printf("Unknown event: %s\n", e.String())
    }
  }
}

func (m *Membership) isLocal(member serf.Member) bool {
  return m.serf.LocalMember().Name == member.Name
}

func (m *Membership) Members() []serf.Member {
  return m.serf.Members()
}

func (m *Membership) Leave() error {
  return m.serf.Leave()
}

func (m *Membership) handleJoin(member serf.Member) {
  m.handler.Join(member.Name, member.Tags["raft_addr"])
}

func (m *Membership) handleLeave(member serf.Member) {
  m.handler.Leave(member.Name)
}
