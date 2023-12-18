<?php
class Color{
    public $ID;
    public $Name;
    public function getID(){
        return $this->ID;
    }
    public function setID($ID){
        $this->ID = $ID;
    }
    public function getName(){
        return $this->Name;
    }
    public function setName($Name){
        $this->Name = $Name;
    }
}
?>