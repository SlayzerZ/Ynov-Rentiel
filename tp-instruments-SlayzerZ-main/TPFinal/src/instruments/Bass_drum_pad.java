package instruments;

import java.util.Scanner;

public class Bass_drum_pad extends Percussion {
public Type type;
	public Bass_drum_pad(int id, int price, String brand, String model, int diameter, String matter, Type type) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ELECTRONIC;
	}
public Bass_drum_pad() {
	this.type = Type.ELECTRONIC;
	}
	@Override
	public void sound() {
		System.out.println("Tac");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Bass_Drum_pad";
	}
}
