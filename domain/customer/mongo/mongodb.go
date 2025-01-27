package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/ddd-go/domain/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() customer.Customer {
	c := customer.Customer{}
	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, connection_string string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection_string))
	if err != nil {
		return nil, err
	}
	db := client.Database("ddd")
	customers := db.Collection("customers")
	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (repo *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := repo.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return customer.Customer{}, err
	}
	return c.ToAggregate(), nil
}

func (repo *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(c)
	_, err := repo.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoRepository) Update(c customer.Customer) error {
	panic("to implement")
}
