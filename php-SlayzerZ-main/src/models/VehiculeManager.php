<?php
class VehiculeManager{
    private $db;
    private $faker;
    private $faker2;

    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
        $this->faker = Faker\Factory::create();
        $this->faker2 = Faker\Factory::create();
    }
    
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Vehicules` (
          `ID` int unsigned not null auto_increment primary key,
          `BrandID` int unsigned not null,
          `SeatID` int unsigned not null,
          `ColorID` int unsigned not null,
          `Image` varchar(255),
          `Location` varchar(255),
          `PriceJ` int,
          `isFree` bool not null,
          Foreign key (`BrandID`) REFERENCES Brand(`ID`),
          Foreign key (`SeatID`) REFERENCES Seat(`ID`),
          Foreign key (`ColorID`) REFERENCES Color(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID;");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $vehicule = new Vehicule();
                $vehicule->setID($row['ID']);
                $vehicule->setBrandID($row['BrandID']);
                $vehicule->setSeatID($row['SeatID']);
                $vehicule->setColorID($row['ColorID']);
                $vehicule->setImage($row['Image']);
                $vehicule->setLocation($row['Location']);
                $vehicule->setPriceJ($row['PriceJ']);
                $vehicule->setisFree($row['isFree']);
                $vehicule->setBrand($row['Brand']);
                $vehicule->setSeat($row['Seat']);
                $vehicule->setColor($row['Color']);
                $vehicules[] = $vehicule;
            };
            return $vehicules;
        }
        return null;
    }

    public function findOnebyID($id){
        $db = $this->db;
        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Vehicules.ID = $id");
        $result->execute();
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $vehicule = new Vehicule();
                $vehicule->setID($row['ID']);
                $vehicule->setBrandID($row['BrandID']);
                $vehicule->setSeatID($row['SeatID']);
                $vehicule->setColorID($row['ColorID']);
                $vehicule->setImage($row['Image']);
                $vehicule->setLocation($row['Location']);
                $vehicule->setPriceJ($row['PriceJ']);
                $vehicule->setisFree($row['isFree']);
                $vehicule->setBrand($row['Brand']);
                $vehicule->setSeat($row['Seat']);
                $vehicule->setColor($row['Color']);
            };
            return $vehicule;
    }

    public function insertOneAuto(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Vehicules` (
          `ID` int unsigned not null auto_increment primary key,
          `BrandID` int unsigned not null,
          `SeatID` int unsigned not null,
          `ColorID` int unsigned not null,
          `Image` varchar(255) not null,
          `Location` varchar(255) not null,
          `PriceJ` int not null,
          `isFree` bool not null,
          Foreign key (`BrandID`) REFERENCES Brand(`ID`),
          Foreign key (`SeatID`) REFERENCES Seat(`ID`),
          Foreign key (`ColorID`) REFERENCES Color(`ID`)
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $table2 = $db->prepare("CREATE table IF NOT EXISTS
        `Cities` (
          `ID` int unsigned not null auto_increment primary key,
          `City` varchar(255) not null)");
        $table2->execute();
        $faker = $this->faker;
        $faker2 = $this->faker2;
        $faker2->addProvider(new Faker\Provider\fr_FR\Address($faker2));
        $brandM = new BrandManager;
        $brand = $brandM->findAll();
        $seatM = new SeatManager;
        $seat = $seatM->findAll();
        $colorM = new ColorManager;
        $color = $colorM->findAll();
        $brandIDs = $faker->randomElement($brand);
        $seatIDs = $faker->randomElement($seat);
        $colorIDs = $faker->randomElement($color);
        $image = array([]);
        if($brandIDs->Name == "Nissan"){
            $image = array(["NissanB.png","NissanW.jpg","NissanR.jpg"]);
        } else if($brandIDs->Name == "Renault"){
            $image = array(["RenaultB.webp","RenaultG.jpg","RenaultR.jpg"]);
        }else if($brandIDs->Name == "Volvo"){
            $image = array(["VolvoB.jpg","VolvoW.jpg"]);
        }else if($brandIDs->Name == "Tesla"){
            $image = array(["TeslaW.jpg","TeslaR.webp"]);
        }else if($brandIDs->Name == "Fiat"){
            $image = array(["FiatB.jpg","FiatW.jpg","FiatR.jpg"]);
        }else if($brandIDs->Name == "Peugeot"){
            $image = array(["PeugeotB.jpg","PeugeotB2.jpg","PeugeotB3.jpg","PeugeotR.jpg"]);
        }else if($brandIDs->Name == "Volkswagen"){
            $image = array(["VolkG.png","VolkW.jpeg","VolkR.jpg"]);
        }else if($brandIDs->Name == "Ferrari"){
            $image = array(["FerrariB.jpg","FerrariG.jpg","FerrariR.png"]);
        }else if($brandIDs->Name == "Hyundai"){
            $image = array(["HyundaiB.png","HyundaiB2.jpg","HyundaiB3.jpg"]);
        }else if($brandIDs->Name == "Kia"){
            $image = array(["KiaB.webp","KiaB2.jpg","KiaW.jpg","PeugeotR.jpg"]);
        }
        $imageUrl = $faker->randomElement($image[0]);
        $price = $faker->numberBetween(20, 200);
        $loc = $faker2->city;
        $result = $db->prepare("INSERT INTO Vehicules(BrandID,SeatID,ColorID,Image,Location,PriceJ,isFree)
          Values($brandIDs->ID,$seatIDs->ID,$colorIDs->ID,'$imageUrl','$loc',$price,1);
          INSERT INTO Cities(City) VALUE('$loc')");
        $result->execute();
    }
    public function insertOneManual($brandid,$seatid,$colorid,$image,$location,$price){
        $db = $this->db;
        $table2 = $db->prepare("CREATE table IF NOT EXISTS
        `Cities` (
          `ID` int unsigned not null auto_increment primary key,
          `City` varchar(255) not null)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table2->execute();
        $result = $db->prepare("INSERT INTO Vehicules(BrandID,SeatID,ColorID,Image,Location,PriceJ,isFree)
          Values(
          $brandid,
          $seatid,
          $colorid,
          '$image',
          '$location',
          $price,
          1);
          INSERT INTO Cities(City) VALUE('$location')");
        $result->execute();
    }
    public function updateOne($id,$brandid,$seatid,$colorid,$image,$location,$price,$free){
        $db = $this->db;
        $result = $db->prepare("UPDATE Vehicules SET BrandID = $brandid, SeatID = $seatid, ColorID = $colorid,
        Image = '$image', Location = '$location', PriceJ = $price, isFree = $free WHERE ID = $id");
        $result->execute();
    }
    public function findAllbyLocation($location){
        $db = $this->db;
        if($location == "Tous"){
            $result = $db->prepare("SELECT Vehicules.*, 
            Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
               JOIN Brand ON Vehicules.BrandID = Brand.ID
               JOIN Seat ON Vehicules.SeatID = Seat.ID
               JOIN Color ON Vehicules.ColorID = Color.ID;");
        } else {
            $result = $db->prepare("SELECT Vehicules.*, 
            Brand.Brand, 
           Seat.Seat,
           Color.Color
           FROM Vehicules
          JOIN Brand ON Vehicules.BrandID = Brand.ID
          JOIN Seat ON Vehicules.SeatID = Seat.ID
          JOIN Color ON Vehicules.ColorID = Color.ID
          WHERE Location = '$location';");
        }
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $vehicule = new Vehicule();
                $vehicule->setID($row['ID']);
                $vehicule->setBrandID($row['BrandID']);
                $vehicule->setSeatID($row['SeatID']);
                $vehicule->setColorID($row['ColorID']);
                $vehicule->setImage($row['Image']);
                $vehicule->setLocation($row['Location']);
                $vehicule->setPriceJ($row['PriceJ']);
                $vehicule->setisFree($row['isFree']);
                $vehicule->setBrand($row['Brand']);
                $vehicule->setSeat($row['Seat']);
                $vehicule->setColor($row['Color']);
                $vehicules[] = $vehicule;
            };
            return $vehicules;
        }
        return null;
    }
    public function findAllbyFilters($brand,$seat,$color,$price){
        $db = $this->db;
        if($price == "Tous"){
            if($brand == "Tous"){
                if($seat == "Tous"){
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID;");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Color = '$color';");
                    }
                } else {
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Seat = '$seat';");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Seat = '$seat' AND Color = '$color';");
                    }
                }
            } else {
                if($seat == "Tous"){
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand';");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand' AND Color = '$color';");
                    }
                } else {
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
        Brand.Brand, 
        Seat.Seat,
        Color.Color
        FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand' AND Seat = '$seat';");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
                        Brand.Brand, 
                       Seat.Seat,
                       Color.Color
                       FROM Vehicules
                      JOIN Brand ON Vehicules.BrandID = Brand.ID
                      JOIN Seat ON Vehicules.SeatID = Seat.ID
                      JOIN Color ON Vehicules.ColorID = Color.ID
                      WHERE Brand = '$brand' AND Seat = '$seat' AND Color = '$color';");
                    }
                }
            }        
        } else {
            if($brand == "Tous"){
                if($seat == "Tous"){
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE PriceJ <= $price;");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Color = '$color' AND PriceJ <= $price;");
                    }
                } else {
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Seat = '$seat' AND PriceJ <= $price;");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Seat = '$seat' AND Color = '$color' AND PriceJ <= $price;");
                    }
                }
            } else {
                if($seat == "Tous"){
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand' AND PriceJ <= $price;");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand' AND Color = '$color' AND PriceJ <= $price;");
                    }
                } else {
                    if($color == "Tous"){
                        $result = $db->prepare("SELECT Vehicules.*, 
             Brand.Brand, 
            Seat.Seat,
            Color.Color
            FROM Vehicules
           JOIN Brand ON Vehicules.BrandID = Brand.ID
           JOIN Seat ON Vehicules.SeatID = Seat.ID
           JOIN Color ON Vehicules.ColorID = Color.ID
           WHERE Brand = '$brand' AND Seat = '$seat' AND PriceJ <= $price;");
                    } else {
                        $result = $db->prepare("SELECT Vehicules.*, 
                        Brand.Brand, 
                       Seat.Seat,
                       Color.Color
                       FROM Vehicules
                      JOIN Brand ON Vehicules.BrandID = Brand.ID
                      JOIN Seat ON Vehicules.SeatID = Seat.ID
                      JOIN Color ON Vehicules.ColorID = Color.ID
                      WHERE Brand = '$brand' AND Seat = '$seat' AND Color = '$color' AND PriceJ <= $price;");
                    }
                }
            }
            
        }
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $vehicule = new Vehicule();
                $vehicule->setID($row['ID']);
                $vehicule->setBrandID($row['BrandID']);
                $vehicule->setSeatID($row['SeatID']);
                $vehicule->setColorID($row['ColorID']);
                $vehicule->setImage($row['Image']);
                $vehicule->setLocation($row['Location']);
                $vehicule->setPriceJ($row['PriceJ']);
                $vehicule->setisFree($row['isFree']);
                $vehicule->setBrand($row['Brand']);
                $vehicule->setSeat($row['Seat']);
                $vehicule->setColor($row['Color']);
                $vehicules[] = $vehicule;
            };
            return $vehicules;
        }
        return null;
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Favoris WHERE VehiculeID = $id;
        DELETE FROM FeedBack WHERE VehiculeID = $id;
        DELETE FROM Vehicules WHERE ID = $id;");
        $result->execute();
    }
}
?>