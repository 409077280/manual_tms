package tmsctl

import (
	"log"
	"fmt"
	"database/sql"
)

func (pg *Pgconfig)GetPackageIdList(data []DataModel) {
	condition := ""
	packageid := sql.NullString{}

	for i := 0; i < len(data); i++ {
		if last := len(data) -1;i == last {
			condition = `'` + data[i].PackageId + `'`
		}
		condition = `'` + data[i].PackageId + `',`
	}
	db := pg.Connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT package_id FROM journey_last_miles WHERE tracking_number in (%s)", condition)
	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()


	for row.Next() {
		e := DataModel{}
		err = row.Scan(&packageid)
		if err != nil {
			log.Fatal(err)
		}
		e.PackageId = packageid.String
		data = append(data,e)
	}

}