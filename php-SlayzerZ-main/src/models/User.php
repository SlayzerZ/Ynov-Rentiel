<?php 
class User{
    public $ID;
    public $Image;
    public $FirstName;
    public $LastName;
    public $PhoneNumber;
    public $Email;
    public $Password;
    public $Admin;

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
}
?>