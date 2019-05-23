package server

import (
	"ELRA/lnd"
	"ELRA/structs"
	"ELRA/tools"
	"encoding/json"
	"log"
	"net/http"

	"github.com/lightningnetwork/lnd/lnrpc"
)

// CreateInvoice creates an invoice with amount and optional memo
func CreateInvoice(response http.ResponseWriter, request *http.Request) {
	SetupCORS(&response, request)
	var invoiceRequest structs.InvoiceRequest
	err := json.NewDecoder(request.Body).Decode(&invoiceRequest)
	if err != nil || invoiceRequest.Amount == 0 {
		log.Print("Invoice Creation attempt with insufficient parameters.")
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	tools.CheckError(err)

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
