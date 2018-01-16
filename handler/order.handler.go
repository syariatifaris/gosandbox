package handler

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/syariatifaris/gosandbox/core/config"
	"github.com/syariatifaris/gosandbox/core/middleware"
	"github.com/syariatifaris/gosandbox/modules/order"
)

type OrderHandler struct {
	baseHandler
	repo order.OrderRepo
}

func NewOrderHandler(cfg *config.ConfigurationData, repo order.OrderRepo) *OrderHandler {
	return &OrderHandler{
		baseHandler: baseHandler{
			config: cfg,
		},
		repo: repo,
	}
}

func (c *OrderHandler) Name() string {
	return "OrderHandler"
}

func (c *OrderHandler) RegisterHandlers(muxRouter *mux.Router) {
	muxRouter.HandleFunc("/order/get/", middleware.Handle(c.GetOrderById)).Methods(http.MethodGet)
	muxRouter.HandleFunc("/order/update", middleware.Handle(c.UpdateOrderHandler)).Methods(http.MethodPost)
}

func (c *OrderHandler) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mOrder order.Order
	err := getPostData(r, &mOrder)

	if err != nil {
		render(w, nil, err)
		return
	}

	err = c.repo.UpdateOrder(mOrder)

	data := fmt.Sprintf("Success updating order with id %d", mOrder.Id)
	if err != nil {
		data = OperationFailed
	}

	render(w, data, err)
}

func (c *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderRes := c.repo.GetOrderById(23768)
	render(w, struct {
		Order order.Order `json:"order"`
	}{
		Order: orderRes,
	}, nil)
}
