package com.example.tamagotchi;

public class Tamagotchi {
    protected String Name;
    protected String imgName;
    protected String deathCase;
    protected double Age;
    protected double Hunger;
    protected double Happy;
    protected double Health;
    protected boolean Alive;
    public Tamagotchi(String newName,String newimgName,String newdeathCase, double newAge, double newHunger,double newHappy,double newHealth, boolean newAlive){
        Name = newName;
        imgName = newimgName;
        deathCase = newdeathCase;
        Age = newAge;
        Hunger = newHunger;
        Health = newHealth;
        Happy = newHappy;
        Alive = newAlive;
    }
    public String getName(){return Name;}
    public String getImgName(){return imgName;}
    public String getDeathCase(){return deathCase;}
    public double getAge(){return Age;}
    public double getHunger(){return Hunger;}
    public double getHealth(){return Health;}
    public double getHappy(){return Happy;}
    public boolean getAlive(){return Alive;}

    public void setName(String name){Name = name;}
    public void setImgName(String name){imgName = name;}
    public void setDeathCase(String name){deathCase = name;}
    public void setAge(double name){
        Age = name;
        if (Age > 100){
            Age = 100;
        } else  if (Age < 0){
            Age = 0;
        }
    }
    public void setHunger(double name){
        Hunger = name;
        if (Hunger > 100){
            Hunger = 100;
        }else  if (Hunger < 0){
            Hunger = 0;
        }
    }
    public void setHealth(double name){
        Health = name;
        if (Health > 100){
            Health = 100;
        }else  if (Health < 0){
            Health = 0;
        }
    }
    public void setHappy(double name){
        Happy = name;
        if (Happy > 100){
            Happy = 100;
        }else  if (Happy < 0){
            Happy = 0;
        }
    }
    public void setAlive(boolean name){Alive = name;}
}
