<?php
class CityManager{
    private $db;

    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Cities` (
          `ID` int unsigned not null auto_increment primary key,
          `City` Varchar(255) not null
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT DISTINCT * FROM Cities");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $City = new City();
                $City->setID($row['ID']);
                $City->setName($row['City']);
                $Citys[] = $City;
            };
            return $Citys;
        }
    }

    public function findOnebyID($id){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Cities WHERE ID = $id");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            $row = $result->fetch(PDO::FETCH_ASSOC);
            return $row['City'];
        } else {
            return null;
        }
    }
    public function insertOne($City){
        $db = $this->db;
        $result = $db->prepare("INSERT INTO Cities(City) VALUE('$City')");
        $result->execute();
    }
    public function updateOne($City,$id){
        $db = $this->db;
        $result = $db->prepare("UPDATE Cities SET City = '$City' WHERE ID = $id");
        $result->execute();
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Cities WHERE ID = $id");
        $result->execute();
    }
}
?>