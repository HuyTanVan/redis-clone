package command

import "my-redis/resp"

// Command handlers for string operations (SET, GET, DEL)

func (d *Dispatcher) ping(args []resp.Value) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: "string", Str: "PONG"}
	}
	return resp.Value{Typ: "bulk", Bulk: args[0].Bulk}
}

func (d *Dispatcher) set(args []resp.Value) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "WAITING FOR MORE ARGUMENTS"}
	}
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for SET"}
	}
	d.store.Set(args[0].Bulk, args[1].Bulk)
	return resp.Value{Typ: "string", Str: "OK"}
}

func (d *Dispatcher) get(args []resp.Value) resp.Value {
	val, ok := d.store.Get(args[0].Bulk)
	if !ok {
		return resp.Value{Typ: "null"} // null
	}
	return resp.Value{Typ: "bulk", Bulk: val}
}

func (d *Dispatcher) del(args []resp.Value) resp.Value {
	d.store.Del(args[0].Bulk)
	return resp.Value{Typ: "string", Str: "OK"}
}
