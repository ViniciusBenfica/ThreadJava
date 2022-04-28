public class Patrocinadora extends Thread {
    private Conta conta;
    private Threads thread1;
    private Threads thread2;
    private Threads thread3;
    Thread t;

    public Patrocinadora(Conta conta, Threads thread1, Threads thread2, Threads thread3) {
        this.conta = conta;
        this.thread1 = thread1;
        this.thread2 = thread2;
        this.thread3 = thread3;
        t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {
        try {
            while (true) {
                if (thread1.pausado && thread2.pausado && thread3.pausado) {
                    conta.deposito(100);
                    System.out.println("Thread Patrocinadora" + ": R$" + conta.getSaldo());
                    thread1.setpausado(false);
                    thread2.setpausado(false);
                    thread3.setpausado(false);
                    synchronized (this.conta) {
                        conta.notifyAll();
                    }
                }
                Threads.sleep(1000);
            }
        } catch (InterruptedException e) {
            System.out.println("deu merda");
        }
    }
}