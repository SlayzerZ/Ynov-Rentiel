package com.example.tamagotchi;

import android.content.Context;
import android.content.Intent;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import java.util.List;
import java.util.concurrent.atomic.AtomicReference;

public class ItemAdapter2 extends BaseAdapter {
    private Context context;
    private List<Item> ItemList;
    private LayoutInflater inflater;
    public static boolean use2 = false;
    public ItemAdapter2(Context context, List<Item> ItemList){
        this.context = context;
        this.ItemList = ItemList;
        this.inflater = LayoutInflater.from(context);
    }
    @Override
    public int getCount() {
        return ItemList.size();
    }

    @Override
    public Object getItem(int position) {
        return ItemList.get(position);
    }

    @Override
    public long getItemId(int position) {
        return 0;
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
        View view = inflater.inflate(R.layout.adapt_list2, null);
        Item currentItem = (Item) getItem(position);
        String itemName = currentItem.getName();
        double itemPrice = currentItem.getPrice();
        String imgName = currentItem.getName().toLowerCase();
        String itemDes = currentItem.getDescription();
        AtomicReference<Button> use = new AtomicReference<>(view.findViewById(R.id.useButton));
        TextView itemNameView = view.findViewById(R.id.item_name2);
        itemNameView.setText(itemName);
        TextView itemPriceView = view.findViewById(R.id.item_price2);
        itemPriceView.setText(itemDes);
        ImageView itemIconView = view.findViewById(R.id.item_icon2);
        String ressourceName = imgName;
        int resId = context.getResources().getIdentifier(ressourceName,"drawable",context.getPackageName());
        itemIconView.setImageResource(resId);
        use.get().setOnClickListener(v -> {
            if (currentItem.getType().equals("Hunger20")){
               double hunger = SecondActivity.Tama.getHunger() + 20;
                SecondActivity.Tama.setHunger(hunger);
            }
            if (currentItem.getType().equals("Hunger50")){
                double hunger = SecondActivity.Tama.getHunger() + 50;
                SecondActivity.Tama.setHunger(hunger);
            }
            if (currentItem.getType().equals("Hunger80")){
                double hunger = SecondActivity.Tama.getHunger() + 80;
                SecondActivity.Tama.setHunger(hunger);
            }
            if (currentItem.getType().equals("Happy20")){
                double happy = SecondActivity.Tama.getHappy() + 20;
                SecondActivity.Tama.setHappy(happy);
            }
            if (currentItem.getType().equals("Happy50")){
                double happy = SecondActivity.Tama.getHappy() + 50;
                SecondActivity.Tama.setHappy(happy);
            }
            if (currentItem.getType().equals("Happy80")){
                double happy = SecondActivity.Tama.getHappy() + 80;
                SecondActivity.Tama.setHappy(happy);
            }
            ItemList.remove(currentItem);
            use2 = true;
            Toast.makeText(context, "Object used",Toast.LENGTH_SHORT).show();
           // Shop.Refresh(context);
        });
        return view;
    }
}
