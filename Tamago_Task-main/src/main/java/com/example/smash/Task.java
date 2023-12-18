package com.example.smash;

import javafx.beans.property.SimpleStringProperty;
import javafx.beans.property.StringProperty;

public class Task {
        private final StringProperty Task;
        private final StringProperty Price;
        private final SimpleStringProperty day;
        private final SimpleStringProperty deadline;
        private final SimpleStringProperty description;
        
        Task(String todo1, String day1, int price1, String deadline1, String description1){
            this.day = new SimpleStringProperty (day1);
            this.deadline = new SimpleStringProperty(deadline1);
            this.description = new SimpleStringProperty(description1);
            this.Task = new SimpleStringProperty(todo1);
            this.Price = new SimpleStringProperty(Integer.toString(price1));
        }
        public final StringProperty TaskProperty() {
            return Task;
        }
         public final StringProperty PriceProperty() {
            return Price;
        }
        public String getToDo() {
            return Task.get();
        }
        
        public void setToDo(String todo1) {
            Task.set(todo1);
        }
        
        public String getDay() {
            return day.get();
        }
        
        public void setDay(String day1) {
            day.set(day1);
        }
        
        public String getPrice(){
            return Price.get();
        }

        public void setPrice(String price1) {
            Price.set(price1);
        }

        public String getDeadline() {
            return deadline.get();
        }
        
        public void setDeadline(String deadline1) {
            deadline.set(deadline1);
        }
        
        public String getDescription() {
            return description.get();
        }
        
        public void setDescription(String deadline1) {
            description.set(deadline1);
        }  
}