package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"

	"github.com/bxcodec/faker/v3"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var SqlStmt = `
CREATE TABLE IF NOT EXISTS 'Picture'
(
	IDPicture	INTEGER PRIMARY KEY NOT NULL,
	Name		VARCHAR(80) NOT NULL,
	Size		INTEGER NOT NULL,
	Type		VARCHAR(10)NOT NULL,
	Bin			LongBLOB NOT NULL
);

CREATE TABLE IF NOT EXISTS 'User'
(
	IDUser		INTEGER PRIMARY KEY NOT NULL,
	LastName	VARCHAR(60) NOT NULL,
	FirstName	VARCHAR(60) NOT NULL,
	Gender		VARCHAR(5) DEFAULT '-',
	BirthDate	Date NOT NULL,
	Email		VARCHAR(255) NOT NULL,
	Mdp	VARCHAR(80) NOT NULL,
	IdPicture	INT	REFERENCES Picture(IDPicture)
);

CREATE TABLE IF NOT EXISTS 'Address'
(
	IDAddress	INTEGER NOT NULL,
	IDUser		INTEGER NOT NULL,
	Country		VARCHAR(42) NOT NULL,
	Town		VARCHAR(63) NOT NULL,
	PostalCode	VARCHAR(15) NOT NULL,
	Addresse	VARCHAR(80) NOT NULL,
	PRIMARY KEY (IDAddress),
	FOREIGN KEY (IDUser) REFERENCES User(IDUser)
);

CREATE TABLE IF NOT EXISTS 'Product'
(
	IDProduct	INTEGER NOT NULL,
	Name		VARCHAR(30) NOT NULL,
	Description	VARCHAR(200) NOT NULL,
	Price		INTEGER NOT NULL,
	IdPicture	INT	NOT NULL,
	PRIMARY KEY (IDProduct),
	FOREIGN KEY (IdPicture) REFERENCES Picture(IDPicture)
);

CREATE TABLE IF NOT EXISTS 'cart_table'
(
	IDCart	INTEGER NOT NULL,
	IdUser	INT NOT NULL,
	PRIMARY KEY (IdCart),
	FOREIGN KEY	(IdUser) REFERENCES User(IdUser)
);

CREATE TABLE IF NOT EXISTS 'command_table'
(
	IDCart		INT NOT NULL,
	IdProduct	INT NOT NULL,
	SameProduct	INTEGER NOT NULL DEFAULT '0',
	FOREIGN KEY (IDCart) REFERENCES cart_table(IDCart),
	FOREIGN KEY (IdProduct) REFERENCES Product(IDProduct)
);

CREATE TABLE IF NOT EXISTS 'invoices_table'
(
	IDFacturation	INTEGER NOT NULL,
	IDCart			INT NOT NULL,
	DateFacturation	Date NOT NULL,
	Facturation		INTEGER NOT NULL,
	NomberOfProduct INTEGER NOT NULL,
	PRIMARY KEY (IDFacturation),
	FOREIGN KEY (IDCart) REFERENCES cart_table(IDCart)
);`
var dB, _ = sql.Open("sqlite3", "e-comerce.db")

func main() {
	fmt.Println("Debut du script")
	insertPicture()
	InitDatabase()
	insertUser()
	insertAddress()
	insertProduct()
	insertCartTable()
	insertCommandTable()
	insertinvoicesTable()
	defer dB.Close()
	defer fmt.Println("FIN")
}

func InitDatabase() {
	dB.Exec(SqlStmt)

}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func insertUser() {
	gender := "Femme"
	if rand.Intn(100) <= 50 {
		gender = "Homme"
	}
	dB.Exec("INSERT OR IGNORE INTO 'User' (IDUser, LastName, FirstName, Gender, BirthDate, Email, Mdp) VALUES (1, 'BOB', 'JACK', 'Homme', '28062003', 'Bob@gmail.com', '" + HashPassword("123") + "')")
	dB.Exec("INSERT OR IGNORE INTO 'User' (IDUser, LastName, FirstName, Gender, BirthDate, Email, Mdp) VALUES (2, '" + faker.LastName() + "', '" + faker.LastName() + "', '" + gender + "', '" + faker.Date() + "', '" + faker.Email() + "', '" + HashPassword(faker.Password()) + "')")
}

func insertPicture() {
	dB.Exec("INSERT OR IGNORE INTO 'Picture' (IDPicture, Name, Size, Type, Bin) VALUES (1, 'default', 5, 'png', '10011')")
	file, _ := os.Open("./profil.png")
	var data []byte
	count,_ := file.Read(data)
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// dB.Exec("INSERT OR IGNORE INTO 'Picture' (IDPicture, Name, Size, Type, Bin) VALUES (2, 'profil', , 'png', '10011')")
}

func insertAddress() {
	dB.Exec("INSERT OR IGNORE INTO 'Address' (IDAddress, IDUser, Country, Town, PostalCode, Addresse) VALUES (1, 1, 'France', 'Colombes', '92700', '12 rue Aviouon')")
	dB.Exec("INSERT OR IGNORE INTO 'Address' (IDAddress, IDUser, Country, Town, PostalCode, Addresse) VALUES (2, 2, 'France', 'Colombes', '92700', '13 rue Aviouon')")
}

func insertProduct() {
	dB.Exec("INSERT OR IGNORE INTO 'Product' (IDProduct, Name, Description, Price, IdPicture) VALUES (1, 'Feuille blanche', 'Une feuilli blanche rien deçu vendu a lunite', 1, 1)")
}

func insertCartTable() {
	dB.Exec("INSERT OR IGNORE INTO 'cart_table' (IDCart, IdUser) VALUES (1, 1)")
	dB.Exec("INSERT OR IGNORE INTO 'cart_table' (IDCart, IdUser) VALUES (2, 2)")
	dB.Exec("INSERT OR IGNORE INTO 'cart_table' (IDCart, IdUser) VALUES (3, 1)")
}

func insertCommandTable() {
	dB.Exec("INSERT OR IGNORE INTO 'command_table' (IDCart, IdProduct, SameProduct) VALUES (1, 1, 1)")
	dB.Exec("INSERT OR IGNORE INTO 'command_table' (IDCart, IdProduct, SameProduct) VALUES (2, 1, 3)")
	dB.Exec("INSERT OR IGNORE INTO 'command_table' (IDCart, IdProduct, SameProduct) VALUES (3, 1, 2)")
}

func insertinvoicesTable() {
	dB.Exec("INSERT OR IGNORE INTO 'invoices_table' (IDFacturation, IDCart, DateFacturation, Facturation, NomberOfProduct) VALUES (1, 1, '" + faker.Date() + "', 1, 1)")
}
