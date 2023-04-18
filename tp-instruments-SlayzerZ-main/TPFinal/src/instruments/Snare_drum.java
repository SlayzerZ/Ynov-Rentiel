package instruments;

public class Snare_drum extends Percussion {
	public Type type;
	public Snare_drum(int id, int price, String brand, String model, int diameter, String matter) {
		super(id, price, brand, model, diameter, matter);
		this.type = Type.ACOUSTIC;
	}
	public Snare_drum() {
		this.type = Type.ACOUSTIC;
	}
	@Override
	public void sound() {
		System.out.println("Bam");
	}
	public void initialize() {
		super.initialize();
	}
	public String print() {
		return super.print() + "\nType : " + this.type + "\nPercussion : Snare_Drum";
	}
}
