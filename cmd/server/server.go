package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	ps "github.com/iztleu/sailet_api/proto"

	"github.com/iztleu/sailet_api/database/mongodb"
	"github.com/iztleu/sailet_api/models"
	"google.golang.org/grpc"
)

//AccountServiceServer is
type AccountServiceServer struct {
	ps.UnimplementedAccountServiceServer
}

//Login is
func (a *AccountServiceServer) Login(ctx context.Context,
	req *ps.LoginRequest) (*ps.Account, error) {

	var err error
	response := new(ps.Account)

	acc, err := model.Login(req.Email, req.Password)

	response.Email = acc.Email
	response.Password = acc.Password
	response.Info = acc.Info

	return response, err
}

//Get is
func (a *AccountServiceServer) Get(ctx context.Context,
	req *ps.GetRequest) (*ps.Account, error) {

	var err error
	response := new(ps.Account)

	acc, err := model.GetAccount(req.Email)

	response.Email = acc.Email
	response.Password = acc.Password
	response.Info = acc.Info

	return response, err
}

//GetAll is
func (a *AccountServiceServer) GetAll(ctx context.Context,
	req *ps.Empty) (*ps.ItemAccount, error) {

	var err error
	response := new(ps.ItemAccount)

	accounts := []ps.Account{}

	accs, err := model.GetAccounts()

	for _, v := range accs {

		account := ps.Account{
			Email:    v.Email,
			Password: v.Password,
			Info:     v.Info,
		}

		accounts = append(accounts, account)
	}

	return response, err
}

//Create is
func (a *AccountServiceServer) Create(ctx context.Context,
	req *ps.Account) (*ps.Account, error) {

	var err error
	response := new(ps.UpdateResponce)

	account := &models.Account{
		Email:    req.Email,
		Password: req.Password,
		Info:     req.Info,
	}

	ok, err := model.CreateOrUpdateAccount(account)

	response.Message = ok

	return req, err
}

//Update is
func (a *AccountServiceServer) Update(ctx context.Context,
	req *ps.Account) (*ps.UpdateResponce, error) {

	var err error
	response := new(ps.UpdateResponce)

	account := &models.Account{
		Email:    req.Email,
		Password: req.Password,
		Info:     req.Info,
	}

	ok, err := model.CreateOrUpdateAccount(account)

	response.Message = ok

	return response, err
}

//Delet is
func (a *AccountServiceServer) Delet(ctx context.Context,
	req *ps.Account) (*ps.DeletResponce, error) {

	var err error
	response := new(ps.DeletResponce)

	msg, err := model.DeleteAccount(req.Email)

	response.Message = msg
	return response, err
}

var model *models.MgModel
var port = flag.Int("port", 50051, "the port to serve on")
var host = flag.String("host", "localhost:27017", "the port to serve on")

func main() {

	flag.Parse()
	fmt.Printf("server starting on port %d...\n", *port)
	fmt.Printf("server starting host db %s...\n", host)
	err := mongodb.InitDatabaseConnection(*host)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mongodb.CloseDatabaseConnection()
	s := mongodb.GetSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	model = models.NewMgModel(s)

	server := grpc.NewServer()
	instance := new(AccountServiceServer)

	ps.RegisterAccountServiceServer(server, instance)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}
	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}

}
