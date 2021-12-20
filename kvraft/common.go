package raftkv

const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Id    int64
	Key   string
	Value string
	Op    string // "Put" or "Append"
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Id  int64
	Key string
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}
