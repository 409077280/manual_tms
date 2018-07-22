package main

import (
//	"bytes"
	"database/sql"
//	"database/sql/driver"
	"fmt"
	_ "github.com/lib/pq"
//	"strconv"
//	"strings"
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

func main() {
	pgurl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "tms_bi_vn", "voh9ahceiDahR2ge", "10.117.0.5", "30947", "tms_db")

	db, err := sql.Open("postgres", pgurl)
	if err != nil {
		panic(fmt.Errorf("Connected database error:%v", err))
	}

	querySql := `SELECT package_id FROM journey_last_miles LIMIT 10`
	rows, err := db.Query(querySql)
	if err != nil {
		panic(fmt.Errorf("Select error:%v", err))
	}

	for rows.Next() {
		agent := &Agent{}
		//agent := &Agent{Coordinate: &Point{}}
		/*
		err = rows.Scan(&agent.Id,
			&agent.Name, &agent.Code,
			&agent.CS_NO, &agent.Channel_id,
			&agent.Address, agent.Coordinate)
		*/
		err = rows.Scan(&agent.Package_id)
		fmt.Println(agent, err)
	}
	/*
	var id int
	err = db.QueryRow("INSERT INTO t_agent (name, code, cs_no, address, coordinate) VALUES($1,$2,$3,$4,$5) RETURNING id",
		"test1", "123457", "2", "111", "(12,43)").Scan(&id)

	fmt.Println("id:", id, "err:", err)
	*/
}