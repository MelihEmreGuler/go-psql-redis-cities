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
		rowsAffected, err := r.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		if rowsAffected == 1 {
			fmt.Println(city.Name + " inserted to database")
		}
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

// Update city in database table cities
func (repo cityRepo) Update(city entity.City) *entity.City {
	stmt, err := repo.db.Prepare("update cities set name = $1, code = $2 where id = $3")
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		r, err := stmt.Exec(city.Name, city.Code, city.Id)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			rowsAffected, err := r.RowsAffected()
			if err != nil {
				fmt.Println(err)
				return nil
			}
			if rowsAffected != 1 {
				fmt.Println("No city updated in database")
				return nil
			} else {
				fmt.Println(city.Name + " updated in database")
				return &city
			}
		}
	}
}

// Delete city from database table cities
func (repo cityRepo) Delete(city entity.City) {
	// We need to print city name before deleting it from database
	cityName := repo.GetById(city.Id).Name

	stmt, err := repo.db.Prepare("delete from cities where id = $1")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		r, err := stmt.Exec(city.Id)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			rowsAffected, err := r.RowsAffected()
			if err != nil {
				fmt.Println(err)
				return
			}
			if rowsAffected != 1 {
				fmt.Println("There is something wrong, no city deleted from database")
				return
			} else {
				fmt.Println(cityName, "deleted from database")
				return
			}
		}
	}
}
