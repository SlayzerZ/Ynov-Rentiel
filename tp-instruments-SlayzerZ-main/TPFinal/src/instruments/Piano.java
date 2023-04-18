package instruments;

import java.util.Scanner;

public class Piano extends Keyboard_instruments implements Tuneable {
	String hammer;
	int tuning_price;
	int tune_days = 26;
	int tune_month = 1;
	int tune_year = 2023;
	public Piano(int id, int price, String brand, String model, int key_number, String hammer) {
		super(id, price, brand, model, key_number);
		this.hammer = hammer;
		this.tuning_price = tuning_price;
		this.tune_days = tune_days;
		this.tune_month = tune_month;
		this.tune_year = tune_year;
	}
	public Piano() {
		
	}
	
	public int getTuning_price() {
		return this.tuning_price;
	}
	public void setTuning_price(int tuning_price) {
		this.tuning_price = tuning_price;
	}
	public String getHammer() {
		return this.hammer;
	}
	public void setHammer(String hammer) {
		this.hammer = hammer;
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
	public String getTune() {
		return  this.tune_days + "/" + this.tune_month + "/" + this.tune_year;
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Tuning Price : ");
		this.tuning_price = input.nextInt();
		input.nextLine();
		System.out.print("Hammer : ");
		this.hammer = input.nextLine();
	}
	@Override
	public void tuneable() {
		Scanner input = new Scanner(System.in);
		String newString = input.nextLine();
		setHammer(newString);
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
	public String print() {
		return super.print() + "\nTuning Price : " + getTuning_price() + "\nHammer : " + getHammer() + "\nTune Date : " + getTune();
	}
	public String print2() {
		return super.print2() + "\nType : Piano"; 
}

}
