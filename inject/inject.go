package inject

import (
	"reflect"

	"github.com/gorilla/mux"
	"github.com/karlkfi/inject"
	"github.com/syariatifaris/gosandbox/core/config"
	rdb "github.com/syariatifaris/gosandbox/core/db"
	"github.com/syariatifaris/gosandbox/handler"
	"github.com/syariatifaris/gosandbox/modules/order"
)

//The injection is currently tightly coupled with kalkfi library

func NewDependencies() inject.Graph {
	var (
		router        *mux.Router
		orderRepo     order.OrderRepo
		orderHandler  *handler.OrderHandler
		relationalDb  rdb.DB
		configuration *config.ConfigurationData
	)
	graph := inject.NewGraph()

	//inject and resolve configuration
	graph.Define(&configuration, inject.NewProvider(config.NewConfiguration))

	//inject database
	graph.Define(&relationalDb, inject.NewProvider(rdb.NewInjectRelationalDBConnection, &configuration))
	//inject router
	graph.Define(&router, inject.NewProvider(mux.NewRouter))

	//define order handler and repositories
	graph.Define(&orderRepo, inject.NewProvider(order.NewOrderRepo, &relationalDb))
	graph.Define(&orderHandler, inject.NewProvider(handler.NewOrderHandler, &configuration, &orderRepo))

	return graph
}

func GetAssignedDependencies(graph inject.Graph, lisPtr interface{}) []reflect.Value {
	return inject.FindAssignable(graph, lisPtr)
}

func GetAssignedDependency(graph inject.Graph, v interface{}) reflect.Value {
	return inject.ExtractAssignable(graph, v)
}
