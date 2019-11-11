/*
Copyright 2019 The LB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"net"
	"os"

	"github.com/boltdb/bolt"
	"google.golang.org/grpc"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	v1 "github.com/thetechnick/lb/pkg/api/v1"
	v1services "github.com/thetechnick/lb/pkg/api/v1/services"
	"github.com/thetechnick/lb/pkg/storage"
	boltdb "github.com/thetechnick/lb/pkg/storage/bolt"
)

const (
	port = ":50051"
)

func main() {
	ctrl.SetLogger(zap.New(func(options *zap.Options) {
		options.Development = true
	}))

	log := ctrl.Log.WithName("agent")

	boltDB, err := bolt.Open("lb.db", 0600, nil)
	if err != nil {
		log.Error(err, "open database")
		os.Exit(1)
	}
	defer boltDB.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(err, "listen")
		os.Exit(1)
	}

	eventHub := storage.NewEventHub()
	go eventHub.Run()

	s := grpc.NewServer()
	serverStorage, err := boltdb.NewServerStorage(log, boltDB, eventHub)
	if err != nil {
		log.Error(err, "creating server storage")
		os.Exit(1)
	}

	v1.RegisterServerServiceServer(s, v1services.NewServerService(log, serverStorage))

	if err := s.Serve(lis); err != nil {
		log.Error(err, "serve")
		os.Exit(1)
	}
}
