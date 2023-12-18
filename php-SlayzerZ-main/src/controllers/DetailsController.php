<?php
session_start();
class DetailsController
{
    public function showDetails()
    {   
        if($_SESSION['ID'] == null){
            header('Location: /sign');
        }
        $id = $_GET['id'];
        $manager = new VehiculeManager();
        $Vehicule = $manager->findOnebyID($id);
        $manager2 = new FavorisManager();
        $Fav = $manager2->findAllbyID($_SESSION['ID'],$id);
        $manager3 = new BookingManager();
        $Feed = $manager3->checkFeedback($_SESSION['ID'],$id);
        $manager4 = new FeedbackManager();
        $Review = $manager4->findAllbyVehiculeID($id);
        $free = "";
        $visible = "";
        if($Vehicule->isFree == 1){
            $free = "✅";
        } else {
            $free = "❌";
        }
        if($free == "✅"){
            $visible = "visible";
        } else {
            $visible = "hidden";
        }
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...✅❌
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Details.twig', ['vehicule' => $Vehicule, 'fav' => $Fav, 'feed' => $Feed, 'reviews' => $Review,
        'free' => $free, 'visible' => $visible, 'co' => $_SESSION['ID']]);
    }
    public function Fav()
    {   
        $id = $_GET['id'];
        $manager = new FavorisManager();
        $manager->insertOne($_SESSION['ID'],$id);
        header("Location: /details?id=$id");
    }
    public function Defav()
    {   
        $id = $_GET['id'];
        $manager = new FavorisManager();
        $manager->deleteOne($_SESSION['ID'],$id);
        header("Location: /details?id=$id");
    }
    public function createFeedback(){
        if($_SERVER["REQUEST_METHOD"] == "POST"){
            $VecID = $_POST['myCar'];
            $Note = $_POST['rate'];
            $Text = $_POST['myFB'];
            $manager = new FeedbackManager();
            if($manager->findAllbyIDs($_SESSION['ID'],$VecID)){
                echo "Tu as déjà commenter, Chien";
            } else {
                $texts = explode("'",$Text);
                if(count($texts)-1 > 0){
                    foreach(explode("'",$Text,-1) as $value){
                        $tab[] = $value . "''";
                    }
                    $text = $tab[0];
                    for($c = 1; $c < count($tab); $c++){
                        $text = $text . $tab[$c];
                    }
                    $text = $text . $texts[count($texts)-1];
                } else {
                    $text = $Text;
                }
                $manager->addOne($_SESSION['ID'],$VecID,$Note,$text);
                header("Location: /details?id=$VecID");
            }
        }
    }
}
