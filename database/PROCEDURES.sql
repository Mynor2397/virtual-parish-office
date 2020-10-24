DROP PROCEDURE IF EXISTS registerperson;

delimiter $ CREATE PROCEDURE registerperson(
    in _ID VARCHAR(36),
    in _Firstname VARCHAR(100),
    in _Secondname VARCHAR(100),
    in _Lastname VARCHAR(100),
    in _Secondlastname VARCHAR(100),
    in _DPI VARCHAR(100),
    in _Sexo VARCHAR(1),
    in _idAddress VARCHAR(36),
    in _Address VARCHAR(100)
) BEGIN DECLARE EXIT HANDLER FOR SQLEXCEPTION BEGIN SHOW ERRORS
LIMIT
    1;

ROLLBACK;

END;

DECLARE EXIT HANDLER FOR SQLWARNING BEGIN SHOW WARNINGS
LIMIT
    1;

ROLLBACK;

END;

START TRANSACTION;

INSERT INTO
    VPO_Address (uuidAddress, address)
VALUES
    (_idAddress, _Address);

INSERT INTO
    VPO_Person (
        idPerson,
        firstName,
        secondName,
        lastName,
        secondLastName,
        DPI,
        sex,
        uuidAddress
    )
VALUES
    (
        _ID,
        _Firstname,
        _Secondname,
        _Lastname,
        _Secondlastname,
        _DPI,
        _Sexo,
        _idAddress
    );

COMMIT;

END $ ----------
--
--
DROP PROCEDURE IF EXISTS registbaptism;

delimiter $ CREATE PROCEDURE registbaptism(
    in _idAddress VARCHAR(36),
    in _Address VARCHAR(100),
    --
    in _idBaptism VARCHAR(36),
    in _idFolio INT(11),
    in _baptismDate date,
    in _idPriest INT(11),
    in _idPlace VARCHAR(36),
    --
    in _ID VARCHAR(36),
    in _Firstname VARCHAR(100),
    in _Secondname VARCHAR(100),
    in _Lastname VARCHAR(100),
    in _Secondlastname VARCHAR(100),
    in _BornDate Date,
    in _DPI VARCHAR(100),
    in _Sexo VARCHAR(1),
    --
    in _idFather VARCHAR(36),
    in _idMother VARCHAR(36),
    in _idGodFather VARCHAR(36),
    in _idGodMother VARCHAR(36),
    in _idManager VARCHAR(36)
) BEGIN DECLARE _count INT;

DECLARE EXIT HANDLER FOR SQLEXCEPTION BEGIN SHOW ERRORS
LIMIT
    1;

ROLLBACK;

END;

DECLARE EXIT HANDLER FOR SQLWARNING BEGIN SHOW WARNINGS
LIMIT
    1;

ROLLBACK;

END;

START TRANSACTION;

SELECT
    COUNT(*) + 1 AS COUNT INTO _count
FROM VPO_Baptism;

INSERT INTO
    VPO_Address (uuidAddress, address)
VALUES
    (_idAddress, _Address);

INSERT INTO
    VPO_Baptism(
        idBaptism,
        numberBaptism,
        idFolio,
        baptismDate,
        idPriest,
        idPlace
    )
VALUES
    (
        _idBaptism,
        _count,
        _idFolio,
        _baptismDate,
        _idPriest,
        _idPlace
    );

INSERT INTO
    VPO_Person (
        idPerson,
        firstName,
        secondName,
        lastName,
        secondLastName,
        bornDate,
        DPI,
        sex,
        uuidAddress,
        idFather,
        idMother,
        idGodFather,
        idGodMother,
        idManager,
        idBaptism
    )
VALUES
    (
        _ID,
        _Firstname,
        _Secondname,
        _Lastname,
        _Secondlastname,
        _BornDate,
        _DPI,
        _Sexo,
        _idAddress,
        _idFather,
        _idMother,
        _idGodFather,
        _idGodMother,
        _idManager,
        _idBaptism
    );

COMMIT;

END $ ----------

DROP PROCEDURE IF EXISTS baptism_delete;

delimiter $
CREATE PROCEDURE baptism_delete(
    in _ID VARCHAR(36)
) BEGIN DECLARE  uuidaddresss VARCHAR(36);
DECLARE EXIT HANDLER FOR SQLEXCEPTION BEGIN SHOW ERRORS
LIMIT
    1;

ROLLBACK;

END;

DECLARE EXIT HANDLER FOR SQLWARNING BEGIN SHOW WARNINGS
LIMIT
    1;

ROLLBACK;

END;

START TRANSACTION;
SELECT
    uuidAddress INTO uuidaddresss
FROM VPO_Person WHERE idBaptism = _ID;

DELETE FROM VPO_Person WHERE idBaptism = _ID;
DELETE FROM VPO_Address WHERE uuidAddress = uuidaddresss;

COMMIT;

END $ ----------