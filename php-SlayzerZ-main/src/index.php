<?php 
include_once('_config.php');
date_default_timezone_set('Europe/Paris');
// echo boolval($pass? 'true' : 'false') ; exit;
MyAutoLoad::start();

$request = $_GET['action'];

if($request == '') {
    $request = 'home';
}

$routeur = new Router($request);
$routeur->renderController();

?>