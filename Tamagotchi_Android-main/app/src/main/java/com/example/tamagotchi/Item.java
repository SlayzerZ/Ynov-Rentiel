package com.example.tamagotchi;

public class Item {
    private String name;
    private double price;
    private String type;
    private String description;
    public Item (String name, double price, String type, String description){
        this.name = name;
        this.price = price;
        this.type = type;
        this.description = description;
    }

    public String getName() {
        return this.name;
    }
    public String getDescription(){return this.description;}
    public double getPrice() {
        return this.price;
    }
    public String getType(){
        return this.type;
    }

}
