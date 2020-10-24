DROP DATABASE IF EXISTS VPO;
CREATE DATABASE VPO;
USE VPO;

-- -----------------------------------------------------
-- Table `VPO`.`vpo_address`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_address` (
  `uuidAddress` VARCHAR(36) NOT NULL,
  `address` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`uuidAddress`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_baptismbook`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_baptismbook` (
  `idBook` INT(11) NOT NULL AUTO_INCREMENT,
  `numberBook` INT(11) NULL DEFAULT NULL,
  `startDate` DATE NULL DEFAULT NULL,
  `endDate` DATE NULL DEFAULT NULL,
  `commentary` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`idBook`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_folio`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_folio` (
  `idFolio` INT(11) NOT NULL AUTO_INCREMENT,
  `numberFolio` INT(11) NOT NULL,
  `idBook` INT(11) NOT NULL,
  PRIMARY KEY (`idFolio`),
  INDEX `VPO_BaptismBook` (`idBook` ASC) VISIBLE,
  CONSTRAINT `VPO_BaptismBook`
    FOREIGN KEY (`idBook`)
    REFERENCES `VPO`.`vpo_baptismbook` (`idBook`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_place`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_place` (
  `idPlace` VARCHAR(36) NOT NULL,
  `place` VARCHAR(500) NULL DEFAULT NULL,
  `description` VARCHAR(36) NULL DEFAULT NULL,
  PRIMARY KEY (`idPlace`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_priest`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_priest` (
  `idPriest` INT(11) NOT NULL AUTO_INCREMENT,
  `firstName` VARCHAR(30) NULL DEFAULT NULL,
  `secondName` VARCHAR(30) NULL DEFAULT NULL,
  `lastName` VARCHAR(30) NULL DEFAULT NULL,
  `secondLastName` VARCHAR(30) NULL DEFAULT NULL,
  `credentials` VARCHAR(20) NULL DEFAULT NULL,
  `parishOrigin` VARCHAR(100) NULL DEFAULT NULL,
  PRIMARY KEY (`idPriest`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_baptism`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_baptism` (
  `idBaptism` VARCHAR(36) NOT NULL,
  `numberBaptism` INT(11) NULL DEFAULT NULL,
  `idFolio` INT(11) NOT NULL,
  `baptismDate` DATE NOT NULL,
  `idPriest` INT(11) NOT NULL,
  `idPlace` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`idBaptism`),
  INDEX `VPO_Folio` (`idFolio` ASC) VISIBLE,
  INDEX `VPO_Place` (`idPlace` ASC) VISIBLE,
  INDEX `VPO_Priest` (`idPriest` ASC) VISIBLE,
  CONSTRAINT `VPO_Folio`
    FOREIGN KEY (`idFolio`)
    REFERENCES `VPO`.`vpo_folio` (`idFolio`),
  CONSTRAINT `VPO_Place`
    FOREIGN KEY (`idPlace`)
    REFERENCES `VPO`.`vpo_place` (`idPlace`),
  CONSTRAINT `VPO_Priest`
    FOREIGN KEY (`idPriest`)
    REFERENCES `VPO`.`vpo_priest` (`idPriest`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_userrole`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_userrole` (
  `uuidRol` INT(11) NOT NULL AUTO_INCREMENT,
  `typeRol` VARCHAR(20) NULL DEFAULT NULL,
  `description` VARCHAR(50) NULL DEFAULT NULL,
  PRIMARY KEY (`uuidRol`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_user` (
  `idUser` VARCHAR(36) NOT NULL,
  `userName` VARCHAR(50) NULL DEFAULT NULL,
  `password` VARCHAR(300) NULL DEFAULT NULL,
  `uuidRol` INT(11) NOT NULL,
  PRIMARY KEY (`idUser`),
  INDEX `VPO_UserRole` (`uuidRol` ASC) VISIBLE,
  CONSTRAINT `VPO_UserRole`
    FOREIGN KEY (`uuidRol`)
    REFERENCES `VPO`.`vpo_userrole` (`uuidRol`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_baptismhistory`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_baptismhistory` (
  `idBaptismHistory` INT(11) NOT NULL AUTO_INCREMENT,
  `dateEmitted` DATE NULL DEFAULT NULL,
  `idUser` VARCHAR(36) NOT NULL,
  `idBaptism` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`idBaptismHistory`),
  INDEX `VPO_Baptism` (`idBaptism` ASC) VISIBLE,
  INDEX `VPO_User` (`idUser` ASC) VISIBLE,
  CONSTRAINT `VPO_Baptism`
    FOREIGN KEY (`idBaptism`)
    REFERENCES `VPO`.`vpo_baptism` (`idBaptism`),
  CONSTRAINT `VPO_User`
    FOREIGN KEY (`idUser`)
    REFERENCES `VPO`.`vpo_user` (`idUser`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `VPO`.`vpo_person`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `VPO`.`vpo_person` (
  `idPerson` VARCHAR(36) NOT NULL,
  `firstName` VARCHAR(30) NULL DEFAULT NULL,
  `secondName` VARCHAR(30) NULL DEFAULT NULL,
  `lastName` VARCHAR(30) NULL DEFAULT NULL,
  `secondLastName` VARCHAR(30) NULL DEFAULT NULL,
  `bornDate` DATE NULL DEFAULT NULL,
  `DPI` VARCHAR(25) NULL DEFAULT NULL,
  `sex` CHAR(1) NULL DEFAULT NULL,
  `uuidAddress` VARCHAR(36) NOT NULL,
  `idFather` VARCHAR(36) NULL DEFAULT NULL,
  `idMother` VARCHAR(36) NULL DEFAULT NULL,
  `idGodFather` VARCHAR(36) NULL DEFAULT NULL,
  `idGodMother` VARCHAR(36) NULL DEFAULT NULL,
  `idManager` VARCHAR(36) NULL DEFAULT NULL,
  `idBaptism` VARCHAR(36) NOT NULL,
  PRIMARY KEY (`idPerson`),
  INDEX `VPO_Address` (`uuidAddress` ASC) VISIBLE,
  INDEX `VPO_Baptismv2` (`idBaptism` ASC) VISIBLE,
  INDEX `VPO_Person` (`idFather` ASC) VISIBLE,
  INDEX `VPO_Personv3` (`idMother` ASC) VISIBLE,
  INDEX `VPO_Personv4` (`idGodFather` ASC) VISIBLE,
  INDEX `VPO_Personv5` (`idGodMother` ASC) VISIBLE,
  INDEX `VPO_Personv6` (`idManager` ASC) VISIBLE,
  CONSTRAINT `VPO_Address`
    FOREIGN KEY (`uuidAddress`)
    REFERENCES `VPO`.`vpo_address` (`uuidAddress`),
  CONSTRAINT `VPO_Baptismv2`
    FOREIGN KEY (`idBaptism`)
    REFERENCES `VPO`.`vpo_baptism` (`idBaptism`),
  CONSTRAINT `VPO_Person`
    FOREIGN KEY (`idFather`)
    REFERENCES `VPO`.`vpo_person` (`idPerson`),
  CONSTRAINT `VPO_Personv3`
    FOREIGN KEY (`idMother`)
    REFERENCES `VPO`.`vpo_person` (`idPerson`),
  CONSTRAINT `VPO_Personv4`
    FOREIGN KEY (`idGodFather`)
    REFERENCES `VPO`.`vpo_person` (`idPerson`),
  CONSTRAINT `VPO_Personv5`
    FOREIGN KEY (`idGodMother`)
    REFERENCES `VPO`.`vpo_person` (`idPerson`),
  CONSTRAINT `VPO_Personv6`
    FOREIGN KEY (`idManager`)
    REFERENCES `VPO`.`vpo_person` (`idPerson`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;