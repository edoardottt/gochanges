package db

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//InsertScrapeTargets inserts into the collection "scrapeTargets"
func (dbc *DatabaseConnection) InsertScrapeTargets(scrapeTargets []ScrapeTarget) {
	databaseAndContext := dbc.connect()
	defer databaseAndContext.cleanup()

	scrapeTargetsCollection := databaseAndContext.database.Collection("scrapeTargets")

	for i := range scrapeTargets {
		target := scrapeTargets[i]
		result, err := scrapeTargetsCollection.ReplaceOne(
			databaseAndContext.ctx,
			bson.D{{"url", target.Url}},
			target,
			options.Replace().SetUpsert(true),
		)
		if err != nil {
			log.Fatal(err)
		}
		if result.MatchedCount != 0 {
			log.Println("matched and replaced an existing document")
			return
		}
		if result.UpsertedCount != 0 {
			log.Printf("inserted a new document with ID %v\n", result.UpsertedID)
		}
	}
}

// GetScrapeTargets gets all the scrapeTargets stored.
func (dbc *DatabaseConnection) GetScrapeTargets() []ScrapeTarget {
	databaseAndContext := dbc.connect()
	defer databaseAndContext.cleanup()

	var scrapeTargets []ScrapeTarget
	collection := databaseAndContext.database.Collection("scrapeTargets")
	filterCursor, err := collection.Find(databaseAndContext.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = filterCursor.All(databaseAndContext.ctx, &scrapeTargets)
	log.Printf("Got %d targets", len(scrapeTargets))
	if err != nil {
		log.Fatal(err)
	}
	return scrapeTargets
}
