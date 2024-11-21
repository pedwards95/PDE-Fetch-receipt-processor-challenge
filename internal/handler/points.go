package handler

import (
	"net/http"

	"github.com/pedwards95/PDE-Fetch-receipt-processor-challenge/internal/errorhandler"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// HandleGetPoints handler function for get points endpoint
func (hl *Handler) HandleGetPoints(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errorhandler.ValidationError("id")
	}
	resp, err := hl.pointmanager.CalculatePoints(ctx, id)
	if err != nil {
		return err
	}
	hl.fastJSON(w, r, resp)
	return nil
}
