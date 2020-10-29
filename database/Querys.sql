CREATE 
VIEW `batismforpdf` AS
    SELECT 
        `bap`.`idBaptism` AS `idBaptism`,
        `bap`.`numberBaptism` AS `numberBaptism`,
        `fol`.`numberFolio` AS `numberFolio`,
        `book`.`numberBook` AS `numberBook`,
        `bau`.`bornDate` AS `borndDate`,
        `bap`.`baptismDate` AS `baptismDate`,
        `bau`.`firstName` AS `firstName`,
        `bau`.`secondName` AS `secondName`,
        `bau`.`lastName` AS `lastname`,
        `bau`.`secondLastName` AS `secondLastname`,
        `bau`.`sex` AS `sexo`,
        `fat`.`firstName` AS `firstNameFather`,
        `fat`.`secondName` AS `secondnameFather`,
        `fat`.`lastName` AS `lastnameFather`,
        `fat`.`secondLastName` AS `secondlastnameFather`,
        `mot`.`firstName` AS `firstNameMother`,
        `mot`.`secondName` AS `secondnameMother`,
        `mot`.`lastName` AS `lastnameMother`,
        `mot`.`secondLastName` AS `secondlastnameMother`,
        `gfat`.`firstName` AS `firstNameGodfather`,
        `gfat`.`secondName` AS `secondnameGodfather`,
        `gfat`.`lastName` AS `lastnameGodfather`,
        `gfat`.`secondLastName` AS `secondlastnameGodfather`,
        `gmot`.`firstName` AS `firstnameGodmother`,
        `gmot`.`secondName` AS `secondnameGodmother`,
        `gmot`.`lastName` AS `lastnameGodmother`,
        `gmot`.`secondLastName` AS `secondlastnameGodmother`,
        `mang`.`firstName` AS `firstnameManager`,
        `mang`.`secondName` AS `secondnameManager`,
        `mang`.`lastName` AS `lastnameManager`,
        `mang`.`secondLastName` AS `secondlastnameManager`,
        `pri`.`firstName` AS `firstnamePriest`,
        `pri`.`secondName` AS `secondnamePriest`,
        `pri`.`lastName` AS `lastnamePriest`,
        `pri`.`secondLastName` AS `secondlastnamePriest`
    FROM
        (((((((((`vpo_baptism` `bap`
        JOIN `vpo_folio` `fol` ON ((`fol`.`idFolio` = `bap`.`idFolio`)))
        JOIN `vpo_baptismbook` `book` ON ((`book`.`idBook` = `fol`.`idBook`)))
        JOIN `vpo_priest` `pri` ON ((`pri`.`idPriest` = `bap`.`idPriest`)))
        JOIN `vpo_person` `bau` ON ((`bau`.`idBaptism` = `bap`.`idBaptism`)))
        LEFT JOIN `vpo_person` `fat` ON ((`fat`.`idPerson` = `bau`.`idFather`)))
        LEFT JOIN `vpo_person` `mot` ON ((`mot`.`idPerson` = `bau`.`idMother`)))
        LEFT JOIN `vpo_person` `gfat` ON ((`gfat`.`idPerson` = `bau`.`idGodFather`)))
        LEFT JOIN `vpo_person` `gmot` ON ((`gmot`.`idPerson` = `bau`.`idGodMother`)))
        LEFT JOIN `vpo_person` `mang` ON ((`mang`.`idPerson` = `bau`.`idManager`)))


--
--
--

select bap.idBaptism, per.firstName, per.secondName, per.lastName, per.secondLastname from VPO_Baptism bap
INNER JOIN VPO_Person per on per.idBaptism = bap.idBaptism
WHERE baptismDate BETWEEN DATE_FORMAT ('2015-01-24', '%Y-%m-%d') AND DATE_FORMAT ('2015-01-25','%Y-%m-%d');
