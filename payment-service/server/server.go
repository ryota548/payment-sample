package main

import (
	"context"
	"log"
	"net"
	"os"

	gpay "payment-sample/payment-service/proto"

	payjp "github.com/payjp/payjp-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {

	//PAIの初期化
	pay := payjp.New(os.Getenv("sk_test_2edddef9d0406bb17b6e5542"), nil)

	// 支払いをします。第一引数に支払い金額、第二引数に支払いの方法や設定を入れます。
	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     true,
		Description: req.Name + ":" + req.Description,
	})
	if err != nil {
		return nil, err
	}
	res := &gpay.PayResponse{
		Paid:     charge.Paid,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpay.RegisterPayManagerServer(s, &server{})

	reflection.Register(s)
	log.Printf("gRPC Server started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
