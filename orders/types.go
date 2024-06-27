package orders

import "context"

type OrdersService interface{

	CreateOrder(context.Context)

}

type OrdersStore interface{

	Create()

}