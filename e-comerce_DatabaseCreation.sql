-- tableau des Images
CREATE TABLE IF NOT EXISTS 'Picture'
(
	IDPicture	INTEGER PRIMARY KEY NOT NULL,
	Name		VARCHAR(80) NOT NULL,
	Size		INTEGER NOT NULL,
	Type		VARCHAR(10)NOT NULL,
	Bin			LongBLOB NOT NULL
);

-- tableau des Utilisateurs
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

-- tableau des Addresse
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

-- tableau des Produits
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

-- tableau du Pannier
CREATE TABLE IF NOT EXISTS 'cart_table'
(
	IDCart	INTEGER NOT NULL,
	IdUser	INT NOT NULL,
	PRIMARY KEY (IdCart),
	FOREIGN KEY	(IdUser) REFERENCES User(IdUser)
);

-- tableau des commandes
CREATE TABLE IF NOT EXISTS 'command_table'
(
	IDCart		INT NOT NULL,
	IdProduct	INT NOT NULL,
	SameProduct	INTEGER NOT NULL DEFAULT '0',
	FOREIGN KEY (IDCart) REFERENCES cart_table(IDCart),
	FOREIGN KEY (IdProduct) REFERENCES Product(IDProduct)
);

-- tableau des Facturations
CREATE TABLE IF NOT EXISTS 'invoices_table'
(
	IDFacturation	INTEGER NOT NULL,
	IDCart			INT NOT NULL,
	DateFacturation	Date NOT NULL,
	Facturation		INTEGER NOT NULL,
	NomberOfProduct INTEGER NOT NULL,
	PRIMARY KEY (IDFacturation),
	FOREIGN KEY (IDCart) REFERENCES cart_table(IDCart)
);