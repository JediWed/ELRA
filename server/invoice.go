package server

import (
	"ELRA/database"
	"ELRA/lnd"
	"ELRA/structs"
	"ELRA/tools"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joncalhoun/qson"
	"github.com/lightningnetwork/lnd/lnrpc"
)

// GetInvoiceEndpoint is the Endpoint to get information about an existing Invoice like if it was settled (paid)
const GetInvoiceEndpoint = "/invoice/{rhash}"

// CreateInvoiceEndpoint is the Endpoint for CreateInvoice
const CreateInvoiceEndpoint = "/invoice/createInvoice"

// GetInvoice returns the settlement status of an invoice
func GetInvoice(response http.ResponseWriter, request *http.Request) {
	ipAddress := tools.ExtractIPAddressFromRequest(request)
	SetupCORS(&response, request)
	database.AccessLog(ipAddress, GetInvoiceEndpoint)
	rhash := mux.Vars(request)["rhash"]

	context, cancel, client, conn := lnd.ConnectLND()
	defer conn.Close()
	defer cancel()

	invoice, err := client.LookupInvoice(context, &lnrpc.PaymentHash{RHashStr: rhash})

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	responseJSON, err := json.Marshal(&structs.InvoiceCheckResponse{Settled: invoice.Settled})
	tools.CheckError(err)
	response.Write(responseJSON)
}

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
	responseJSON, err := json.Marshal(&structs.InvoiceResponse{PaymentRequest: invoice.GetPaymentRequest(), RHash: hex.EncodeToString(invoice.GetRHash())})
	tools.CheckError(err)
	response.Write(responseJSON)
}
