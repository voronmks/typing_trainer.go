# typing_trainer.rb — Ruby версия

def generate_sequence(length)
  chars = '0123456789'
  length.times.map { chars[rand(chars.length)] }.join
end

def print_highlight(target, user_input)
  print "\r\033[K"
  target.each_char.with_index do |ch, i|
    color = if i < user_input.length
              if user_input[i] == ch
                "\033[32m"
              else
                "\033[31m"
              end
            else
              "\033[33m"
            end
    print "#{color}#{ch}\033[0m "
  end
  puts
end

puts "\033[36m⌨️  Тренажёр печати (цифры) (Ruby)\033[0m"
print "Длина упражнения (по умолч. 20): "
length_input = gets.chomp
length = length_input.empty? ? 20 : length_input.to_i

puts "Нажмите Enter, когда будете готовы..."
gets

target = generate_sequence(length)
puts "\033[33mВведите:\033[0m #{target.chars.join(' ')}"

puts "Вводите цифры последовательно (Enter после каждой):"
user_input = ''
start = Time.now

while user_input.length < target.length
  print "> "
  ch = gets.chomp
  next if ch.empty?
  user_input << ch[0]
  print_highlight(target, user_input)
end

elapsed = Time.now - start
correct = 0
user_input.each_char.with_index do |u, i|
  correct += 1 if i < target.length && u == target[i]
end
accuracy = (correct.to_f / target.length) * 100
wpm = target.length / (elapsed / 60.0)

puts "\n\033[36mСтатистика:\033[0m"
puts "  Время: #{elapsed.round(1)} сек"
puts "  Скорость: #{wpm.round(1)} зн/мин"
puts "  Точность: #{accuracy.round(1)}%"
puts "  Ошибок: #{target.length - correct}"

File.open("typing_stats.txt", "a") do |f|
  f.puts "Ruby\t#{elapsed.round(1)}\t#{wpm.round(1)}\t#{accuracy.round(1)}\t#{target.length - correct}"
end
