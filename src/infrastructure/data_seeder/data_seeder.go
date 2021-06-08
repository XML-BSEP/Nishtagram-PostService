package data_seeder

import (
	"fmt"
	"github.com/gocql/gocql"
	"post-service/repository"
	"time"
)

func SeedData(cassandraSession *gocql.Session) {
	SeedAllReportType(cassandraSession)
	SeedAllReportStatus(cassandraSession)
	SeedPosts(cassandraSession)
	SeedLikes(cassandraSession)
	SeedCollection(cassandraSession)
	SeedFavorites(cassandraSession)
	SeedReports(cassandraSession)
}

const (
	CreateReportTypesTable = "CREATE TABLE IF NOT EXISTS post_keyspace.ReportType (name text, PRIMARY KEY (name));"
	CreateReportStatusTable = "CREATE TABLE IF NOT EXISTS post_keyspace.ReportStatus (name text, PRIMARY KEY (name));"
	InsertIntoReportTypes = "INSERT INTO post_keyspace.ReportType (name) VALUES (?) IF NOT EXISTS;"
	InsertIntoReportStatus = "INSERT INTO post_keyspace.ReportStatus (name) VALUES (?) IF NOT EXISTS;"
)

func SeedAllReportStatus(session *gocql.Session) {
	err := session.Query(CreateReportStatusTable).Exec()

	if err != nil {
		fmt.Println(err)
	}

	err = session.Query(InsertIntoReportStatus, "CREATED").Exec()
	err = session.Query(InsertIntoReportStatus, "DECLINED").Exec()
	err = session.Query(InsertIntoReportStatus, "APPROVED").Exec()

	if err != nil {
		fmt.Println(err)
	}
}

func SeedAllReportType(session *gocql.Session) {
	err := session.Query(CreateReportTypesTable).Exec()

	if err != nil {
		fmt.Println(err)
	}

	err = session.Query(InsertIntoReportTypes, "NUDITY").Exec()
	err = session.Query(InsertIntoReportTypes, "HATE_SPEECH").Exec()
	err = session.Query(InsertIntoReportTypes, "VIOLENCE_ORG").Exec()
	err = session.Query(InsertIntoReportTypes, "ILLEGAL_SALES").Exec()
	err = session.Query(InsertIntoReportTypes, "BULLYING").Exec()
	err = session.Query(InsertIntoReportTypes, "VIOLATION_IP").Exec()
	err = session.Query(InsertIntoReportTypes, "SCAM").Exec()
	err = session.Query(InsertIntoReportTypes, "SELF_HARM").Exec()
	err = session.Query(InsertIntoReportTypes, "FALSE_INFO").Exec()

	if err != nil {
		fmt.Println(err)
	}

}

