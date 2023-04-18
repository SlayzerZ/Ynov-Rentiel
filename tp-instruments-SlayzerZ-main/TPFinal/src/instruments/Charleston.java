package instruments;

public class Charleston extends Percussion {
	public Type type;
	public Charleston(int id, int price, String brand, String model, int diameter, String matter) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ACOUSTIC;
	}
public Charleston() {
	this.type = Type.ACOUSTIC;
	}
	@Override
	public void sound() {
		System.out.println("Dam");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Charleston";
	}
}
