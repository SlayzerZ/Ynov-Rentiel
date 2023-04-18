package instruments;

import java.util.ArrayList;
import java.util.Scanner;

public class Menu {
	private static int find(ArrayList<Instruments> instruments, int id) {
		int i = 0;
		for (Instruments instruments1 : instruments) {
			if (instruments1.getId() == id) {
				return i;
			}
			i++;
		}
		return -1;
	}

	public static void main(String[] args) {
		Scanner input = new Scanner(System.in);

		ArrayList<Instruments> instruments = new ArrayList<Instruments>();

		Guitar guitar = new Guitar(1, 54, "Mario", "Zeldo", 55, "Wood", 5, "Min", "Back");

		instruments.add(guitar);

		int choice;

		do {
			System.out.println("1 View all instruments");
			System.out.println("2 View a instruments");
			System.out.println("3 Add a instruments");
			System.out.println("4 Modify a instruments");
			System.out.println("5 Delete a instruments");
			System.out.println("6 Quit");

			choice = input.nextInt();
			switch (choice) {
			case 1:
				System.out.println("Total : " + instruments.size());
				System.out.println();
				for (Instruments instruments1 : instruments) {
					System.out.println(instruments1.print2());
				}
				break;
			case 2:
				System.out.print("ID : ");
				int id = input.nextInt();
				int position = find(instruments, id);
				if (-1 == position) {
					System.out.println("That ID doesn't exist");
				} else {
					System.out.println(instruments.get(position).print());
				}
				break;
			case 3:
				System.out.println("Do you want to add :");
				System.out.println("1 A String instruments ?");
				System.out.println("2 A keyboard instruments ?");
				System.out.println("3 A Drum ?");
				System.out.println("4 Quit ?");
				choice = input.nextInt();
				Instruments newInstrument;
				switch (choice) {
				case 1:
					System.out.println("Do you want to add :");
					System.out.println("1 A Guitar ?");
					System.out.println("2 A Violin ?");
					System.out.println("3 Quit ?");
					choice = input.nextInt();
					switch (choice) {
					case 1:
						newInstrument = new Guitar();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					case 2:
						newInstrument = new Violin();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					}
					break;
				case 2:
					System.out.println("Do you want to add :");
					System.out.println("1 A Piano ?");
					System.out.println("2 A Arranger ?");
					System.out.println("3 Quit ?");
					choice = input.nextInt();
					switch (choice) {
					case 1:
						newInstrument = new Piano();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					case 2:
						newInstrument = new Arranger();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					}
					break;
				case 3:
					System.out.println("Do you want to add :");
					System.out.println("1 A Acoustic Drum ?");
					System.out.println("2 A Electronic Drum ?");
					System.out.println("3 Quit ?");
					choice = input.nextInt();
					switch (choice) {
					case 1:
						newInstrument = new Acoustic_drum();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					case 2 :
						newInstrument = new Electronic_drum();
						newInstrument.initialize();
						newInstrument.id = instruments.size() +1;
						instruments.add(newInstrument);
						break;
					}
				break;
				}
				System.out.println("Added");
				break;
			case 4 :
				System.out.print("ID : ");
				id = input.nextInt();
				position = find(instruments, id);
				if(-1 == position) {
					System.out.println("That ID doesn't exist");
				}else {
					instruments.get(position).modify();
				}
				break;
			case 5 :
				System.out.print("ID : ");
				id = input.nextInt();
				int position1 = find(instruments, id);
				if(-1 == position1) {
					System.out.println("That ID doesn't exist");
				}else {
					instruments.remove(position1);
					System.out.println("Deleted");
				}
				break;
			}
		} while (choice != 6);
		input.close();
}
}
