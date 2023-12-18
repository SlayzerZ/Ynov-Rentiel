<?php
class Vehicule{
    public $ID;
    public $BrandID;
    public $SeatID;
    public $ColorID;
    public $Image;
    public $Location;
    public $PriceJ;
    public $isFree;
    public $Brand;
    public $Seat;
    public $Color;

    public function getID(){
        return $this->ID;
    }
    public function setID($ID){
        $this->ID = $ID;
    }
    public function getBrandID(){
        return $this->BrandID;
    }
    public function setBrandID($BrandID){
        $this->BrandID = $BrandID;
    }
    public function getSeatID(){
        return $this->SeatID;
    }
    public function setSeatID($SeatID){
        $this->SeatID = $SeatID;
    }
    public function getColorID(){
        return $this->ColorID;
    }
    public function setColorID($ColorID){
        $this->ColorID = $ColorID;
    }
    public function getImage(){
        return $this->Image;
    }
    public function setImage($Image){
        $this->Image = $Image;
    }
    public function getLocation(){
        return $this->Location;
    }
    public function setLocation($Location){
        $this->Location = $Location;
    }
    public function getPriceJ(){
        return $this->PriceJ;
    }
    public function setPriceJ($PriceJ){
        $this->PriceJ = $PriceJ;
    }
    public function getisFree(){
        return $this->isFree;
    }
    public function setisFree($isFree){
        $this->isFree = $isFree;
    }
    public function getBrand(){
        return $this->Brand;
    }
    public function setBrand($Brand){
        $this->Brand = $Brand;
    }
    public function getSeat(){
        return $this->Seat;
    }
    public function setSeat($Seat){
        $this->Seat = $Seat;
    }
    public function getColor() {
		return $this->Color;
	}
	public function setColor($Color) {
		$this->Color = $Color;
	}
}
?>