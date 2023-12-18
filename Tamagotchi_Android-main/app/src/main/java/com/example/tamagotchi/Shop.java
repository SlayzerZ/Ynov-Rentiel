package com.example.tamagotchi;

import android.app.ActionBar;
import android.content.Context;
import android.content.Intent;
import android.media.MediaPlayer;
import android.os.Bundle;
import android.widget.Button;
import android.widget.ListView;
import android.widget.TextView;

import androidx.appcompat.app.AppCompatActivity;

import java.util.ArrayList;
import java.util.List;

public class Shop extends AppCompatActivity {
    static TextView money_meter;
    int money = SecondActivity.getMoney();
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.list2);
        SecondActivity.shopM.start();
        SecondActivity.shopM.setLooping(true);
        money_meter = findViewById(R.id.money_meter2);
        money_meter.setText(": " + money);
        List<Item> ItemList = new ArrayList<>();
        List<Item> ItemList2 = new ArrayList<>();
        ItemList.add(new Item("Heart2", 120,"Hunger20","Un coeur cru, ton Tamagotchi va adorer. Faim +20"));
        ItemList.add(new Item("Mushroom", 300,"Happy20","Un p'tit champi sa ne fait de mal à personne. Joie +20"));
        ItemList.add(new Item("Seringue", 700,"Happy50","Un cocktail de stimulants naturels qui boost vos humeurs (Garanti sans effets secondaires). Joie +50"));
        ItemList.add(new Item("Lonlon", 800,"Hunger50","Le bon lait frais & délicieux directement sorti du Ranch Lon Lon. Faim +50"));
        ItemList.add(new Item("Tomato", 1150,"Hunger80","Une tomate tous les matins et plus besoin de médecin. Faim +80"));
        ItemList.add(new Item("Nanotech", 1500,"Happy80","Besoin de faire le plein de Nanotech? Joie +80"));
        ListView shopListView = findViewById(R.id.shop_list_view2);
        shopListView.setAdapter(new ItemAdapter(this,ItemList,ItemList2));
    }
    static void setMoney_meter(String name){money_meter.setText(name);}

}

