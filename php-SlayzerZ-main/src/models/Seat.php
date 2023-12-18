<?php
class Seat{
    public $ID;
    public $Seat;
    public function getID(){
        return $this->ID;
    }
    public function setID($ID){
        $this->ID = $ID;
    }
    public function getSeat(){
        return $this->Seat;
    }
    public function setSeat($Seat){
        $this->Seat = $Seat;
    }
}
?>