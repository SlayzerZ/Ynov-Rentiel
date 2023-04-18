package instruments;

import java.util.Scanner;

public abstract class Keyboard_instruments extends Instruments {
private int key_number;

public Keyboard_instruments(int id, int price, String brand, String model, int key_number) {
	super(id, price, brand, model);
	this.key_number = key_number;
}
public Keyboard_instruments() {
	
}

public int getKey_number() {
	return key_number;
}
public void initialize() {
	super.initialize();
	Scanner input = new Scanner(System.in);
	System.out.print("Key Number : ");
	this.key_number = input.nextInt();
	input.nextLine();
}
public String print() {
	return super.print() + "\nKey Number : " + getKey_number();
}
}
