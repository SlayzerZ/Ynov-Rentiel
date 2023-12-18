package com.example.tamagotchi;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;
import android.widget.Toast;

import java.util.List;

public class ItemAdapter extends BaseAdapter {
    private Context context;
    private List<Item> ItemList;
    private List<Item> ItemList2;
    private LayoutInflater inflater;
    int price = SecondActivity.getMoney();
    public ItemAdapter(Context context, List<Item> ItemList, List<Item> ItemList2){
        this.context = context;
        this.ItemList = ItemList;
        this.ItemList2 = ItemList2;
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
        View view = inflater.inflate(R.layout.adapt_list3, null);
        Item currentItem = (Item) getItem(position);

        String itemName = currentItem.getName();
        double itemPrice = currentItem.getPrice();
        String imgName = currentItem.getName().toLowerCase();
        String type = currentItem.getType();
        String itemDes = currentItem.getDescription();

        TextView itemNameView = view.findViewById(R.id.item_name);
        itemNameView.setText(itemName);
        TextView itemPriceView = view.findViewById(R.id.item_price);
        itemPriceView.setText(itemPrice + "r");
      //  TextView itemDesView = view.findViewById(R.id.item_price3);
       // itemDesView.setText(itemDes);
        ImageView itemIconView = view.findViewById(R.id.item_icon);
        String ressourceName = imgName;
        int resId = context.getResources().getIdentifier(ressourceName,"drawable",context.getPackageName());
        itemIconView.setImageResource(resId);
        Button buy = view.findViewById(R.id.useButton);
        buy.setOnClickListener(v -> {
            Item ajout = new Item(imgName,itemPrice,type,itemDes);
            int a = (int) (price - currentItem.getPrice());
            int b = (int) (currentItem.getPrice() - price);
            if (a < 0){
                Toast.makeText(context, "Il vous manque " + b + "r",Toast.LENGTH_SHORT).show();
            } else {
                SecondActivity.setMoney(a);
                price = SecondActivity.getMoney();
                Shop.setMoney_meter(": " + price);
                SecondActivity.inventoryList.add(ajout);
               // SecondActivity.inventoryList = ItemList2;
            }
        });
        return view;
    }
}
