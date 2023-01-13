package database

import (
	"context"
	"log"

	"Github/amc-apis-go/internal/data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBClient        *mongo.Client
	MeezanDB        *mongo.Collection
	DBContext       context.Context
	DBContextCancel context.CancelFunc
)

// const uri = "mongodb+srv://sherry:1234asdf@amc.4qjwjcc.mongodb.net/?retryWrites=true&w=majority"

// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// clientOptions := options.Client().
//     ApplyURI("mongodb+srv://sherry:1234asdf@amc.4qjwjcc.mongodb.net/?retryWrites=true&w=majority").
//     SetServerAPIOptions(serverAPIOptions)
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, clientOptions)
// if err != nil {
//     log.Fatal(err)
// }

// func main() {
// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()
// 	// Ping the primary
// 	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Successfully connected and pinged.")
// }

func ConnectToDB(uri string, database string, collection string) error {
	DBContext = context.Background()
	DBClient, err := mongo.Connect(DBContext, options.Client().ApplyURI(uri))
	MeezanDB = DBClient.Database(database).Collection(collection)

	if err != nil {
		panic(err)
		log.Println(" COULD NOT CONNECT TO MONGO DB ")
		log.Fatal(err)
		return err
	}

	if err := DBClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err, nil)
		return err
	}

	log.Println(" CONNECTED AND PINGED DATABASE SUCCESSFULLY ")

	return nil
}

func GetAccOpeningRecord(cnic string) (status string, err error) {

	var record data.AccountOpeningDBStruct
	record.Cnic = cnic

	err = MeezanDB.FindOne(DBContext, bson.M{"_id": cnic}).Decode(&record)

	if err != nil {
		log.Println("ERR: ", err)
		if err == mongo.ErrNoDocuments {

			var newRecord data.AccountOpeningDBStruct
			newRecord.Cnic = cnic
			newRecord.Status = "Pending"
			_, err = MeezanDB.InsertOne(DBContext, newRecord)

			if err != nil {
				log.Fatal(err)
			}

			return "Pending", nil
		}
	}

	log.Println("Record: ", record)

	return record.Status, nil
}

func UpdateAccOpeningRecord(cnic string, status string) (err error) {
	var record data.AccountOpeningDBStruct
	record.Cnic = cnic

	err = MeezanDB.FindOne(DBContext, bson.M{"_id": cnic}).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("RECORD FOUND: ", record)

	_, err = MeezanDB.UpdateOne(
		DBContext,
		bson.M{"_id": cnic},
		bson.M{"$set": bson.M{"status": status}},
		// options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DOCUMENT UPDATED SUCCESSFULLY")

	return nil
}

func GetInvestMentRecord(cnic string, fundId string) (status string, err error) {

	var record data.InvestmentDBStruct
	record.Cnic = cnic
	record.FundId = fundId

	err = MeezanDB.FindOne(DBContext, bson.M{"_id": cnic, "fund_id": fundId}).Decode(&record)

	if err != nil {
		log.Println("ERR: ", err)
		if err == mongo.ErrNoDocuments {
			return "NA", err
		}
	}

	log.Println("Record: ", record)

	return record.Status, nil
}

func CreateInvestmentRecord(cnic string, fundId string, units float64) (err error) {
	newRecord := data.InvestmentDBStruct{
		Cnic:   cnic,
		Status: "pending",
		FundId: fundId,
		Units:  units,
	}
	_, err = MeezanDB.InsertOne(DBContext, newRecord)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func UpdateInvestmentRecord(cnic string, status string) (err error) {
	var record data.InvestmentDBStruct
	record.Cnic = cnic

	err = MeezanDB.FindOne(DBContext, bson.M{"_id": cnic}).Decode(&record)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("RECORD FOUND: ", record)

	_, err = MeezanDB.UpdateOne(
		DBContext,
		bson.M{"_id": cnic},
		bson.M{"$set": bson.M{"status": status}},
		// options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("DOCUMENT UPDATED SUCCESSFULLY")

	return nil
}

// func main() {
// 	const uri = "mongodb+srv://sherry:1234asdf@amc.4qjwjcc.mongodb.net/?retryWrites=true&w=majority"
// 	connect(uri)
// 	getAccOpeningRecord("4230167123667")
// 	updateAccOpeningRecord("4230167123667", "completed")
// 	// getRecord("4230167123667")
// }
