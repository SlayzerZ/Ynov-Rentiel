<?php 
class FeedbackManager{
    private $db;
    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
public function findAll(){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )");
    $table->execute();
    $result = $db->prepare("SELECT FeedBack.ID as 'fib', FeedBack.*, Users.*, Vehicules.*, Brand.Brand,
    Seat.Seat,
    Color.Color FROM FeedBack
        JOIN Users ON FeedBack.UserID = Users.ID
        JOIN Vehicules ON FeedBack.VehiculeID = Vehicules.ID
        JOIN Brand ON Vehicules.BrandID = Brand.ID
        JOIN Seat ON Vehicules.SeatID = Seat.ID
        JOIN Color ON Vehicules.ColorID = Color.ID");
    $result->execute();
    if($result->rowCount() > 0){
        while ($row = $result->fetch(PDO::FETCH_ASSOC)){
            $feed = new Feedback();
            $feed->setID($row['fib']);
            $feed->setUserID($row['UserID']);
            $feed->setVehiculeID($row['VehiculeID']);
            $feed->setNote($row['Note']);
            $feed->setText($row['Text']);
            $feed->setFirstName($row['FirstName']);
            $feed->setLastName($row['LastName']);
            $feed->setPhoneNumber($row['PhoneNumber']);
            $feed->setEmail($row['Email']);
            $feed->setPassword($row['MDP']);
            $feed->setAdmin($row['Admin']);
            $feed->setBrandID($row['BrandID']);
            $feed->setSeatID($row['SeatID']);
            $feed->setColorID($row['ColorID']);
            $feed->setImage($row['Image']);
            $feed->setLocation($row['Location']);
            $feed->setPriceJ($row['PriceJ']);
            $feed->setisFree($row['isFree']);
            $feed->setBrand($row['Brand']);
            $feed->setSeat($row['Seat']);
            $feed->setColor($row['Color']);
            $feeds[] = $feed;
        };
        return $feeds;
    }
    return null;
}
public function findAllbyUserID($id){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )");
    $table->execute();
    $result = $db->prepare("SELECT FeedBack.*,Vehicules.*, Brand.Brand,
    Seat.Seat,
    Color.Color FROM FeedBack
        JOIN Vehicules ON FeedBack.VehiculeID = Vehicules.ID
        JOIN Brand ON Vehicules.BrandID = Brand.ID
        JOIN Seat ON Vehicules.SeatID = Seat.ID
        JOIN Color ON Vehicules.ColorID = Color.ID WHERE UserID = $id");
    $result->execute();
    if($result->rowCount() > 0){
        while ($row = $result->fetch(PDO::FETCH_ASSOC)){
            $feed = new Feedback();
            $feed->setID($row['ID']);
            $feed->setUserID($row['UserID']);
            $feed->setVehiculeID($row['VehiculeID']);
            $feed->setNote($row['Note']);
            $feed->setText($row['Text']);
            $feed->setBrandID($row['BrandID']);
            $feed->setSeatID($row['SeatID']);
            $feed->setColorID($row['ColorID']);
            $feed->setImage($row['Image']);
            $feed->setLocation($row['Location']);
            $feed->setPriceJ($row['PriceJ']);
            $feed->setisFree($row['isFree']);
            $feed->setBrand($row['Brand']);
            $feed->setSeat($row['Seat']);
            $feed->setColor($row['Color']);
            $feeds[] = $feed;
        };
        return $feeds;
    }
    return null;
}
public function findAllbyVehiculeID($id){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )");
    $table->execute();
    $result = $db->prepare("SELECT FeedBack.*, Users.* FROM FeedBack JOIN Users ON FeedBack.UserID = Users.ID
    WHERE VehiculeID = $id");
    $result->execute();
    if($result->rowCount() > 0){
        while ($row = $result->fetch(PDO::FETCH_ASSOC)){
            $feed = new Feedback();
            $feed->setID($row['ID']);
            $feed->setUserID($row['UserID']);
            $feed->setVehiculeID($row['VehiculeID']);
            $feed->setNote($row['Note']);
            $feed->setText($row['Text']);
            $feed->setFirstName($row['FirstName']);
            $feed->setLastName($row['LastName']);
            $feed->setPhoneNumber($row['PhoneNumber']);
            $feed->setEmail($row['Email']);
            $feed->setPassword($row['MDP']);
            $feed->setAdmin($row['Admin']);
            $feeds[] = $feed;
        };
        return $feeds;
    }
    return null;
}
public function findAllbyIDs($Userid,$VehiculeID){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )");
    $table->execute();
    $result = $db->prepare("SELECT * FROM FeedBack WHERE UserID = $Userid AND VehiculeID = $VehiculeID");
    $result->execute();
    if($result->rowCount() > 0){
        return 1;
    } else {
        return 0;
    }
    return null;
}
public function addOne($UserID,$VehiculeID,$note,$text){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
    $table->execute();
    $result = $db->prepare("INSERT INTO FeedBack(UserID,VehiculeID,Note,Text) 
    Values($UserID,$VehiculeID,$note,'$text');");
    // var_dump($result);
    try{
        $result->execute();
        return "Avis Enregistrer avec success";
    } catch(Exception $e){
        return $e->getMessage();
    }
}
public function updateOne($id,$UserID,$VehiculeID,$note,$text){
    $db = $this->db;
    $result = $db->prepare("UPDATE FeedBack SET UserID = $UserID, VehiculeID = $VehiculeID, Note = $note, text = '$text' 
    WHERE ID = $id");
    $result->execute();
}
public function deleteOne($UserID){
    $db = $this->db;
    $table = $db->prepare("CREATE table IF NOT EXISTS
    `FeedBack` (
      `ID` int unsigned not null auto_increment primary key,
      `UserID` INT UNSIGNED not null,
      `VehiculeID` INT UNSIGNED not null,
      `Note` INT not null,
      `Text` varchar(255) not null,
      Foreign key (`UserID`) REFERENCES Users(`ID`),
      Foreign key (`VehiculeID`) REFERENCES Vehicules(`ID`)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
    $table->execute();
    $result = $db->prepare("DELETE FROM FeedBack WHERE ID = $UserID;");
    $result->execute();
}
}
?>