package lnd

import (
	"ELRA/globals"
	"ELRA/tools"
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

// SetupLND sets up all needed information for connecting lnd and creates a connection
func SetupLND() {
	log.Print("Preparing LND Connection")

	context, cancel, client, conn := ConnectLND()
	defer conn.Close()
	defer cancel()

	info, err := client.GetInfo(context, &lnrpc.GetInfoRequest{})
	tools.CheckError(err)

	log.Print("LND Connection established.")
	log.Print("LND Version: " + info.GetVersion())
	log.Print("LND Node: " + info.GetAlias())
	log.Print("LND ID: " + info.GetIdentityPubkey())
}

// ConnectLND connects to LND Server
func ConnectLND() (context.Context, context.CancelFunc, lnrpc.LightningClient, *grpc.ClientConn) {
	tlsCreds, err := credentials.NewClientTLSFromFile(globals.Config.TLS, "")
	tools.CheckError(err)

	macaroonBytes, err := ioutil.ReadFile(globals.Config.Macaroon)
	tools.CheckError(err)

	mac := &macaroon.Macaroon{}
	err = mac.UnmarshalBinary(macaroonBytes)
	tools.CheckError(err)

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(macaroons.NewMacaroonCredential(mac)),
	}

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(context, globals.Config.LightningServer+":"+strconv.Itoa(globals.Config.LightninggRPCPort), options...)
	tools.CheckError(err)

	client := lnrpc.NewLightningClient(conn)

	return context, cancel, client, conn

}
