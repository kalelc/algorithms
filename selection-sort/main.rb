# @param elements array
# #return index integer
def find_smallest(elements)
  smallest = elements[0] # get tue smallest number to start with comparing
  smallest_index = 0 # start index value in the first index element to allow comparing
  i = 1

  while i < elements.size
    if elements[i] < smallest
      smallest = elements[i]
      smallest_index = i
    end

    i += 1
  end

  smallest_index
end

# selection sort create new space array @(2n)
# also it runs in O(n^2) time complexity
def selection_sort(elements)
  new_elements = []
  i = 0
  while i < elements.size
    smallest = find_smallest(elements)
    new_elements << elements[smallest]

    elements.delete_at(smallest)
    puts elements.inspect
    i += 1
  end

  puts new_elements.inspect

  new_elements
end

puts [5, 3, 6, 2, 10].inspect
selection_sort([5, 3, 6, 2, 10])
