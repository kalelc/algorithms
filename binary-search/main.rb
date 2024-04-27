def binary_search(elements, value)
  low = 0

  high = elements.size - 1 # the most high index in the elements (array)

  while low <= high
    middle = (low + high) / 2

    guess = elements[middle] # get the middle-middle

    if guess == value # check if the middle-value is equals to value searched
      return middle
    elsif guess > value
      high = middle - 1
    else
      low = middle + 1
    end
  end
  nil
end

print 'Enter the number to search: '
number = gets.chomp.to_i

puts '(if not enter value, the defult list is: 5,10,15,20,50,100)'
print 'enter array to search number example: 1,2,3,4,5,6: '
list = gets.chomp

numbers = list.nil? ? [1, 2, 3, 4, 5, 6] : list.split(',').map { |x| x.to_i }

# get de position where elements exist
position = binary_search(numbers, number)

if position.nil?
  puts 'the value does not exist in the list'
else
  puts "the value is the position #{position}"
end
