package main

import (
	"flag"

	"github.com/serg-pe/blockchainrest/internal/app/apiserver"
	"github.com/serg-pe/blockchainrest/pkg/logging"
)

var (
	configPath string
)

func main() {
	logger := logging.GetLogger()
	defer logger.Sync()

	flag.StringVar(&configPath, "config", "./configs/config.json", "Path to config file")
	flag.Parse()

	server := apiserver.NewServer(configPath)
	server.Start()

	// ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/2fe50b37b89d452499992df1cba8e410")
	// if err != nil {
	// 	logger.Fatal(
	// 		"Can't connect to Eth Client",
	// 		zap.String("err", err.Error()),
	// 	)
	// }
	// defer ethClient.Close()
	// header, err := ethClient.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	logger.Fatal(err.Error())
	// }
	// logger.Info(
	// 	"Latest block number",
	// 	zap.String("header", header.Number.String()),
	// )

	// configFile, err := os.Open(configPath)
	// defer configFile.Close()
	// if err != nil {
	// 	logger.Fatal("Can't open config file")
	// }
	// jsonDecoder := json.NewDecoder(configFile)
	// serverConfig := apiserver.Config{}
	// err = jsonDecoder.Decode(&serverConfig)
	// if err != nil {
	// 	log.Fatal("Can't decode config file")
	// }

	// server := apiserver.NewServer(&serverConfig)
	// err := server.Start()
	// if err != nil {

	// }
}
