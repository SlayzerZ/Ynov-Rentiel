package instruments;

import java.util.Scanner;

public class Guitar extends String_instruments {
	public String microphone;
	public String etring;
	public Guitar(int id, int price, String brand, String model, int string_number, String body_type, int tune_number, String microphone, String etring) {
		super(id, price, brand, model, string_number, body_type);
		this.microphone = microphone;
		this.etring = etring;
	}
public Guitar() {
		
	}
	public String getMicro() {
		return microphone;
	}
	public void setMicro(String microphone) {
		this.microphone = microphone;
	}
	
	public String getString() {
		return etring;
	}
	public void setString(String etring) {
		this.etring = etring;
	}
	public String print() {
		return super.print() + "\nMicrophone : " + getMicro() + "\nString : " + getString();
	}
	public String print2() {
		return super.print2() + "\nType : Guitar";
	}
	public void initialize() {
		super.initialize();
		Scanner input = new Scanner(System.in);
		System.out.print("Microphone : ");
		this.microphone = input.nextLine();
		System.out.print("String : ");
		this.etring = input.nextLine();
	}
	@Override
	public void tuneable() {
		Scanner input = new Scanner(System.in);
		System.out.print("String : ");
		String newString = input.nextLine();
		setString(newString);
		update_tuneable();
	}
	public void modify() {
		super.modify();
		Scanner input = new Scanner(System.in);
		System.out.println("1 Change Mic");
		System.out.println("2 Tune");
		System.out.println("3 Quit");
		choice = input.nextInt();
		input.nextLine();
		switch (choice) {
		case 1:
			System.out.print("New Mic : ");
			setMicro(input.nextLine());
			break;
		case 2:
			tuneable();
			break;
	}
	}
}
