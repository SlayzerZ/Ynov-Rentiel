package com.example.tamagotchi;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.MotionEvent;
import android.view.View;

import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.snackbar.BaseTransientBottomBar;
import com.google.android.material.snackbar.Snackbar;

import java.util.LinkedList;
import java.util.List;

public class MorpionGame2 extends AppCompatActivity implements View.OnTouchListener {
    private MorpionBoard plateau;
    private int taille;
    int coupJ;
    int coupT;
    int money = SecondActivity.getMoney();
    Nombre N1 = new Nombre(0);
    Nombre N2 = new Nombre(0);
    private List<Pion> historique = new LinkedList<>();
    private boolean joueurCourant =false;

    public static final String EXTRA_SIZE = "oc.demos.morpion.taille";

    @Override
    protected void onCreate(Bundle savedInstanceState) {

            super.onCreate(savedInstanceState);
            // Récupération du paramètre d'appel
            Intent intent = getIntent();
            taille = intent.getIntExtra(EXTRA_SIZE, 3);
            coupJ = 0;
            coupT = taille*taille;
            plateau = new MorpionBoard(this, taille, historique);
            setContentView(plateau);
            // Toast.makeText(this, taille, Toast.LENGTH_LONG).show();
            plateau.setOnTouchListener(this);

    }


    /** Ajoute un pion en coordonnées (x,y) pour le joueur j.
     *  retourne vrai ssi cet ajout a pu être fait (la case était libre).
     */
    public boolean addPion(int x, int y, boolean j) {
        // On ne peut pas jouer dans une place occupée
        for(Pion c : historique){
            if(c.x==x && c.y==y) return false;
        }

        // On mémorise le coup
        historique.add(new Pion(x, y, j));
        coupJ++;
        return true;
    }
    public void addPion2(boolean j) {
        // On ne peut pas jouer dans une place occupée
        boolean same = false;
            N1.RandomNumber(0,taille-1);
            N2.RandomNumber(0,taille-1);
            int x = N1.getNumber();
            int y = N2.getNumber();
            for(Pion c : historique){
                if(c.x==x && c.y==y){
                    same = true;
               } else {
                    same = false;
                   historique.add(new Pion(x, y, j));
                }
            }
        // On mémorise le coup

        //joueurCourant = !joueurCourant;
    }

    /** Lorsqu'une touche de l'écran est détectée, un coup est joué si c'est possible.
     *  Cette méthode provoque également le raffraichissement de l'affichage du plateau
     *  et le passage au joueur suivant si le coup proposé est valide et que le jeu
     *  n'est pas fini.
     * */
    public boolean onTouch(View view, MotionEvent motionEvent) {
        if(motionEvent.getAction()==MotionEvent.ACTION_UP) {
            int x = (int)(motionEvent.getX()*taille/plateau.getWidth());
            int y = (int)(motionEvent.getY()*taille/plateau.getHeight());
           // Toast.makeText(this, taille, Toast.LENGTH_LONG).show();
            if(addPion(x, y, joueurCourant)) {
                if (!estFini(x, y)) {
                    joueurCourant = !joueurCourant;
                    //addPion2(joueurCourant);
                    //joueurCourant = !joueurCourant;
                }
            }
        }

        plateau.postInvalidate();
        return true;
    }


    /** retourne vrai ssi il existe un pion en (x,y) et qu'il est
     *  du même symbole que celui du joueur courant. */
    private boolean estMemeQueCourant(int x, int y) {
        for(Pion c:historique) {
            if (c.x==x && c.y==y & c.estCroix==joueurCourant) return true;
        }
        return false;
    }

    /** Retourne le nombre de pions du joueur courant qui sont alignés
     *  par rapport à une direction donnée (dx, dy).
     *  La valeur renvoyée considère les deux sens d'une même direction.
     */
    private int compteAlignementDeCourants(int x, int y, int dx, int dy) {
        int c = 0;
        x+=dx;
        y+=dy;
        while(estMemeQueCourant(x,y)){
            x+=dx;
            y+=dy;
            c++;
        }
        return c;
    }

    /** Retourne vrai ssi un alignement de 3 existe pour le joueur courant.
     Si le jeu est fini, cette méthode désactive l'écouteur de touche (bloque le jeu)
     et transmet la solution trouvée au plateau.
     */
    private boolean estFini(int x, int y) {
        int[][] directions = {
                {1,0}, // Horizontal
                {0,1}, // Vertical
                {1,1}, // Diagonale \
                {-1,1} // Diagonale /
        };

        for(int d = 0;d<directions.length; d++) {
            int dx=directions[d][0];
            int dy=directions[d][1];
            int c1 = compteAlignementDeCourants(x,y,dx,dy);
            int c2 = compteAlignementDeCourants(x,y,-dx,-dy);
            int c=1+c1+c2;
            Log.d("estFini","Direction "+d+" : "+c);
            if(c>=taille) {
                int[] solution = {x+dx*c1, y+dy*c1,
                        x-dx*c2, y-dy*c2};

                plateau.setVainqueur(joueurCourant, solution);
                plateau.setOnTouchListener(null);
                if (joueurCourant){
                    try {
                        money = money + (3 * taille);
                        SecondActivity.setMoney(money);
                        SecondActivity.Tama.setHunger(SecondActivity.Tama.getHunger() - coupJ);
                        //Intent nextscreen = new Intent(this, SecondActivity.class);
                        finish();
                        //startActivity(nextscreen);
                    } catch (Exception e){
                        Snackbar mySnackbar = Snackbar.make(findViewById(R.id.myCoordinatorLayout),e.getMessage(), BaseTransientBottomBar.LENGTH_LONG);
                        mySnackbar.show();
                    }
                } else {
                    try {
                        money = money + (300 * taille);
                        SecondActivity.setMoney(money);
                        SecondActivity.Tama.setHunger(SecondActivity.Tama.getHunger() - coupJ);
                        //Intent nextscreen = new Intent(this, SecondActivity.class);
                        finish();
                        //startActivity(nextscreen);
                    } catch (Exception e){
                        Snackbar mySnackbar = Snackbar.make(findViewById(R.id.myCoordinatorLayout),e.getMessage(), BaseTransientBottomBar.LENGTH_LONG);
                        mySnackbar.show();
                    }
                }
                return true;
            } else if (coupJ >= coupT) {
                try {
                    money = money + (30 * taille);
                    SecondActivity.setMoney(money);
                    SecondActivity.Tama.setHunger(SecondActivity.Tama.getHunger() - coupJ);
                    //Intent nextscreen = new Intent(this, SecondActivity.class);
                    finish();
                    //startActivity(nextscreen);
                } catch (Exception e){
                    Snackbar mySnackbar = Snackbar.make(findViewById(R.id.myCoordinatorLayout),e.getMessage(), BaseTransientBottomBar.LENGTH_LONG);
                    mySnackbar.show();
                }
            }
        }
        return false;
    }
}
