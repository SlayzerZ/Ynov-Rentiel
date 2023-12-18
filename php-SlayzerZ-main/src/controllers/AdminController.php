<?php 
session_start();
class AdminController
{
    public function showPanel()
    {
        if($_SESSION['ID'] == null){
            header("Location: /sign");
        }
        $manager = new UserManager();
        $user = $manager->findOnebyID($_SESSION['ID']);
        if($user->getAdmin() && isset($_SESSION['ID'])){
            $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Admin.twig', ['co' => $_SESSION['ID']]);
        } else {
            echo "Tu n'est pas Admin, Dégage Bouffon!";
        }
    }
    public function showAccess()
    {
        if($_SESSION['ID'] == null){
            header("Location: /sign");
        }
        $manager = new UserManager();
        $user = $manager->findOnebyID($_SESSION['ID']);
        if($user->getAdmin() && isset($_SESSION['ID'])){
            $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        $type = $_GET['type'];
        $action = $_GET['Action'];
        if($action == "modify") {
            if($type == "Brand"){
                $manager = new BrandManager();
            } elseif($type == "Seat"){
                $manager = new SeatManager();
            } elseif($type == "Color"){
                $manager = new ColorManager();
            } elseif($type == "Users"){
                $manager = new UserManager();
            } elseif($type == "Vehicules"){
                $manager = new VehiculeManager();
            } elseif($type == "FeedBack"){
                $manager = new FeedbackManager();
            }
            $item = $manager->findAll();
        } else {
            $item = "";
        }
        echo $twig->render('Access.twig', ['type' => $type, 'action' => $action, 'items' => $item, 'co' => $_SESSION['ID']]);
        } else {
            echo "Tu n'est pas Admin, Dégage Bouffon!";
        }
    }
    public function createItem(){
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
        'debug' => true,
        // ...
    ]);
    $twig->addExtension(new \Twig\Extension\DebugExtension());
        $type = $_POST['type'];
        if($type == "Brand"){
            $manager = new BrandManager();
            $item = $_POST['brand'];
            $manager->insertOne($item);
        } elseif($type == "Seat"){
            $manager = new SeatManager();
            $item = $_POST['seat'];
            $manager->insertOne($item);
        } elseif($type == "Color"){
            $manager = new ColorManager();
            $item = $_POST['color'];
            $manager->insertOne($item);
        } elseif($type == "Users"){
            $manager = new UserManager();
            $id = $_POST['id'];
            $image = $_POST['image'];
            $firstname = $_POST['firstname'];
            $lastname = $_POST['lastname'];
            $email = $_POST['email'];
            $phone = $_POST['phonenumber'];
            $password = $_POST['password'];
            if(empty($id) || empty($firstname) || empty($lastname) || empty($email) || empty($phone) || empty($password)){
                echo "Hormis Image, Faut Remplir tous les champs!";
            } else {
                $password2 = password_hash($password, PASSWORD_DEFAULT);
                try{
                    $manager->insertOne($id,$image,$firstname,$lastname,$phone,$email,$password2);
                } catch(Exception $e){
                    echo $e->getMessage();
                }
            }
        } elseif($type == "Vehicules"){
            $manager = new VehiculeManager();
            $brandid = $_POST['brandid'];
            $seatid = $_POST['seatid'];
            $colorid = $_POST['colorid'];
            $image = $_POST['image'];
            $location = $_POST['location'];
            $price = $_POST['price'];
            try{
                $manager->insertOneManual($brandid,$seatid,$colorid,$image,$location,$price);
            } catch(Exception $e){
                echo $e->getMessage();
            }
        } elseif($type == "FeedBack"){
            $manager = new FeedbackManager();
            $userid = $_POST['userid'];
            $vehiculeid = $_POST['vehiculeid'];
            $note = $_POST['note'];
            $text = $_POST['text'];

            $texts = explode("'",$text);
            if(count($texts)-1 > 0){
                foreach(explode("'",$text,-1) as $value){
                    $tab[] = $value . "''";
                }
                $text = $tab[0];
                for($c = 1; $c < count($tab); $c++){
                    $text = $text . $tab[$c];
                }
                $text = $text . $texts[count($texts)-1];
            }

            try{
                $manager->addOne($userid,$vehiculeid,$note,$text);
            } catch(Exception $e){
                echo $e->getMessage();
            }
        }
        echo $twig->render('Admin.twig', ['co' => $_SESSION['ID']]);
    }

    public function modifyItem(){
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        $type = $_POST['type'];
        $id = $_POST['id'];
        if($type == "Brand"){
            $manager = new BrandManager();
            $item = $_POST['brand'];
            $manager->updateOne($item,$id);
        } 
        elseif($type == "Seat"){
            $manager = new SeatManager();
            $item = $_POST['seat'];
            $manager->updateOne($item,$id);
        } 
        elseif($type == "Color"){
            $manager = new ColorManager();
            $item = $_POST['color'];
            $manager->updateOne($item,$id);
        } 
        elseif($type == "Users"){
            $manager = new UserManager();
            $image = $_POST['image'];
            $firstname = $_POST['firstname'];
            $lastname = $_POST['lastname'];
            $email = $_POST['email'];
            $phone = $_POST['phonenumber'];
            $password = $_POST['password'];
            if(isset($_POST['admin'])){
                $admin = 1;
            } else {
                $admin = 0;
            }
            if($password == ""){
                try{
                    $manager->updateOne($id,$image,$firstname,$lastname,$phone,$email,$admin);
                } catch(Exception $e){
                    print_r($e->getMessage());
                }  
            } else {
                $password2 = password_hash($password, PASSWORD_DEFAULT);
                try{
                    $manager->updateOneP($id,$image,$firstname,$lastname,$phone,$email,$password2,$admin);
                } catch(Exception $e){
                    print_r($e->getMessage());
                }  
            }
        }
        elseif($type == "Vehicules"){
            $manager = new VehiculeManager();
            $brandid = $_POST['brandid'];
            $seatid = $_POST['seatid'];
            $colorid = $_POST['colorid'];
            $image = $_POST['image'];
            $location = $_POST['location'];
            $price = $_POST['price'];
            if(isset($_POST['free'])){
                $free = 1;
            } else {
                $free = 0;
            }
            try{
                $manager->updateOne($id,$brandid,$seatid,$colorid,$image,$location,$price,$free);
            } catch(Exception $e){
                echo $e->getMessage();
            }
        } 
        elseif($type == "FeedBack"){
            $manager = new FeedbackManager();
            $userid = $_POST['userid'];
            $vehiculeid = $_POST['vehiculeid'];
            $note = $_POST['note'];
            $text = $_POST['text'];

            $texts = explode("'",$text);
            if(count($texts)-1 > 0){
                foreach(explode("'",$text,-1) as $value){
                    $tab[] = $value . "''";
                }
                $text = $tab[0];
                for($c = 1; $c < count($tab); $c++){
                    $text = $text . $tab[$c];
                }
                $text = $text . $texts[count($texts)-1];
            }

            try{
                $manager->updateOne($id,$userid,$vehiculeid,$note,$text);
            } catch(Exception $e){
                echo $e->getMessage();
            }
        }
        echo $twig->render('Admin.twig', ['co' => $_SESSION['ID']]);
    }
    public function deleteItem(){
        $type = $_GET['type'];
        $id = $_GET['id'];
        if($type == "Brand"){
            $manager = new BrandManager();
            $manager->deleteOne($id);
        } elseif($type == "Seat"){
            $manager = new SeatManager();
            $manager->deleteOne($id);
        } elseif($type == "Color"){
            $manager = new ColorManager();
            $manager->deleteOne($id);
        } elseif($type == "Users"){
            $manager = new UserManager();
            $manager->deleteOne($id);
        } elseif($type == "Vehicules"){
            $manager = new VehiculeManager();
            $manager->deleteOne($id);
        } elseif($type == "FeedBack"){
            $manager = new FeedbackManager();
            $manager->deleteOne($id);
        }
        header("Location: /adminPanel");
    }
    public function truncateTable(){
        $type = $_GET['type'];
        $db = new PDO('mysql:host=172.20.27.212;dbname=my_database', "my_user", "my_password");
        if($type == "Vehicules"){
            $result = $db->prepare("TRUNCATE TABLE `FeedBack`;");
            $result->execute();
            $result = $db->prepare("TRUNCATE TABLE `Favoris`;");
            $result->execute();
            $result = $db->prepare("TRUNCATE TABLE `Booking`;");
            $result->execute();
            $result = $db->prepare("TRUNCATE TABLE `Cities`;");
            $result->execute();
            $result = $db->prepare("DELETE FROM `Vehicules`;");
            $result->execute();
        } else {
            $result = $db->prepare("TRUNCATE TABLE `$type`");
            $result->execute();
        }
        header("Location: /adminPanel");
    }
}
?>
