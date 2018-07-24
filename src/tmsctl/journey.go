package tmsctl

import (
	"log"
	"fmt"
	"database/sql"
)

func (pg *Pgconfig)GetPackageIdList(data []DataModel,list []string) {
	condition := ""
	packageid := sql.NullString{}
	trackNb := sql.NullString{}

	for i := 0; i < len(list); i++ {
		if last := len(list) -1;i == last {
			condition = `'` + list[i] + `'`
		}
		condition = `'` + list[i] + `',`
	}
	db := pg.connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT package_id FROM journey_last_miles WHERE tracking_number in (%s)", condition)
	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()


	for row.Next() {
		e := DataModel{}
		err = row.Scan(&packageid, &trackNb)
		if err != nil {
			log.Fatal(err)
		}
		e.PackageId = packageid.String
		e.TrackNumber = trackNb.String
		data = append(data,e)
	}

}