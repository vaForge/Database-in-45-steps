package db0103

import "bytes"

type KV struct {
	log Log
	mem map[string][]byte
}

func (kv *KV) Open() error {
	if err := kv.log.Open(); err != nil {
		return err
	}
	kv.mem = make(map[string][]byte)
	for {
		var ent Entry
		eof, err := kv.log.Read(&ent)
		if err != nil {
			return err
		}
		if eof {
			break
		}
		if ent.deleted {
			delete(kv.mem, string(ent.key))
		} else {
			kv.mem[string(ent.key)] = ent.val
		}
	}
	return nil
}

func (kv *KV) Close() error { return kv.log.Close() }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {
	val, ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	prev, ok := kv.mem[string(key)]
	kv.mem[string(key)] = val
	updated = !ok || !bytes.Equal(prev, val)
	ent := &Entry{key: key, val: val, deleted: false}
	err = kv.log.Write(ent)
	return
}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	_, deleted = kv.mem[string(key)]
	if !deleted {
		return
	}
	ent := &Entry{key: key, deleted: deleted}
	err = kv.log.Write(ent)
	delete(kv.mem, string(key))
	return
}

// QzBQWVJJOUhU https://trialofcode.org/
