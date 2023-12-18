<?php
session_start();
class SignController
{
    public function showSign()
    {   
        $manager = new UserManager();
        $Users = $manager->findAll();
        if(empty($Users)){
            $manager->firstAdmin();
        }
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Sign.twig');
    }
    public function SignUp()
    {   
        $manager = new UserManager();
        if($_SERVER["REQUEST_METHOD"] == "POST"){
            $FirstName = $_POST['myFN'];
            $LastName = $_POST['myLN'];
            $PhoneNumber = $_POST['myPN'];
            $Email = $_POST['myEM'];
            $Password = $_POST['myPS'];
        if(empty($FirstName) || empty($LastName) || empty($PhoneNumber) || empty($Email) || empty($Password)){
            $User = "Veuillez remplir tous les champs";
            $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Sign.twig', ['user' => $User]);
        } else {
        $Password2 = password_hash($Password, PASSWORD_DEFAULT);
        $faker = Faker\Factory::Create();
        $image = array(["sonic.png","mario.png","Bip_Bip.png",
        "Blue_Falcon.png","FlashRender.png","kizaru.png","Lucci_CP0.png","minato.png","pit.png","pit2.png",
        "Rainbow_Dash.png","saitama.png","speedy.png","Whellie_Rider.png","warp.png"]);
        $id = $faker->randomNumber(4, true);
        $imageUrl = $faker->randomElement($image[0]);
        
        $User = $manager->insertOne($id,$imageUrl,$FirstName,$LastName,$PhoneNumber,$Email,$Password2);
       

        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        if($User == "success"){
        echo $twig->render('Home.twig', ['co' => $_SESSION['ID']]);
        } else {
            echo $twig->render('Sign.twig', ['user' => $User]);
        }
        }
        }
    }
    public function SignIn(){
        if($_SERVER["REQUEST_METHOD"] == "POST"){
            $Email = $_POST['myEM'];
            $Password = $_POST['myPS'];
            if(empty($Email) || empty($Password)){
                $User = "Veuillez remplir tous les champs";
                $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
                'debug' => true,
                // ...
            ]);
            $twig->addExtension(new \Twig\Extension\DebugExtension());
            echo $twig->render('Sign.twig', ['user2' => $User]);
            } else {
            $manager = new UserManager();
            $tab = $manager->findOnebyConnection($Email,$Password);
            $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
                'debug' => true,
                // ...
            ]);
            $twig->addExtension(new \Twig\Extension\DebugExtension());
            if($tab == "true"){
                echo "Connexion Réussi";
                echo $twig->render('Home.twig', ['co' => $_SESSION['ID']]);
            } elseif($tab == "false"){
                $echec = "Echec de Connexion";
                echo $twig->render('Sign.twig', ['echec' => $echec, 'co' => $_SESSION['ID']]);
            }
            }
        }
    }
    public function SignOut()
    {
        $_SESSION['ID'] = NULL;
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo "Déconnexion Réussi";
        echo $twig->render('Home.twig', ['co' => $_SESSION['ID']]);
    }
}
