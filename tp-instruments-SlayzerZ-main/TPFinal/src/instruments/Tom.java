package instruments;

public class Tom extends Percussion {
	public Type type;
	public Tom(int id, int price, String brand, String model, int diameter, String matter) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ACOUSTIC;
	}
	public Tom() {
		this.type = Type.ACOUSTIC;
	}
	@Override
	public void sound() {
		System.out.println("Bim");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Tom";
	}
}
