public class DiningPhiloshophersND {
    //PROBLEMA SEM DEADLOCK
    public static void main(String[] args){

        Filosofo1[] filosofos = new Filosofo1[5];//cria os filosofos
        Garfo1[] garfos = new Garfo1[5]; //cria os garfos
        for (int i=0; i < garfos.length; i++){
            garfos[i] = new Garfo1();//instanciando os garfos
        }
        for (int i=0; i < filosofos.length; i++){
            //criando os atributos do filosofo
            Garfo1 left = garfos[i];
            Garfo1 right = garfos[(i+1)%garfos.length];
            String nome = "Filosofo-" + i;
            if (i == filosofos.length-1){
                //resolvemos o deadlock fazendo um filoso pegar primeiro seu
                //garfo da direita
                filosofos[i] = new Filosofo1(nome,right, left);
            }else{
                filosofos[i] = new Filosofo1(nome,left, right);
            }


            Thread t = new Thread(filosofos[i], nome);
            t.start();
        }
    }

}

class Garfo1{

}

class Filosofo1 implements Runnable{
    private Garfo1 left;
    private Garfo1 right;
    private String nome;

    public Filosofo1(String nome, Garfo1 garfo1, Garfo1 garfo2){
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
        Thread.sleep( (((int)(Math.random() * 100))%3000));
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
