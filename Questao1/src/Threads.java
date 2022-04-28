public class Threads extends Thread {
    private Conta conta;
    private int valor;
    private int millis;
    private String name;
    private int saques = 0;
    public boolean pausado = false;
    Thread t;

    public Threads(Conta conta, int valor, int millis, String name) {
        this.conta = conta;
        this.valor = valor;
        this.millis = millis;
        this.name = name;
        t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {
        try {
            while (conta.getSaldo() >= 0) {
                synchronized (this.conta) {
                    if (conta.saque(this.valor)) {
                        System.out.println("Thread " + this.name + ": R$" + conta.getSaldo() + "");
                        this.saques = this.saques + 1;
                    } else {
                        System.out.println("PARANDO A THREAD " + this.name + "\nSAQUES " + this.saques);
                        this.pausado = true;
                        conta.wait();
                    }
                }
                Threads.sleep(this.millis);
            }
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }


    public void setpausado(Boolean pausado) {
        this.pausado = pausado;
    }
}