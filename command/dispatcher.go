package command

import (
	"my-redis/persistence"
	"my-redis/resp"
	"my-redis/store"
	"strings"
)

type HandlerFunc func(args []resp.Value) resp.Value

type Dispatcher struct {
	store    *store.Store
	aof      *persistence.Aof
	handlers map[string]HandlerFunc
}

// Dispatcher is responsible for routing commands to their respective handlers based on the command name.
func NewDispatcher(s *store.Store, aof *persistence.Aof) *Dispatcher {
	d := &Dispatcher{
		store:    s,
		aof:      aof,
		handlers: make(map[string]HandlerFunc),
	}
	d.registerHandlers()
	return d
}

func (d *Dispatcher) registerHandlers() {
	d.handlers["SET"] = d.set
	d.handlers["GET"] = d.get
	d.handlers["DEL"] = d.del
	d.handlers["PING"] = d.ping
}

func (d *Dispatcher) Dispatch(v resp.Value) resp.Value {
	if len(v.Array) == 0 {
		return resp.Value{Typ: "error", Str: "ERR empty command"}
	}

	// convert CMD to UPPERCASE
	cmd := strings.ToUpper(v.Array[0].Bulk)
	handler, ok := d.handlers[cmd]
	if !ok {
		return resp.Value{Typ: "error", Str: "ERR unknown command " + cmd}
	}
	if cmd == "SET" || cmd == "HSET" {
		d.aof.Write(v)
	}
	return handler(v.Array[1:])
}
