package instruments;

import java.util.Scanner;

public abstract class Percussion extends Instruments {
	public int diameter;
	public String matter;

	public Percussion(int id, int price, String brand, String model, int diameter, String matter) {
		super(id, price, brand, model);
		this.diameter = diameter;
		this.matter = matter;
	}
	
	public Percussion() {
		
	}

	public abstract void sound();

	public int getDiameter() {
		return diameter;
	}

	public String getMatter() {
		return matter;
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Diameter : ");
		this.diameter = input.nextInt();
		input.nextLine();
		System.out.print("Matter : ");
		this.matter = input.nextLine();
	}
	public String print() {
		return super.print() + "\nDiameter : " + getDiameter() + "\nMatter : " + getMatter();
	}
}