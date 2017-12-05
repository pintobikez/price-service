package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gen "github.com/pintobikez/price-service/api/structures"
	cnfs "github.com/pintobikez/price-service/config/structures"
	"log"
	"strconv"
	"time"
)

const (
	IsEmpty = "% is empty"
)

type Client struct {
	config *cnfs.DatabaseConfig
	db     *sql.DB
}

func New(cnfg *cnfs.DatabaseConfig) (*Client, error) {
	if cnfg == nil {
		return nil, fmt.Errorf("Client configuration not loaded")
	}

	return &Client{config: cnfg}, nil
}

//Connect to the mysql database
func (r *Client) Connect() error {

	urlString, err := r.buildStringConnection()
	if err != nil {
		return err
	}

	r.db, err = sql.Open("mysql", urlString)
	if err != nil {
		return err
	}
	return nil
}

//Disconnect from the mysql database
func (r *Client) Disconnect() {
	r.db.Close()
}

func (r *Client) GetChannels() (map[string]int64, error) {

	ret := make(map[string]int64)

	rows, err := r.db.Query("SELECT id, name FROM channels")

	if err != nil {
		return ret, err
	}

	for rows.Next() {
		var name_ string
		var id_ int64

		err = rows.Scan(&id_, &name_)
		if err != nil {
			return ret, fmt.Errorf("Error reading rows: %s", err.Error())
		}

		ret[name_] = id_
	}
	rows.Close()

	return ret, nil
}

//FindProduct Product value and Retrives an ProductResponse
func (r *Client) FindProduct(ID string) (*gen.Product, error) {

	// TODO
	resp := new(gen.Product)

	q := "SELECT p.id as id,c.name as channel,price,special_price,special_from_date,special_to_date,updated_at,c.id as channel_id"
	q += " FROM price p INNER JOIN channels c on c.id=p.fk_channel WHERE p.id=?"

	rows, err := r.db.Query(q, ID)

	if err != nil {
		return resp, err
	}

	//var arr []*gen.ProductPrices

	for rows.Next() {
		var (
			id_ string
			p   float64
			sp  float64
			spf time.Time
			spt time.Time
			ch  string
			upd time.Time
			cid int64
		)
		log.Println("fasdas")
		err = rows.Scan(&id_, &ch, &p, &sp, &spf, &spt, &upd, &cid)
		if err != nil {
			log.Println("%s", err.Error())
			return resp, fmt.Errorf("Error reading rows: %s", err.Error())
		}

		aux := new(gen.ProductPrices)
		aux.Price = p
		aux.SpecialPrice = sp
		aux.SpecialFrom = spf
		aux.SpecialTo = spt
		aux.Channel = ch
		aux.UpdatedAt = upd
		aux.ChannelID = cid
		log.Printf("%v\n", aux)
		resp.Prices = append(resp.Prices, aux)

		resp.ID = id_
	}

	rows.Close()

	if resp.ID == "" {
		return resp, fmt.Errorf("%s not found", ID)
	}

	return resp, nil
}

// Updates the given Product
func (r *Client) PutProduct(s *gen.Product) (int64, error) {

	var af int64 = 0
	q := "INSERT INTO price (id,fk_channel,price,special_price,special_from_date,special_to_date)"
	q += " VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE price=?, special_price=?,special_from_date=?,special_to_date=?"

	stmt, err := r.db.Prepare(q)
	if err != nil {
		return 0, fmt.Errorf("Error in prepared statement: %s", err.Error())
	}

	for _, p := range s.Prices {
		res, err := stmt.Exec(s.ID, p.ChannelID, p.Price, p.SpecialPrice, p.SpecialFrom, p.SpecialTo, p.Price, p.SpecialPrice, p.SpecialFrom, p.SpecialTo)
		if err != nil {
			stmt.Close()
			return 0, fmt.Errorf("Could not insert/update Product %s", s.ID)
		}
		aff, err := res.RowsAffected()
		if err != nil {
			return 0, fmt.Errorf("Could not insert/update Product %s", s.ID)
		}
		stmt.Close()
		af += aff
	}

	return af, nil
}

// Health Endpoint of the Client
func (r *Client) Health() error {

	str, err := r.buildStringConnection()
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", str)
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func (r *Client) buildStringConnection() (string, error) {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if r.config == nil {
		return "", fmt.Errorf("Client configuration not loaded")
	}
	if r.config.Driver.User == "" {
		return "", fmt.Errorf(IsEmpty, "User")
	}
	if r.config.Driver.Pw == "" {
		return "", fmt.Errorf(IsEmpty, "Password")
	}
	if r.config.Driver.Host == "" {
		return "", fmt.Errorf(IsEmpty, "Host")
	}
	if r.config.Driver.Port <= 0 {
		return "", fmt.Errorf(IsEmpty, "Port")
	}
	if r.config.Driver.Schema == "" {
		return "", fmt.Errorf(IsEmpty, "Schema")
	}

	stringConn := r.config.Driver.User + ":" + r.config.Driver.Pw
	stringConn += "@tcp(" + r.config.Driver.Host + ":" + strconv.Itoa(r.config.Driver.Port) + ")"
	stringConn += "/" + r.config.Driver.Schema + "?charset=utf8"

	return stringConn, nil
}
