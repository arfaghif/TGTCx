package service

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/arfaghif/TGTCx/backend/database"
	"github.com/arfaghif/TGTCx/backend/dictionary"
)

func UploadBanner(banner dictionary.Banner) (err error) {
	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// // construct query
	query := `
		INSERT INTO banners(name, description, image_path, start_date, end_date) VALUES($1, $2, $3, $4, $5) RETURNING id
	`
	err = tx.QueryRowContext(ctx, query, banner.Name, banner.Description, banner.ImgPath, banner.StartDate, banner.EndDate).Scan(&banner.ID)

	log.Println(query)

	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		log.Println("failed insert banner")
		return
	}

	query = `SELECT id FROM tags WHERE tag IN (`
	var tags []interface{}
	for i, tag := range banner.Tags {
		query += "$" + strconv.Itoa(i+1)
		tags = append(tags, tag)
		if i == len(banner.Tags)-1 {
			query += ")"
		} else {
			query += ","
		}
	}

	query2 := `INSERT INTO tags(tag) VALUES`

	// query2 := `INSERT INTO banner_tags(banner_id, tag_id) VALUES`
	// var tagIDs []interface{}

	for i := 1; i <= len(banner.Tags); i++ {
		query2 += "($" + strconv.Itoa(i) + ")"
		if i != len(banner.Tags) {
			query2 += ","
		}
	}

	// // read query result, and assign to variable(s)
	query2 += " ON CONFLICT DO NOTHING RETURNING id"
	log.Println(query2)
	rows, err := tx.QueryContext(ctx, query2, tags...)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(query)

	query2 = `INSERT INTO banner_tags(banner_id, tag_id) VALUES`
	var tagIDs []interface{}
	i := 1
	for rows.Next() {
		var tagID int
		if err = rows.Scan(&tagID); err != nil {
			return
		}
		query2 += "($" + strconv.Itoa(2*i-1) + ",$" + strconv.Itoa((2 * i)) + ")"
		i++
		tagIDs = append(tagIDs, banner.ID)
		tagIDs = append(tagIDs, tagID)
		if i == len(banner.Tags) {
			query2 += ","
		}
	}
	// if err != nil {
	// 	// Incase we find any error in the query execution, rollback the transaction

	// 	tx.Rollback()
	// 	log.Println(query)
	// 	log.Println("failed insert tag")
	// 	return
	// }
	log.Println(query2)
	_, err = tx.ExecContext(ctx, query2, tagIDs...)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		log.Println("failed insert banner tag")
		return
	}

	return tx.Commit()
}

func AddTagBanner(id int, tags []string) (err error) {
	// // you can connect and
	// // get current database connection
	db := database.GetDB()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query2 := `INSERT INTO tags(tag) VALUES`

	// query2 := `INSERT INTO banner_tags(banner_id, tag_id) VALUES`
	// var tagIDs []interface{}

	var queryArgs []interface{}

	for i := 1; i <= len(tags); i++ {
		query2 += "($" + strconv.Itoa(i) + ")"
		if i != len(tags) {
			query2 += ","
		}
		queryArgs = append(queryArgs, tags[i-1])
	}

	// // read query result, and assign to variable(s)
	query2 += " ON CONFLICT DO NOTHING RETURNING id"
	log.Println(query2)
	rows, err := tx.QueryContext(ctx, query2, queryArgs...)
	if err != nil {
		log.Println(err.Error())
	}

	query2 = `INSERT INTO banner_tags(banner_id, tag_id) VALUES`
	var tagIDs []interface{}
	i := 1
	for rows.Next() {
		var tagID int
		if err = rows.Scan(&tagID); err != nil {
			return
		}
		query2 += "($" + strconv.Itoa(2*i-1) + ",$" + strconv.Itoa((2 * i)) + ")"
		i++
		tagIDs = append(tagIDs, id)
		tagIDs = append(tagIDs, tagID)
		if i == len(tags) {
			query2 += ","
		}
	}
	// if err != nil {
	// 	// Incase we find any error in the query execution, rollback the transaction

	// 	tx.Rollback()
	// 	log.Println(query)
	// 	log.Println("failed insert tag")
	// 	return
	// }
	log.Println(query2)
	query2 += " ON CONFLICT DO NOTHING"
	log.Println(query2)
	_, err = tx.ExecContext(ctx, query2, tagIDs...)
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		log.Println("failed insert banner tag")
		return
	}

	return tx.Commit()
}

func UpdateBanner(id int, name string, description string, start_date time.Time, end_date time.Time) (err error) {
	db := database.GetDB()

	get_query := `
		SELECT name, description, start_date, end_date
		FROM banners
		WHERE id = $1
	`

	row := db.QueryRow(get_query, id)

	var banner dictionary.Banner
	// // read query result, and assign to variable(s)
	err = row.Scan(&banner.Name, &banner.Description, &banner.StartDate, &banner.EndDate)

	log.Println("MULA: ", banner.ID, banner.Name, banner.Description, banner.StartDate, banner.EndDate)

	if err != nil {
		log.Println(err.Error())
	}

	if name == "" {
		name = banner.Name
	}

	if description == "" {
		description = banner.Description
	}

	if start_date.IsZero() {
		start_date = banner.StartDate
	}

	if end_date.IsZero() {
		end_date = banner.EndDate
	}

	log.Println("Logging: ", id, name, description, start_date, end_date)

	query := `
		UPDATE banners
		SET name = $2, description = $3, start_date = $4, end_date = $5
		WHERE id = $1
	`

	_, err = db.Exec(query, id, name, description, start_date, end_date)

	row = db.QueryRow(get_query, id)
	err = row.Scan(&banner.Name, &banner.Description, &banner.StartDate, &banner.EndDate)

	log.Println(banner.Name, banner.Description, banner.StartDate, banner.EndDate)
	return nil
}
