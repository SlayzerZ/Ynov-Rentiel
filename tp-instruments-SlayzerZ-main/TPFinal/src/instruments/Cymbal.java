package instruments;

public class Cymbal extends Percussion {
	public Type type;
	public Cymbal(int id, int price, String brand, String model, int diameter, String matter) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ACOUSTIC;
	}
	public Cymbal() {
		this.type = Type.ACOUSTIC;
	}
	@Override
	public void sound() {
		System.out.println("Splash");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Cymbal";
	}
}
