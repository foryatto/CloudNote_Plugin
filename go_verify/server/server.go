package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go_verify/library"
	pb "go_verify/proto"
	"log"
)

type Server struct {
	pb.UnimplementedVerifyServer
}

var redisClient *redis.Client

func init() {
	redisDb, err := library.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}
	redisClient = redisDb
}

func (s *Server) SendToUserEmail(ctx context.Context, in *pb.EmailVerifyReq) (*pb.EmailVerifyResp, error) {
	verifyInfo := library.NewRandomNums(8)
	//log.Printf("Send: %v, to: %v, by user: %v", verifyInfo, in.GetEmail(), in.GetUid())

	// todo
	// send verifyInfo to user email

	if redisClient == nil {
		return nil, errors.New("redis-server is not enable")
	}
	redisClient.Set(fmt.Sprintf("%s:%s", in.GetUid, in.GetType()), verifyInfo, 0)
	return &pb.EmailVerifyResp{Code: 0, Msg: verifyInfo}, nil
}

func (s *Server) VerifyUserEmail(ctx context.Context, in *pb.EmailVerifyReq) (*pb.EmailVerifyResp, error) {
	if redisClient == nil {
		return nil, errors.New("redis-server is not enable")
	}
	result := redisClient.Get(fmt.Sprintf("%s:%s", in.GetUid, in.GetType()))
	log.Printf(result.String(), result.String() == in.GetInfo())
	return &pb.EmailVerifyResp{}, nil
}
