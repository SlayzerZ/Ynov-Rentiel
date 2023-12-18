package com.example.tamagotchi;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.SeekBar;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.snackbar.BaseTransientBottomBar;
import com.google.android.material.snackbar.Snackbar;

public class MorpionMenu extends AppCompatActivity {
    private Button Btn1,Btn2;
    private SeekBar choixTaille;
    private TextView tvTaille;

    private final int MIN_TAILLE = 3;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.morpion_main);

        tvTaille = (TextView)findViewById(R.id.taille);
        choixTaille = (SeekBar) findViewById(R.id.choixTaille);
        Btn1 = (Button)findViewById(R.id.Btn1);
        Btn2 = (Button)findViewById(R.id.Btn2);
        tvTaille.setText((MIN_TAILLE)+"x"+(MIN_TAILLE));
        choixTaille.setProgress(0);

        choixTaille.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {
                tvTaille.setText((MIN_TAILLE +i)+"x"+(MIN_TAILLE +i));
            }
            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {}
            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {}
        });
        Btn2.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                int val = choixTaille.getProgress()+ MIN_TAILLE;
                go2(val);
            }
        });

        Btn1.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                int val = choixTaille.getProgress()+ MIN_TAILLE;
                go(val);
            }
        });
    }

    public void go(int taille) {
        try {
            Intent caller = new Intent(this, MorpionGame.class);
            caller.putExtra(MorpionGame.EXTRA_SIZE, taille);
            startActivity(caller);
            finish();
        }catch (Exception e){

        }

    }
    public void go2(int taille) {
        try {
            Intent caller = new Intent(this, MorpionGame2.class);
            caller.putExtra(MorpionGame.EXTRA_SIZE, taille);
            startActivity(caller);
            finish();
        }catch (Exception e){

        }

    }
}
