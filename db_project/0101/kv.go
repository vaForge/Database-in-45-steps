package db0101

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = make(map[string][]byte) // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) ([]byte, bool, error) {
	skey := string(key)
	val, ok := kv.mem[skey]
	if !ok {
		return nil, false, nil
	}
	return val, ok, nil
}

func (kv *KV) Set(key []byte, val []byte) (bool, error) {
	skey := string(key)
	kv.mem[skey] = val
	return true, nil
}

func (kv *KV) Del(key []byte) (bool, error) {
	skey := string(key)
	if _, ok := kv.mem[skey]; !ok {
		return false, nil
	}
	delete(kv.mem, skey)
	return true, nil
}

// QzBQWVJJOUhU https://trialofcode.org/
