package com.example.tamagotchi;

import androidx.appcompat.app.AppCompatActivity;

import android.annotation.SuppressLint;
import android.content.Intent;
import android.os.Bundle;
import android.widget.Button;
import android.widget.Toast;

import org.json.JSONException;
import org.json.JSONObject;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class MainActivity extends AppCompatActivity {
   Button NG, Continue, Help;
    public static Tamagotchi Tama1;
   public static boolean isLoad = false;
    public static boolean getLoad(){return isLoad;}
    @SuppressLint("MissingInflatedId")
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.start_page);
        NG = findViewById(R.id.ng);
        Continue = findViewById(R.id.continue1);
        Help = findViewById(R.id.help);
        NG.setOnClickListener(v -> {
            Intent nextScreen = new Intent(getApplicationContext(), StartSetup.class);
            startActivity(nextScreen);
            //finish();
        });
        Continue.setOnClickListener(v -> {
            try {
                isLoad = true;
                File file = getDataDir();
                File file1 = new File(file.getAbsolutePath() + "/tam.json");
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
                JSONObject jsonObject  = null;
                try {
                    jsonObject = new JSONObject(responce);
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
//Java Object
                try {
                    Tama1 = new Tamagotchi(jsonObject.get("Name").toString(),
                            jsonObject.get("imgName").toString(),
                            jsonObject.get("deathCase").toString(),
                            Integer.parseInt(jsonObject.get("Age").toString()),
                            Integer.parseInt(jsonObject.get("Hunger").toString()),
                            Integer.parseInt(jsonObject.get("Happy").toString()),
                            Integer.parseInt(jsonObject.get("Health").toString()),
                            Boolean.parseBoolean(jsonObject.get("Alive").toString()))
                    ;
                    SecondActivity.setMoney(Integer.parseInt(jsonObject.get("Money").toString()));
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                if (Tama1.getAlive()){
                    Intent nextScreen = new Intent(getApplicationContext(), SecondActivity.class);
                    startActivity(nextScreen);
                } else {
                    Toast.makeText(getApplicationContext(), Tama1.getName() + " est parti(e), il/elle ne reviendra pas", Toast.LENGTH_SHORT).show();
                }
                //finish();
            } catch (Exception e){
                Toast.makeText(getApplicationContext(), "Pas de sauvegarde", Toast.LENGTH_SHORT).show();
                Toast.makeText(getApplicationContext(), e.getMessage(), Toast.LENGTH_LONG).show();
            }

        });
        Help.setOnClickListener(v -> {
            Toast.makeText(getApplicationContext(), "Il y a pas d'aide pour le moment, démerdez-vous!",Toast.LENGTH_LONG).show();
        });
    }
}