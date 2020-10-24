CREATE  VIEW batismforpdf AS
SELECT 
bap.idBaptism,
bap.numberBaptism, 
fol.numberFolio, 
book.numberBook,
bap.bornDate,
bap.baptismDate,  
bau.firstName,
bau.secondName,
fat.firstName as firstNameFather, 
fat.secondName as secondnameFather,
fat.lastName as lastnameFather,
fat.secondLastName as secondlastnameFather,
mot.firstName as firstNameMother, 
mot.secondName as secondnameMother,
mot.lastName as lastnameMother,
mot.secondLastName as secondlastnameMother,
gfat.firstName as firstNameGodfather, 
gfat.secondName as secondnameGodfather,
gfat.lastName as lastnameGodfather,
gfat.secondLastName as secondlastnameGodfather,
gmot.firstName as firstnameGodmother,
gmot.secondName as secondnameGodmother,
gmot.lastName as lastnameGodmother,
gmot.secondLastName as secondlastnameGodmother,
mang.firstName as firstnameManager,
mang.secondName as secondnameManager,
mang.lastName as lastnameManager,
mang.secondLastName as secondlastnameManager,
pri.firstName as firstnamePriest, 
pri.secondName as secondnamePriest,
pri.lastName as lastnamePriest,
pri.secondLastName as secondlastnamePriest
FROM vpo_baptism bap
INNER JOIN vpo_folio fol ON fol.idFolio = bap.idFolio
INNER JOIN vpo_baptismbook book ON book.idBook = fol.idBook
INNER JOIN vpo_priest pri ON pri.idPriest = bap.idPriest
INNER JOIN vpo_person bau ON bau.idBaptism = bap.idBaptism
LEFT JOIN vpo_person fat ON fat.idPerson = bau.idFather
LEFT JOIN VPO_Person mot ON mot.idPerson = bau.idMother
LEFT JOIN VPO_Person gfat ON gfat.idPerson = bau.idGodFather
LEFT JOIN VPO_Person gmot ON gmot.idPerson = bau.idGodMother
LEFT JOIN VPO_Person mang ON mang.idPerson = bau.idManager


--
--
--

select bap.idBaptism, per.firstName, per.secondName, per.lastName, per.secondLastname from VPO_Baptism bap
INNER JOIN VPO_Person per on per.idBaptism = bap.idBaptism
WHERE baptismDate BETWEEN DATE_FORMAT ('2015-01-24', '%Y-%m-%d') AND DATE_FORMAT ('2015-01-25','%Y-%m-%d');
