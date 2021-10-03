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
	db := database.GetDB()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	query := `
		INSERT INTO banners(name, description, image_path, start_date, end_date) VALUES($1, $2, $3, $4, $5) RETURNING id
	`
	err = tx.QueryRowContext(ctx, query, banner.Name, banner.Description, banner.ImgPath, banner.StartDate, banner.EndDate).Scan(&banner.ID)

	if err != nil {
		tx.Rollback()
		log.Println("failed insert banner")
		return
	}

	var tags []interface{}

	query2 := `INSERT INTO tags(tag) VALUES`

	for i := 1; i <= len(banner.Tags); i++ {
		query2 += "($" + strconv.Itoa(i) + ")"
		if i != len(banner.Tags) {
			query2 += ","
		}
		tags = append(tags, banner.Tags[i-1])
	}

	query2 += " ON CONFLICT DO NOTHING RETURNING id"
	rows, err := tx.QueryContext(ctx, query2, tags...)
	if err != nil {
		log.Println("failed insert tags")
		log.Println(query2)
		return
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
		tagIDs = append(tagIDs, banner.ID)
		tagIDs = append(tagIDs, tagID)
		if i == len(banner.Tags) {
			query2 += ","
		}
	}

	_, err = tx.ExecContext(ctx, query2, tagIDs...)
	if err != nil {
		tx.Rollback()
		log.Println("failed insert banner tag")
		return
	}

	return tx.Commit()
}

func AddTagBanner(id int, tags []string) (banner dictionary.Banner, err error) {
	db := database.GetDB()

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return

	}

	query := `SELECT id, name, description, image_path, start_date, end_date FROM banners WHERE id = $1`
	err = tx.QueryRowContext(ctx, query, id).Scan(&banner.ID, &banner.Name, &banner.Description, &banner.ImgPath, &banner.StartDate, &banner.EndDate)
	if err != nil {
		tx.Rollback()
		log.Println("failed select banner")
		return
	}

	query2 := `INSERT INTO tags(tag) VALUES`

	var queryArgs []interface{}

	for i := 1; i <= len(tags); i++ {
		query2 += "($" + strconv.Itoa(i) + ")"
		if i != len(tags) {
			query2 += ","
		}
		queryArgs = append(queryArgs, tags[i-1])
	}

	query2 += " ON CONFLICT DO NOTHING RETURNING id"
	rows, err := tx.QueryContext(ctx, query2, queryArgs...)
	if err != nil {
		log.Println(err.Error())
		return
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

	query2 += " ON CONFLICT DO NOTHING"
	_, err = tx.ExecContext(ctx, query2, tagIDs...)
	if err != nil {
		tx.Rollback()
		log.Println("failed insert banner tag")
		return
	}

	return banner, tx.Commit()
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

	if err != nil {
		log.Println(err.Error())
	}

	if name == "" {
		name = banner.Name
	}

	if description == "" {
		description = banner.Description
	}

	if start_date.IsZero() || start_date.After(banner.StartDate) {
		start_date = banner.StartDate
	}

	if end_date.IsZero() {
		end_date = banner.EndDate
	}

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

func GetBannerUser(id int) (banners []dictionary.Banner, err error) {
	db := database.GetDB()

	query := `
	SELECT banners.id, banners.name, banners.description, banners.image_path, banners.start_date, banners.end_date 
	FROM banners NATURAL JOIN banner_tags, tags, tag_users, users
	WHERE users.id = $1
	GROUP BY banners.id
	ORDER BY COUNT(banners.id) DESC
	LIMIT 3
	`

	banners = []dictionary.Banner{}
	rows, err := db.Query(query, id)
	for rows.Next() {
		banner := dictionary.Banner{}
		if err = rows.Scan(&banner.ID, &banner.Name, &banner.Description, &banner.ImgPath, &banner.StartDate, &banner.EndDate); err != nil {
			return
		}
		banners = append(banners, banner)
	}

	return
}
