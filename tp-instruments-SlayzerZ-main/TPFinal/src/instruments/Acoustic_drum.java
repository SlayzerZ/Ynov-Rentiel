package instruments;

import java.util.ArrayList;
import java.util.Scanner;

public class Acoustic_drum extends Drum {
	Scanner input = new Scanner(System.in);
	public Mount mount;
	public Acoustic_drum() {
		this.mount = Mount.UNMOUNT;
	}
	int choice;
	public void Add() {
		Percussion newPercussion;
		System.out.println("1 Drum bass");
		System.out.println("2 Snare bass");
		System.out.println("3 Tom");
		System.out.println("4 Charleston");
		System.out.println("5 Cymbal");
		System.out.println("6 Quit");
		choice = input.nextInt();
		switch(choice) {
		case 1 :
			newPercussion = new Bass_drum();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 2 :
			newPercussion = new Snare_drum();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 3 :
			newPercussion = new Tom();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 4 :
			newPercussion = new Charleston();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 5 :
			newPercussion = new Cymbal();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		}
		
	}
	public void Remove() {
		System.out.print("ID : ");
		int id2 = input.nextInt();
		int position = find(percussions, id2);
		if(-1 == position) {
			System.out.println("Cet ID n'existe pas");
		}else {
			percussions.remove(position);
		}
	}
	public void initialize() {
		super.initialize();
	}
	public void modify() {
		super.modify();
		Scanner input = new Scanner(System.in);
		System.out.println("1 Modify Decoration");
		System.out.println("2 Add percussion");
		System.out.println("3 Remove percussion");
		System.out.println("4 Quit");
		choice = input.nextInt();
		input.nextLine();
		switch (choice) {
		case 1:
			for (Percussion percussion1 : percussions) {
				if (percussion1 instanceof Bass_drum) {
					System.out.println("New Decoration : ");
					((Bass_drum) percussion1).setDecoration(input.nextLine());
				}
			}
			break;
		case 2:
			Add();
			System.out.println("Percussion added");
			break;
		case 3:
			Remove();
			System.out.println("Percussion removed");
			break;
	}
	}
	public String print() {
		return super.print();
	}
	public String print2() {
		return super.print2() + "\nType : Acoustic Drum";
	}
}
