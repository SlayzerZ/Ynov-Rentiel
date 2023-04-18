package instruments;

import java.util.Scanner;

public abstract class String_instruments extends Instruments implements Tuneable  {
	private int string_number;
	private String body_type;
	int tune_days = 26;
	int tune_month = 1;
	int tune_year = 2023;
	public String_instruments(int id, int price, String brand, String model, int string_number, String body_type) {
		super(id, price, brand, model);
		this.string_number = string_number;
		this.body_type = body_type;
		this.tune_days = tune_days;
		this.tune_month = tune_month;
		this.tune_year = tune_year;
	}
	public String_instruments() {
		
	}
	public void update_tuneable() {
	if (tune_days == 31 && tune_month == 12) {
		tune_days = 1;
		tune_month = 1;
		tune_year++;
	}else if (tune_days == 31) {
		tune_days = 1;
		tune_month++;
	}else {
		tune_days++;
	}
		System.out.println("Tuning over!");
		this.tune_days = tune_days;
		this.tune_month = tune_month;
		this.tune_year = tune_year;
	}
	public int getSN() {
		return string_number;
	}
	public String getBT() {
		return body_type;
	}
	public String getTune() {
		return  this.tune_days + "/" + this.tune_month + "/" + this.tune_year;
	}
	
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("String number : ");
		this.string_number = input.nextInt();
		input.nextLine();
		System.out.print("Body type : ");
		this.body_type = input.nextLine();

	}
//	public void modify() {
//	super.modify();
//		Scanner input = new Scanner(System.in);
//		System.out.println("1 Tune");
//		System.out.println("2 Quit");
//		choice = input.nextInt();
//		switch (choice) {
//		case 1:
//			tuneable();
//			break;
//	}
//	}
	public String print() {
		return super.print() + "\nString Number : " + getSN() + "\nBody Type : " + getBT() + "\nTune Date : " + getTune();
	}
}
