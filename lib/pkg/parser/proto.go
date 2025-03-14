package parser

import "google.golang.org/protobuf/proto"

func ProtoToBytes(p proto.Message) ([]byte, error) {
	return proto.Marshal(p)
}

func BytesToProto(b []byte, m proto.Message) error {
	return proto.Unmarshal(b, m)
}
