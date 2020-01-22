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
	ipAddress := tools.ExtractIPAddressFromRequest(request)
	SetupCORS(&response, request)

	limitExceeded, err := database.AccessLimitExceeded(ipAddress, CreateInvoiceEndpoint, 3, 10)

	if err != nil || limitExceeded {
		log.Print("Limit Exceeded for IP Address " + ipAddress + " at endpoint " + CreateInvoiceEndpoint)
		if err != nil {
			log.Print(err.Error())
		}
		response.WriteHeader(http.StatusTooManyRequests)
		return
	}

	database.AccessLog(ipAddress, CreateInvoiceEndpoint)

	var invoiceRequest structs.InvoiceRequest

	err = qson.Unmarshal(&invoiceRequest, request.URL.Query().Encode())
	if err != nil || invoiceRequest.Amount == 0 {
		log.Print("Invoice Creation attempt with insufficient parameters.")
		log.Print(err.Error())
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	context, cancel, client, conn := lnd.ConnectLND()
	defer conn.Close()
	defer cancel()

	if invoiceRequest.Expiry <= 0 {
		invoiceRequest.Expiry = 3600
	} else if invoiceRequest.Expiry > 86400 {
		invoiceRequest.Expiry = 86400
	}

	invoice, err := client.AddInvoice(context, &lnrpc.Invoice{Memo: invoiceRequest.Memo, Value: invoiceRequest.Amount, Expiry: invoiceRequest.Expiry})
	tools.CheckError(err)

	response.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(&structs.InvoiceResponse{PaymentRequest: invoice.GetPaymentRequest()})
	tools.CheckError(err)
	response.Write(responseJSON)
}
