class Animal {
  constructor(nome, idade, especie) {
      this.nome = nome;
      this.idade = idade;
      this.especie = especie;
  }

  printInfo() {
      console.log(`Nome: ${this.nome}, Idade: ${this.idade}, Espécie: ${this.especie}`);
  }
}

class Cachorro extends Animal {
  #raca;

  constructor(nome, idade, especie, raca) {
      super(nome, idade, especie);
      this.#raca = raca;
  }

  printInfo() {
      super.printInfo();
      console.log(`Raça: ${this.#raca}`);
  }
}

class Gato extends Animal {
  constructor(nome, idade, especie, cores) {
      super(nome, idade, especie);
      this.cores = [...cores];
  }

  printInfo() {
      super.printInfo();
      console.log(`Cores: ${this.cores.join(", ")}`);
  }
}

const animal = new Animal("Leão", 5, "Felino");
const cachorro = new Cachorro("Rex", 3, "Cachorro", "Labrador");
const gato = new Gato("Miau", 2, "Gato", ["Branco", "Cinza", "Preto"]);

animal.printInfo();
cachorro.printInfo();
gato.printInfo();
