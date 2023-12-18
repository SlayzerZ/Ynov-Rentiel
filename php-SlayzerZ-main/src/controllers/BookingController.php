<?php
session_start(); 
class BookingController{
public function showBooking(){
    if(isset($_GET['id'])){
    $id = $_GET['id'];
    $manager = new VehiculeManager();
    $Vehicule = $manager->findOnebyID($id);
    if($Vehicule->isFree){
        $currentDateTime = new DateTime('now, +1 days'); 
        $currentDate = $currentDateTime->format('Y-m-d');
        for ($n = 2; $n < 40; $n++) {
            $nextDateTime = new DateTime('now, +'.$n.' days'); 
            $nextDate = $nextDateTime->format('Y-m-d');
            $nextDates[] = $nextDate;
        }
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
                'debug' => true,
                // ...✅❌
            ]);
            $twig->addExtension(new \Twig\Extension\DebugExtension());
            echo $twig->render('Booking.twig', ['vehicule' => $Vehicule, 'currentDate' => $currentDate, 'nextDates' => $nextDates, 'co' => $_SESSION['ID']]);
    }
    } else {
        header("Location: /home");
    }
}
public function haveBook(){
    if($_SERVER["REQUEST_METHOD"] == "POST"){
        $VecID = $_POST['myCar'];
        $Start = $_POST['Start'];
        $End = $_POST['End'];
        $Start_date = new DateTime($Start);
        $End_date = new DateTime($End);
        $interval = $Start_date->diff($End_date);
        if($interval->format("%R%a") > 0) {
        $manager = new BookingManager();
        $booker = $manager->createOne($_SESSION['ID'],$VecID,$Start_date->format('Y-m-d H:i:s'),$End_date->format('Y-m-d H:i:s'),$interval->format("%R%a"));
        echo $booker;
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...✅❌
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Home.twig', ['co' => $_SESSION['ID']]);
    } else {
        echo "Make Sense Bitch!!";
    }
        }
}
}
?>