package instruments;

import java.util.ArrayList;
import java.util.Scanner;

public abstract class Drum extends Instruments {
	public Mount mount;
	protected static int find(ArrayList<Percussion> percussion, int id) {
		int i = 0;
		for(Percussion percussion1 : percussion) {
			if(percussion1.getId() == id) {
				return i;
			}
			i++;
		}
		return -1;
	}
	public Drum() {
		this.mount = Mount.UNMOUNT;
	}
	ArrayList<Percussion> percussions = new ArrayList<Percussion>();
	public void setMount() {
		if (this.mount == Mount.UNMOUNT) {
			this.mount = Mount.MOUNT;
		}else {
			this.mount = Mount.UNMOUNT;
		}
		System.out.println("Status changed");
	}
	public void Demo() {
		if (this.mount == Mount.MOUNT) {
			for(Percussion percussion1 : percussions) {
				percussion1.sound();
			}	
		}else {
			System.out.println("Drum Unmounted!!");
		}
		}
	public void initialize() {
		super.initialize();
		this.mount = Mount.UNMOUNT;
	}
	public void modify() {
		super.modify();
		Scanner input = new Scanner(System.in);
		System.out.println("1 Mount");
		System.out.println("2 Show percussion");
		System.out.println("3 Demo");
		System.out.println("4 Quit");
		choice = input.nextInt();
		input.nextLine();
		switch (choice) {
		case 1:
			setMount();
			break;
		case 2:
			for (Percussion percussion1 : percussions) {
				System.out.println(percussion1.print());
			}
			break;
		case 3:
			Demo();
			break;
	}
	}
	public String print() {
		return super.print() + "\nMount : " + this.mount;
	}
}
