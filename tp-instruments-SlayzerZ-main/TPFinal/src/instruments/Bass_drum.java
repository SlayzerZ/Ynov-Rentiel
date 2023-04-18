package instruments;

import java.util.Scanner;

public class Bass_drum extends Percussion {
	public Type type;
	public String texte;
	
	public Bass_drum(int id, int price, String brand, String model, int diameter, String matter, String texte) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ACOUSTIC;
		this.texte = texte;
	}

	public Bass_drum() {
		this.type = Type.ACOUSTIC;
	}
	public void setDecoration(String texte) {
		this.texte = texte;
	}
	@Override
	public void sound() {
		System.out.println("Boom");
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Decoration : ");
		this.texte = input.nextLine();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nDecoration : " + this.texte + "\nPercussion : Bass_Drum";
	}
}
