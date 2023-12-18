package com.example.tamagotchi;

import android.app.Dialog;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.view.animation.Animation;
import android.view.animation.AnimationUtils;
import android.widget.Button;
import android.widget.Chronometer;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.fragment.app.DialogFragment;

public class ChronoActivity extends AppCompatActivity {
    Animation happy2;
    TextView operation,editAnswer;
    ImageView Tama;
    Button Answer;
    Chronometer chronometer;
    Nombre N1 = new Nombre(0);
    Nombre N2 = new Nombre(0);
    Nombre N3 = new Nombre(0);
    private int counter = 15;
    private double attempt =0;
    private double correct =0;
    int happy = (int) SecondActivity.Tama.getHappy();
    int hunger = (int) SecondActivity.Tama.getHunger();
    int health = (int) SecondActivity.Tama.getHealth();
    double money = SecondActivity.getMoney();
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.calcul);
        operation = findViewById(R.id.operation);
        editAnswer = findViewById(R.id.editAnswer);
        Tama = findViewById(R.id.Tama2);
        Answer = findViewById(R.id.Answer);
        chronometer = findViewById(R.id.chrono);
        int resId = this.getResources().getIdentifier(SecondActivity.Tama.getImgName(),"drawable",this.getPackageName());
        Tama.setImageResource(resId);
        reset();
        chronometer.setText(counter + "s");
        chronometer.setOnChronometerTickListener(chronometer -> onChronometerTickHandler());
        doStart();
        Answer.setOnClickListener(v -> {
            attempt++;
            int i = Integer.parseInt(editAnswer.getText().toString());
            if (i == N3.getNumber()){
                happy2 = AnimationUtils.loadAnimation(getApplicationContext(),R.anim.happy2);
                Tama.setVisibility(View.VISIBLE);
                Tama.startAnimation(happy2);
                Toast.makeText(getApplicationContext(), "Correct", Toast.LENGTH_SHORT).show();
                reset();
                correct++;
            } else {
                Toast.makeText(getApplicationContext(), "Nullos", Toast.LENGTH_SHORT).show();
                reset();
            }
        });
    }
    void reset(){
        N1.RandomNumber(0,10);
        N2.RandomNumber(0,10);
        N3.setNumber(N1.getNumber() * N2.getNumber());
        operation.setText(N1.getNumber() + " x " + N2.getNumber());
        editAnswer.setText("");
    }
    private void onChronometerTickHandler()  {
        if(this.counter < 0) {
            double a = (correct*attempt) *10;
            double a2 = money + a;
          // Toast.makeText(getApplicationContext(), Integer.toString((int) money), Toast.LENGTH_SHORT).show();
            this.counter = (int) money;
            this.chronometer.setText(counter + "s");
            Toast.makeText(getApplicationContext(), "Ton score est de " + correct + " sur " + attempt, Toast.LENGTH_LONG).show();
            try {
               SecondActivity.setMoney((int) a2);
                hunger = (int) (hunger - attempt);
                SecondActivity.Tama.setHunger(hunger);
                if (correct >= attempt - correct){
                    happy = (int) (happy + (correct * attempt));
                    SecondActivity.Tama.setHappy(happy);
                } else {
                    happy = (int) (happy - (correct * attempt));
                    SecondActivity.Tama.setHappy(happy);
                }
                health = (happy*hunger)/100;
                SecondActivity.Tama.setHealth(health);
                SecondActivity.updateTama();
                Intent nextScreen = new Intent(getApplicationContext(), SecondActivity.class);
                //startActivity(nextScreen);
                finish();
            } catch (Exception e) {
                Toast.makeText(getApplicationContext(), e.getMessage(), Toast.LENGTH_LONG).show();
            }

        } else {
            this.chronometer.setText(counter + "s");
            this.counter--;
        }
    }
    private void doStart()  {
        this.chronometer.start();
    }
}
