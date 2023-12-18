package com.example.tamagotchi;

import android.annotation.SuppressLint;
import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.Writer;
import java.util.ArrayList;
import java.util.List;

public class StartSetup extends AppCompatActivity {
    TextView Name;
    public static Tamagotchi tama1;
    ImageView mario,sonic,kirby,clamiral,amy;
    public static List<Item> inventoryList = new ArrayList<>();
    Button Continue;
    String img;
    @SuppressLint("MissingInflatedId")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.start_setup);
        Name = findViewById(R.id.editName);
        Continue = findViewById(R.id.button);
        mario = findViewById(R.id.link);
        sonic = findViewById(R.id.bokoblin);
        kirby = findViewById(R.id.kirby);
        clamiral = findViewById(R.id.clamiral);
        amy = findViewById(R.id.amy);
        mario.setOnClickListener(v -> {
            img = "link";
        });
        sonic.setOnClickListener(v -> {
            img = "bokoblin";
        });
        kirby.setOnClickListener(v -> {
            img = "kirby";
        });
        clamiral.setOnClickListener(v -> {
            img = "clamiral";
        });
        amy.setOnClickListener(v -> {
            img = "amy";
        });
        Continue.setOnClickListener(v ->{
            if (img != null ){
                if (Name.getText().toString().equals("")){
                    Name.setText("Unknown");
                }
                tama1 = new Tamagotchi(Name.getText().toString(),img,"",1,100,100,100,true);
                JSONArray obj2 = new JSONArray();
                JSONObject obj;
                JSONObject Tama1 = new JSONObject();
                try {
                    Tama1.put("Name",tama1.getName());
                    Tama1.put("imgName",tama1.getImgName());
                    Tama1.put("deathCase",tama1.getDeathCase());
                    Tama1.put("Age",tama1.getAge());
                    Tama1.put("Hunger",tama1.getHunger());
                    Tama1.put("Happy",tama1.getHappy());
                    Tama1.put("Health",tama1.getHealth());
                    Tama1.put("Alive",tama1.getAlive());
                    Tama1.put("Money",0);
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
                 //   Toast.makeText(getApplicationContext(), "Tamagotchi data saved", Toast.LENGTH_LONG).show();
                    // Toast.makeText(getApplicationContext(), file1.getAbsolutePath(), Toast.LENGTH_LONG).show();

                } catch (Exception e) {
                    Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
                }
                //   if (!inventoryList.isEmpty()){
                for (int j = 0; j < inventoryList.size(); j++){
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
                 //   Toast.makeText(getApplicationContext(), "Inventory data saved", Toast.LENGTH_LONG).show();

                } catch (Exception e) {
                    Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
                }
                //   }
                Intent nextScreen = new Intent(getApplicationContext(), SecondActivity.class);
                startActivity(nextScreen);
                finish();
            } else {
                Toast.makeText(getApplicationContext(), "Erreur: Veuillez choisir une image", Toast.LENGTH_LONG).show();
            }

        });
    }
}