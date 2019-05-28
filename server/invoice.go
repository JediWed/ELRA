package server

import (
	"ELRA/database"
	"ELRA/lnd"
	"ELRA/structs"
	"ELRA/tools"
	"encoding/json"
	"log"
	"net/http"

	"github.com/joncalhoun/qson"
	"github.com/lightningnetwork/lnd/lnrpc"
)

// CreateInvoiceEndpoint is the Endpoint for CreateInvoice
const CreateInvoiceEndpoint = "/invoice/createInvoice"

// CreateInvoice creates an invoice with amount and optional memo
func CreateInvoice(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(tools.ExtractIPAddressFromRemoteAddr(request.RemoteAddr), CreateInvoiceEndpoint)
	SetupCORS(&response, request)
	var invoiceRequest structs.InvoiceRequest

	b, errQSON := qson.ToJSON(request.URL.Query().Encode())
	errJSON := json.Unmarshal(b, &invoiceRequest)
	if errQSON != nil || errJSON != nil || invoiceRequest.Amount == 0 {
		log.Print("Invoice Creation attempt with insufficient parameters.")
		log.Print(errQSON.Error())
		log.Print(errJSON.Error())
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	context, cancel, client, conn := lnd.ConnectLND()
	defer conn.Close()
	defer cancel()

	invoice, err := client.AddInvoice(context, &lnrpc.Invoice{Memo: invoiceRequest.Memo, Value: invoiceRequest.Amount})
	tools.CheckError(err)

	response.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(&structs.InvoiceResponse{PaymentRequest: invoice.GetPaymentRequest()})
	tools.CheckError(err)
	response.Write(responseJSON)
}
