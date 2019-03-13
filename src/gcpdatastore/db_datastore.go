package gcpdatastore

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"rest-api/src/structs"
)

var projectID string = "trusty-sol-212013"

func getContextClient() (context.Context, *datastore.Client) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return ctx, client

}




func GetAll() []*structs.Person {

	var ctx, client = getContextClient()

	var people []*structs.Person
	keys, err := client.GetAll(ctx, datastore.NewQuery("Person"), &people)
	if err != nil {
		log.Fatalf("Error performing query: %v", err)
	}

	for i, key := range keys {
		fmt.Println(key)
		fmt.Println(people[i])
	}

	return people
}

func SavePerson(person structs.Person){
	var kind = "Person"
	var personId = uuid.New();
	print("UUID %v", personId.String())

	var ctx, client = getContextClient()

	personKey := datastore.NameKey(kind, personId.String(), nil)

	if _, err := client.Put(ctx, personKey, &person); err != nil {
		log.Fatalf("Failed to save person: %v", err)
	}

	fmt.Printf("Saved %v: %v\n", personKey, person.Lastname)

}
