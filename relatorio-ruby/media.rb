class CalculadoraMedia
  def calcular_media
    numeros = []
    3.times do |i|
      puts "Digite o número #{i + 1}:"
      numeros << gets.to_f
    end
    media = numeros.sum / numeros.size
    puts "A média dos números é: #{media}"
  end
end

calculadora = CalculadoraMedia.new
calculadora.calcular_media