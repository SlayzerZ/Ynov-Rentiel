<?php 
class FavorisManager{
    private $db;
    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAllbyUserID($Userid){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Favoris` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` int unsigned not null,
          `VehiculeID` int unsigned not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT Favoris.*, Vehicules.*,Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Favoris
           JOIN Vehicules ON Favoris.VehiculeID = Vehicules.ID
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE UserID = $Userid;");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $fav = new Favoris();
                $fav->setID($row['ID']);
                $fav->setUserID($row['UserID']);
                $fav->setVehiculeID($row['VehiculeID']);
                $fav->setBrandID($row['BrandID']);
                $fav->setSeatID($row['SeatID']);
                $fav->setColorID($row['ColorID']);
                $fav->setImage($row['Image']);
                $fav->setLocation($row['Location']);
                $fav->setPriceJ($row['PriceJ']);
                $fav->setisFree($row['isFree']);
                $fav->setBrand($row['Brand']);
                $fav->setSeat($row['Seat']);
                $fav->setColor($row['Color']);
                $Favs[] = $fav;
            };
            return $Favs;
        }
        return null;
    }
    public function findAllbyID($Userid,$VehiculeID){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Favoris` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` int unsigned not null,
          `VehiculeID` int unsigned not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT Favoris.*, Vehicules.*,Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Favoris
           JOIN Vehicules ON Favoris.VehiculeID = Vehicules.ID
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE UserID = $Userid AND VehiculeID = $VehiculeID;");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $fav = new Favoris();
                $fav->setID($row['ID']);
                $fav->setUserID($row['UserID']);
                $fav->setVehiculeID($row['VehiculeID']);
                $fav->setBrandID($row['BrandID']);
                $fav->setSeatID($row['SeatID']);
                $fav->setColorID($row['ColorID']);
                $fav->setImage($row['Image']);
                $fav->setLocation($row['Location']);
                $fav->setPriceJ($row['PriceJ']);
                $fav->setisFree($row['isFree']);
                $fav->setBrand($row['Brand']);
                $fav->setSeat($row['Seat']);
                $fav->setColor($row['Color']);
                $Favs[] = $fav;
            };
            return $Favs;
        }
        return null;
    }
    public function insertOne($UserID,$VehiculeID){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Favoris` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` int unsigned not null,
          `VehiculeID` int unsigned not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("INSERT INTO Favoris(UserID,VehiculeID) Values($UserID,$VehiculeID);");
        $result->execute();
    }
    public function deleteOne($UserID,$VehiculeID){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Favoris` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` int unsigned not null,
          `VehiculeID` int unsigned not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )");
        $table->execute();
        $result = $db->prepare("DELETE FROM Favoris WHERE UserID = $UserID AND VehiculeID = $VehiculeID;");
        $result->execute();
    }
}
?>