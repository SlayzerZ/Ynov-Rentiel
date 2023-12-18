<?php 
class UserManager{
    private $db;
    public function __construct()
    {
        $this->db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        $this->db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_WARNING);
    }
    public function findAll(){
        $db = $this->db;
        $table = $db->prepare("CREATE TABLE IF NOT EXISTS
        `Users` (
          `ID` int unsigned not null primary key,
          `Image` VARCHAR(255),
          `FirstName` VARCHAR(255),
          `LastName` varchar(255),
          `PhoneNumber` varchar(255),
          `Email` varchar(255)UNIQUE,
          `MDP` varchar(255),
          `Admin` bool
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("SELECT * FROM Users;");
        $result->execute();
        $count = $result->rowCount();
        if($count > 0){
            while ($row = $result->fetch(PDO::FETCH_ASSOC)){
                $user = new User();
                $user->setID($row['ID']);
                $user->setImage($row['Image']);
                $user->setFirstName($row['FirstName']);
                $user->setLastName($row['LastName']);
                $user->setPhoneNumber($row['PhoneNumber']);
                $user->setEmail($row['Email']);
                $user->setPassword($row['MDP']);
                $user->setAdmin($row['Admin']);
                $users[] = $user;
            };
            return $users;
        }
        return null;
    }
    public function insertOne($id,$image,$FirstName,$LastName,$PhoneNumber,$Email,$Password){
        $test = "";
        $db = $this->db;
        $table = $db->prepare("CREATE TABLE IF NOT EXISTS
        `Users` (
          `ID` int unsigned not null primary key,
          `Image` VARCHAR(255),
          `FirstName` VARCHAR(255),
          `LastName` varchar(255),
          `PhoneNumber` varchar(255),
          `Email` varchar(255)UNIQUE,
          `MDP` varchar(255),
          `Admin` bool
        )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci");
        $table->execute();
        $result = $db->prepare("INSERT INTO Users(ID,Image,FirstName,LastName,PhoneNumber,Email,MDP,Admin)
          Values(
        $id,
        '$image',
        '$FirstName',
        '$LastName',
        '$PhoneNumber',
        '$Email',
        '$Password',
        0)");
        try{
        $result->execute();
        $test = "success";
        } catch (Exception $e){
            echo $e->getMessage();
            $test = "fail";
        }
        return $test;
    }
    public function findOnebyConnection($email,$password){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Users WHERE Email = '$email' OR PhoneNumber = '$email'");
        $result->execute();
        if($result->rowCount() > 0){
            $data = $result->fetchAll();
            if(password_verify($password, $data[0]['MDP'])){
                $_SESSION['ID'] = $data[0]['ID'];
                return "true";
            } else {
                $_SESSION['ID'] = Null;
                return "false";
            }
        } else {
            echo "no User found";
        }
    }
    public function findOnebyID($id){
        $db = $this->db;
        $result = $db->prepare("SELECT * FROM Users WHERE ID = $id");
        $result->execute();
        if($result->rowCount() > 0){
            $data = $result->fetchAll();
            $User = new User();
            $User->setID($data[0]['ID']);
            $User->setImage($data[0]['Image']);
            $User->setFirstName($data[0]['FirstName']);
            $User->setLastName($data[0]['LastName']);
            $User->setPhoneNumber($data[0]['PhoneNumber']);
            $User->setEmail($data[0]['Email']);
            $User->setPassword($data[0]['MDP']);
            $User->setAdmin($data[0]['Admin']);
            return $User;
        } 
    }
    public function updateOneP($id,$image,$FirstName,$LastName,$PhoneNumber,$email,$password,$admin){
        $db = $this->db;
        $result = $db->prepare("UPDATE Users SET Image = '$image', FirstName = '$FirstName', LastName = '$LastName',
        PhoneNumber = '$PhoneNumber', Email = '$email', MDP = '$password', Admin = $admin WHERE ID = $id");
        $result->execute();
    }
    public function updateOne($id,$image,$FirstName,$LastName,$PhoneNumber,$email,$admin){
        $db = $this->db;
        $result = $db->prepare("UPDATE Users SET Image = '$image', FirstName = '$FirstName', LastName = '$LastName',
        PhoneNumber = '$PhoneNumber', Email = '$email', Admin = $admin WHERE ID = $id");
        $result->execute();
    }
    public function deleteOne($id){
        $db = $this->db;
        $result = $db->prepare("DELETE FROM Users WHERE ID = $id");
        $result->execute();
    }
    public function firstAdmin(){
        $db = $this->db;
        $Password = password_hash("ratchet", PASSWORD_DEFAULT);
        $result = $db->prepare("INSERT INTO Users(ID,Image,FirstName,LastName,PhoneNumber,Email,MDP,Admin)
          Values(9000,'sonic.png','Ibrahim','Ari Malla','+33625957705','ibrahimallari84@gmail.com','$Password',1)");
        $result->execute();
    }
}
?>