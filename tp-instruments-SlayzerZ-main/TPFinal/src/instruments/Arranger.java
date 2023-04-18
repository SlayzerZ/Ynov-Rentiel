package instruments;

import java.util.Scanner;

public class Arranger extends Keyboard_instruments {
private int sound_number;
	public Arranger(int id, int price, String brand, String model, int key_number, int sound_number) {
		super(id, price, brand, model, key_number);
		this.sound_number = sound_number;
	}
public Arranger() {
		
	}
	public int getsound_number() {
		return sound_number;
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Sound Number : ");
		this.sound_number = input.nextInt();
	}
	public String print() {
		return super.print() + "\nSound Number : " + getsound_number();
	}
	public String print2() {
		return super.print2() + "\nType : Arranger";
	}
	
}
