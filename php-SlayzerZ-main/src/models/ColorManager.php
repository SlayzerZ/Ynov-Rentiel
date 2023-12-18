<?php
class ColorManager{
    private $db;

    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Color` (
          `ID` int unsigned not null auto_increment primary key,
          `Color` Varchar(255) not null
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT * FROM Color");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $color = new Color();
                $color->setID($row['ID']);
                $color->setName($row['Color']);
                $colors[] = $color;
            };
            return $colors;
        } else {
            $Color = array(["Rouge","Blanc","Gris","Noir","Noir-Mat","Bleu-turquoise","Bleu-Marine"]);
            foreach($Color[0] as $b){
                $insert = $db->prepare("INSERT INTO Color(Color) VALUE('$b')");
                $insert->execute();
            }
            $result = $db->prepare("SELECT * FROM Color");
            $result->execute();
            $count = $result->rowCount();
            if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $color = new Color();
                $color->setID($row['ID']);
                $color->setName($row['Color']);
                $colors[] = $color;
            };
            return $colors;
        }
        }
    }
    public function findOnebyID($id){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Color WHERE ID = $id");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            $row = $result->fetch(PDO::FETCH_ASSOC);
            return $row['Name'];
        } else {
            return null;
        }
    }
    public function insertOne($color){
        $db = $this->db;
        $result = $db->prepare("INSERT INTO Color(Colro) VALUE('$color')");
        $result->execute();
    }
    public function updateOne($color,$id){
        $db = $this->db;
        $result = $db->prepare("UPDATE Color SET Color = '$color' WHERE ID = $id");
        $result->execute();
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Color WHERE ID = $id");
        $result->execute();
    }
}
?>