func SeedReports(session *gocql.Session) {
	err := session.Query(repository.InsertReportStatement, "9331c882-f72c-427b-bdfc-3918d90dc364", "4752f49f-3011-44af-9c62-2a6f4086233d", time.Now(),
		"424935b1-766c-4f99-b306-9263731518bc", "e2b5f92e-c31b-11eb-8529-0242ac130003", "NUDITY", "CREATED").Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func SeedFavorites(session *gocql.Session) {
	media := make(map[string]string, 3)

	media["4752f49f-3011-44af-9c62-2a6f4086233d"] = "e2b5f92e-c31b-11eb-8529-0242ac130003"
	media["d459e0f2-ab61-48e8-a593-29933ce99525"] = "424935b1-766c-4f99-b306-9263731518bc"
	media["adfee6f4-fe45-40ad-8f8e-760ec861a35e"] = "43420055-3174-4c2a-9823-a8f060d644c3"

	err := session.Query(repository.InsertFavoriteStatement, "a2c2f993-dc32-4a82-82ed-a5f6866f7d03", time.Now(), media).Exec()
	if err != nil {
		fmt.Println(err)
	}
}

func SeedCollection(session *gocql.Session) {
	media := make(map[string]string, 1)

	media["adfee6f4-fe45-40ad-8f8e-760ec861a35e"] = "43420055-3174-4c2a-9823-a8f060d644c3"
	err := session.Query(repository.InsertCollectionStatement, "1f1280aa-8048-4d5d-a950-ab52752f9672", "e2b5f92e-c31b-11eb-8529-0242ac130003",
		"Ide Gas", time.Now(), media).Exec()
	if err !=  nil {
		fmt.Println(err)
	}
}

func SeedLikes(session *gocql.Session) {
	err := session.Query(repository.InsertLikeStatement, "4752f49f-3011-44af-9c62-2a6f4086233d", time.Now(), "424935b1-766c-4f99-b306-9263731518bc").Exec()
	if err != nil {
		fmt.Println(err)
	}
	var id string
	var profile_id string
	var timestamp time.Time

	iter := session.Query(repository.GetPrimaryKeysById, "4752f49f-3011-44af-9c62-2a6f4086233d").Iter()
	for iter.Scan(&id, &profile_id, &timestamp) {
		fmt.Println(id + profile_id + timestamp.String())
	}

	err = session.Query(repository.AddLikeToPost, 1, "4752f49f-3011-44af-9c62-2a6f4086233d", "e2b5f92e-c31b-11eb-8529-0242ac130003").Exec()
	if err != nil {
		fmt.Println(err)
	}

	err = session.Query(repository.InsertDislikeStatement, "4752f49f-3011-44af-9c62-2a6f4086233d", time.Now(), "43420055-3174-4c2a-9823-a8f060d644c3").Exec()

	if err != nil {
		fmt.Println(err)
	}
	err = session.Query(repository.AddDislikeToPost, 1, "4752f49f-3011-44af-9c62-2a6f4086233d", "e2b5f92e-c31b-11eb-8529-0242ac130003").Exec()
	if err != nil {
		fmt.Println(err)
	}


	err = session.Query(repository.InsertDislikeStatement, "d459e0f2-ab61-48e8-a593-29933ce99525", time.Now(), "424935b1-766c-4f99-b306-9263731518bc").Exec()
	iter = session.Query(repository.GetPrimaryKeysById, "d459e0f2-ab61-48e8-a593-29933ce99525").Iter()
	for iter.Scan(&id, &profile_id, &timestamp) {
		fmt.Println(id + profile_id + timestamp.String())
	}
	if err != nil {
		fmt.Println(err)
	}

	err = session.Query(repository.AddDislikeToPost, 1, "d459e0f2-ab61-48e8-a593-29933ce99525", "424935b1-766c-4f99-b306-9263731518bc").Exec()
	if err != nil {
		fmt.Println(err)
	}

}

func SeedPosts(session *gocql.Session) {

	hashtags := [2]string{"tbt", "idegasnamax"}
	media := [1]string{"stefan_smeker.jpg"}
	mentions := []string {""}
	timestamp := time.Now()
	var ifExists int
	session.Query(repository.SeeIfPostExists, "4752f49f-3011-44af-9c62-2a6f4086233d", "e2b5f92e-c31b-11eb-8529-0242ac130003").Iter().Scan(&ifExists)

	if ifExists == 0 {
		err := session.Query(repository.InsertIntoPostsTimestampTable, "4752f49f-3011-44af-9c62-2a6f4086233d", "e2b5f92e-c31b-11eb-8529-0242ac130003", timestamp).Exec()
		if err != nil {
			fmt.Println(err)
		}
	}
	err := session.Query(repository.InsertIntoPostTable, "4752f49f-3011-44af-9c62-2a6f4086233d", "e2b5f92e-c31b-11eb-8529-0242ac130003",
		"Da se podsetimo!", timestamp, 0, 0, 0, false, "KI", 0.0, 0.0, mentions, hashtags, media, "IMAGE", false).Exec()
	if err != nil {
		fmt.Println(err)
	}



	hashtag2 := [1]string{"idegasnamax"}
	media = [1]string{"shone.jpg"}

	timestamp = time.Now()
	session.Query(repository.SeeIfPostExists, "d459e0f2-ab61-48e8-a593-29933ce99525", "424935b1-766c-4f99-b306-9263731518bc").Iter().Scan(&ifExists)

	if ifExists == 0 {
		err = session.Query(repository.InsertIntoPostsTimestampTable, "d459e0f2-ab61-48e8-a593-29933ce99525", "424935b1-766c-4f99-b306-9263731518bc", timestamp).Exec()
	}
	err = session.Query(repository.InsertIntoPostTable, "d459e0f2-ab61-48e8-a593-29933ce99525", "424935b1-766c-4f99-b306-9263731518bc",
		"Bitno da je nekad bilo lepo...", timestamp, 0, 0, 0, false, "KI", 0.0, 0.0, mentions, hashtag2, media, "IMAGE", false).Exec()
	if err != nil {
		fmt.Println(err)
	}




	hashtag2 = [1]string{"idegasnamax"}
	media = [1]string{"pablo.jpg"}
	timestamp = time.Now()
	session.Query(repository.SeeIfPostExists, "1ea5b7bc-94eb-40c0-98fd-7858e197e3b2", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03").Iter().Scan(&ifExists)

	if ifExists == 0 {
		err = session.Query(repository.InsertIntoPostsTimestampTable, "1ea5b7bc-94eb-40c0-98fd-7858e197e3b2", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03", timestamp).Exec()
	}
	err = session.Query(repository.InsertIntoPostTable, "1ea5b7bc-94eb-40c0-98fd-7858e197e3b2", "a2c2f993-dc32-4a82-82ed-a5f6866f7d03",
		"Bitno da je nekad bilo lepo...", timestamp, 0, 0, 0, false, "SM", 0.0, 0.0, mentions, hashtag2, media, "IMAGE", false).Exec()
	if err != nil {
		fmt.Println(err)
	}




	hashtag2 = [1]string{"idegasnamax"}
	media = [1]string{"pablo.jpg"}

	timestamp = time.Now()
	session.Query(repository.SeeIfPostExists, "adfee6f4-fe45-40ad-8f8e-760ec861a35e", "43420055-3174-4c2a-9823-a8f060d644c3").Iter().Scan(&ifExists)

	if ifExists == 0 {
		err = session.Query(repository.InsertIntoPostsTimestampTable, "adfee6f4-fe45-40ad-8f8e-760ec861a35e", "43420055-3174-4c2a-9823-a8f060d644c3", timestamp).Exec()
	}

	err = session.Query(repository.InsertIntoPostTable, "adfee6f4-fe45-40ad-8f8e-760ec861a35e", "43420055-3174-4c2a-9823-a8f060d644c3",
		"Bitno da je nekad bilo lepo...", timestamp, 0, 0, 0, false, "NS", 0.0, 0.0, mentions, hashtag2, media, "IMAGE", false).Exec()
	if err != nil {
		fmt.Println(err)
	}



}
