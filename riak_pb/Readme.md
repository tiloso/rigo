
*.proto files are taken from src/ of github.com/basho/riak_pb

regenerate *.pg.go files with
  `protoc --go_out=. riak.proto`
  `protoc --go_out=. riak_kv.proto`
