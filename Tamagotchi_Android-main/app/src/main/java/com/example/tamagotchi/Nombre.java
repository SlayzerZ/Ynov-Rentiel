package com.example.tamagotchi;

import java.util.Random;

public class Nombre {
    protected int Number;
    public Nombre(int Number) {
        this.Number = Number;
    }
    public void RandomNumber(int min, int max) {
        if (min >= max) {
            throw new IllegalArgumentException("max must be greater than min");
        }

        Random r = new Random();
        Number = r.nextInt((max - min) + 1) + min;
    }
    public void setNumber(int Number1) {
        this.Number = Number1;
    }
    public int getNumber() {
        return Number;
    }
}
