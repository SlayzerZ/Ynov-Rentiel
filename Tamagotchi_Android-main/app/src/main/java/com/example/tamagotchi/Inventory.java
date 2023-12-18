package com.example.tamagotchi;

import android.os.Bundle;
import android.widget.Chronometer;
import android.widget.ListView;

import androidx.appcompat.app.AppCompatActivity;

import java.util.ArrayList;
import java.util.List;

public class Inventory extends AppCompatActivity {
    int counter =0;
    Chronometer Chronos;
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.list);
        ListView shopListView = findViewById(R.id.shop_list_view);
        Chronos = findViewById(R.id.chrono3);
        Chronos.setOnChronometerTickListener(chronometer -> onChronometerTickHandler());
        this.Chronos.start();
        shopListView.setAdapter(new ItemAdapter2(this,SecondActivity.inventoryList));

    }
    private void onChronometerTickHandler()  {
        if(this.counter > 1) {
            this.counter = 0;
            if (ItemAdapter2.use2){
                ItemAdapter2.use2 = false;
                finish();
                startActivity(getIntent());
            }
        }
        this.counter++;
        //chronometer.setText(counter);
    }
}
