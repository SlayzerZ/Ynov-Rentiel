package com.example.smash;

import javafx.beans.property.SimpleStringProperty;
import javafx.beans.property.StringProperty;

public class Item {
    private final StringProperty name;
    private final StringProperty price;
    private final StringProperty description;
    Item(String name1, int price1, String description1){
        this.name = new SimpleStringProperty(name1);
        this.price = new SimpleStringProperty(Integer.toString(price1));
        this.description = new SimpleStringProperty(description1);
    }
    public final StringProperty nameProperty() {
        return name;
    }
    public String getToDo() {
        return name.get();
    }
    
    public void setToDo(String todo1) {
        name.set(todo1);
    }
    public String getPrice(){
        return price.get();
    }

    public void setPrice(String price1) {
        price.set(price1);
    }
    public String getDescription() {
        return description.get();
    }
    public void setDescription(String deadline1) {
        description.set(deadline1);
    }
}
