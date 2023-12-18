<?php
class BrandManager{
    private $db;

    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Brand` (
          `ID` int unsigned not null auto_increment primary key,
          `Brand` Varchar(255) not null
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT * FROM Brand");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $brand = new Brand();
                $brand->setID($row['ID']);
                $brand->setName($row['Brand']);
                $brands[] = $brand;
            };
            return $brands;
        } else {
            $Brands = array(["Nissan","Renault","Volvo","Tesla","Fiat","Peugeot","Volkswagen","Ferrari","Hyundai","Kia"]);
            foreach($Brands[0] as $b){
                $insert = $db->prepare("INSERT INTO Brand(Brand) VALUE('$b')");
                $insert->execute();
            }
            $result = $db->prepare("SELECT * FROM Brand");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $brand = new Brand();
                $brand->setID($row['ID']);
                $brand->setName($row['Brand']);
                $brands[] = $brand;
            };
            return $brands;
        }
        }
    }

    public function findOnebyID($id){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Brand WHERE ID = $id");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            $row = $result->fetch(PDO::FETCH_ASSOC);
            return $row['Brand'];
        } else {
            return null;
        }
    }
    public function insertOne($brand){
        $db = $this->db;
        $result = $db->prepare("INSERT INTO Brand(Brand) VALUE('$brand')");
        $result->execute();
    }
    public function updateOne($brand,$id){
        $db = $this->db;
        $result = $db->prepare("UPDATE Brand SET Brand = '$brand' WHERE ID = $id");
        $result->execute();
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Brand WHERE ID = $id");
        $result->execute();
    }
}
?>