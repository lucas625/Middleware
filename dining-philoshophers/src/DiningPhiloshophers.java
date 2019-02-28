public class DiningPhiloshophers {
//PROBLEM WITH DEADLOCK
    public static void main(String[] args){

        Filosofo[] filosofos = new Filosofo[5];//cria os filosofos
        Garfo[] garfos = new Garfo[5]; //cria os garfos
        for (int i=0; i < garfos.length; i++){
            garfos[i] = new Garfo();//instanciando os garfos
        }
        for (int i=0; i < filosofos.length; i++){
            //criando os atributos do filosofo
            Garfo left = garfos[i];
            Garfo right = garfos[(i+1)%garfos.length];
            String nome = "Filosofo-" + i;

            filosofos[i] = new Filosofo(nome,left, right);

            Thread t = new Thread(filosofos[i], nome);
            t.start();
        }
    }

}

class Garfo {

}

class Filosofo implements Runnable{
    private Garfo left;
    private Garfo right;
    private String nome;

    public Filosofo(String nome, Garfo garfo1, Garfo garfo2){
        this.left = garfo1;
        this.right = garfo2;
        this.nome = nome;
    }

    public String getNome(){
        //retorna o nome do filosofo
        return this.nome;
    }

    public void acao(String fazendo) throws InterruptedException{
        // aqui é ele comendo ou pensando
        System.out.println(this.nome + " " + fazendo);
        Thread.sleep( 200);
    }
    
    @Override
    public void run(){
        try{
            while(true){
                acao(System.nanoTime() + ": Pensando");//aqui ele pensa

                synchronized (this.left){
                    acao(System.nanoTime()+": "+this.nome+" pegou o garfo esquerdo");//aqui ele pega um garfo
                    synchronized (this.right) {
                        acao(System.nanoTime() + ": " + this.nome + " pegou o garfo direito e está comendo");//aqui ele pegou outro garfo
                        acao(System.nanoTime() + ": "+this.nome+" soltou o garfo direito");//aqui ele solta o garfo
                    }
                    acao(System.nanoTime() + ": "+this.nome+" soltou o garfo esquerdo");//aqui ele solta o garfo
                }
            }

        }catch (InterruptedException e){
            Thread.currentThread().interrupt();
            return;
        }
    }
}
