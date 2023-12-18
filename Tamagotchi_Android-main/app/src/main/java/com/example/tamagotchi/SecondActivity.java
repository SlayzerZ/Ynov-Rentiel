package com.example.tamagotchi;

import android.annotation.SuppressLint;
import android.content.Context;
import android.content.Intent;
import android.media.MediaPlayer;
import android.os.Bundle;
import android.view.View;
import android.view.animation.Animation;
import android.view.animation.AnimationUtils;
import android.widget.Button;
import android.widget.Chronometer;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.Writer;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;
import java.util.List;
import java.util.ListIterator;
import java.util.Map;

public class SecondActivity extends AppCompatActivity {
    Animation happy;
    boolean charged = false;
    Context context;
    Button setHappy, setHappy2;
    ImageView tama,save,shop,inventory;
    TextView displayName;
    Item nItem;
    public static TextView age_meter;
    public static TextView hunger_meter;
    public static TextView happy_meter;
    public static TextView heatlh_meter;
    public static Tamagotchi Tama;
    private Chronometer chronometer;
    public static int money;
    int counter =0;
    int counter2 =0;
    public static int getMoney(){return money;}
    public static void setMoney(int money1){
        money = money1;
    }
    public static TextView moneyView;
    public static List<Item> inventoryList = new ArrayList<>();
    public int inventoryCount() {return inventoryList.size();}
    public static MediaPlayer shopM,mainM;
    //@SuppressLint("MissingInflatedId")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        displayName = findViewById(R.id.displayName);
        shopM = MediaPlayer.create(getApplicationContext(),R.raw.boutique);
        mainM = MediaPlayer.create(getApplicationContext(),R.raw.exam);
       // mainM.start();
        mainM.setLooping(true);
        context = getBaseContext();
        if (!MainActivity.getLoad())
        {
            Tama = StartSetup.tama1;
            inventoryList = StartSetup.inventoryList;
        } else {
            Tama = MainActivity.Tama1;
            MainActivity.isLoad = false;
            if (!charged){
                charged = true;
                File file = getDataDir();
                File file1 = new File(file.getAbsolutePath() + "/itemlist.json");
                FileReader fileReader = null;
                try {
                    fileReader = new FileReader(file1);
                } catch (FileNotFoundException e) {
                    throw new RuntimeException(e);
                }
                BufferedReader bufferedReader = new BufferedReader(fileReader);
                StringBuilder stringBuilder = new StringBuilder();
                String line = null;
                try {
                    line = bufferedReader.readLine();
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
                while (line != null){
                    stringBuilder.append(line).append("\n");
                    try {
                        line = bufferedReader.readLine();
                    } catch (IOException e) {
                        throw new RuntimeException(e);
                    }
                }
                try {
                    bufferedReader.close();
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
// This responce will have Json Format String
                String responce = stringBuilder.toString();
                JSONArray jsonObject  = null;
                try {
                    jsonObject = new JSONArray(responce);
                    // Toast.makeText(getApplicationContext(), jsonObject.toString(), Toast.LENGTH_SHORT).show();
                } catch (Exception e) {
                    throw new RuntimeException(e);
                }
//Java Object
                try {
                        JSONObject mario;
                        Double d;
                    for (int i =0; i < jsonObject.length();i++){
                        mario = new JSONObject(jsonObject.get(i).toString());
                        d = new Double(Integer.parseInt(mario.get("Price").toString()));//first way
                        nItem = new Item(mario.get("Name").toString(), d,mario.get("Type").toString(),mario.get("Description").toString());
                        // Item nItem2 = new Item("", 0,"","");
                       // Toast.makeText(getApplicationContext(), nItem.getDescription(),Toast.LENGTH_SHORT).show();
                        inventoryList.add(nItem);
                    }
                } catch (Exception e) {
                    Toast.makeText(getApplicationContext(), e.getMessage(),Toast.LENGTH_SHORT).show();
                    displayName.setText(e.getMessage());
                }
            }
        }
        tama = findViewById(R.id.Tama);
        save = findViewById(R.id.save);
        shop = findViewById(R.id.shopIcon);
        inventory = findViewById(R.id.inventoryIcon);
        age_meter = findViewById(R.id.age_meter);
        hunger_meter = findViewById(R.id.hunger_meter);
        happy_meter = findViewById(R.id.happy_meter);
        heatlh_meter = findViewById(R.id.health_meter);
        moneyView = findViewById(R.id.money_meter);
        setHappy = findViewById(R.id.game1);
        setHappy2 = findViewById(R.id.sethappy2);
        chronometer = findViewById(R.id.chrono2);
        int resId = this.getResources().getIdentifier(Tama.getImgName(),"drawable",this.getPackageName());
        tama.setImageResource(resId);
        displayName.setText(Tama.getName());
       // chronometer.setText(counter);
        chronometer.setOnChronometerTickListener(chronometer -> onChronometerTickHandler());
        this.chronometer.start();
        shop.setOnClickListener(v -> {
            mainM.pause();
            Intent nextScreen = new Intent(context, Shop.class);
            startActivity(nextScreen);
           // shopM.start();
        });
        inventory.setOnClickListener(v -> {
            Intent nextScreen = new Intent(getApplicationContext(), Inventory.class);
            startActivity(nextScreen);
        });
        setHappy.setOnClickListener(v -> {
            Intent nextScreen = new Intent(getApplicationContext(), ChronoActivity.class);
            startActivity(nextScreen);

        });
        setHappy2.setOnClickListener(v -> {
            Intent nextScreen = new Intent(getApplicationContext(), MorpionMenu.class);
            startActivity(nextScreen);

        });
        save.setOnClickListener(v -> {
            JSONArray obj2 = new JSONArray();
            JSONObject obj;
            JSONObject Tama1 = new JSONObject();
            try {
                Tama1.put("Name",Tama.getName());
                Tama1.put("imgName",Tama.getImgName());
                Tama1.put("deathCase",Tama.getDeathCase());
                Tama1.put("Age",Tama.getAge());
                Tama1.put("Hunger",Tama.getHunger());
                Tama1.put("Happy",Tama.getHappy());
                Tama1.put("Health",Tama.getHealth());
                Tama1.put("Alive",Tama.getAlive());
                Tama1.put("Money",getMoney());
            } catch (JSONException e) {
                throw new RuntimeException(e);
            }
            try {
                Writer output = null;
                File file = getDataDir();
                File file1 = new File(file.getAbsolutePath() + "/tam.json");
                output = new BufferedWriter(new FileWriter(file1));
                output.write(Tama1.toString());
                output.close();
                Toast.makeText(getApplicationContext(), "Tamagotchi data saved", Toast.LENGTH_LONG).show();
               // Toast.makeText(getApplicationContext(), file1.getAbsolutePath(), Toast.LENGTH_LONG).show();

            } catch (Exception e) {
                Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
            }
         //   if (!inventoryList.isEmpty()){
                for (int j = 0; j < inventoryCount(); j++){
                    try {
                        obj = new JSONObject();
                        obj.put("Name",inventoryList.get(j).getName());
                        obj.put("Price",inventoryList.get(j).getPrice());
                        obj.put("Type",inventoryList.get(j).getType());
                        obj.put("Description",inventoryList.get(j).getDescription());
                        obj2.put(obj);
                    } catch (JSONException e) {
                        throw new RuntimeException(e);
                    }
                }
                try {
                    Writer output2 = null;
                    File file2 = getDataDir();
                    File file3 = new File(file2.getAbsolutePath() + "/itemlist.json");
                    output2 = new BufferedWriter(new FileWriter(file3));
                    output2.write(obj2.toString());
                    output2.close();
                    Toast.makeText(getApplicationContext(), "Inventory data saved", Toast.LENGTH_LONG).show();

                } catch (Exception e) {
                    Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
                }
         //   }
        });
    }
    public static void updateTama(){
        age_meter.setText(": " + (int) Tama.getAge());
        hunger_meter.setText(": " + (int) Tama.getHunger());
        happy_meter.setText(": " + (int) Tama.getHappy());
        heatlh_meter.setText(": " + (int) Tama.getHealth());
        moneyView.setText(": " + getMoney());
    }
    private void onChronometerTickHandler()  {
        if(this.counter > 25) {
            this.counter = 0;
            Tama.setAge(Tama.getAge()+1);
        } else if (this.counter == 10) {
            Tama.setHunger(Tama.getHunger()-1);
        } else if (this.counter == 20) {
            Tama.setHappy(Tama.getHappy()-1);
        }
        Tama.setHealth((Tama.getHappy()*Tama.getHunger())/100);
        //Tama.setHappy(0);
        if ((int) Tama.getHealth() == 0){
            mainM.pause();
            Tama.setAlive(false);
           // Tama.setDeathCase("Vous êtes une merde, Vôtre " + Tama.getName() +
           //             " à développer des troubles par votre faute. C'est pourtant pas compliquer de gérer un équilibre entre faim et bonheur. " +
           //             "La garde nationale vous l'a retiré pour maltraitance et il même aujourd'hui une existence bien plus paisible.");
            if (Tama.getHappy() < Tama.getHunger()){
                Tama.setDeathCase("Vous n'avez pas su apporter joie et bonheur à " + Tama.getName() +
                        " Vous êtes très certainement un(e) mauvais(e) père/mère. Face à votre incompétence, il a preférer en finir lui-même en se pendant à une corde. " +
                        "Une âme chère nous a quittez Aujourd'hui, vous devriez avoir honte.");
            }
            if (Tama.getHunger() < Tama.getHappy()){
                Tama.setDeathCase("Vôtre " + Tama.getName() + " est mort(e) le jeûne aura eu raison de lui/d'elle." +
                        " Pour quelle raison ne l'avez vous pas nourri? Vous êtes quelqu'un d'occupé(e) ou c'est juste que vous en aviez plus rien à foutre? " +
                        "Dans tous les cas vous êtes le/la seul(e) fautif/ve et tôt ou tard vous finirez par le regrettez");
            }
            Intent nextscreen = new Intent(getApplicationContext(), RIP.class);
            startActivity(nextscreen);
            finish();
        }
        if (counter % 5 == 0){
            if (shopM.isPlaying()){
                shopM.pause();
                // shopM.release();
            }
            if (this.context == getBaseContext()){
        //        mainM.start();
            }
        }

        updateTama();
        this.counter++;
        //chronometer.setText(counter);
    }
}