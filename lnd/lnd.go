package lnd

import (
	"ELRA/globals"
	"ELRA/tools"
	"context"
	"io/ioutil"
	"strconv"

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

	conn, err := grpc.Dial(globals.Config.LightningServer+":"+strconv.Itoa(globals.Config.LightninggRPCPort), options...)
	tools.CheckError(err)
	defer conn.Close()

	Client = lnrpc.NewLightningClient(conn)

	Context := context.Background()

	_ = Context

	// getInfoResp, err := Client.GetInfo(Context, &lnrpc.GetInfoRequest{})
	// tools.CheckError(err)

	// log.Println(getInfoResp.GetVersion())
}
