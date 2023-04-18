package instruments;

public class Tom_pad extends Percussion {
public Type type;
	public Tom_pad(int id, int price, String brand, String model, int diameter, String matter, Type type) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ELECTRONIC;
	}
	public Tom_pad() {
		this.type = Type.ELECTRONIC;
	}
	@Override
	public void sound() {
		System.out.println("Dom");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Tom_pad";
	}
}
