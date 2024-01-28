package service

import (
	"log"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/product"
	"github.com/sylphritz/go-api-server/pkg/config"
	"github.com/sylphritz/go-api-server/pkg/db/schema"
)

func GetSubscriptionPlans() {
	stripe.Key = config.GetConfig().StripeSecret
	params := &stripe.ProductListParams{}

	it := product.List(params)
	for it.Next() {
		log.Printf("%+v\n", it.Product())
	}

	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}

func CreateCustomer(user *schema.User) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(user.Email),
	}

	c, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func FindOrCreateCustomer(user *schema.User) (*stripe.Customer, error) {
	var c *stripe.Customer

	if user.CustomerID == "" {
		c, err := CreateCustomer(user)
		if err != nil {
			return nil, err
		}

		return c, nil
	}

	c, err := customer.Get(user.CustomerID, &stripe.CustomerParams{})
	if err != nil {
		return nil, err
	}

	user.CustomerID = c.ID
	err = UpdateOrCreateUser(user)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func CreateCheckoutSession() {
	// session.New()
}
