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

var Client lnrpc.LightningClient
var Context context.Context

// SetupLND sets up all needed information for connecting lnd and creates a connection
func SetupLND() {
	log.Print("Preparing LND Connection")

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

	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(Context, globals.Config.LightningServer+":"+strconv.Itoa(globals.Config.LightninggRPCPort), options...)
	defer conn.Close()
	defer cancel()
	tools.CheckError(err)

	Client = lnrpc.NewLightningClient(conn)

	info, err := Client.GetInfo(Context, &lnrpc.GetInfoRequest{})
	tools.CheckError(err)

	log.Print("LND Connection established.")
	log.Print("LND Version: " + info.GetVersion())
	log.Print("LND Node: " + info.GetAlias())
	log.Print("LND ID: " + info.GetIdentityPubkey())

}
