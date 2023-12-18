<?php
session_start();
class HomeController
{
    public function showHome()
    {
        if(empty($_SESSION['ID'])){
            $_SESSION['ID'] = null;
        }
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Home.twig', ['co' => $_SESSION['ID']]);
    }
}
