module github.com/realzhangm/go/crond

go 1.14

// replace geecache => ./geecache
replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0

//require go.etcd.md.io/bbolt v1.3.4 // indirect

require (
	github.com/coreos/etcd v3.3.19+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.14.1 // indirect
	google.golang.org/genproto v0.0.0-20200326112834-f447254575fd // indirect
)
