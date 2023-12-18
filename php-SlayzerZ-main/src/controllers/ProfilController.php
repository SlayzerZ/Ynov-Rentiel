<?php
session_start();
class ProfilController
{
    public function showProfil()
    {
        try {
            if($_SESSION['ID'] == null){
                header('Location: /sign');
            }
            $manager = new UserManager();
            $User = $manager->findOnebyID($_SESSION['ID']);
            $manager2 = new FavorisManager();
            $Fav = $manager2->findAllbyUserID($_SESSION['ID']);
            $manager3 = new BookingManager();
            $Book = $manager3->findAllbyUserID($_SESSION['ID']);
            $manager4 = new FeedbackManager();
            $Review = $manager4->findAllbyUserID($_SESSION['ID']);
            $loader = new \Twig\Loader\FilesystemLoader(VIEW);
            $twig = new \Twig\Environment($loader, [
                'debug' => true,
                // ...
            ]);
            $twig->addExtension(new \Twig\Extension\DebugExtension());
            echo $twig->render('Profil.twig', ['user' => $User, 'favs' => $Fav, 'books' => $Book, 'reviews' => $Review,
            'co' => $_SESSION['ID']]);
        } catch(Exception $e){
            echo $e->getMessage();
        }
        
    }
}
