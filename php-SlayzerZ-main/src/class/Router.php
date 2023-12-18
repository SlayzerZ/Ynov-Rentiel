<?php
class Router{
    private $request;
    private $routes = [ 
        "home" => ["controller" => "HomeController", "method" => "showHome"],
        "search" => ["controller" => "SearchController", "method" => "showSearch"],
        "profil" => ["controller" => "ProfilController", "method" => "showProfil"],
        "sign" => ["controller" => "SignController", "method" => "showSign"],
        "details" => ["controller" => "DetailsController", "method" => "showDetails"],
        "applyFilters" => ["controller" => "SearchController", "method" => "useFilters"],
        "signup" => ["controller" => "SignController", "method" => "SignUp"],
        "signin" => ["controller" => "SignController", "method" => "SignIn"],
        "signout" => ["controller" => "SignController", "method" => "SignOut"],
        "booking" => ["controller" => "BookingController", "method" => "showBooking"],
        "booker" => ["controller" => "BookingController", "method" => "haveBook"],
        "Fav" => ["controller" => "DetailsController", "method" => "Fav"],
        "Defav" => ["controller" => "DetailsController", "method" => "Defav"],
        "feedback" => ["controller" => "DetailsController", "method" => "createFeedback"],
        "findLocation" => ["controller" => "SearchController", "method" => "findLocation"],
        "adminPanel" => ["controller" => "AdminController", "method" => "showPanel"],
        "access" => ["controller" => "AdminController", "method" => "showAccess"],
        "create" => ["controller" => "AdminController", "method" => "createItem"],
        "modify" => ["controller" => "AdminController", "method" => "modifyItem"],
        "delete" => ["controller" => "AdminController", "method" => "deleteItem"],
        "truncate" => ["controller" => "AdminController", "method" => "truncateTable"],
    ];
    public function __construct($request)
    {
        $this->request = $request;
    }
    public function renderController(){
        $request = $this->request;
        if(key_exists($request, $this->routes)){
            $controller = $this->routes[$request]['controller'];
            $method = $this->routes[$request]['method'];
            $currentController = new $controller();
            $currentController->$method();
        } else {
            include(CONTROLLER.'404Controller.php');
        }
    }
}
?>