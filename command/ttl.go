package command

import "my-redis/resp"

// get remaining time (seconds) to live of a key
// eg: TTL mykey
func (d *Dispatcher) ttl(args []resp.Value) resp.Value {
	if len(args) != 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for TTL"}
	}
	// Implementation for TTL command
	return resp.Value{Typ: "error", Str: "ERR not implemented"}
}

// delete a key after a certain time
// eg: EXPIRE mykey 10
func (d *Dispatcher) expire(args []resp.Value) resp.Value {
	// Implementation for EXPIRE command
	return resp.Value{Typ: "error", Str: "ERR not implemented"}
}
