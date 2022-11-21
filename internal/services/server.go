package services

import (
	"fmt"
	"github.com/spf13/viper"
	"gitlab.com/dh-backend/delivery-service/config"
	"log"
	"net"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"gitlab.com/dh-backend/delivery-service/proto"
	//"gitlab.com/grpc-buffer/proto/go/pkg/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedDeliveryServiceServer
}

func Start() {
	configs := config.ReadConfigs(".")

	port := configs.GrpcPort

	PORT := fmt.Sprintf(":%s", port)
	if PORT == ":" {
		PORT += "8080"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0%v", PORT))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpc server
	grpcServer := grpc.NewServer()

	healthpb.RegisterHealthServer(grpcServer, health.NewServer())

	env, err := config.VaultSecrets(viper.GetString("VAULT_ADDR"), viper.GetString("VAULT_AUTH_TOKEN"), viper.GetString("VAULT_SECRET_PATH"))
	if err != nil {
		log.Println("couldn't load secrets")
		return
	}

	client, err := config.NewConsulClient(env.ConsulAddress)
	if err != nil {
		log.Println("couldn't connect to consul", err)
	}
	fmt.Println("consul address: ", env.ConsulAddress)

	go func() {

		client.ServiceRegistryWithConsul("delivery-grpc", "delivery", port, []string{"GRPC", "backend"})
		client.ServiceRegistryWithConsul("delivery-http", "delivery", ":8981", []string{"HTTP", "envoy"})
	}()

	deliveryServer := &Server{}

	proto.RegisterDeliveryServiceServer(grpcServer, deliveryServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
