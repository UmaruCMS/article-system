package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sony/sonyflake"
	"google.golang.org/grpc"
	// MySQL database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/UmaruCMS/article-system/rpc/client/auth"
	"github.com/UmaruCMS/article-system/rpc/client/user"
)

type config struct {
	MySQLAddr string `json:"mysql_addr"`
	RootPath  string `json:"root_path"`
	GRPCAddr  string `json:"grpc_addr"`
}

// Database Instance
var Database *gorm.DB

// RootPath is data's root path
var RootPath = ""

// UIDGenerator generates UID
var UIDGenerator *sonyflake.Sonyflake

// RPC Clients
var RPC = &struct {
	conn       *grpc.ClientConn
	AuthClient auth.AuthClient
	UserClient user.UserClient
}{}

func initDatabase(cfg *config) {
	var err error
	Database, err = gorm.Open("mysql", cfg.MySQLAddr)
	if err != nil {
		panic(err.Error())
	}
}

func initDataFolder(cfg *config) {
	if RootPath = cfg.RootPath; RootPath == "" {
		panic("root path unsepcified")
	}

	if _, err := os.Stat(RootPath); err != nil {
		err = os.MkdirAll(RootPath, 0755)
		if err != nil {
			panic(err.Error())
		}
	} else {
		return
	}

	// check again
	if _, err := os.Stat(RootPath); err != nil {
		panic(err.Error())
	}
}

func initUIDGenerator() {
	UIDGenerator = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func initRPCClients(cfg *config) {
	conn, err := grpc.Dial(cfg.GRPCAddr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	RPC.conn = conn
	RPC.AuthClient = auth.NewAuthClient(conn)
	RPC.UserClient = user.NewUserClient(conn)
}

func getConfig() *config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err.Error())
	}
	cfg := &config{}
	json.Unmarshal(raw, cfg)
	return cfg
}

func init() {
	cfg := getConfig()
	initDatabase(cfg)
	initDataFolder(cfg)
	initUIDGenerator()
	initRPCClients(cfg)
}

// Release all resources
func Release() {
	if Database != nil {
		Database.Close()
	}
	if RPC.conn != nil {
		RPC.conn.Close()
	}
}
