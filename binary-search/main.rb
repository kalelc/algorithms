def binary_search(elements, value)
  low = 0

  high = elements.size - 1 # the most high index in the elements (array)

  while low <= high
    middle = (low + high) / 2

    puts middle

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

elements = [1,2,3,4,5,10,25,50,100]

# get de position where elements exist
puts binary_search(elements, 10)
