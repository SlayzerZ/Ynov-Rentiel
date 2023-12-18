<?php
class SeatManager{
    private $db;

    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE table IF NOT EXISTS
        `Seat` (
          `ID` int unsigned not null auto_increment primary key,
          `Seat` int not null
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT * FROM Seat");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $seat = new Seat();
                $seat->setID($row['ID']);
                $seat->setSeat($row['Seat']);
                $seats[] = $seat;
            };
            return $seats;
        } else {
            $Seat = array([2,3,4,5,7,9]);
            foreach($Seat[0] as $b){
                $insert = $db->prepare("INSERT INTO Seat(Seat) VALUE('$b')");
                $insert->execute();
            }
            $result = $db->prepare("SELECT * FROM Seat");
            $result->execute();
            $count = $result->rowCount();
            if($count > 0){
                while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                    $seat = new Seat();
                    $seat->setID($row['ID']);
                    $seat->setSeat($row['Seat']);
                    $seats[] = $seat;
                };
                return $seats;
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
            return $row['Seat'];
        } else {
            return null;
        }
    }
    public function insertOne($seat){
        $db = $this->db;
        $result = $db->prepare("INSERT INTO Seat(Seat) VALUE($seat)");
        $result->execute();
    }
    public function updateOne($seat,$id){
        $db = $this->db;
        $result = $db->prepare("UPDATE Seat SET Seat = '$seat WHERE ID = $id");
        $result->execute();
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Seat WHERE ID = $id");
        $result->execute();
    }
}
?>