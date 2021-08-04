module github.com/SEAPUNK/horahora/archiver

go 1.16

replace github.com/horahoradev/horahora/video_service => ../video_service

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.2
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.0
)
