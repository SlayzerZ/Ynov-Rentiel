package instruments;

public class Cymbal_pad extends Percussion {
public Type type;
	public Cymbal_pad(int id, int price, String brand, String model, int diameter, String matter, Type type) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ELECTRONIC;
	}
public Cymbal_pad() {
	this.type = Type.ELECTRONIC;
	}
	@Override
	public void sound() {
		System.out.println("Dong");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Cymbal_pad";
	}
}
