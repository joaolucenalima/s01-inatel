class Carro
  def descrever
    "Este é um carro genérico."
  end
end

class CarroEsportivo < Carro
  def descrever
    "Este é um carro esportivo, rápido e elegante."
  end
end

class CarroSedan < Carro
  def descrever
    "Este é um carro sedan, confortável e espaçoso."
  end
end

carro = Carro.new
carro_esportivo = CarroEsportivo.new
carro_sedan = CarroSedan.new

puts carro.descrever
puts carro_esportivo.descrever
puts carro_sedan.descrever