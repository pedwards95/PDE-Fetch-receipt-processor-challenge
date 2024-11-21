package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/errorhandler"
	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/models"
)

// HandleProcessReceipt handler function for post receipt endpoint
func (hl *Handler) HandleProcessReceipt(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errorhandler.ValidationError("request.Body", "HandleProcessReceipt; Could not read body")
	}
	bodyReceiptModel := &models.Receipt{}
	err = json.Unmarshal(data, bodyReceiptModel)
	if err != nil {
		return errorhandler.ValidationError("receipt")
	}
	uid, err := hl.receiptmanager.ProcessReceipt(ctx, bodyReceiptModel)
	if err != nil {
		return err
	}
	_, err = hl.pointmanager.CalculatePoints(ctx, uid.ID)
	if err != nil {
		return err
	}

	hl.fastJSON(w, r, uid)
	return nil
}
