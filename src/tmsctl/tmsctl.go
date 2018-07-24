package tmsctl

import (
	//	"bytes"
	"database/sql"
	//	"database/sql/driver"
	"fmt"
	_ "github.com/lib/pq"
)

/*
// Implements driver.Valuer interface
func (p *Point) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "(%f %f)", p.X, p.Y)
	return buf.Bytes(), nil
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v %v)", p.X, p.Y)
}
// Implements sql.Scanner interface
func (p *Point) Scan(val interface{}) (err error) {
	if bb, ok := val.([]uint8); ok {
		tmp := bb[1 : len(bb)-1]
		coors := strings.Split(string(tmp[:]), ",")
		if p.X, err = strconv.ParseFloat(coors[0], 64); err != nil {
			return err
		}
		if p.Y, err = strconv.ParseFloat(coors[1], 64); err != nil {
			return err
		}
	}
	return nil
}
*/

type Agent struct {
	Package_id string  `json:"package_id"`
	/*	Id             `json:"id"`
		Name       string `json:"name"`
		Code       string `json:"code"`
		CS_NO      string `json:"cs_no"`
		Channel_id int    `json:"channel_id"`
		Address    string `json:"address"`
		Coordinate *Point `json:"point"` */
}
/*
//  Agent's child object
type Point struct {
	X float64 `json:"lat"`
	Y float64 `json:"lng"`
}
*/
type Pgconfig struct {
	User	string
	Pass	string
	Host	string
	Port	string
	Dbname	string
	Sslmodel string
}

type DataModel struct {
	PackageId string
	TrackNumber string
}

func (pg *Pgconfig)connect() *sql.DB {

	pgurl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%S", pg.User, pg.Pass, pg.Host, pg.Port, pg.Sslmodel, )

	db, err := sql.Open("postgres", pgurl)
	if err != nil {
		panic(fmt.Errorf("Connected database error:%v", err))
	}
	return db
}