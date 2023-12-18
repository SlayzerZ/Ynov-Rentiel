<?php 
class Booking{
    public $ID;
    public $UserID;
    public $VehiculeID;
    public $Start;
    public $End;
    public $BrandID;
    public $SeatID;
    public $ColorID;
    public $Image;
    public $Location;
    public $PriceT;
    public $isFree;
    public $Brand;
    public $Seat;
    public $Color;

    /**
     * Get the value of ID
     */
    public function getID()
    {
        return $this->ID;
    }

    /**
     * Set the value of ID
     */
    public function setID($ID): self
    {
        $this->ID = $ID;

        return $this;
    }

    /**
     * Get the value of UserID
     */
    public function getUserID()
    {
        return $this->UserID;
    }

    /**
     * Set the value of UserID
     */
    public function setUserID($UserID): self
    {
        $this->UserID = $UserID;

        return $this;
    }

    /**
     * Get the value of VehiculeID
     */
    public function getVehiculeID()
    {
        return $this->VehiculeID;
    }

    /**
     * Set the value of VehiculeID
     */
    public function setVehiculeID($VehiculeID): self
    {
        $this->VehiculeID = $VehiculeID;

        return $this;
    }

    /**
     * Get the value of Start
     */
    public function getStart()
    {
        return $this->Start;
    }

    /**
     * Set the value of Start
     */
    public function setStart($Start): self
    {
        $this->Start = $Start;

        return $this;
    }

    /**
     * Get the value of End
     */
    public function getEnd()
    {
        return $this->End;
    }

    /**
     * Set the value of End
     */
    public function setEnd($End): self
    {
        $this->End = $End;

        return $this;
    }

    /**
     * Get the value of BrandID
     */
    public function getBrandID()
    {
        return $this->BrandID;
    }

    /**
     * Set the value of BrandID
     */
    public function setBrandID($BrandID): self
    {
        $this->BrandID = $BrandID;

        return $this;
    }

    /**
     * Get the value of SeatID
     */
    public function getSeatID()
    {
        return $this->SeatID;
    }

    /**
     * Set the value of SeatID
     */
    public function setSeatID($SeatID): self
    {
        $this->SeatID = $SeatID;

        return $this;
    }

    /**
     * Get the value of ColorID
     */
    public function getColorID()
    {
        return $this->ColorID;
    }

    /**
     * Set the value of ColorID
     */
    public function setColorID($ColorID): self
    {
        $this->ColorID = $ColorID;

        return $this;
    }

    /**
     * Get the value of Image
     */
    public function getImage()
    {
        return $this->Image;
    }

    /**
     * Set the value of Image
     */
    public function setImage($Image): self
    {
        $this->Image = $Image;

        return $this;
    }

    /**
     * Get the value of Location
     */
    public function getLocation()
    {
        return $this->Location;
    }

    /**
     * Set the value of Location
     */
    public function setLocation($Location): self
    {
        $this->Location = $Location;

        return $this;
    }

    /**
     * Get the value of PriceJ
     */
    public function getPriceT()
    {
        return $this->PriceT;
    }

    /**
     * Set the value of PriceJ
     */
    public function setPriceT($PriceT): self
    {
        $this->PriceT = $PriceT;

        return $this;
    }

    /**
     * Get the value of isFree
     */
    public function getIsFree()
    {
        return $this->isFree;
    }

    /**
     * Set the value of isFree
     */
    public function setIsFree($isFree): self
    {
        $this->isFree = $isFree;

        return $this;
    }

    /**
     * Get the value of Brand
     */
    public function getBrand()
    {
        return $this->Brand;
    }

    /**
     * Set the value of Brand
     */
    public function setBrand($Brand): self
    {
        $this->Brand = $Brand;

        return $this;
    }

    /**
     * Get the value of Seat
     */
    public function getSeat()
    {
        return $this->Seat;
    }

    /**
     * Set the value of Seat
     */
    public function setSeat($Seat): self
    {
        $this->Seat = $Seat;

        return $this;
    }

    /**
     * Get the value of Color
     */
    public function getColor()
    {
        return $this->Color;
    }

    /**
     * Set the value of Color
     */
    public function setColor($Color): self
    {
        $this->Color = $Color;

        return $this;
    }
}
?>