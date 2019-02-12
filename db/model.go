//Пакет по работе с БД
package db

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//Структура БД
type Dostup101 struct {
	tableName      struct{} `sql:"dostup101"`
	ID             string   `sql:"id"`
	Diapazon       string   `sql:"diapazon"`
	Adres          string   `sql:"adres"`
	Infrastructura string   `sql:"infrastructura"`
	Prinadlegnost  string   `sql:"prinadlegnost"`
	Dostup         string   `sql:"dostup"`
	Kluch          string   `sql:"kluch"`
	Aparatnaa      string   `sql:"aparatnaa"`
	AD             string   `sql:"ad"`
	Contact        string   `sql:"contact"`
	WGS            string   `sql:"wgs"`
	Primechanie    string   `sql:"primechanie"`
}

//Функция создания таблицы в базе данных
func CreateMyTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Dostup101{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table Dostup, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table Dostup creatrd successfully.\n")
	return nil
}

func (pl *Dostup101) Save(db *pg.DB) error {

	insertErr := db.Insert(pl)
	if insertErr != nil {
		log.Printf("Error if you added new PL in DB, Reason: v/n", insertErr)
		return insertErr
	}
	log.Printf("new PL  %s insert in DB. \n", pl.ID)
	return nil
}

func (pl *Dostup101) SaveReturn(db *pg.DB) (*Dostup101, error) {

	InsertResult, insertErr := db.Model(pl).Returning("*").Insert()
	if insertErr != nil {
		log.Printf("Error insering pl in DB, Reason: v/n", insertErr)
		return nil, insertErr
	}
	log.Printf("new PL  %s insert in DB. \n", pl.ID)
	log.Printf("new PL  %s return new value. \n", InsertResult.RowsAffected)
	return pl, nil
}

func (pl *Dostup101) SaveMulti(db *pg.DB, items []*Dostup101) error {

	_, insertErr := db.Model(items[0], items[1]).Insert()
	if insertErr != nil {
		log.Printf("Error insert many new PL in DB, Reason: v/n", insertErr)
		return insertErr
	}
	log.Printf("Many new PL  %s insert in DB. \n")
	return nil
}

func (pl *Dostup101) DeletePl(db *pg.DB) error {

	//db.Delete(pl)
	_, deleteErr := db.Model(pl).Where("id = ?id").Delete()

	if deleteErr != nil {
		log.Printf("Error delete PL in DB, Reason: v/n", deleteErr)
		return deleteErr
	}
	log.Printf("delete PL  %s  in DB. \n", pl.ID)
	return nil
}

func (pl *Dostup101) UpdateDiapazon(db *pg.DB) error {

	tx, txErr := db.Begin()
	if txErr != nil {
		log.Printf("Error while openning tx, Reason: %v\n", txErr)
		return txErr
	}

	_, updateErr := tx.Model(pl).Set("diapazon = ?diapazon").Where("id = ?id").Update()

	if updateErr != nil {
		log.Printf("Error update PL in DB, Reason: v/n", updateErr)
		tx.Rollback()
		return updateErr
	}
	tx.Commit()
	log.Printf("update PL  %s  in DB. \n", pl.Diapazon)
	return nil
}

func (pl *Dostup101) GetPl(db *pg.DB) error {

	getErr := db.Select(pl)

	if getErr != nil {
		log.Printf("Error get PL in DB, Reason: v/n", getErr)
		return getErr
	}
	log.Printf("get PL  %s  in DB. \n", *pl)
	return nil
}
