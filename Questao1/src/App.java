public class App {
    public static void main(String[] args) throws Exception {

        Conta conta = new Conta(123, "titular");
        Threads aGastadora = new Threads(conta, 50, 6000, "AGastadora");
        Threads aEsperta = new Threads(conta, 10, 3000, "AEsperta");
        Threads aEconomica = new Threads(conta, 5, 12000, "AEconomica");
        Patrocinadora aPatrocinadora = new Patrocinadora(conta, aGastadora, aEsperta, aEconomica);
        
        try {
            aGastadora.t.join();
            aEsperta.t.join();
            aEconomica.t.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        // System.out.println(conta);
    }
}
