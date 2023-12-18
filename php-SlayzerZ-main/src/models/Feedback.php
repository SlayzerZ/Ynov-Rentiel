<?php 
class Feedback{
    public $ID;
    public $UserID;
    public $VehiculeID;
    public $Note;
    public $Text;
    public $FirstName;
    public $LastName;
    public $PhoneNumber;
    public $Email;
    public $Password;
    public $Admin;
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
     * Get the value of Note
     */
    public function getNote()
    {
        return $this->Note;
    }

    /**
     * Set the value of Note
     */
    public function setNote($Note): self
    {
        $this->Note = $Note;

        return $this;
    }

    /**
     * Get the value of Text
     */
    public function getText()
    {
        return $this->Text;
    }

    /**
     * Set the value of Text
     */
    public function setText($Text): self
    {
        $this->Text = $Text;

        return $this;
    }

    /**
     * Get the value of FirstName
     */
    public function getFirstName()
    {
        return $this->FirstName;
    }

    /**
     * Set the value of FirstName
     */
    public function setFirstName($FirstName): self
    {
        $this->FirstName = $FirstName;

        return $this;
    }

    /**
     * Get the value of LastName
     */
    public function getLastName()
    {
        return $this->LastName;
    }

    /**
     * Set the value of LastName
     */
    public function setLastName($LastName): self
    {
        $this->LastName = $LastName;

        return $this;
    }

    /**
     * Get the value of PhoneNumber
     */
    public function getPhoneNumber()
    {
        return $this->PhoneNumber;
    }

    /**
     * Set the value of PhoneNumber
     */
    public function setPhoneNumber($PhoneNumber): self
    {
        $this->PhoneNumber = $PhoneNumber;

        return $this;
    }

    /**
     * Get the value of Email
     */
    public function getEmail()
    {
        return $this->Email;
    }

    /**
     * Set the value of Email
     */
    public function setEmail($Email): self
    {
        $this->Email = $Email;

        return $this;
    }

    /**
     * Get the value of Password
     */
    public function getPassword()
    {
        return $this->Password;
    }

    /**
     * Set the value of Password
     */
    public function setPassword($Password): self
    {
        $this->Password = $Password;

        return $this;
    }

    /**
     * Get the value of Admin
     */
    public function getAdmin()
    {
        return $this->Admin;
    }

    /**
     * Set the value of Admin
     */
    public function setAdmin($Admin): self
    {
        $this->Admin = $Admin;

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
    public function getPriceJ()
    {
        return $this->PriceJ;
    }

    /**
     * Set the value of PriceJ
     */
    public function setPriceJ($PriceJ): self
    {
        $this->PriceJ = $PriceJ;

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