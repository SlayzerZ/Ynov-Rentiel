<?php
session_start();
class SearchController
{
    public function showSearch()
    {   
        if(empty($_SESSION['ID'])){
            $_SESSION['ID'] = null;
        }

        $manager3 = new BrandManager();
        $brands = $manager3->findAll();
        $manager4 = new SeatManager();
        $seats = $manager4->findAll();
        $manager5 = new ColorManager();
        $colors = $manager5->findAll();

        $manager = new VehiculeManager();
        if(is_null($manager->findAll())){
            $manager->insertOneAuto();
            do{$manager->insertOneAuto();}while(count($manager->findAll()) < 12);
        }
        $manager7 = new UserManager();
        $Users = $manager7->findAll();
        if(!empty($Users)){
            $datenow = new DateTime('now,+1 days');
            $manager2 = new BookingManager();
            $manager2->restore($datenow->format("Y-m-d"));
        }
        $Vehicules = $manager->findAll();
        $manager6 = new CityManager();
        $cities = $manager6->findAll();
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Search.twig', ['vehicules' => $Vehicules, 'brands' => $brands, 'seats' => $seats, 'colors' => $colors, 'cities' => $cities, 'co' => $_SESSION['ID']]);
    }
    public function useFilters()
    {
        if($_SERVER["REQUEST_METHOD"] == "POST"){
            $brand = $_POST['myBrand'];
            $seat = $_POST['mySeat'];
            $color = $_POST['myColor'];
            $price = $_POST['myPrice'];
        $manager = new VehiculeManager();
        $Vehicules = $manager->findAllbyFilters($brand,$seat,$color,$price);
        // var_dump($Vehicules);
        $manager3 = new BrandManager();
        $brands = $manager3->findAll();
        $manager4 = new SeatManager();
        $seats = $manager4->findAll();
        $manager5 = new ColorManager();
        $colors = $manager5->findAll();
        $manager6 = new CityManager();
        $cities = $manager6->findAll();
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Search.twig', ['vehicules' => $Vehicules, 'brands' => $brands, 'seats' => $seats, 'colors' => $colors, 'cities' => $cities, 'co' => $_SESSION['ID']]);
        }
    }
    public function findLocation()
    {
        if($_SERVER["REQUEST_METHOD"] == "POST"){
        $location = $_POST['myLocation'];
        $manager = new VehiculeManager();
        $Vehicules = $manager->findAllbyLocation($location);
        $manager3 = new BrandManager();
        $brands = $manager3->findAll();
        $manager4 = new SeatManager();
        $seats = $manager4->findAll();
        $manager5 = new ColorManager();
        $colors = $manager5->findAll();
        $manager6 = new CityManager();
        $cities = $manager6->findAll();
        $loader = new \Twig\Loader\FilesystemLoader(VIEW);
        $twig = new \Twig\Environment($loader, [
            'debug' => true,
            // ...
        ]);
        $twig->addExtension(new \Twig\Extension\DebugExtension());
        echo $twig->render('Search.twig', ['vehicules' => $Vehicules, 'brands' => $brands, 'seats' => $seats, 'colors' => $colors, 'cities' => $cities, 'co' => $_SESSION['ID']]);
        }
    }
}
