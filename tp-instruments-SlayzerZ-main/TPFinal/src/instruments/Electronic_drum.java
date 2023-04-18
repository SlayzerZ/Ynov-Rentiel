package instruments;

import java.util.ArrayList;
import java.util.Scanner;

public class Electronic_drum extends Drum {
	Scanner input = new Scanner(System.in);
	public Mount mount;
	public Plug plug;
	public Electronic_drum() {
		this.mount = Mount.UNMOUNT;
		this.plug = Plug.UNPLUG;
	}
	int choice;
	public void setPlug() {
		if (this.plug == Plug.UNPLUG) {
			this.plug = Plug.PLUG;
		}else {
			this.plug = Plug.UNPLUG;
		}
		System.out.println("Status changed");
	}
	public void Add() {
		Percussion newPercussion;
		System.out.println("1 Drum bass pad");
		System.out.println("2 Tom pad");
		System.out.println("3 Cymbal pad");
		System.out.println("4 Quit");
		choice = input.nextInt();
		switch(choice) {
		case 1 :
			newPercussion = new Bass_drum_pad();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 2 :
			newPercussion = new Tom_pad();
			newPercussion.initialize();
			newPercussion.id = percussions.size() +10;
			percussions.add(newPercussion);
			break;
		case 3 :
			newPercussion = new Cymbal_pad();
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
		System.out.println("1 Plug");
		System.out.println("2 Add percussion");
		System.out.println("3 Remove percussion");
		System.out.println("4 Quit");
		choice = input.nextInt();
		input.nextLine();
		switch (choice) {
		case 1:
			setPlug();
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
		return super.print() + "\nPlug : " + this.plug;
	}
	public String print2() {
		return super.print2() + "\nType : Electronic Drum";
	}
}
