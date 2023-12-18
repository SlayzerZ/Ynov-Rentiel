<?php 
class BookingManager{
    private $db;
    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function createOne($UserID,$VehiculeID,$start,$end,$diff){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Booking` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` INT UNSIGNED not null,
          `VehiculeID` INT UNSIGNED not null,
          `Start` DATETIME not null,
          `End` DATETIME not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $manager = new VehiculeManager();
        $vechicule = $manager->findOnebyID($VehiculeID);
        $result = $db->prepare("INSERT INTO Booking(UserID,VehiculeID,Start,End) 
        Values($UserID,$VehiculeID,'$start','$end');");
        // var_dump($result);
        try{
            $result->execute();
            $block = $db->prepare("UPDATE Vehicules
            SET isFree = 0
            WHERE ID = $VehiculeID");
            $block->execute();
            $price = $vechicule->PriceJ * $diff;
            return "La voiture est booker vous aurez à payer $price €";
        } catch(Exception $e){
            return $e->getMessage();
        }
    }
    public function findAllbyUserID($Userid){
        $db = $this->db;
        $result = $db->prepare("SELECT Booking.*, Vehicules.*,Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Booking
           JOIN Vehicules ON Booking.VehiculeID = Vehicules.ID
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE UserID = $Userid;");
        $result->execute();
        if($result->rowCount()>0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $book = new Booking();
                $book->setID($row['ID']);
                $book->setUserID($row['UserID']);
                $book->setVehiculeID($row['VehiculeID']);

                $heo = new DateTime($row['Start']);
                $book->setStart($heo->format("Y-m-d"));
                $heo2 = new DateTime($row['End']);
                $book->setEnd($heo2->format("Y-m-d"));

                $book->setBrandID($row['BrandID']);
                $book->setSeatID($row['SeatID']);
                $book->setColorID($row['ColorID']);
                $book->setImage($row['Image']);
                $book->setLocation($row['Location']);

                $interval = $heo->diff($heo2);
                $book->setPriceT($row['PriceJ'] * $interval->format("%R%a"));
                $book->setisFree($row['isFree']);
                $book->setBrand($row['Brand']);
                $book->setSeat($row['Seat']);
                $book->setColor($row['Color']);
                $books[] = $book;
            };
            return $books;
        }
            
    }
    public function checkFeedback($Userid,$VehiculeID){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Booking WHERE UserID = $Userid AND VehiculeID = $VehiculeID;");
        $result->execute();
        if($result->rowCount()>0){
            return 1;
        } else {
            return 0;
        }
    }
    public function restore($End){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Booking` (
          `ID` int unsigned not null auto_increment primary key,
          `UserID` INT UNSIGNED not null,
          `VehiculeID` INT UNSIGNED not null,
          `Start` DATETIME not null,
          `End` DATETIME not null,
          Foreign key (`UserID`) REFERENCES Users(`ID`),
          Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT * FROM Booking WHERE DATE(End) = '$End'");
        $result->execute();
        if($result->rowCount()>0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $vec = $row['VehiculeID'];
                $exec = $db->prepare("UPDATE Vehicules SET isFree = 1 WHERE ID = $vec");
                $exec->execute();
            }
        }
    }
}
?>