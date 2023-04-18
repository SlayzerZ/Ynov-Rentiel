package instruments;

import java.util.Scanner;

public class Violin extends String_instruments {
	public String bowtype;
	public Violin(int id, int price, String brand, String model, int string_number, String body_type, int tune_number, String microphone, String bowtype) {
		super(id, price, brand, model, string_number, body_type);
		this.bowtype = bowtype;
	}
	public Violin() {
		
	}
	public String getBowtype() {
		return bowtype;
	}
	public void setBowtype(String bowtype) {
		this.bowtype = bowtype;
	}
	public String print() {
		return super.print() + "\nBow type : " + getBowtype();
	}
	public String print2() {
		return super.print2() + "\nType : Violon";
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Bow type : ");
		this.bowtype = input.nextLine();
	}
	@Override
	public void tuneable() {
		Scanner input = new Scanner(System.in);
		System.out.print("Bow type : ");
		String newBowtype = input.nextLine();
		setBowtype(newBowtype);
		update_tuneable();
	}
	public void modify() {
		super.modify();
		Scanner input = new Scanner(System.in);
		System.out.println("1 Tune");
		System.out.println("2 Quit");
		choice = input.nextInt();
		input.nextLine();
		switch (choice) {
		case 1:
			tuneable();
			break;
	}
	}
}
