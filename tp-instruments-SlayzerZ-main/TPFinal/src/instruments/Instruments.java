package instruments;

import java.util.Scanner;

public abstract class Instruments {
	protected int price;
	private String brand;
	private String model;
	public Exposition exposition;
	public int id;
	int choice;

	public Instruments(int id, int price, String brand, String model) {
		this.id = id;
		this.price = price;
		this.brand = brand;
		this.model = model;
		this.exposition = Exposition.RESERVE;
	}

	public Instruments() {

	}

	public int getPrice() {
		return price;
	}

	public int getId() {
		return id;
	}

	public String getBrand() {
		return brand;
	}

	public String getModel() {
		return model;
	}

	public Exposition getExposition() {
		return exposition;
	}

	public void setPrice(int price) {
		this.price = price;
	}

	public void setExposition() {
		if (this.exposition == Exposition.RESERVE) {
			this.exposition = Exposition.EXPOSITION;
		}else {
			this.exposition = Exposition.RESERVE;
		}
		System.out.println("Status changed");
	}

	public void initialize() {
		Scanner input = new Scanner(System.in);
		System.out.println("Price : ");
		this.price = input.nextInt();
		input.nextLine();

		System.out.println("Brand : ");
		this.brand = input.nextLine();

		System.out.println("Model : ");
		this.model = input.nextLine();
		this.exposition = Exposition.RESERVE;

	}

	public void modify() {
		Scanner input = new Scanner(System.in);
		System.out.println("1 Change Price");
		System.out.println("2 Change Exposition");
		System.out.println("3 Quit");
		choice = input.nextInt();
		switch (choice) {
		case 1:
			System.out.print("New Price : ");
			setPrice(input.nextInt());
			break;
		case 2:
			setExposition();
			break;
		}
	}

	public String print() {
		return "ID : " + getId() + "\nPrice : " + getPrice() + "\nBrand : " + getBrand() + "\nModel : " + getModel()
				+ "\nExposition : " + this.exposition;
	}
	public String print2() {
		return "ID : " + getId() + "\nPrice : " + getPrice() + "\nExposition : " + this.exposition;
	}
	
}
