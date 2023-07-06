package repository

import (
	"database/sql"
	"fmt"
	"github.com/MelihEmreGuler/go-psql-redis-cities/entity"
)

type cityRepo struct {
	db *sql.DB
}

var CityRepo *cityRepo // Singleton pattern (only one instance of this struct)

func NewRepo(db *sql.DB) {
	CityRepo = &cityRepo{
		db: db,
	}
}

/*
	Create -> Insert
	Read   -> Select
	Update -> Update
	Delete -> Delete
*/

// Insert city to database table cities (name, code)
func (repo cityRepo) Insert(city entity.City) {
	stmt, err := repo.db.Prepare("insert into cities (name, code) values ($1, $2)") // Prepare returns sql.Stmt (we need to execute it in stmt.Exec)

	r, err := stmt.Exec(city.Name, city.Code) // Exec returns sql.Result
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(r.RowsAffected()) // Returns number of rows affected by the query
	}
}

// List Select city from database table cities
func (repo cityRepo) List() []entity.City {
	var cityList []entity.City

	rows, err := repo.db.Query("select id, name, code from cities") // Query returns sql.Rows
	if err != nil {
		fmt.Println(err)
		return cityList
	} else {
		for rows.Next() {
			var city entity.City
			err = rows.Scan(&city.Id, &city.Name, &city.Code) // Scan returns error
			if err != nil {
				fmt.Println(err)
			}
			cityList = append(cityList, city)
		}
		err = rows.Close()
		if err != nil {
			fmt.Println(err)
		}
		return cityList
	}
}

// GetById get city from database table cities
func (repo cityRepo) GetById(id int) *entity.City {
	//careful, this is a vulnerability (SQL injection)
	var city entity.City

	formattedSql := fmt.Sprintf("select id, name, code from cities where id = %d", id)

	err := repo.db.QueryRow(formattedSql).Scan(&city.Id, &city.Name, &city.Code) // QueryRow returns sql.Row
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return &city
	}
}

// GetByName get city from database table cities
func (repo cityRepo) GetByName(name string) *entity.City {
	stmt, err := repo.db.Prepare("select id, name, code from cities where name = $1")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		var city entity.City
		err = stmt.QueryRow(name).Scan(&city.Id, &city.Name, &city.Code)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			return &city
		}
	}

}
