package com.example.tamagotchi;

import static com.example.tamagotchi.SecondActivity.Tama;
import static com.example.tamagotchi.SecondActivity.getMoney;
import static com.example.tamagotchi.SecondActivity.inventoryList;

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

public class RIP extends AppCompatActivity {
    TextView RIP,CJ,death;
    ImageView tama5;
    Button menu;
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.rip_screen);
        RIP = findViewById(R.id.RIP);
        CJ = findViewById(R.id.CJ);
        death = findViewById(R.id.death);
        tama5 = findViewById(R.id.tama5);
        menu = findViewById(R.id.menu);
        int resId = this.getResources().getIdentifier(Tama.getImgName(),"drawable",this.getPackageName());
        tama5.setImageResource(resId);
        CJ.setText(Tama.getName() + " Ami(e)s de tous, parti(e) trop tôt.");
        death.setText(Tama.getDeathCase());
        menu.setOnClickListener(v -> {
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
               // Toast.makeText(getApplicationContext(), "Tamagotchi data saved", Toast.LENGTH_LONG).show();
                // Toast.makeText(getApplicationContext(), file1.getAbsolutePath(), Toast.LENGTH_LONG).show();

            } catch (Exception e) {
                Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
            }
            for (int j = 0; j < inventoryList.size(); j++){
                try {
                    obj = new JSONObject();
                    obj.put("Name","");
                    obj.put("Price",0);
                    obj.put("Type","");
                    obj.put("Description","");
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
               // Toast.makeText(getApplicationContext(), "Inventory data saved", Toast.LENGTH_LONG).show();

            } catch (Exception e) {
                Toast.makeText(getBaseContext(), e.getMessage(), Toast.LENGTH_LONG).show();
            }
            finish();
        });
    }
}
