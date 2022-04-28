
public class Conta {
    private int numero;
    private String titular;
    private double saldo = 1000.00;

    public Conta(int numero, String titular) {
        this.numero = numero;
        this.titular = titular;
    }

    public int getNumero() {
        return numero;
    }

    public double getSaldo() {
        return saldo;
    }

    public String getTitular() {
        return titular;
    }

    public void setTitular(String titular) {
        this.titular = titular;
    }

    public void setNumero(int numero) {
        this.numero = numero;
    }

    public synchronized double deposito(double valor) {
        return this.saldo += valor;
    }

    public boolean saque(double valor) {
        System.out.println(valor);
        if (this.saldo - valor >= 0) {
            this.saldo = this.saldo - valor;
            return true;
        }
        return false;
    }
}
